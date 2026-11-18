package bot

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/go-telegram/bot/models"
)

// WebhookHandler returns an HTTP handler for updates processed by StartWebhook.
// It returns 503 when StartWebhook is not accepting updates, so Telegram
// retries them later. An unreadable request body returns 400.
func (b *Bot) WebhookHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if b.webhookSecretToken != "" && req.Header.Get("X-Telegram-Bot-Api-Secret-Token") != b.webhookSecretToken {
			b.error("invalid webhook secret token received from update")
			return
		}

		body, errReadBody := io.ReadAll(req.Body)
		if errReadBody != nil {
			b.error("error read request body, %w", errReadBody)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		update := &models.Update{}

		errDecode := json.Unmarshal(body, update)
		if errDecode != nil {
			b.error("error decode request body, %s, %w", body, errDecode)
			return
		}

		if b.isDebug {
			b.debugHandler("webhook request '%s'", body)
		}

		if !b.queueWebhookUpdate(req.Context(), update) {
			w.WriteHeader(http.StatusServiceUnavailable)
		}
	}
}

// queueWebhookUpdate reports whether the update was accepted. It checks
// webhookCtx and sends under updatesMx, which StartWebhook also holds to stop
// admission, and a request woken while waiting for space checks again. So an
// update is never accepted after StartWebhook's ctx is canceled.
func (b *Bot) queueWebhookUpdate(ctx context.Context, update *models.Update) bool {
	var waitedRun *updatesRun
	for {
		b.updatesMx.Lock()
		run := b.updatesRun
		if waitedRun != nil && run != waitedRun {
			// The run this request waited on has ended; that is a shutdown, not a misconfiguration.
			b.updatesMx.Unlock()
			return false
		}
		if run == nil {
			b.updatesMx.Unlock()
			b.error("webhook update rejected: StartWebhook is not running")
			return false
		}
		if run.webhookCtx == nil || run.webhookCtx.Err() != nil || ctx.Err() != nil {
			b.updatesMx.Unlock()
			return false
		}
		if run.webhookQueued < cap(run.queue) {
			// The reservation guarantees buffer space, so the send does not block
			// and StartWebhook can close the queue once it stops admission.
			run.webhookQueued++
			run.queue <- queuedUpdate{update: update}
			b.updatesMx.Unlock()
			return true
		}
		if run.queueSpace == nil {
			run.queueSpace = make(chan struct{})
		}
		space := run.queueSpace
		stopping := run.webhookCtx.Done()
		waitedRun = run
		b.updatesMx.Unlock()

		select {
		case <-ctx.Done():
			return false
		case <-stopping:
			return false
		case <-space:
		}
	}
}
