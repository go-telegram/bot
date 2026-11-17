package bot

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"net/url"
	"sort"
	"strings"
	"testing"
)

func TestEscapeMarkdown(t *testing.T) {
	res := EscapeMarkdown("foo \\! _*[]()~`>#+-=|{}.! bar")
	if res != "foo \\\\! \\_\\*\\[\\]\\(\\)\\~\\`\\>\\#\\+\\-\\=\\|\\{\\}\\.\\! bar" {
		t.Fatalf("unexpected result %q", res)
	}
}

func TestEscapeMarkdownUnescaped(t *testing.T) {
	res := EscapeMarkdownUnescaped("foo \\! _*[]()~`>#+-=|{}.! bar")
	if res != "foo \\! \\_\\*\\[\\]\\(\\)\\~\\`\\>\\#\\+\\-\\=\\|\\{\\}\\.\\! bar" {
		t.Fatalf("unexpected result %q", res)
	}
}

func webAppHash(token string, pairs map[string]string) string {
	keys := make([]string, 0, len(pairs))
	for k := range pairs {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	lines := make([]string, len(keys))
	for i, k := range keys {
		lines[i] = k + "=" + pairs[k]
	}

	mac := hmac.New(sha256.New, []byte("WebAppData"))
	mac.Write([]byte(token))
	secret := mac.Sum(nil)

	mac = hmac.New(sha256.New, secret)
	mac.Write([]byte(strings.Join(lines, "\n")))
	return fmt.Sprintf("%x", mac.Sum(nil))
}

func signWebAppQuery(token string, pairs map[string]string) string {
	values := url.Values{}
	for k, v := range pairs {
		values.Set(k, v)
	}
	values.Set("hash", webAppHash(token, pairs))
	return values.Encode()
}

func TestValidateWebappRequest_Issue311Plus(t *testing.T) {
	const token = "123456:token"
	raw := "auth_date=1770000000" +
		"&query_id=AAEAAAE" +
		"&user=%7B%22id%22%3A42%2C%22first_name%22%3A%22A%2BB%22%7D" +
		"&hash=9ad374c06d5e48711811310bf935f5caf9ca5e80cc44c10aaa15a2ad1b1b9cb8"

	values, err := url.ParseQuery(raw)
	if err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(values.Get("user"), `"first_name":"A+B"`) {
		t.Fatalf("fixture was decoded before validation: %q", values.Get("user"))
	}

	user, ok := ValidateWebappRequest(values, token)
	if !ok {
		t.Fatal("expected a valid signature")
	}
	if user.ID != 42 || user.FirstName != "A+B" {
		t.Fatalf("user %+v", user)
	}
	if values.Get("hash") == "" {
		t.Fatal("hash was removed from the caller values")
	}
}

func TestValidateWebappRequest_SpacePercent20(t *testing.T) {
	const token = "123456:token"
	const userJSON = `{"id":8,"first_name":"A B"}`

	hash := webAppHash(token, map[string]string{
		"auth_date": "1770000000",
		"user":      userJSON,
	})
	raw := "auth_date=1770000000&user=" + strings.ReplaceAll(url.QueryEscape(userJSON), "+", "%20") + "&hash=" + hash
	if !strings.Contains(raw, "%20") || strings.Contains(raw, "+") {
		t.Fatalf("raw query must encode the space as %%20: %s", raw)
	}

	values, err := url.ParseQuery(raw)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(values.Get("user"), `"first_name":"A B"`) {
		t.Fatalf("ParseQuery user = %q", values.Get("user"))
	}

	user, ok := ValidateWebappRequest(values, token)
	if !ok {
		t.Fatal("expected a valid signature")
	}
	if user.ID != 8 || user.FirstName != "A B" {
		t.Fatalf("user %+v", user)
	}
}

func TestValidateWebappRequest_SecondCall(t *testing.T) {
	const token = "123456:token"
	values, err := url.ParseQuery(signWebAppQuery(token, map[string]string{
		"auth_date": "1770000000",
		"user":      `{"id":1,"first_name":"Ann"}`,
	}))
	if err != nil {
		t.Fatal(err)
	}
	hash := values.Get("hash")

	for i := 0; i < 2; i++ {
		user, ok := ValidateWebappRequest(values, token)
		if !ok || user.ID != 1 || user.FirstName != "Ann" {
			t.Fatalf("call %d: ok=%v user=%+v", i, ok, user)
		}
		if values.Get("hash") != hash {
			t.Fatalf("call %d removed hash", i)
		}
	}
}

func TestValidateWebappRequest(t *testing.T) {
	const token = "123456:token"

	tests := []struct {
		name     string
		pairs    map[string]string
		mutate   func(url.Values)
		wantOK   bool
		wantName string
		wantID   int64
	}{
		{
			name: "plain name",
			pairs: map[string]string{
				"auth_date": "1770000000",
				"user":      `{"id":1,"first_name":"Ann"}`,
			},
			wantOK:   true,
			wantName: "Ann",
			wantID:   1,
		},
		{
			name: "percent",
			pairs: map[string]string{
				"auth_date": "1770000000",
				"user":      `{"id":7,"first_name":"100%"}`,
			},
			wantOK:   true,
			wantName: "100%",
			wantID:   7,
		},
		{
			name: "space encoded as plus",
			pairs: map[string]string{
				"auth_date": "1770000000",
				"user":      `{"id":8,"first_name":"A B"}`,
			},
			wantOK:   true,
			wantName: "A B",
			wantID:   8,
		},
		{
			name: "bad hash",
			pairs: map[string]string{
				"auth_date": "1770000000",
				"user":      `{"id":1,"first_name":"Ann"}`,
			},
			mutate: func(v url.Values) {
				v.Set("hash", "0000000000000000000000000000000000000000000000000000000000000000")
			},
			wantOK: false,
		},
		{
			name: "user is not json",
			pairs: map[string]string{
				"auth_date": "1770000000",
				"user":      "not-json",
			},
			wantOK: false,
		},
		{
			name: "no user",
			pairs: map[string]string{
				"auth_date": "1770000000",
				"query_id":  "AAEAAAE",
			},
			wantOK: true,
		},
		{
			name: "empty value slice",
			pairs: map[string]string{
				"auth_date": "1770000000",
				"user":      `{"id":1,"first_name":"Ann"}`,
			},
			mutate: func(v url.Values) {
				v["auth_date"] = []string{}
			},
			wantOK: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			values, err := url.ParseQuery(signWebAppQuery(token, tt.pairs))
			if err != nil {
				t.Fatal(err)
			}
			if tt.mutate != nil {
				tt.mutate(values)
			}

			user, ok := ValidateWebappRequest(values, token)
			if ok != tt.wantOK {
				t.Fatalf("ok=%v, want %v", ok, tt.wantOK)
			}
			if !tt.wantOK {
				if user != nil {
					t.Fatalf("user %+v", user)
				}
				return
			}
			if user.ID != tt.wantID || user.FirstName != tt.wantName {
				t.Fatalf("user %+v", user)
			}
		})
	}
}

func TestRandomString(t *testing.T) {
	res := map[string]struct{}{}
	for i := 0; i < 100; i++ {
		r := RandomString(10)
		if _, ok := res[r]; ok {
			t.Fatalf("value already exists")
		}
		res[r] = struct{}{}
	}

	for i := 10; i < 50; i++ {
		r := RandomString(i)
		if len(r) != i {
			t.Fatalf("unexpected len")
		}
	}
}
