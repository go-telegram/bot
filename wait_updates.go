package bot

import (
	"context"
	"sync"
)

// waitUpdates listen Updates channel and run ProcessUpdate
func (b *Bot) waitUpdates(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		case upd := <-b.updates:
			b.ProcessUpdate(ctx, upd)
		}
	}
}
