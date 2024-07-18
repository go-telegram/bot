package bot

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-telegram/bot/models"
)

type mockDebugHandler struct {
	messages []string
}

func (h *mockDebugHandler) Handle(format string, args ...interface{}) {
	h.messages = append(h.messages, format)
}

type mockErrorsHandler struct {
	errors []error
}

func (h *mockErrorsHandler) Handle(err error) {
	h.errors = append(h.errors, err)
}

type errReaderStruct struct {
	err error
}

func (e *errReaderStruct) Read(p []byte) (int, error) {
	return 0, e.err
}

func (e *errReaderStruct) Close() error {
	return nil
}

func errReader(err error) io.ReadCloser {
	return &errReaderStruct{err: err}
}

func TestWebhookHandler_Success(t *testing.T) {
	debugHandler := &mockDebugHandler{}
	errorsHandler := &mockErrorsHandler{}

	bot := &Bot{
		updates: make(chan *models.Update, 1),
		isDebug: true,
		debugHandler: func(format string, args ...interface{}) {
			debugHandler.Handle(format, args...)
		},
		errorsHandler: func(err error) {
			errorsHandler.Handle(err)
		},
	}

	update := &models.Update{
		ID: 12345,
	}
	updateBody, _ := json.Marshal(update)

	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(updateBody))
	w := httptest.NewRecorder()

	handler := bot.WebhookHandler()
	handler(w, req)

	select {
	case upd := <-bot.updates:
		if upd.ID != update.ID {
			t.Fatalf("Expected update ID %d, got %d", update.ID, upd.ID)
		}
	default:
		t.Fatal("Expected update to be sent to bot.updates channel")
	}
}

func TestWebhookHandler_ReadBodyError(t *testing.T) {
	debugHandler := &mockDebugHandler{}
	errorsHandler := &mockErrorsHandler{}

	bot := &Bot{
		debugHandler: func(format string, args ...interface{}) {
			debugHandler.Handle(format, args...)
		},
		errorsHandler: func(err error) {
			errorsHandler.Handle(err)
		},
	}

	req := httptest.NewRequest(http.MethodPost, "/", errReader(errors.New("read error")))
	w := httptest.NewRecorder()

	handler := bot.WebhookHandler()
	handler(w, req)

	if len(errorsHandler.errors) == 0 {
		t.Fatal("Expected an error, but none occurred")
	}

	if capturedError := errorsHandler.errors[0]; capturedError == nil || !containsString(capturedError.Error(), "read error") {
		t.Fatalf("Expected read body error, got %v", capturedError)
	}
}

func TestWebhookHandler_DecodeError(t *testing.T) {
	debugHandler := &mockDebugHandler{}
	errorsHandler := &mockErrorsHandler{}

	bot := &Bot{
		debugHandler: func(format string, args ...interface{}) {
			debugHandler.Handle(format, args...)
		},
		errorsHandler: func(err error) {
			errorsHandler.Handle(err)
		},
	}

	invalidJSON := []byte("{invalid json}")
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(invalidJSON))
	w := httptest.NewRecorder()

	handler := bot.WebhookHandler()
	handler(w, req)

	if len(errorsHandler.errors) == 0 {
		t.Fatal("Expected an error, but none occurred")
	}

	if capturedError := errorsHandler.errors[0]; capturedError == nil || !containsString(capturedError.Error(), "error decode request body") {
		t.Fatalf("Expected decode error, got %v", capturedError)
	}
}

func TestWebhookHandler_ContextDone(t *testing.T) {
	debugHandler := &mockDebugHandler{}
	errorsHandler := &mockErrorsHandler{}

	bot := &Bot{
		updates: make(chan *models.Update, 1),
		debugHandler: func(format string, args ...interface{}) {
			debugHandler.Handle(format, args...)
		},
		errorsHandler: func(err error) {
			errorsHandler.Handle(err)
		},
	}

	update := &models.Update{
		ID: 12345,
	}
	updateBody, _ := json.Marshal(update)

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(updateBody)).WithContext(ctx)
	w := httptest.NewRecorder()

	handler := bot.WebhookHandler()
	handler(w, req)

	select {
	case <-bot.updates:
		t.Fatal("Did not expect update to be sent to bot.updates channel")
	default:
		// Expected outcome, context was cancelled before sending the update
	}

	if len(errorsHandler.errors) == 0 {
		t.Fatal("Expected an error, but none occurred")
	}

	if capturedError := errorsHandler.errors[0]; capturedError == nil || !containsString(capturedError.Error(), "some updates lost, ctx done") {
		t.Fatalf("Expected context done error, got %v", capturedError)
	}
}

func containsString(s, substr string) bool {
	return bytes.Contains([]byte(s), []byte(substr))
}
