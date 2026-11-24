package bot

import (
	"crypto/ed25519"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/go-telegram/bot/models"
)

const (
	WebAppPublicKeyHex     = "e7bf03a2fa4602af4580703d88dda5bb59f32ed8b02a56c187fe7d34caed242d"
	WebAppTestPublicKeyHex = "40055058a4ee38156a06562e52eece92a771bcd8346a8c4615cb7376eddf72ec"
)

var (
	errEmptyWebAppBotToken    = errors.New("empty bot token")
	errInvalidWebAppPublicKey = errors.New("invalid ed25519 public key size")
	errInvalidWebAppBotID     = errors.New("invalid bot id")
)

// WebAppChat https://core.telegram.org/bots/webapps#webappchat
type WebAppChat struct {
	ID       int64           `json:"id"`
	Type     models.ChatType `json:"type"`
	Title    string          `json:"title"`
	Username string          `json:"username,omitempty"`
	PhotoURL string          `json:"photo_url,omitempty"`
}

// WebAppInitData https://core.telegram.org/bots/webapps#webappinitdata
//
// User, Receiver and Chat are nil when the field is absent or null.
type WebAppInitData struct {
	QueryID                string      `json:"query_id,omitempty"`
	ChatJoinRequestQueryID string      `json:"chat_join_request_query_id,omitempty"`
	User                   *WebAppUser `json:"user,omitempty"`
	Receiver               *WebAppUser `json:"receiver,omitempty"`
	Chat                   *WebAppChat `json:"chat,omitempty"`
	ChatType               string      `json:"chat_type,omitempty"`
	ChatInstance           string      `json:"chat_instance,omitempty"`
	StartParam             string      `json:"start_param,omitempty"`
	CanSendAfter           int         `json:"can_send_after,omitempty"` // seconds
	AuthDate               int         `json:"auth_date"`
	Hash                   string      `json:"hash"`
	Signature              string      `json:"signature"`
}

// ValidateWebAppInitData parses Mini App init data and checks its hash with the bot token.
//
// Pass raw Telegram.WebApp.initData without a leading "?" or "#tgWebAppData=".
// Unknown fields are included in the check but not returned.
// Decoded field names containing '=' or LF, or values containing LF, are malformed.
//
// An empty token returns a configuration error before parsing initData.
// On failure it returns nil data. Use errors.Is for ErrorWebAppInitDataMalformed,
// ErrorWebAppInitDataUnsigned and ErrorWebAppInitDataInvalid.
func ValidateWebAppInitData(initData, token string) (*WebAppInitData, error) {
	if token == "" {
		return nil, errEmptyWebAppBotToken
	}

	fields, err := parseWebAppFields(initData)
	if err != nil {
		return nil, err
	}

	hash := fields["hash"]
	if hash == "" {
		return nil, ErrorWebAppInitDataUnsigned
	}

	check := make(map[string]string, len(fields))
	for k, v := range fields {
		if k != "hash" {
			check[k] = v
		}
	}
	expected := webAppHMAC(token, webAppDataCheckString(check))
	if !hmac.Equal([]byte(expected), []byte(hash)) {
		return nil, ErrorWebAppInitDataInvalid
	}

	return newWebAppInitData(fields)
}

// ValidateWebAppInitDataSignature verifies Mini App init data with Telegram's
// Ed25519 public key. Decode WebAppPublicKeyHex or WebAppTestPublicKeyHex with hex.DecodeString.
//
// Signatures use canonical base64url, with or without padding.
// The returned Hash is unauthenticated.
//
// An invalid public key length or nonpositive bot ID returns a configuration error
// before parsing. See ValidateWebAppInitData for the input format and data errors.
// On failure it returns nil data.
func ValidateWebAppInitDataSignature(initData string, botID int64, publicKey ed25519.PublicKey) (*WebAppInitData, error) {
	if len(publicKey) != ed25519.PublicKeySize {
		return nil, errInvalidWebAppPublicKey
	}
	if botID <= 0 {
		return nil, errInvalidWebAppBotID
	}

	fields, err := parseWebAppFields(initData)
	if err != nil {
		return nil, err
	}

	signature := fields["signature"]
	if signature == "" {
		return nil, ErrorWebAppInitDataUnsigned
	}
	sig, err := decodeWebAppSignature(signature)
	if err != nil {
		return nil, ErrorWebAppInitDataInvalid
	}

	check := make(map[string]string, len(fields))
	for k, v := range fields {
		if k != "hash" && k != "signature" {
			check[k] = v
		}
	}
	dcs := strconv.FormatInt(botID, 10) + ":WebAppData\n" + webAppDataCheckString(check)
	if !ed25519.Verify(publicKey, []byte(dcs), sig) {
		return nil, ErrorWebAppInitDataInvalid
	}

	return newWebAppInitData(fields)
}

