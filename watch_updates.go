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

// https://core.telegram.org/bots/api#getupdates

type getUpdatesParams struct {
	Offset         int64    `json:"offset,omitempty"`
	Limit          int      `json:"limit,omitempty"`
	Timeout        int      `json:"timeout,omitempty"`
	AllowedUpdates []string `json:"allowed_updates,omitempty"`
}

func (p getUpdatesParams) Validate() error {
	return nil
}

func (getUpdatesParams) GetReplyMarkup() models.ReplyMarkup {
	return nil
}

func (b *Bot) watchUpdates(ctx context.Context, wg *sync.WaitGroup, updatesChan chan<- *models.Update) {
	defer wg.Done()

	var timeoutAfterError time.Duration

	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		if timeoutAfterError > 0 {
			b.debug("wait after error, %v", timeoutAfterError)
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

		var updates []*models.Update

		errRequest := RawRequest(ctx, b, "getUpdates", params, &updates)
		if errRequest != nil {
			if errors.Is(errRequest, context.Canceled) {
				return
			}
			b.error("error get updates, %v", errRequest)
			timeoutAfterError = incErrTimeout(timeoutAfterError)
			continue
		}

		timeoutAfterError = 0

		for _, upd := range updates {
			atomic.StoreInt64(&b.lastUpdateID, upd.ID)
			select {
			case updatesChan <- upd:
			default:
				b.error("error send update to processing, channel is full")
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
