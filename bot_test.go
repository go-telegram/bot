package bot

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/go-telegram/bot/models"
)

type serverMock struct {
	s          *httptest.Server
	custom     map[string]any
	hooks      map[string]func(body []byte) any
	hooksCalls map[string]int
	updateIdx  int
	updates    []*models.Update
	token      string
}

type getUpdatesResponse struct {
	OK     bool             `json:"ok"`
	Result []*models.Update `json:"result"`
}

func (s *serverMock) handler(rw http.ResponseWriter, req *http.Request) {
	if req.URL.String() == "/bot"+s.token+"/getMe" {
		_, err := rw.Write([]byte(`{"ok":true,"result":{}}`))
		if err != nil {
			panic(err)
		}
		return
	}
	if req.URL.String() == "/bot"+s.token+"/getUpdates" {
		s.handlerGetUpdates(rw)
		return
	}

	reqBody, errReadBody := io.ReadAll(req.Body)
	if errReadBody != nil {
		panic(errReadBody)
	}
	defer func() {
		if err := req.Body.Close(); err != nil {
			panic(err)
		}
	}()

	hook, okHook := s.hooks[req.URL.String()]
	if okHook {
		s.hooksCalls[req.URL.String()]++
		resp, errData := json.Marshal(hook(reqBody))
		if errData != nil {
			panic(errData)
		}
		_, err := rw.Write(resp)
		if err != nil {
			panic(err)
		}
		return
	}

	d, ok := s.custom[req.URL.String()]
	if !ok {
		panic("answer not found for request: " + req.URL.String())
	}

	resp, errData := json.Marshal(d)
	if errData != nil {
		panic(errData)
	}
	_, err := rw.Write(resp)
	if err != nil {
		panic(err)
	}
}

func (s *serverMock) Close() {
	s.s.Close()
}

func (s *serverMock) handlerGetUpdates(rw http.ResponseWriter) {
	if s.updateIdx >= len(s.updates) {
		_, err := rw.Write([]byte(`{"ok":true,"result":[]}`))
		if err != nil {
			panic(err)
		}
		return
	}

	s.updates[s.updateIdx].ID = int64(s.updateIdx + 1)

	r := getUpdatesResponse{
		OK:     true,
		Result: []*models.Update{s.updates[s.updateIdx]},
	}

	s.updateIdx++

	d, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}
	_, err = rw.Write(d)
	if err != nil {
		panic(err)
	}
}

func (s *serverMock) URL() string {
	return s.s.URL
}

func newServerMock(token string) *serverMock {
	s := &serverMock{
		token:      token,
		custom:     map[string]any{},
		hooks:      map[string]func([]byte) any{},
		hooksCalls: map[string]int{},
	}

	s.s = httptest.NewServer(http.HandlerFunc(s.handler))

	return s
}

func TestNew_error_getMe(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		_, err := rw.Write([]byte(`{"ok":false,"description":"err1","error_code":400}`))
		if err != nil {
			panic(err)
		}
	}))
	defer server.Close()

	_, err := New("xxx", WithServerURL(server.URL))
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	if err.Error() != "error call getMe, bad request, err1" {
		t.Fatalf("wrong error message %q", err.Error())
	}
}

func TestNew_emptyToken(t *testing.T) {
	_, err := New("")
	if err == nil {
		t.Fatalf("expect error %q", err)
	}
	if err.Error() != "empty token" {
		t.Fatalf("exexpected error %q", err)
	}
}

func TestNew(t *testing.T) {
	s := newServerMock("xxx")
	defer s.Close()

	_, err := New("xxx", WithServerURL(s.URL()))
	if err != nil {
		t.Fatalf("unexpected error %q", err)
	}
}

