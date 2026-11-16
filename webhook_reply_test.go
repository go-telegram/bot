package bot

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"mime"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-telegram/bot/models"
)

func newWebhookReplyTestBot(handler HandlerFunc, errorsHandler *mockErrorsHandler) *Bot {
	return &Bot{
		defaultHandlerFunc: handler,
		debugHandler:       func(format string, args ...any) {},
		errorsHandler: func(err error) {
			errorsHandler.Handle(err)
		},
	}
}

func serveWebhookReply(t *testing.T, b *Bot, update *models.Update) *httptest.ResponseRecorder {
	t.Helper()

	updateBody, err := json.Marshal(update)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(updateBody))
	w := httptest.NewRecorder()

	b.WebhookReplyHandler()(w, req)

	return w
}

func readWebhookReplyForm(t *testing.T, w *httptest.ResponseRecorder) map[string][]string {
	t.Helper()

	mediaType, mediaParams, err := mime.ParseMediaType(w.Header().Get("Content-Type"))
	if err != nil {
		t.Fatalf("unexpected error parsing content type: %v", err)
	}
	if mediaType != "multipart/form-data" {
		t.Fatalf("expected multipart/form-data, got %s", mediaType)
	}

	form, err := multipart.NewReader(w.Body, mediaParams["boundary"]).ReadForm(1 << 20)
	if err != nil {
		t.Fatalf("unexpected error reading form: %v", err)
	}

	return form.Value
}

func TestWebhookReplyHandler_Reply(t *testing.T) {
	errorsHandler := &mockErrorsHandler{}

	b := newWebhookReplyTestBot(func(ctx context.Context, b *Bot, update *models.Update) {
		err := b.WebhookReply(ctx, "answerInlineQuery", &AnswerInlineQueryParams{
			InlineQueryID: update.InlineQuery.ID,
			Results: []models.InlineQueryResult{
				&models.InlineQueryResultArticle{
					ID:                  "1",
					Title:               "Foo",
					InputMessageContent: &models.InputTextMessageContent{MessageText: "foo"},
				},
			},
		})
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	}, errorsHandler)

	w := serveWebhookReply(t, b, &models.Update{
		ID:          1,
		InlineQuery: &models.InlineQuery{ID: "42", Query: "foo"},
	})

	if w.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, w.Code)
	}

	form := readWebhookReplyForm(t, w)

	if got := form["method"]; len(got) != 1 || got[0] != "answerInlineQuery" {
		t.Fatalf("expected method answerInlineQuery, got %v", got)
	}
	if got := form["inline_query_id"]; len(got) != 1 || got[0] != "42" {
		t.Fatalf("expected inline_query_id 42, got %v", got)
	}

	var results []map[string]any
	if err := json.Unmarshal([]byte(form["results"][0]), &results); err != nil {
		t.Fatalf("unexpected error decoding results: %v", err)
	}
	if len(results) != 1 || results[0]["type"] != "article" || results[0]["id"] != "1" {
		t.Fatalf("unexpected results: %v", results)
	}

	if len(errorsHandler.errors) != 0 {
		t.Fatalf("unexpected errors: %v", errorsHandler.errors)
	}
}

func TestWebhookReplyHandler_NoReply(t *testing.T) {
	errorsHandler := &mockErrorsHandler{}

	var called bool
	b := newWebhookReplyTestBot(func(ctx context.Context, b *Bot, update *models.Update) {
		called = true
	}, errorsHandler)

	w := serveWebhookReply(t, b, &models.Update{ID: 1})

	if !called {
		t.Fatal("expected handler to be called")
	}
	if w.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, w.Code)
	}
	if w.Body.Len() != 0 {
		t.Fatalf("expected empty body, got %q", w.Body.String())
	}
	if w.Header().Get("Content-Type") != "" {
		t.Fatalf("expected no content type, got %q", w.Header().Get("Content-Type"))
	}
}

func TestWebhookReplyHandler_ReplyWithoutParams(t *testing.T) {
	errorsHandler := &mockErrorsHandler{}

	b := newWebhookReplyTestBot(func(ctx context.Context, b *Bot, update *models.Update) {
		if err := b.WebhookReply(ctx, "logOut", nil); err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	}, errorsHandler)

	w := serveWebhookReply(t, b, &models.Update{ID: 1})

	form := readWebhookReplyForm(t, w)

	if len(form) != 1 || form["method"][0] != "logOut" {
		t.Fatalf("expected only method logOut, got %v", form)
	}
}

