package bot

import (
	"context"
	"encoding/json"
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

		var updates []json.RawMessage

		errRequest := b.rawRequest(ctx, "getUpdates", params, &updates)
		if errRequest != nil {
			if errors.Is(errRequest, context.Canceled) {
				return
			}
			b.error("error get updates, %w", errRequest)

			var tooManyRequestsErr *TooManyRequestsError
			if errors.As(errRequest, &tooManyRequestsErr) && tooManyRequestsErr.RetryAfter > 0 {
				timeoutAfterError = time.Duration(tooManyRequestsErr.RetryAfter) * time.Second
			} else {
				timeoutAfterError = incErrTimeout(timeoutAfterError)
			}

			continue
		}

		// offsetStuck is set when the last update of the batch left the offset where it
		// was, so the very same batch comes back on the next request.
		offsetStuck := false

		for _, raw := range updates {
			upd, errDecode := decodeUpdate(raw)
			if upd.ID == 0 {
				if errDecode == nil {
					errDecode = errMissingUpdateID
				}
				b.error("error decode update, %s, %w", raw, errDecode)
				offsetStuck = true
				continue
			}

			offsetStuck = false
			atomic.StoreInt64(&b.lastUpdateID, upd.ID)

			if errDecode != nil {
				b.error("error decode update %d, skipped, %s, %w", upd.ID, raw, errDecode)
				continue
			}

			select {
			case <-ctx.Done():
				b.error("some updates lost, ctx done")
				return
			case b.updates <- upd:
			}
		}

		if offsetStuck {
			// Back off as on a failed request instead of re-requesting the same
			// batch at full speed.
			timeoutAfterError = incErrTimeout(timeoutAfterError)
		} else {
			timeoutAfterError = 0
		}
	}
}

var errMissingUpdateID = errors.New("missing update_id")

// decodeUpdate decodes one update on its own, so a single undecodable update does not
// reject the whole batch. On failure the id is still read when possible, so the offset
// can move past it. A zero id means it could not be read.
func decodeUpdate(raw json.RawMessage) (*models.Update, error) {
	upd := &models.Update{}
	errDecode := json.Unmarshal(raw, upd)
	if errDecode == nil {
		return upd, nil
	}

	head := struct {
		ID int64 `json:"update_id"`
	}{}
	_ = json.Unmarshal(raw, &head)

	return &models.Update{ID: head.ID}, errDecode
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
