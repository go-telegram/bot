package bot

import (
	"context"
	"sync"

	"github.com/go-telegram/bot/models"
)

type queuedUpdate struct {
	update *models.Update
	// claim is set for polled updates.
	claim *pollClaim
	// dispatched and requestDone are set for webhook updates. dispatched is an
	// unbuffered handoff to the waiting request.
	dispatched  chan<- struct{}
	requestDone <-chan struct{}
}

type updatesRun struct {
	ctx     context.Context
	webhook bool
	queue   chan queuedUpdate
	done    chan struct{}
}

func (b *Bot) start(ctx context.Context, webhook bool) {
	if ctx.Err() != nil {
		return
	}
	if b.workers <= 0 {
		// The errors handler is user code. Run it apart so that a block or a
		// runtime.Goexit there cannot hold up or end the caller.
		go b.error("cannot start bot without update workers")
		return
	}

	runCtx, cancel := context.WithCancel(ctx)
	defer cancel()
	run := &updatesRun{ctx: runCtx, webhook: webhook, done: make(chan struct{})}
	if !b.beginUpdates(run) {
		return
	}
	defer b.endUpdates()

	var polling, workers, handlers sync.WaitGroup
	if !webhook {
		polling.Add(1)
		go func() {
			defer polling.Done()
			defer cancel()
			b.getUpdates(runCtx, run.queue)
		}()
	}
	workers.Add(b.workers)
	for i := 0; i < b.workers; i++ {
		go func() {
			defer workers.Done()
			b.waitUpdates(run, &handlers)
		}()
	}
	workers.Wait()
	// The last worker may have exited through a user callback. Wake any
	// producer waiting for a handoff before waiting for running handlers.
	cancel()
	polling.Wait()
	handlers.Wait()
}

func (b *Bot) beginUpdates(run *updatesRun) bool {
	for run.ctx.Err() == nil {
		b.updatesMx.Lock()
		previous := b.updatesRun
		if previous == nil {
			run.queue = make(chan queuedUpdate, b.updatesChanCap)
			b.updatesRun = run
			b.updatesMx.Unlock()
			return true
		}
		b.updatesMx.Unlock()
		select {
		case <-previous.done:
		case <-run.ctx.Done():
			return false
		}
	}
	return false
}

func (b *Bot) endUpdates() {
	b.updatesMx.Lock()
	run := b.updatesRun
	b.updatesRun = nil
	b.updatesMx.Unlock()
	close(run.done)
}

func (b *Bot) waitUpdates(run *updatesRun, handlers *sync.WaitGroup) {
	for {
		select {
		case <-run.ctx.Done():
			return
		case item := <-run.queue:
			handler, ok := b.prepareUpdate(run, item)
			if !ok {
				continue
			}
			invoke := func() {
				defer handlers.Done()
				handler(run.ctx, b, item.update)
			}
			handlers.Add(1)
			if b.notAsyncHandlers {
				invoke()
			} else {
				go invoke()
			}
		}
	}
}

func (b *Bot) prepareUpdate(run *updatesRun, item queuedUpdate) (HandlerFunc, bool) {
	if run.webhook {
		// The request answers 200 once it receives this handoff, so from here
		// on the update must reach its handler even if the run stops.
		select {
		case item.dispatched <- struct{}{}:
		case <-run.ctx.Done():
			return nil, false
		case <-item.requestDone:
			return nil, false
		}
	} else if run.ctx.Err() != nil || !item.claim.take() {
		// Leave it to be polled again: the run is stopping, or the poller has
		// already given up on it.
		return nil, false
	}

	return applyMiddlewares(b.findHandler(item.update), b.middlewares...), true
}