func TestWebhookReplyHandler_SecondReply(t *testing.T) {
	errorsHandler := &mockErrorsHandler{}

	var errSecond error
	b := newWebhookReplyTestBot(func(ctx context.Context, b *Bot, update *models.Update) {
		if err := b.WebhookReply(ctx, "sendMessage", &SendMessageParams{ChatID: 1, Text: "first"}); err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		errSecond = b.WebhookReply(ctx, "sendMessage", &SendMessageParams{ChatID: 1, Text: "second"})
	}, errorsHandler)

	w := serveWebhookReply(t, b, &models.Update{ID: 1})

	if !errors.Is(errSecond, ErrorWebhookReplyUnavailable) {
		t.Fatalf("expected ErrorWebhookReplyUnavailable, got %v", errSecond)
	}

	form := readWebhookReplyForm(t, w)

	if got := form["text"]; len(got) != 1 || got[0] != "first" {
		t.Fatalf("expected the first reply, got %v", got)
	}
}

func TestWebhookReplyHandler_ReplyAfterReturn(t *testing.T) {
	errorsHandler := &mockErrorsHandler{}

	var replyCtx context.Context
	b := newWebhookReplyTestBot(func(ctx context.Context, b *Bot, update *models.Update) {
		replyCtx = ctx
	}, errorsHandler)

	w := serveWebhookReply(t, b, &models.Update{ID: 1})

	err := b.WebhookReply(replyCtx, "sendMessage", &SendMessageParams{ChatID: 1, Text: "late"})
	if !errors.Is(err, ErrorWebhookReplyUnavailable) {
		t.Fatalf("expected ErrorWebhookReplyUnavailable, got %v", err)
	}
	if w.Body.Len() != 0 {
		t.Fatalf("expected empty body, got %q", w.Body.String())
	}
}

type countingReadError struct {
	reads int
	err   error
}

func (r *countingReadError) Read(_ []byte) (int, error) {
	r.reads++
	return 0, r.err
}

func TestWebhookReply_OutsideWebhookReplyHandler(t *testing.T) {
	b := &Bot{}

	err := b.WebhookReply(context.Background(), "sendMessage", &SendMessageParams{ChatID: 1, Text: "foo"})
	if !errors.Is(err, ErrorWebhookReplyUnavailable) {
		t.Fatalf("expected ErrorWebhookReplyUnavailable, got %v", err)
	}
}

func TestWebhookReply_UnsupportedParamsType(t *testing.T) {
	n := 1

	tests := []struct {
		name   string
		params any
	}{
		{name: "struct value", params: SendMessageParams{ChatID: 1, Text: "foo"}},
		{name: "pointer to non-struct", params: &n},
		{name: "int", params: 1},
		{name: "string", params: "foo"},
		{name: "map", params: map[string]any{"chat_id": 1}},
		{name: "slice", params: []string{"foo"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reply := &webhookReply{}
			ctx := context.WithValue(context.Background(), webhookReplyKey{}, reply)
			b := &Bot{}

			err := b.WebhookReply(ctx, "sendMessage", tt.params)
			if err == nil {
				t.Fatal("expected an error for unsupported params type")
			}
			if errors.Is(err, ErrorWebhookReplyUnavailable) {
				t.Fatalf("expected a params type error, got %v", err)
			}
			if reply.body != nil {
				t.Fatalf("expected no reply to be recorded, got %q", reply.body)
			}
		})
	}
}

func TestWebhookReplyHandler_ReplyWithTypedNilParams(t *testing.T) {
	errorsHandler := &mockErrorsHandler{}

	b := newWebhookReplyTestBot(func(ctx context.Context, b *Bot, update *models.Update) {
		var params *SendMessageParams
		if err := b.WebhookReply(ctx, "logOut", params); err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	}, errorsHandler)

	w := serveWebhookReply(t, b, &models.Update{ID: 1})

	form := readWebhookReplyForm(t, w)

	if len(form) != 1 || form["method"][0] != "logOut" {
		t.Fatalf("expected only method logOut, got %v", form)
	}
}

