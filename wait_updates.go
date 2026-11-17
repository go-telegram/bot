package bot

import (
	"context"
	"sync"
	"sync/atomic"

	"github.com/go-telegram/bot/models"
)

// queuedUpdate carries an update to a worker. dispatched, when set, receives a
// signal once the worker hands the update to its handler. It must be buffered
// for the whole batch because the worker signals without waiting.
type queuedUpdate struct {
	update     *models.Update
	dispatched chan<- struct{}
}

type updatesRun struct {
	// dropped counts updates the workers discarded after cancellation. Only
	// StartWebhook reports it: unconfirmed polled updates are delivered again.
	// It stays first so 32-bit platforms align it for atomic access.
	dropped int64

	queue      chan queuedUpdate
	webhookCtx context.Context
	senders    sync.WaitGroup
	// done is closed when the run ends, so a later Start can proceed.
	done chan struct{}
}

func (b *Bot) beginUpdates(ctx context.Context, webhook bool) *updatesRun {
	if ctx.Err() != nil {
		return nil
	}
	if b.workers <= 0 {
		b.error("cannot start bot without update workers")
		return nil
	}

	for {
		if ctx.Err() != nil {
			return nil
		}
		b.updatesMx.Lock()
		previous := b.updatesRun
		if previous == nil {
			run := &updatesRun{queue: make(chan queuedUpdate, b.updatesChanCap), done: make(chan struct{})}
			if webhook {
				run.webhookCtx = ctx
			}
			b.updatesRun = run
			b.updatesMx.Unlock()
			return run
		}
		b.updatesMx.Unlock()
		// The previous run can still be waiting for its handlers.
		select {
		case <-previous.done:
		case <-ctx.Done():
			return nil
		}
	}
}

func (b *Bot) endUpdates() {
	b.updatesMx.Lock()
	run := b.updatesRun
	b.updatesRun = nil
	b.updatesMx.Unlock()
	close(run.done)
}

func (b *Bot) startUpdateWorkers(ctx context.Context, run *updatesRun) *sync.WaitGroup {
	workers := &sync.WaitGroup{}
	workers.Add(b.workers)
	for i := 0; i < b.workers; i++ {
		go func() {
			defer workers.Done()
			b.waitUpdates(ctx, run)
		}()
	}
	return workers
}

// waitUpdates dispatches queued updates until the queue is closed, then waits
// for the handlers it started.
func (b *Bot) waitUpdates(ctx context.Context, run *updatesRun) {
	var handlers sync.WaitGroup
	defer handlers.Wait()
	for item := range run.queue {
		if ctx.Err() != nil {
			// Keep draining so StartWebhook can report accepted updates that were never dispatched.
			atomic.AddInt64(&run.dropped, 1)
			continue
		}

		if item.dispatched != nil {
			item.dispatched <- struct{}{}
		}
		handlers.Add(1)
		b.dispatchUpdate(ctx, item.update, handlers.Done)
	}
}
