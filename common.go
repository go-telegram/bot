package bot

import (
	"crypto/hmac"
	"encoding/json"
	"math/rand"
	"net/url"
	"strings"
	"sync"
	"time"
)

// escape special symbols in text for MarkdownV2 parse mode

var shouldBeEscaped = "_*[]()~`>#+-=|{}.!"

// EscapeMarkdown escapes plain text for Telegram MarkdownV2 syntax.
func EscapeMarkdown(s string) string {
	var result []rune
	for _, r := range s {
		if r == '\\' || strings.ContainsRune(shouldBeEscaped, r) {
			result = append(result, '\\')
		}
		result = append(result, r)
	}
	return string(result)
}

// EscapeMarkdownUnescaped escapes unescaped special symbols for Telegram Markdown v2 syntax
func EscapeMarkdownUnescaped(s string) string {
	var result []rune
	var escaped bool
	for _, r := range s {
		if r == '\\' {
			escaped = !escaped
			result = append(result, r)
			continue
		}
		if strings.ContainsRune(shouldBeEscaped, r) && !escaped {
			result = append(result, '\\')
		}
		escaped = false
		result = append(result, r)
	}
	return string(result)
}

// log functions

// random string generator

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var randSrc = rand.NewSource(time.Now().UnixNano())
var randSrcMx sync.Mutex

// RandomString returns random a-zA-Z string with n length
func RandomString(n int) string {
	b := make([]byte, n)
	// A randSrc.Int63() generates 63 random bits, enough for letterIdxMax characters!
	randSrcMx.Lock()
	ch := randSrc.Int63()
	randSrcMx.Unlock()
	for i, remain := n-1, letterIdxMax; i >= 0; {
		if remain == 0 {
			randSrcMx.Lock()
			ch, remain = randSrc.Int63(), letterIdxMax
			randSrcMx.Unlock()
		}
		if idx := int(ch & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		ch >>= letterIdxBits
		remain--
	}

	return string(b)
}

// WebAppUser represents user model from webapp request
type WebAppUser struct {
	ID                    int64  `json:"id"`
	IsBot                 bool   `json:"is_bot"`
	FirstName             string `json:"first_name"`
	LastName              string `json:"last_name"`
	Username              string `json:"username"`
	LanguageCode          string `json:"language_code"`
	IsPremium             bool   `json:"is_premium"`
	AddedToAttachmentMenu bool   `json:"added_to_attachment_menu"`
	AllowsWriteToPM       bool   `json:"allows_write_to_pm"`
	PhotoURL              string `json:"photo_url"`
}

// ValidateWebappRequest validates Mini App init data.
// values must already be decoded, as returned by url.ParseQuery or http.Request.URL.Query.
// It uses the first value per key and leaves values unchanged.
// An empty token, LF in signed names or values, or '=' in signed names returns nil, false.
func ValidateWebappRequest(values url.Values, token string) (user *WebAppUser, ok bool) {
	fields := prepareWebAppRequest(values, token)
	if fields == nil {
		return nil, false
	}

	expected := webAppHMAC(token, webAppDataCheckString(fields))
	if !hmac.Equal([]byte(expected), []byte(values.Get("hash"))) {
		return nil, false
	}

	var u WebAppUser
	if v, ok := fields["user"]; ok {
		if err := json.Unmarshal([]byte(v), &u); err != nil {
			return nil, false
		}
	}

	return &u, true
}

func prepareWebAppRequest(values url.Values, token string) map[string]string {
	if token == "" {
		return nil
	}

	fields := make(map[string]string, len(values))
	for k, v := range values {
		if k == "hash" {
			continue
		}
		if len(v) == 0 || !isUnambiguousWebAppField(k, v[0]) {
			return nil
		}
		fields[k] = v[0]
	}
	return fields
}