func TestWebhookReplyHandler_RetryAfterFailedReply(t *testing.T) {
	errorsHandler := &mockErrorsHandler{}

	errRead := errors.New("read failed")

	var errFirst, errRetry error
	b := newWebhookReplyTestBot(func(ctx context.Context, b *Bot, update *models.Update) {
		errFirst = b.WebhookReply(ctx, "sendDocument", &SendDocumentParams{
			ChatID:   1,
			Document: &models.InputFileUpload{Filename: "broken.txt", Data: &countingReadError{err: errRead}},
		})
		errRetry = b.WebhookReply(ctx, "sendMessage", &SendMessageParams{ChatID: 1, Text: "retry"})
	}, errorsHandler)

	w := serveWebhookReply(t, b, &models.Update{ID: 1})

	if !errors.Is(errFirst, errRead) {
		t.Fatalf("expected the read error, got %v", errFirst)
	}
	if errors.Is(errFirst, ErrorWebhookReplyUnavailable) {
		t.Fatalf("expected failed reply to leave the reply available, got %v", errFirst)
	}
	if errRetry != nil {
		t.Fatalf("unexpected error for retry: %v", errRetry)
	}

	form := readWebhookReplyForm(t, w)

	if got := form["method"]; len(got) != 1 || got[0] != "sendMessage" {
		t.Fatalf("expected method sendMessage, got %v", got)
	}
	if got := form["text"]; len(got) != 1 || got[0] != "retry" {
		t.Fatalf("expected text retry, got %v", got)
	}
}

func TestWebhookReplyHandler_SecondReplyDoesNotReadRejectedParams(t *testing.T) {
	errorsHandler := &mockErrorsHandler{}

	var errSecond error
	reader := &countingReadError{err: errors.New("second read failed")}
	b := newWebhookReplyTestBot(func(ctx context.Context, b *Bot, update *models.Update) {
		if err := b.WebhookReply(ctx, "sendDocument", &SendDocumentParams{
			ChatID:   1,
			Document: &models.InputFileUpload{Filename: "first.txt", Data: bytes.NewReader([]byte("first"))},
		}); err != nil {
			t.Fatalf("unexpected error for first reply: %v", err)
		}
		errSecond = b.WebhookReply(ctx, "sendDocument", &SendDocumentParams{
			ChatID:   1,
			Document: &models.InputFileUpload{Filename: "second.txt", Data: reader},
		})
	}, errorsHandler)

	serveWebhookReply(t, b, &models.Update{ID: 1})

	if !errors.Is(errSecond, ErrorWebhookReplyUnavailable) {
		t.Fatalf("expected ErrorWebhookReplyUnavailable, got %v", errSecond)
	}
	if reader.reads != 0 {
		t.Fatalf("expected rejected second reply to avoid reading upload data, got %d reads", reader.reads)
	}
}

func TestWebhookReplyHandler_ReplyAfterReturnDoesNotReadRejectedParams(t *testing.T) {
	errorsHandler := &mockErrorsHandler{}

	var replyCtx context.Context
	b := newWebhookReplyTestBot(func(ctx context.Context, b *Bot, update *models.Update) {
		replyCtx = ctx
	}, errorsHandler)

	w := serveWebhookReply(t, b, &models.Update{ID: 1})

	reader := &countingReadError{err: errors.New("late read failed")}
	err := b.WebhookReply(replyCtx, "sendDocument", &SendDocumentParams{
		ChatID:   1,
		Document: &models.InputFileUpload{Filename: "late.txt", Data: reader},
	})
	if !errors.Is(err, ErrorWebhookReplyUnavailable) {
		t.Fatalf("expected ErrorWebhookReplyUnavailable, got %v", err)
	}
	if reader.reads != 0 {
		t.Fatalf("expected rejected late reply to avoid reading upload data, got %d reads", reader.reads)
	}
	if w.Body.Len() != 0 {
		t.Fatalf("expected empty body, got %q", w.Body.String())
	}
}

