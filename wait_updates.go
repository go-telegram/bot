package bot

import (
	"context"
	"sync"
)

// waitUpdates listen Updates channel and run ProcessUpdate
func (b *Bot) waitUpdates(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		if b.processReceivedUpdatesBeforeShutdown {
			upd, ok := <-b.updates
			if !ok {
				return
			}
			b.ProcessUpdate(ctx, upd)
		} else {
			select {
			case <-ctx.Done():
				return
			case upd, ok := <-b.updates:
				if !ok {
					return
				}
				b.ProcessUpdate(ctx, upd)
			}
		}
	}

}
