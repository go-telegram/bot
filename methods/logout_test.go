package methods

import (
	"context"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/go-telegram/bot"
)

func TestBot_Logout(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"ok":true,"result":true}`))
	}))

	b := bot.New("", bot.WithHTTPClient(time.Second, s.Client()), bot.WithServerURL(s.URL))

	res, err := Logout(context.Background(), b)
	if err != nil {
		t.Errorf("Expected no error, got: %s", err.Error())
	}

	if !reflect.DeepEqual(res, true) {
		t.Errorf("Expected: %v, got: %v", true, res)
	}
}
