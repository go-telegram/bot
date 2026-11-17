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

// queueWebhookUpdate reports whether the update was accepted. Registering as a
// sender under updatesMx lets StartWebhook wait for in-flight sends before it
// closes the queue.
func (b *Bot) queueWebhookUpdate(ctx context.Context, update *models.Update) bool {
	b.updatesMx.Lock()
	run := b.updatesRun
	if run == nil {
		b.updatesMx.Unlock()
		b.error("webhook update rejected: StartWebhook is not running")
		return false
	}
	if run.webhookCtx == nil || run.webhookCtx.Err() != nil || ctx.Err() != nil {
		b.updatesMx.Unlock()
		return false
	}
	stopping := run.webhookCtx.Done()
	run.senders.Add(1)
	b.updatesMx.Unlock()
	defer run.senders.Done()

	select {
	case <-ctx.Done():
		return false
	case <-stopping:
		return false
	case run.queue <- queuedUpdate{update: update}:
		return true
	}
}
