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
	"github.com/stretchr/testify/assert"
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

func TestWebhookHandler_Success(t *testing.T) {
	bot := &Bot{
		updates: make(chan *models.Update, 1),
		isDebug: true,
		debugHandler: func(format string, args ...interface{}) {
			// Mock debug handler
		},
		errorsHandler: func(err error) {
			// Mock error handler
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
		assert.Equal(t, update.ID, upd.ID)
	default:
		t.Fatal("Expected update to be sent to bot.updates channel")
	}
}

func TestWebhookHandler_ReadBodyError(t *testing.T) {
	var capturedError error
	bot := &Bot{
		errorsHandler: func(err error) {
			capturedError = err
		},
	}

	req := httptest.NewRequest(http.MethodPost, "/", errReader(errors.New("read error")))
	w := httptest.NewRecorder()

	handler := bot.WebhookHandler()
	handler(w, req)

	assert.Error(t, capturedError)
	assert.Contains(t, capturedError.Error(), "read error")
}

func TestWebhookHandler_DecodeError(t *testing.T) {
	var capturedError error
	bot := &Bot{
		errorsHandler: func(err error) {
			capturedError = err
		},
	}

	invalidJSON := []byte("{invalid json}")
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(invalidJSON))
	w := httptest.NewRecorder()

	handler := bot.WebhookHandler()
	handler(w, req)

	assert.Error(t, capturedError)
	assert.Contains(t, capturedError.Error(), "error decode request body")
}

func TestWebhookHandler_ContextDone(t *testing.T) {
	var capturedError error
	bot := &Bot{
		updates: make(chan *models.Update, 1),
		errorsHandler: func(err error) {
			capturedError = err
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
	}

	assert.Error(t, capturedError)
	assert.Contains(t, capturedError.Error(), "some updates lost, ctx done")
}

func errReader(err error) io.ReadCloser {
	return &errReaderStruct{err: err}
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
