package bot

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/go-telegram/bot/models"
)

type clientMock struct {
	requestURI string
	gotReq     *http.Request
}

func (c *clientMock) Do(req *http.Request) (*http.Response, error) {
	c.requestURI = req.URL.RequestURI()
	c.gotReq = req
	resp := http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(strings.NewReader(`{"ok":true}`)),
	}
	return &resp, nil
}

func Test_rawRequest_url(t *testing.T) {
	cm := &clientMock{}
	b := &Bot{
		token:  "XXX",
		client: cm,
	}

	err := b.rawRequest(context.Background(), "foo", nil, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if cm.requestURI != "/botXXX/foo" {
		t.Fatalf("unexpected requestURI: %s", cm.requestURI)
	}
}

func Test_rawRequest_url_testEnv(t *testing.T) {
	cm := &clientMock{}
	b := &Bot{
		token:           "XXX",
		client:          cm,
		testEnvironment: true,
	}

	err := b.rawRequest(context.Background(), "foo", nil, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if cm.requestURI != "/botXXX/test/foo" {
		t.Fatalf("unexpected requestURI: %s", cm.requestURI)
	}
}

// The outgoing request must be replayable: a non-nil GetBody and a known
// ContentLength, with GetBody returning the same bytes as the body. That is what
// http2.Transport relies on to retry a POST after the server sends GOAWAY.
func Test_rawRequest_setsGetBody(t *testing.T) {
	cm := &clientMock{}
	b := &Bot{token: "XXX", client: cm}

	err := b.rawRequest(context.Background(), "sendMessage", &SendMessageParams{
		ChatID: 123,
		Text:   "hello",
	}, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	req := cm.gotReq
	if req == nil {
		t.Fatal("clientMock didn't observe a request")
	}
	if req.GetBody == nil {
		t.Fatal("Request.GetBody is nil — http2.Transport cannot retry on GOAWAY")
	}
	if req.ContentLength <= 0 {
		t.Fatalf("Request.ContentLength is %d, expected > 0", req.ContentLength)
	}

	original, err := io.ReadAll(req.Body)
	if err != nil {
		t.Fatalf("read original body: %v", err)
	}
	if int64(len(original)) != req.ContentLength {
		t.Fatalf("body length %d != ContentLength %d", len(original), req.ContentLength)
	}

	// Two GetBody calls must each return the full body — this is what
	// http2.Transport does on retry, and it can be invoked more than
	// once if a connection is dropped multiple times in quick succession.
	for i := 0; i < 2; i++ {
		rc, err := req.GetBody()
		if err != nil {
			t.Fatalf("GetBody call %d returned error: %v", i, err)
		}
		replay, err := io.ReadAll(rc)
		if err != nil {
			t.Fatalf("read replayed body call %d: %v", i, err)
		}
		if err := rc.Close(); err != nil {
			t.Fatalf("close replayed body call %d: %v", i, err)
		}
		if !bytes.Equal(original, replay) {
			t.Fatalf("replayed body call %d differs from original (lens %d vs %d)",
				i, len(replay), len(original))
		}
	}

	// The body should also contain the expected multipart fields,
	// proving GetBody isn't just returning empty bytes that happen
	// to match an empty original.
	if !bytes.Contains(original, []byte(`name="chat_id"`)) {
		t.Errorf("body missing chat_id field")
	}
	if !bytes.Contains(original, []byte(`name="text"`)) {
		t.Errorf("body missing text field")
	}
	if !bytes.Contains(original, []byte("hello")) {
		t.Errorf("body missing text value")
	}
}

// A method without fields (getMe, or a params struct whose fields are all omitted)
// must send neither a body nor a multipart Content-Type. Local telegram-bot-api
// servers reject an empty multipart body with a bare 400 (#285).
func Test_rawRequest_noBodyWithoutFields(t *testing.T) {
	type emptyParams struct {
		A string `json:"a,omitempty"`
	}

	cases := []struct {
		name   string
		params any
	}{
		{"nil params", nil},
		{"typed nil params", (*emptyParams)(nil)},
		{"all fields omitted", &emptyParams{}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			cm := &clientMock{}
			b := &Bot{token: "XXX", client: cm}

			if requestErr := b.rawRequest(context.Background(), "getMe", c.params, nil); requestErr != nil {
				t.Fatalf("unexpected error: %v", requestErr)
			}

			req := cm.gotReq
			if got := req.Header.Get("Content-Type"); got != "" {
				t.Fatalf("expected no Content-Type, got %q", got)
			}
			if req.ContentLength != 0 {
				t.Fatalf("expected ContentLength 0, got %d", req.ContentLength)
			}
			if req.Body != http.NoBody {
				t.Fatalf("expected http.NoBody, got %T", req.Body)
			}
		})
	}
}

// A request whose only field is a custom-marshaled or InputMedia value is not empty:
// the field is written to the form and must be counted, or the body would be dropped.
func Test_rawRequest_customMarshalAndInputMediaFieldsCount(t *testing.T) {
	cases := []struct {
		name   string
		params any
		part   string
	}{
		{
			"custom marshal field",
			&struct {
				Scope *models.BotCommandScopeDefault `json:"scope,omitempty"`
			}{Scope: &models.BotCommandScopeDefault{}},
			`name="scope"`,
		},
		{
			"custom marshal behind interface",
			&struct {
				Scope models.BotCommandScope `json:"scope,omitempty"`
			}{Scope: &models.BotCommandScopeDefault{}},
			`name="scope"`,
		},
		{
			"input media field",
			&struct {
				Media *models.InputMediaPhoto `json:"media,omitempty"`
			}{Media: &models.InputMediaPhoto{Media: "file_id"}},
			`name="media"`,
		},
		{
			"input media behind interface",
			&struct {
				Media models.InputMedia `json:"media,omitempty"`
			}{Media: &models.InputMediaPhoto{Media: "file_id"}},
			`name="media"`,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			cm := &clientMock{}
			b := &Bot{token: "XXX", client: cm}

			if requestErr := b.rawRequest(context.Background(), "editMessageMedia", c.params, nil); requestErr != nil {
				t.Fatalf("unexpected error: %v", requestErr)
			}

			req := cm.gotReq
			if got := req.Header.Get("Content-Type"); !strings.HasPrefix(got, "multipart/form-data") {
				t.Fatalf("expected a multipart Content-Type, got %q", got)
			}
			body, readErr := io.ReadAll(req.Body)
			if readErr != nil {
				t.Fatal(readErr)
			}
			if !bytes.Contains(body, []byte(c.part)) {
				t.Fatalf("body must contain %s, got:\n%s", c.part, body)
			}
		})
	}
}