func TestWebhookReplyHandler_Middlewares(t *testing.T) {
	errorsHandler := &mockErrorsHandler{}

	var calls []string
	b := newWebhookReplyTestBot(func(ctx context.Context, b *Bot, update *models.Update) {
		calls = append(calls, "handler")
	}, errorsHandler)
	b.middlewares = []Middleware{
		func(next HandlerFunc) HandlerFunc {
			return func(ctx context.Context, b *Bot, update *models.Update) {
				calls = append(calls, "middleware")
				next(ctx, b, update)
			}
		},
	}

	serveWebhookReply(t, b, &models.Update{ID: 1})

	if len(calls) != 2 || calls[0] != "middleware" || calls[1] != "handler" {
		t.Fatalf("expected middleware then handler, got %v", calls)
	}
}

func TestWebhookReplyHandler_RegisteredHandler(t *testing.T) {
	errorsHandler := &mockErrorsHandler{}

	b := newWebhookReplyTestBot(func(ctx context.Context, b *Bot, update *models.Update) {
		t.Error("expected the registered handler, got the default handler")
	}, errorsHandler)

	var called bool
	b.RegisterHandlerInlineQuery(func(ctx context.Context, b *Bot, update *models.Update) {
		called = true
	})

	serveWebhookReply(t, b, &models.Update{
		ID:          1,
		InlineQuery: &models.InlineQuery{ID: "42"},
	})

	if !called {
		t.Fatal("expected the registered handler to be called")
	}
}

func TestWebhookReplyHandler_InvalidSecretToken(t *testing.T) {
	errorsHandler := &mockErrorsHandler{}

	b := newWebhookReplyTestBot(func(ctx context.Context, b *Bot, update *models.Update) {
		t.Error("expected handler not to be called")
	}, errorsHandler)
	b.webhookSecretToken = "secret"

	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBufferString(`{"update_id":1}`))
	req.Header.Set("X-Telegram-Bot-Api-Secret-Token", "wrong")
	w := httptest.NewRecorder()

	b.WebhookReplyHandler()(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Fatalf("expected status %d, got %d", http.StatusUnauthorized, w.Code)
	}
	if len(errorsHandler.errors) == 0 {
		t.Fatal("expected an error, but none occurred")
	}
}

func TestWebhookReplyHandler_ValidSecretToken(t *testing.T) {
	errorsHandler := &mockErrorsHandler{}

	var called bool
	b := newWebhookReplyTestBot(func(ctx context.Context, b *Bot, update *models.Update) {
		called = true
	}, errorsHandler)
	b.webhookSecretToken = "secret"

	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBufferString(`{"update_id":1}`))
	req.Header.Set("X-Telegram-Bot-Api-Secret-Token", "secret")
	w := httptest.NewRecorder()

	b.WebhookReplyHandler()(w, req)

	if !called {
		t.Fatal("expected handler to be called")
	}
	if w.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, w.Code)
	}
}

func TestWebhookReplyHandler_ReadBodyError(t *testing.T) {
	errorsHandler := &mockErrorsHandler{}

	b := newWebhookReplyTestBot(func(ctx context.Context, b *Bot, update *models.Update) {
		t.Error("expected handler not to be called")
	}, errorsHandler)

	req := httptest.NewRequest(http.MethodPost, "/", errReader(errors.New("read error")))
	w := httptest.NewRecorder()

	b.WebhookReplyHandler()(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, w.Code)
	}
	if len(errorsHandler.errors) == 0 || !containsString(errorsHandler.errors[0].Error(), "read error") {
		t.Fatalf("expected read body error, got %v", errorsHandler.errors)
	}
}

func TestWebhookReplyHandler_DecodeError(t *testing.T) {
	errorsHandler := &mockErrorsHandler{}

	b := newWebhookReplyTestBot(func(ctx context.Context, b *Bot, update *models.Update) {
		t.Error("expected handler not to be called")
	}, errorsHandler)

	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBufferString("{invalid json}"))
	w := httptest.NewRecorder()

	b.WebhookReplyHandler()(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, w.Code)
	}
	if len(errorsHandler.errors) == 0 || !containsString(errorsHandler.errors[0].Error(), "error decode request body") {
		t.Fatalf("expected decode error, got %v", errorsHandler.errors)
	}
}
