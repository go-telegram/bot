package bot

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/go-telegram/bot/models"
)

func (b *Bot) WebhookHandler() http.HandlerFunc {
	return func(_ http.ResponseWriter, req *http.Request) {
		if b.webhookSecretToken != "" && req.Header.Get("X-Telegram-Bot-Api-Secret-Token") != b.webhookSecretToken {
			b.error("invalid webhook secret token received from update")
			return
		}

		body, errReadBody := io.ReadAll(req.Body)
		if errReadBody != nil {
			b.error("error read request body, %w", errReadBody)
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

		select {
		case <-req.Context().Done():
			b.error("some updates lost, ctx done")
			return
		default:
		}

		select {
		case b.updates <- update:
		case <-req.Context().Done():
			b.error("failed to send update, ctx done")
		}
	}
}
