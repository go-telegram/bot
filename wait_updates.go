package bot

import (
	"context"
	"time"

	"github.com/go-telegram/bot/models"
)

// waitUpdates listen Updates channel and spawn goroutines if needed. It's a simple worker pool
func (b *Bot) waitUpdates(ctx context.Context) {
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
					b.processUpdate(ctx, upd)

					const cleanupDuration = 10 * time.Second
					cleanupTicker := time.NewTicker(cleanupDuration)
					defer cleanupTicker.Stop()

					for {
						select {
						case <-ctx.Done():
							return
						case upd := <-taskQueue:
							b.processUpdate(ctx, upd)
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