func TestBot_StartWebhookWithCorrectSecret(t *testing.T) {
	s := newServerMock("xxx")
	defer s.Close()

	opts := []Option{WithServerURL(s.URL()), WithWebhookSecretToken("zzzz")}
	b, err := New("xxx", opts...)
	if err != nil {
		t.Fatalf("unexpected error %q", err)
	}

	mx := sync.Mutex{}
	var called bool

	b.defaultHandlerFunc = func(ctx context.Context, bot *Bot, update *models.Update) {
		if update.Message.ID != 1 {
			t.Errorf("unexpected message id")
		}
		mx.Lock()
		called = true
		mx.Unlock()
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go b.StartWebhook(ctx)

	time.Sleep(time.Millisecond * 100)

	req, errReq := http.NewRequest(http.MethodPost, "", strings.NewReader(`{"update_id":1,"message":{"message_id":1}}`))
	if errReq != nil {
		t.Error(errReq)
		return
	}
	req.Header.Set("X-Telegram-Bot-Api-Secret-Token", "zzzz")

	b.WebhookHandler().ServeHTTP(nil, req)

	cancel()

	time.Sleep(time.Millisecond * 100)

	mx.Lock()
	if !called {
		t.Errorf("not called default handler")
	}
	mx.Unlock()
}

func TestBot_StartWebhookWithNoSecret(t *testing.T) {
	s := newServerMock("xxx")
	defer s.Close()

	opts := []Option{WithServerURL(s.URL())}
	b, err := New("xxx", opts...)
	if err != nil {
		t.Fatalf("unexpected error %q", err)
	}

	mx := sync.Mutex{}
	var called bool

	b.defaultHandlerFunc = func(ctx context.Context, bot *Bot, update *models.Update) {
		if update.Message.ID != 1 {
			t.Errorf("unexpected message id")
		}
		mx.Lock()
		called = true
		mx.Unlock()
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go b.StartWebhook(ctx)

	time.Sleep(time.Millisecond * 100)

	req, errReq := http.NewRequest(http.MethodPost, "", strings.NewReader(`{"update_id":1,"message":{"message_id":1}}`))
	if errReq != nil {
		t.Error(errReq)
		return
	}

	b.WebhookHandler().ServeHTTP(nil, req)

	cancel()

	time.Sleep(time.Millisecond * 100)

	mx.Lock()
	if !called {
		t.Errorf("not called default handler")
	}
	mx.Unlock()
}

func TestBot_StartWebhookWithWrongSecret(t *testing.T) {
	s := newServerMock("xxx")
	defer s.Close()

	opts := []Option{WithServerURL(s.URL()), WithWebhookSecretToken("zzzz")}
	b, err := New("xxx", opts...)
	if err != nil {
		t.Fatalf("unexpected error %q", err)
	}

	mx := sync.Mutex{}
	var called bool

	b.defaultHandlerFunc = func(ctx context.Context, bot *Bot, update *models.Update) {
		if update.Message.ID != 1 {
			t.Errorf("unexpected message id")
		}
		mx.Lock()
		called = true
		mx.Unlock()
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go b.StartWebhook(ctx)

	time.Sleep(time.Millisecond * 100)

	req, errReq := http.NewRequest(http.MethodPost, "", strings.NewReader(`{"update_id":1,"message":{"message_id":1}}`))
	if errReq != nil {
		t.Error(errReq)
		return
	}
	req.Header.Set("X-Telegram-Bot-Api-Secret-Token", "wrong_secret")

	b.WebhookHandler().ServeHTTP(nil, req)

	cancel()

	time.Sleep(time.Millisecond * 100)

	mx.Lock()
	if called {
		t.Errorf("not supposed to call the default handler with wrong token")
	}
	mx.Unlock()
}

func TestBot_Start(t *testing.T) {
	s := newServerMock("xxx")
	defer s.Close()

	s.updates = append(s.updates, &models.Update{Message: &models.Message{ID: 42}})

	b, err := New("xxx", WithServerURL(s.URL()))
	if err != nil {
		t.Fatalf("unexpected error %q", err)
	}

	var called bool

	b.defaultHandlerFunc = func(ctx context.Context, bot *Bot, update *models.Update) {
		if update.Message.ID != 42 {
			t.Errorf("unexpected message id %d", update.Message.ID)
		}
		called = true
	}

	ctx, cancel := context.WithCancel(context.Background())

	go b.Start(ctx)

	time.Sleep(time.Millisecond * 100)

	cancel()

	time.Sleep(time.Millisecond * 100)

	if !called {
		t.Errorf("not called default handler")
	}
}
