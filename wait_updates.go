package bot

import (
	"context"
	"sync"
	"time"

	"github.com/go-telegram/bot/models"
)

// waitUpdates listen Updates channel and spawn goroutines if needed. It's a simple worker pool
func (b *Bot) waitUpdates(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	taskQueue := make(chan *models.Update)

	for {
		select {
		case <-ctx.Done():
			return
		case upd := <-b.updates:
			select {
			case taskQueue <- upd:
			default:
				wg.Add(1)
				go func(ctx context.Context, wg *sync.WaitGroup, taskQueue chan *models.Update) {
					defer wg.Done()

					b.ProcessUpdate(ctx, upd)

					const cleanupDuration = 10 * time.Second
					cleanupTicker := time.NewTicker(cleanupDuration)
					defer cleanupTicker.Stop()

					for {
						select {
						case <-ctx.Done():
							return
						case upd := <-taskQueue:
							b.ProcessUpdate(ctx, upd)
							cleanupTicker.Reset(cleanupDuration)
						case <-cleanupTicker.C:
							return
						}
					}
				}(ctx, wg, taskQueue)
			}
		}
	}
}
