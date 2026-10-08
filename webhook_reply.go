package bot

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"reflect"
	"sync"

	"github.com/go-telegram/bot/models"
)

type webhookReplyKey struct{}

// webhookReply holds the method call that answers an update in the response to
// the webhook request that delivered it.
type webhookReply struct {
	mx          sync.Mutex
	closed      bool
	body        []byte
	contentType string
}

// WebhookReplyHandler returns an HTTP handler for webhook updates that lets
// handlers answer an update in the webhook response with WebhookReply.
//
// WebhookReplyHandler runs the handler for each update synchronously, in the
// goroutine serving the request, and responds once the handler returns: with
// the reply if the handler made one, or with an empty 200 OK otherwise. A
// request that fails the webhook secret token check is answered with 401
// Unauthorized, and one with an unreadable body with 400 Bad Request.
func (b *Bot) WebhookReplyHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if b.webhookSecretToken != "" && req.Header.Get("X-Telegram-Bot-Api-Secret-Token") != b.webhookSecretToken {
			b.error("invalid webhook secret token received from update")
			w.WriteHeader(http.StatusUnauthorized)
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
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if b.isDebug {
			b.debugHandler("webhook request '%s'", body)
		}

		reply := &webhookReply{}
		ctx := context.WithValue(req.Context(), webhookReplyKey{}, reply)

		h := applyMiddlewares(b.findHandler(update), b.middlewares...)
		h(ctx, b, update)

		reply.mx.Lock()
		defer reply.mx.Unlock()

		reply.closed = true

		if reply.body == nil {
			w.WriteHeader(http.StatusOK)
			return
		}

		w.Header().Set("Content-Type", reply.contentType)
		w.WriteHeader(http.StatusOK)
		if _, errWrite := w.Write(reply.body); errWrite != nil {
			b.error("error write webhook reply, %w", errWrite)
		}
	}
}

// WebhookReply answers the update being handled by sending a Bot API method
// call in the response to the webhook request, Telegram does not report whether
// the response succeeded.
//
// The method is the Bot API method name, such as "answerInlineQuery", and the
// params are its parameters, such as *AnswerInlineQueryParams.
//
// WebhookReply works only in a handler run by WebhookReplyHandler. It must be
// called before the handler returns. It can only be called once per update.
// Otherwise, WebhookReply returns an error wrapping
// ErrorWebhookReplyUnavailable.
//
// See https://core.telegram.org/bots/api#making-requests-when-getting-updates
func (b *Bot) WebhookReply(ctx context.Context, method string, params any) error {
	reply, ok := ctx.Value(webhookReplyKey{}).(*webhookReply)
	if !ok {
		return fmt.Errorf("%w, update not handled by WebhookReplyHandler", ErrorWebhookReplyUnavailable)
	}

	var bodyBuf bytes.Buffer
	form := multipart.NewWriter(&bodyBuf)

	if errMethod := form.WriteField("method", method); errMethod != nil {
		return fmt.Errorf("error build webhook reply for method %s, %w", method, errMethod)
	}

	if params != nil && !reflect.ValueOf(params).IsNil() {
		if _, errFormData := buildRequestForm(form, params); errFormData != nil {
			return fmt.Errorf("error build webhook reply for method %s, %w", method, errFormData)
		}
	}

	if errFormClose := form.Close(); errFormClose != nil {
		return fmt.Errorf("error form close for method %s, %w", method, errFormClose)
	}

	reply.mx.Lock()
	defer reply.mx.Unlock()

	if reply.closed {
		return fmt.Errorf("%w, webhook request already answered", ErrorWebhookReplyUnavailable)
	}

	if reply.body != nil {
		return fmt.Errorf("%w, update already replied to", ErrorWebhookReplyUnavailable)
	}

	if b.isDebug {
		replyDebugData, _ := json.Marshal(params)
		b.debugHandler("webhook reply method: %s, payload: %s", method, replyDebugData)
	}

	reply.body = bodyBuf.Bytes()
	reply.contentType = form.FormDataContentType()

	return nil
}
