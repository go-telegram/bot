package bot

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/url"
	"sort"
	"strings"
	"sync"
	"time"
)

// escape special symbols in text for MarkdownV2 parse mode

var shouldBeEscaped = "_*[]()~`>#+-=|{}.!"

// EscapeMarkdown escapes special symbols for Telegram MarkdownV2 syntax
func EscapeMarkdown(s string) string {
	var result []rune
	for _, r := range s {
		if strings.ContainsRune(shouldBeEscaped, r) {
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

// ValidateWebappRequest validates request from webapp
func ValidateWebappRequest(values url.Values, token string) (user *WebAppUser, ok bool) {
	h := values.Get("hash")
	values.Del("hash")

	var vals []string

	var u WebAppUser

	for k, v := range values {
		vv, _ := url.QueryUnescape(v[0])
		vals = append(vals, k+"="+vv)
		if k == "user" {
			errDecodeUser := json.Unmarshal([]byte(vv), &u)
			if errDecodeUser != nil {
				return nil, false
			}
		}
	}

	sort.Slice(vals, func(i, j int) bool {
		return vals[i] < vals[j]
	})

	hmac1 := hmac.New(sha256.New, []byte("WebAppData"))
	hmac1.Write([]byte(token))
	r1 := hmac1.Sum(nil)

	data := []byte(strings.Join(vals, "\n"))

	hmac2 := hmac.New(sha256.New, r1)
	hmac2.Write(data)
	r2 := hmac2.Sum(nil)

	if h != fmt.Sprintf("%x", r2) {
		return nil, false
	}

	return &u, true
}
