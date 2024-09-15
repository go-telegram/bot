package bot

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
	"time"

	"github.com/go-telegram/bot/models"
)

const (
	maxTimeoutAfterError = time.Second * 5
)

type getUpdatesParams struct {
	Offset         int64          `json:"offset,omitempty"`
	Limit          int            `json:"limit,omitempty"`
	Timeout        int            `json:"timeout,omitempty"`
	AllowedUpdates AllowedUpdates `json:"allowed_updates,omitempty"`
}

type AllowedUpdates []string

// GetUpdates https://core.telegram.org/bots/api#getupdates
func (b *Bot) getUpdates(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	var timeoutAfterError time.Duration

	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		if timeoutAfterError > 0 {
			if b.isDebug {
				b.debugHandler("wait after error, %v", timeoutAfterError)
			}
			select {
			case <-ctx.Done():
				return
			case <-time.After(timeoutAfterError):
			}
		}

		params := &getUpdatesParams{
			Timeout: int((b.pollTimeout - time.Second).Seconds()),
			Offset:  atomic.LoadInt64(&b.lastUpdateID) + 1,
		}

		if b.allowedUpdates != nil {
			params.AllowedUpdates = b.allowedUpdates
		}

		var updates []*models.Update

		errRequest := b.rawRequest(ctx, "getUpdates", params, &updates)
		if errRequest != nil {
			if errors.Is(errRequest, context.Canceled) {
				return
			}
			b.error("error get updates, %w", errRequest)
			timeoutAfterError = incErrTimeout(timeoutAfterError)
			continue
		}

		timeoutAfterError = 0

		for _, upd := range updates {
			atomic.StoreInt64(&b.lastUpdateID, upd.ID)
			select {
			case <-ctx.Done():
				b.error("some updates lost, ctx done")
				return
			case b.updates <- upd:
			}
		}
	}
}

func incErrTimeout(timeout time.Duration) time.Duration {
	if timeout == 0 {
		return time.Millisecond * 100 // first timeout
	}
	timeout *= 2
	if timeout > maxTimeoutAfterError {
		return maxTimeoutAfterError
	}
	return timeout
}
