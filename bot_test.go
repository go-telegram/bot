package bot

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
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

func TestNewRejectsNegativeUpdatesChannelCap(t *testing.T) {
	b, err := New("xxx", WithSkipGetMe(), WithUpdatesChannelCap(-1))
	if b != nil || err == nil {
		t.Fatalf("New returned (%v, %v), want a configuration error", b, err)
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
	handled := make(chan struct{}, 1)

	b.defaultHandlerFunc = func(ctx context.Context, bot *Bot, update *models.Update) {
		if update.Message.ID != 1 {
			t.Errorf("unexpected message id")
		}
		mx.Lock()
		called = true
		mx.Unlock()
		handled <- struct{}{}
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	stopped := startWebhookForTest(t, b, ctx)

	req, errReq := http.NewRequest(http.MethodPost, "", strings.NewReader(`{"update_id":1,"message":{"message_id":1}}`))
	if errReq != nil {
		t.Error(errReq)
		return
	}
	req.Header.Set("X-Telegram-Bot-Api-Secret-Token", "zzzz")

	b.WebhookHandler().ServeHTTP(httptest.NewRecorder(), req)

	waitForUpdateTest(t, handled, "default handler to run")
	cancel()

	waitForUpdateTest(t, stopped, "webhook to stop")

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
	handled := make(chan struct{}, 1)

	b.defaultHandlerFunc = func(ctx context.Context, bot *Bot, update *models.Update) {
		if update.Message.ID != 1 {
			t.Errorf("unexpected message id")
		}
		mx.Lock()
		called = true
		mx.Unlock()
		handled <- struct{}{}
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	stopped := startWebhookForTest(t, b, ctx)

	req, errReq := http.NewRequest(http.MethodPost, "", strings.NewReader(`{"update_id":1,"message":{"message_id":1}}`))
	if errReq != nil {
		t.Error(errReq)
		return
	}

	b.WebhookHandler().ServeHTTP(httptest.NewRecorder(), req)

	waitForUpdateTest(t, handled, "default handler to run")
	cancel()

	waitForUpdateTest(t, stopped, "webhook to stop")

	mx.Lock()
	if !called {
		t.Errorf("not called default handler")
	}
	mx.Unlock()
}

func TestBot_StartWebhookWithWrongSecret(t *testing.T) {
	s := newServerMock("xxx")
	defer s.Close()

	reported := make(chan error, 4)
	opts := []Option{WithServerURL(s.URL()), WithWebhookSecretToken("zzzz"),
		WithErrorsHandler(func(err error) { reported <- err })}
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

	stopped := startWebhookForTest(t, b, ctx)

	req, errReq := http.NewRequest(http.MethodPost, "", strings.NewReader(`{"update_id":1,"message":{"message_id":1}}`))
	if errReq != nil {
		t.Error(errReq)
		return
	}
	req.Header.Set("X-Telegram-Bot-Api-Secret-Token", "wrong_secret")

	b.WebhookHandler().ServeHTTP(httptest.NewRecorder(), req)

	if got := waitForUpdateTest(t, reported, "secret token rejection"); !strings.Contains(got.Error(), "invalid webhook secret token") {
		t.Errorf("reported %v, want the secret token rejection", got)
	}
	cancel()

	waitForUpdateTest(t, stopped, "webhook to stop")

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

func TestBot_ID(t *testing.T) {
	type fields struct {
		token string
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		{name: "empty token", fields: fields{token: ""}, want: 0},
		{name: "no colon", fields: fields{token: "xxx"}, want: 0},
		{name: "bad value", fields: fields{token: "123xxx:xxx"}, want: 0},
		{name: "bad value", fields: fields{token: ":xxx"}, want: 0},
		{name: "ok", fields: fields{token: "123456:xxx"}, want: 123456},
		{name: "two colon", fields: fields{token: "123456:5678:xxx"}, want: 123456},
		{name: "max int64", fields: fields{token: "9223372036854775807:xxx"}, want: 9223372036854775807},
		{name: "min int64", fields: fields{token: "-9223372036854775808:xxx"}, want: -9223372036854775808},
		{name: "positive overflow", fields: fields{token: "9223372036854775808:xxx"}, want: 0},
		{name: "negative overflow", fields: fields{token: "-9223372036854775809:xxx"}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Bot{
				token: tt.fields.token,
			}
			if got := b.ID(); got != tt.want {
				t.Errorf("ID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBot_Token(t *testing.T) {
	b := &Bot{token: "123456:xxx"}

	token := b.Token()

	if token != "123456:xxx" {
		t.Errorf("Token() = %s, want %s", token, "123456:xxx")
	}
}

func TestBot_SetToken(t *testing.T) {
	b := &Bot{}

	b.SetToken("123456:xxx")

	if b.token != "123456:xxx" {
		t.Errorf("SetToken() = %s, want %s", b.token, "123456:xxx")
	}
}

func getMeClient(t *testing.T, username string) HttpClient {
	return getUpdatesClientFunc(func(req *http.Request) (*http.Response, error) {
		if !strings.HasSuffix(req.URL.Path, "/getMe") {
			t.Errorf("unexpected request %s", req.URL.Path)
		}
		return getUpdatesJSONResponse(`{"ok":true,"result":{"id":1,"is_bot":true,"first_name":"bot","username":"` + username + `"}}`), nil
	})
}

func addressedCommandUpdate(text string) *models.Update {
	return &models.Update{Message: &models.Message{
		Text:     text,
		Entities: []models.MessageEntity{{Type: models.MessageEntityTypeBotCommand, Offset: 0, Length: len(text)}},
	}}
}

func TestBot_ProcessUpdateMatchesCommandAddressedToBot(t *testing.T) {
	var handled []string
	client := getUpdatesClientFunc(func(*http.Request) (*http.Response, error) {
		return getUpdatesJSONResponse(`{"ok":true,"result":{"id":1,"username":"my_bot","active_usernames":["alias_bot"]}}`), nil
	})
	b, err := New("1:token", WithSkipGetMe(), WithHTTPClient(time.Second, client), WithNotAsyncHandlers(),
		WithDefaultHandler(func(_ context.Context, _ *Bot, update *models.Update) {
			handled = append(handled, "default "+update.Message.Text)
		}))
	if err != nil {
		t.Fatal(err)
	}
	b.RegisterHandler(HandlerTypeMessageText, "start", MatchTypeCommandStartOnly, func(_ context.Context, _ *Bot, update *models.Update) {
		handled = append(handled, "start "+update.Message.Text)
	})

	me, err := b.GetMe(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	me.ActiveUsernames[0] = "other_bot" // The caller must not change routing by mutating the result.
	b.ProcessUpdate(context.Background(), addressedCommandUpdate("/start@my_bot"))
	b.ProcessUpdate(context.Background(), addressedCommandUpdate("/start@ALIAS_bot"))
	b.ProcessUpdate(context.Background(), addressedCommandUpdate("/start@other_bot"))
	b.SetToken("2:new_bot")
	b.ProcessUpdate(context.Background(), addressedCommandUpdate("/start@ALIAS_bot"))

	want := []string{"start /start@my_bot", "start /start@ALIAS_bot", "default /start@other_bot", "default /start@ALIAS_bot"}
	if !reflect.DeepEqual(handled, want) {
		t.Errorf("handled %q, want %q", handled, want)
	}
}

func TestBot_SetTokenForgetsUsernameOfAnotherBot(t *testing.T) {
	b, err := New("1:token", WithHTTPClient(time.Second, getMeClient(t, "my_bot")))
	if err != nil {
		t.Fatal(err)
	}
	b.SetToken("1:revoked-and-replaced")
	if got := b.botUsernames(); !reflect.DeepEqual(got, []string{"my_bot"}) {
		t.Errorf("username after a new token of the same bot = %q, want my_bot", got)
	}
	b.SetToken("2:token")
	if got := b.botUsernames(); len(got) != 0 {
		t.Errorf("username after switching to another bot = %q, want empty", got)
	}
}

func TestBot_GetMeIgnoresUsernameOfAnotherBot(t *testing.T) {
	var b *Bot
	// GetMe is sent with bot 2's token, and the token is switched back to bot 1
	// before the reply arrives, so the reply describes another bot.
	b, err := New("1:token", WithSkipGetMe(), WithHTTPClient(time.Second, getUpdatesClientFunc(func(req *http.Request) (*http.Response, error) {
		b.SetToken("1:token")
		return getUpdatesJSONResponse(`{"ok":true,"result":{"id":2,"is_bot":true,"first_name":"bot","username":"other_bot"}}`), nil
	})))
	if err != nil {
		t.Fatal(err)
	}
	b.SetToken("2:token")
	if _, err := b.GetMe(context.Background()); err != nil {
		t.Fatal(err)
	}
	if got := b.botUsernames(); len(got) != 0 {
		t.Errorf("username = %q, want empty for a GetMe answered by another bot", got)
	}
}

func TestBot_GetMeFailureKeepsUsername(t *testing.T) {
	fail := false
	b, err := New("1:token", WithHTTPClient(time.Second, getUpdatesClientFunc(func(req *http.Request) (*http.Response, error) {
		if fail {
			return nil, errors.New("network down")
		}
		return getUpdatesJSONResponse(`{"ok":true,"result":{"id":1,"is_bot":true,"first_name":"bot","username":"my_bot"}}`), nil
	})))
	if err != nil {
		t.Fatal(err)
	}
	fail = true
	if _, err := b.GetMe(context.Background()); err == nil {
		t.Fatal("GetMe succeeded, want an error")
	}
	if got := b.botUsernames(); !reflect.DeepEqual(got, []string{"my_bot"}) {
		t.Errorf("username after a failed GetMe = %q, want my_bot", got)
	}
}

func TestBot_GetMeDecodeErrorRecordsNoUsername(t *testing.T) {
	b, err := New("1:token", WithSkipGetMe(), WithHTTPClient(time.Second, getUpdatesClientFunc(func(*http.Request) (*http.Response, error) {
		// The id and username decode even though is_bot fails.
		return getUpdatesJSONResponse(`{"ok":true,"result":{"id":1,"username":"my_bot","is_bot":"yes"}}`), nil
	})))
	if err != nil {
		t.Fatal(err)
	}
	if _, err := b.GetMe(context.Background()); err == nil {
		t.Fatal("GetMe succeeded, want a decode error")
	}
	if got := b.botUsernames(); len(got) != 0 {
		t.Errorf("username = %q, want empty after a failed GetMe", got)
	}
}

// Run with -race: the username is read while it is set and while the token
// changes.
func TestBot_UsernameConcurrentAccess(t *testing.T) {
	b := &Bot{token: "1:token"}
	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(3)
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				b.setUsernames(&models.BotFullInfo{ID: 1, Username: "my_bot"})
			}
		}()
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				b.SetToken("2:token")
				b.SetToken("1:token")
			}
		}()
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				if got := b.botUsernames(); len(got) != 0 && !reflect.DeepEqual(got, []string{"my_bot"}) {
					t.Errorf("username = %q", got)
				}
			}
		}()
	}
	wg.Wait()
}