func decodeWebAppSignature(signature string) ([]byte, error) {
	if strings.ContainsAny(signature, "\r\n") {
		return nil, ErrorWebAppInitDataInvalid
	}
	if strings.HasSuffix(signature, "=") {
		return base64.URLEncoding.Strict().DecodeString(signature)
	}
	return base64.RawURLEncoding.Strict().DecodeString(signature)
}

// CheckAuthDate returns ErrorWebAppInitDataExpired if the data is older than maxAge
// or if maxAge is not positive. With a positive maxAge, future timestamps pass.
func (d *WebAppInitData) CheckAuthDate(maxAge time.Duration) error {
	return d.checkAuthDate(time.Now(), maxAge)
}

func (d *WebAppInitData) checkAuthDate(now time.Time, maxAge time.Duration) error {
	if maxAge <= 0 {
		return ErrorWebAppInitDataExpired
	}
	cutoff := now.Add(-maxAge)
	cutoffSeconds := cutoff.Unix()
	authDate := int64(d.AuthDate)
	if authDate < cutoffSeconds || authDate == cutoffSeconds && cutoff.Nanosecond() != 0 {
		return ErrorWebAppInitDataExpired
	}
	return nil
}

func parseWebAppFields(initData string) (map[string]string, error) {
	values, err := url.ParseQuery(initData)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrorWebAppInitDataMalformed, err)
	}

	fields := make(map[string]string, len(values))
	for k, v := range values {
		if len(v) != 1 {
			return nil, fmt.Errorf("%w: field %q: repeated", ErrorWebAppInitDataMalformed, k)
		}
		if !isUnambiguousWebAppField(k, v[0]) {
			return nil, fmt.Errorf("%w: field %q: invalid delimiter", ErrorWebAppInitDataMalformed, k)
		}
		fields[k] = v[0]
	}
	return fields, nil
}

// The data-check-string format has no escaping for field separators.
func isUnambiguousWebAppField(key, value string) bool {
	return !strings.ContainsAny(key, "=\n") && !strings.ContainsRune(value, '\n')
}

func newWebAppInitData(fields map[string]string) (*WebAppInitData, error) {
	d := &WebAppInitData{
		QueryID:                fields["query_id"],
		ChatJoinRequestQueryID: fields["chat_join_request_query_id"],
		ChatType:               fields["chat_type"],
		ChatInstance:           fields["chat_instance"],
		StartParam:             fields["start_param"],
		Hash:                   fields["hash"],
		Signature:              fields["signature"],
	}

	authDate, ok := fields["auth_date"]
	if !ok {
		return nil, fmt.Errorf("%w: field auth_date: missing", ErrorWebAppInitDataMalformed)
	}
	var err error
	if d.AuthDate, err = strconv.Atoi(authDate); err != nil {
		return nil, fmt.Errorf("%w: field auth_date: %v", ErrorWebAppInitDataMalformed, err)
	}

	if v, ok := fields["can_send_after"]; ok {
		if d.CanSendAfter, err = strconv.Atoi(v); err != nil {
			return nil, fmt.Errorf("%w: field can_send_after: %v", ErrorWebAppInitDataMalformed, err)
		}
	}

	objects := []struct {
		key string
		dst any
	}{
		{"user", &d.User},
		{"receiver", &d.Receiver},
		{"chat", &d.Chat},
	}
	for _, o := range objects {
		v, ok := fields[o.key]
		if !ok {
			continue
		}
		if err := json.Unmarshal([]byte(v), o.dst); err != nil {
			return nil, fmt.Errorf("%w: field %s: %v", ErrorWebAppInitDataMalformed, o.key, err)
		}
	}

	return d, nil
}

func webAppDataCheckString(fields map[string]string) string {
	keys := make([]string, 0, len(fields))
	for key := range fields {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	lines := make([]string, len(keys))
	for i, key := range keys {
		lines[i] = key + "=" + fields[key]
	}
	return strings.Join(lines, "\n")
}

func webAppHMAC(token, dataCheckString string) string {
	secret := hmac.New(sha256.New, []byte("WebAppData"))
	secret.Write([]byte(token))

	mac := hmac.New(sha256.New, secret.Sum(nil))
	mac.Write([]byte(dataCheckString))
	return hex.EncodeToString(mac.Sum(nil))
}
