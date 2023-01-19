package bot

import (
	"math/rand"
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
