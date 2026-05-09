package bot

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"strings"
	"testing"
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

// Test_rawRequest_setsGetBody verifies that the outgoing *http.Request
// is replayable: it has a non-nil GetBody and a known ContentLength,
// and GetBody returns bytes identical to the request body. This is
// what http2.Transport relies on when it has to retry a POST after the
// server sends a GOAWAY frame mid-flight — without these set, the
// retry returns "cannot retry err [http2: ... GOAWAY] after
// Request.Body was written; define Request.GetBody to avoid this
// error" and every Telegram POST issued on a draining HTTP/2
// connection fails (long-poll getUpdates is GET-with-no-body, so it
// keeps working — the bot keeps receiving updates while silently
// failing every reply).
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
