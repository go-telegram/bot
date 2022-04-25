package methods

import (
	"context"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func TestBot_GetMe(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"ok":true,"result":{"id":123456789,"is_bot":true,"first_name":"firstName","username":"username","language_code":"en-US"}}`))
	}))

	b := bot.New("", bot.WithHTTPClient(time.Second, s.Client()), bot.WithServerURL(s.URL))

	user, err := GetMe(context.Background(), b)
	if err != nil {
		t.Errorf("Expected no error, got: %s", err.Error())
	}

	expect := &models.User{
		ID:           123456789,
		IsBot:        true,
		FirstName:    "firstName",
		Username:     "username",
		LanguageCode: "en-US",
	}

	if !reflect.DeepEqual(user, expect) {
		t.Errorf("Expected user: %+v, got: %+v", expect, user)
	}
}
