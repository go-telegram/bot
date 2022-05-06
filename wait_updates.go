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
				go func(ctx context.Context, taskQueue chan *models.Update) {
					wg.Add(1)
					b.processUpdate(ctx, wg, upd)

					const cleanupDuration = 10 * time.Second
					cleanupTicker := time.NewTicker(cleanupDuration)
					defer cleanupTicker.Stop()

					for {
						select {
						case <-ctx.Done():
							return
						case upd := <-taskQueue:
							wg.Add(1)
							b.processUpdate(ctx, wg, upd)
							cleanupTicker.Reset(cleanupDuration)
						case <-cleanupTicker.C:
							return
						}
					}
				}(ctx, taskQueue)
			}
		}
	}
}
