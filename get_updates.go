package bot

import (
	"context"
	"encoding/json"
	"errors"
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
func (b *Bot) getUpdates(ctx context.Context, queue chan<- queuedUpdate) {
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
		batch := make([]batchUpdate, 0, len(updates))

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
			if params.Offset < 0 && len(batch) == 0 {
				// Telegram has dropped the updates before this batch. Resume from
				// it rather than from the negative offset, which would drop any
				// update of the batch that was not dispatched.
				atomic.StoreInt64(&b.lastUpdateID, upd.ID-1)
			}

			if errDecode != nil {
				b.error("error decode update %d, skipped, %s, %w", upd.ID, raw, errDecode)
				batch = append(batch, batchUpdate{id: upd.ID})
				continue
			}

			claim := &pollClaim{dispatched: make(chan struct{})}
			select {
			case <-ctx.Done():
				b.confirmDispatched(ctx, batch)
				return
			case queue <- queuedUpdate{update: upd, claim: claim}:
				batch = append(batch, batchUpdate{id: upd.ID, claim: claim})
			}
		}

		if !b.confirmDispatched(ctx, batch) {
			return
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

// batchUpdate is an update of a polled batch. claim is nil for an update that
// was skipped and needs no dispatch.
type batchUpdate struct {
	id    int64
	claim *pollClaim
}

const (
	claimPending int32 = iota
	claimTaken
	claimAbandoned
)

// pollClaim settles whether a polled update reaches user code. A worker takes
// it before running any, and the poller abandons it on cancellation unless a
// worker took it first, so both sides agree on what was dispatched.
type pollClaim struct {
	state      int32
	dispatched chan struct{}
}

func (c *pollClaim) take() bool {
	if !atomic.CompareAndSwapInt32(&c.state, claimPending, claimTaken) {
		return false
	}
	close(c.dispatched)
	return true
}

func (c *pollClaim) abandon() bool {
	return atomic.CompareAndSwapInt32(&c.state, claimPending, claimAbandoned)
}

// confirmDispatched confirms the batch once all of it is dispatched and
// reports whether it was. If ctx is canceled first, it confirms only the
// updates before the first one not yet dispatched, so polling resumes there.
func (b *Bot) confirmDispatched(ctx context.Context, batch []batchUpdate) bool {
	settled := 0
	for settled < len(batch) && batch[settled].wait(ctx) {
		settled++
	}
	if settled > 0 {
		atomic.StoreInt64(&b.lastUpdateID, batch[settled-1].id)
	}
	return settled == len(batch)
}

// wait reports whether the update was dispatched. On cancellation it abandons
// the update unless a worker has already taken it.
func (u batchUpdate) wait(ctx context.Context) bool {
	if u.claim == nil {
		return true
	}
	select {
	case <-u.claim.dispatched:
		return true
	case <-ctx.Done():
		return !u.claim.abandon()
	}
}
