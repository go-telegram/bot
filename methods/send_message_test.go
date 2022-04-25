package methods

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func TestSendMessageParams_Validate(t *testing.T) {
	p := SendMessageParams{}
	err := p.Validate()
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	if !errors.Is(err, bot.ErrEmptyChatID) {
		t.Errorf("Expected error %v, got %v", bot.ErrEmptyChatID, err)
	}

	p = SendMessageParams{ChatID: "foo"}
	err = p.Validate()
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	if !errors.Is(err, bot.ErrEmptyText) {
		t.Errorf("Expected error %v, got %v", bot.ErrEmptyText, err)
	}

	p = SendMessageParams{ChatID: "foo", Text: "bar"}
	err = p.Validate()
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
}

func TestSendMessage(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"ok":true,"result":{"message_id":12}}`))
	}))

	b := bot.New("", bot.WithHTTPClient(time.Second, s.Client()), bot.WithServerURL(s.URL))

	res, err := SendMessage(context.Background(), b, &SendMessageParams{
		ChatID: "foo",
		Text:   "bar",
	})
	if err != nil {
		t.Errorf("Expected no error, got: %s", err.Error())
	}

	expect := &models.Message{
		ID: 12,
	}

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Expected: %v, got: %v", expect, res)
	}
}
