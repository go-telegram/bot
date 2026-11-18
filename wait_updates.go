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

	queue     chan queuedUpdate
	isWebhook bool

	// The webhook admission fields are guarded by Bot.updatesMx. webhookCtx is
	// nil once StartWebhook stops accepting updates. webhookQueued counts
	// accepted updates that workers have not yet taken for dispatch. queueSpace,
	// when set, is closed and cleared when a worker frees space.
	webhookCtx    context.Context
	webhookQueued int
	queueSpace    chan struct{}

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
			queueCap := b.updatesChanCap
			if webhook && queueCap < 1 {
				// Webhook admission sends under updatesMx and needs buffer space.
				queueCap = 1
			}
			run := &updatesRun{queue: make(chan queuedUpdate, queueCap), isWebhook: webhook, done: make(chan struct{})}
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
		if !b.takeQueued(ctx, run) {
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

// takeQueued reports whether a received update should be dispatched. A webhook
// update releases its slot only when it is dispatched, so shutdown does not wake
// waiting WebhookHandler requests just to reject them.
func (b *Bot) takeQueued(ctx context.Context, run *updatesRun) bool {
	if !run.isWebhook {
		return ctx.Err() == nil
	}

	b.updatesMx.Lock()
	defer b.updatesMx.Unlock()
	if ctx.Err() != nil {
		return false
	}
	run.webhookQueued--
	if run.queueSpace != nil {
		close(run.queueSpace)
		run.queueSpace = nil
	}
	return true
}
