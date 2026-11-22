package bot

import (
	"crypto/ed25519"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"net/url"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/go-telegram/bot/models"
)

const testWebAppToken = "123456:token"
const fullWebAppUserID = 10000000001

var webAppSentinels = []error{
	ErrorWebAppInitDataMalformed,
	ErrorWebAppInitDataUnsigned,
	ErrorWebAppInitDataInvalid,
	ErrorWebAppInitDataExpired,
}

func fullWebAppPairs() map[string]string {
	return map[string]string{
		"query_id":                   "AAHdF6IQAAAAAN0XohDhrOrc",
		"chat_join_request_query_id": "JRQ-1",
		"user":                       `{"id":10000000001,"first_name":"Alice","last_name":"Smith","username":"alice_test","language_code":"ru","is_premium":true,"added_to_attachment_menu":true,"allows_write_to_pm":true,"photo_url":"https://t.me/i/userpic/320/a.svg"}`,
		"receiver":                   `{"id":42,"is_bot":true,"first_name":"Bot"}`,
		"chat":                       `{"id":-1001234567890,"type":"supergroup","title":"Group","username":"grp","photo_url":"https://t.me/i/chat.jpg"}`,
		"chat_type":                  "supergroup",
		"chat_instance":              "-8318574285046036213",
		"start_param":                "ref_42",
		"can_send_after":             "10",
		"auth_date":                  "1770000000",
		"signature":                  "c2lnbmF0dXJl",
	}
}

func testWebAppKey() ed25519.PrivateKey {
	seed := make([]byte, ed25519.SeedSize)
	for i := range seed {
		seed[i] = byte(i + 1)
	}
	return ed25519.NewKeyFromSeed(seed)
}

func webAppTestSignature(priv ed25519.PrivateKey, botID int64, pairs map[string]string) string {
	keys := make([]string, 0, len(pairs))
	for k := range pairs {
		if k == "hash" || k == "signature" {
			continue
		}
		keys = append(keys, k)
	}
	sort.Strings(keys)

	lines := make([]string, len(keys))
	for i, k := range keys {
		lines[i] = k + "=" + pairs[k]
	}
	msg := fmt.Sprintf("%d:WebAppData\n", botID) + strings.Join(lines, "\n")
	return base64.RawURLEncoding.EncodeToString(ed25519.Sign(priv, []byte(msg)))
}

func signEd25519WebAppQuery(priv ed25519.PrivateKey, botID int64, pairs map[string]string) string {
	values := url.Values{}
	for k, v := range pairs {
		values.Set(k, v)
	}
	values.Set("signature", webAppTestSignature(priv, botID, pairs))
	values.Set("hash", webAppHash(testWebAppToken, pairs))
	return values.Encode()
}

func reencode(t *testing.T, raw string, mutate func(url.Values)) string {
	t.Helper()
	values, err := url.ParseQuery(raw)
	if err != nil {
		t.Fatal(err)
	}
	mutate(values)
	return values.Encode()
}

func requireWebAppError(t *testing.T, data *WebAppInitData, err, want error) {
	t.Helper()
	if !errors.Is(err, want) {
		t.Fatalf("err = %v, want %v", err, want)
	}
	if data != nil {
		t.Fatalf("data = %+v, want nil", data)
	}
}

func requireWebAppConfigurationError(t *testing.T, data *WebAppInitData, err error) {
	t.Helper()
	if err == nil || data != nil {
		t.Fatalf("data = %+v, err = %v; want nil data and a configuration error", data, err)
	}
	for _, sentinel := range webAppSentinels {
		if errors.Is(err, sentinel) {
			t.Fatalf("err %v matches data error %v", err, sentinel)
		}
	}
}

func TestValidateWebAppInitData_Full(t *testing.T) {
	pairs := fullWebAppPairs()
	raw := signWebAppQuery(testWebAppToken, pairs)

	data, err := ValidateWebAppInitData(raw, testWebAppToken)
	if err != nil {
		t.Fatal(err)
	}

	want := &WebAppInitData{
		QueryID:                "AAHdF6IQAAAAAN0XohDhrOrc",
		ChatJoinRequestQueryID: "JRQ-1",
		User: &WebAppUser{
			ID:                    fullWebAppUserID,
			FirstName:             "Alice",
			LastName:              "Smith",
			Username:              "alice_test",
			LanguageCode:          "ru",
			IsPremium:             true,
			AddedToAttachmentMenu: true,
			AllowsWriteToPM:       true,
			PhotoURL:              "https://t.me/i/userpic/320/a.svg",
		},
		Receiver: &WebAppUser{ID: 42, IsBot: true, FirstName: "Bot"},
		Chat: &WebAppChat{
			ID:       -1001234567890,
			Type:     models.ChatTypeSupergroup,
			Title:    "Group",
			Username: "grp",
			PhotoURL: "https://t.me/i/chat.jpg",
		},
		ChatType:     "supergroup",
		ChatInstance: "-8318574285046036213",
		StartParam:   "ref_42",
		CanSendAfter: 10,
		AuthDate:     1770000000,
		Hash:         webAppHash(testWebAppToken, pairs),
		Signature:    "c2lnbmF0dXJl",
	}
	if !reflect.DeepEqual(data, want) {
		t.Fatalf("data = %+v\nwant %+v", data, want)
	}
}

func TestValidateWebAppInitData_Minimal(t *testing.T) {
	pairs := map[string]string{"auth_date": "1770000000"}
	data, err := ValidateWebAppInitData(signWebAppQuery(testWebAppToken, pairs), testWebAppToken)
	if err != nil {
		t.Fatal(err)
	}
	want := &WebAppInitData{AuthDate: 1770000000, Hash: webAppHash(testWebAppToken, pairs)}
	if !reflect.DeepEqual(data, want) {
		t.Fatalf("data = %+v, want %+v", data, want)
	}
}

func TestValidateWebAppInitData_AuthDateIntRange(t *testing.T) {
	const maxInt32 = 1<<31 - 1
	for _, authDate := range []int64{maxInt32, maxInt32 + 1} {
		value := strconv.FormatInt(authDate, 10)
		t.Run(value, func(t *testing.T) {
			raw := signWebAppQuery(testWebAppToken, map[string]string{"auth_date": value})
			data, err := ValidateWebAppInitData(raw, testWebAppToken)
			if strconv.IntSize == 32 && authDate > maxInt32 {
				requireWebAppError(t, data, err, ErrorWebAppInitDataMalformed)
				return
			}
			if err != nil {
				t.Fatal(err)
			}
			if int64(data.AuthDate) != authDate {
				t.Fatalf("auth_date = %d, want %d", data.AuthDate, authDate)
			}
		})
	}
}

func TestValidateWebAppInitData_SignedNumbers(t *testing.T) {
	const botID = 7342037359
	priv := testWebAppKey()
	pub := priv.Public().(ed25519.PublicKey)

	tests := []struct {
		name                      string
		authDate, canSend         string
		wantAuthDate, wantCanSend int
	}{
		{"positive auth date, negative delay", "+1", "-5", 1, -5},
		{"negative auth date, positive delay", "-1", "+5", -1, 5},
		{"zero delay", "1", "0", 1, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pairs := map[string]string{"auth_date": tt.authDate, "can_send_after": tt.canSend}

			check := func(t *testing.T, data *WebAppInitData, err error) {
				t.Helper()
				if err != nil {
					t.Fatal(err)
				}
				if data.AuthDate != tt.wantAuthDate || data.CanSendAfter != tt.wantCanSend {
					t.Fatalf("auth_date = %d, can_send_after = %d; want %d, %d", data.AuthDate, data.CanSendAfter, tt.wantAuthDate, tt.wantCanSend)
				}
			}
			t.Run("HMAC", func(t *testing.T) {
				data, err := ValidateWebAppInitData(signWebAppQuery(testWebAppToken, pairs), testWebAppToken)
				check(t, data, err)
			})
			t.Run("Ed25519", func(t *testing.T) {
				data, err := ValidateWebAppInitDataSignature(signEd25519WebAppQuery(priv, botID, pairs), botID, pub)
				check(t, data, err)
			})
		})
	}
}

func TestValidateWebAppInitData_EmptyToken(t *testing.T) {
	tests := map[string]string{
		"valid":     signWebAppQuery(testWebAppToken, fullWebAppPairs()),
		"unsigned":  "",
		"malformed": "auth_date=%zz",
	}
	for name, input := range tests {
		t.Run(name, func(t *testing.T) {
			data, err := ValidateWebAppInitData(input, "")
			requireWebAppConfigurationError(t, data, err)
		})
	}
}

func TestValidateWebAppInitData_Tampered(t *testing.T) {
	raw := signWebAppQuery(testWebAppToken, fullWebAppPairs())

	tests := []struct {
		name   string
		mutate func(url.Values)
	}{
		{"change user", func(v url.Values) { v.Set("user", `{"id":1,"first_name":"Alice"}`) }},
		{"change auth_date", func(v url.Values) { v.Set("auth_date", "1770000001") }},
		{"change signature", func(v url.Values) { v.Set("signature", "c2lnbmF0dXJm") }},
		{"remove query_id", func(v url.Values) { v.Del("query_id") }},
		{"remove signature", func(v url.Values) { v.Del("signature") }},
		{"add field", func(v url.Values) { v.Set("extra", "1") }},
		{"uppercase hash", func(v url.Values) { v.Set("hash", strings.ToUpper(v.Get("hash"))) }},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := ValidateWebAppInitData(reencode(t, raw, tt.mutate), testWebAppToken)
			requireWebAppError(t, data, err, ErrorWebAppInitDataInvalid)
		})
	}

	t.Run("wrong token", func(t *testing.T) {
		data, err := ValidateWebAppInitData(raw, "654321:other")
		requireWebAppError(t, data, err, ErrorWebAppInitDataInvalid)
	})
}

func TestValidateWebAppInitData_Encoding(t *testing.T) {
	const startParam = "a+b=100% 你好 🙂"
	const userJSON = `{"id":5,"first_name":"A+B C 100% 你好 🙂\nnext line"}`
	pairs := map[string]string{
		"auth_date":   "1770000000",
		"start_param": startParam,
		"user":        userJSON,
	}
	plus := signWebAppQuery(testWebAppToken, pairs)
	if !strings.Contains(plus, "+") || !strings.Contains(plus, "%2B") {
		t.Fatalf("fixture must encode spaces as + and + as %%2B: %s", plus)
	}
	percent20 := strings.ReplaceAll(plus, "+", "%20")

	for name, raw := range map[string]string{"plus": plus, "percent20": percent20} {
		t.Run(name, func(t *testing.T) {
			data, err := ValidateWebAppInitData(raw, testWebAppToken)
			if err != nil {
				t.Fatal(err)
			}
			if data.StartParam != startParam {
				t.Fatalf("start_param = %q", data.StartParam)
			}
			if data.User == nil || data.User.ID != 5 || data.User.FirstName != "A+B C 100% 你好 🙂\nnext line" {
				t.Fatalf("user = %+v", data.User)
			}
		})
	}
}

func TestValidateWebAppInitData_Lexicographic(t *testing.T) {
	pairs := map[string]string{
		"auth_date":                  "1770000000",
		"can_send_after":             "5",
		"chat":                       `{"id":-100,"type":"group","title":"G"}`,
		"chat_instance":              "123",
		"chat_join_request_query_id": "JRQ",
		"chat_type":                  "group",
		"custom_field":               "x",
		"query_id":                   "Q",
		"signature":                  "sig",
		"start_param":                "s",
		"user":                       `{"id":1,"first_name":"Ann"}`,
	}
	raw := signWebAppQuery(testWebAppToken, pairs)

	data, err := ValidateWebAppInitData(raw, testWebAppToken)
	if err != nil {
		t.Fatal(err)
	}
	if data.ChatInstance != "123" || data.ChatType != "group" || data.CanSendAfter != 5 {
		t.Fatalf("data = %+v", data)
	}

	tampered := reencode(t, raw, func(v url.Values) { v.Set("custom_field", "y") })
	data, err = ValidateWebAppInitData(tampered, testWebAppToken)
	requireWebAppError(t, data, err, ErrorWebAppInitDataInvalid)
}

func TestValidateWebAppInitData_Unsigned(t *testing.T) {
	raw := signWebAppQuery(testWebAppToken, map[string]string{"auth_date": "1770000000"})

	tests := map[string]string{
		"empty":      "",
		"no hash":    reencode(t, raw, func(v url.Values) { v.Del("hash") }),
		"empty hash": reencode(t, raw, func(v url.Values) { v.Set("hash", "") }),
	}
	for name, input := range tests {
		t.Run(name, func(t *testing.T) {
			data, err := ValidateWebAppInitData(input, testWebAppToken)
			requireWebAppError(t, data, err, ErrorWebAppInitDataUnsigned)
		})
	}
}

func TestValidateWebAppInitData_MalformedQuery(t *testing.T) {
	raw := signWebAppQuery(testWebAppToken, map[string]string{"auth_date": "1770000000"})
	hash := webAppHash(testWebAppToken, map[string]string{"auth_date": "1770000000"})

	tests := map[string]string{
		"repeated key":       "auth_date=1&auth_date=2&hash=" + hash,
		"unsigned duplicate": "auth_date=1&auth_date=2",
		"repeated hash":      raw + "&hash=" + hash,
		"semicolon":          "auth_date=1770000000;hash=" + hash,
		"bad percent escape": "auth_date=%zz&hash=" + hash,
	}
	for name, input := range tests {
		t.Run(name, func(t *testing.T) {
			data, err := ValidateWebAppInitData(input, testWebAppToken)
			requireWebAppError(t, data, err, ErrorWebAppInitDataMalformed)
		})
	}
}

func TestWebAppValidators_RejectFieldSeparators(t *testing.T) {
	const botID = 7342037359
	pub := testWebAppKey().Public().(ed25519.PublicKey)

	tests := []struct {
		name  string
		key   string
		value string
	}{
		{"equals in key", "extra=field", "value"},
		{"line feed in key", "extra\nfield", "value"},
		{"line feed in value", "extra", "line\nbreak"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			values := url.Values{
				"auth_date": {"1770000000"},
				tt.key:      {tt.value},
			}
			raw := values.Encode()
			data, err := ValidateWebAppInitData(raw, testWebAppToken)
			requireWebAppError(t, data, err, ErrorWebAppInitDataMalformed)

			data, err = ValidateWebAppInitDataSignature(raw, botID, pub)
			requireWebAppError(t, data, err, ErrorWebAppInitDataMalformed)

			if fields := prepareWebAppRequest(values, testWebAppToken); fields != nil {
				t.Fatalf("legacy signed fields = %v, want nil", fields)
			}
			if user, ok := ValidateWebappRequest(values, testWebAppToken); ok || user != nil {
				t.Fatalf("legacy: ok=%v user=%+v", ok, user)
			}
		})
	}
}

func TestValidateWebAppInitData_MalformedFields(t *testing.T) {
	tests := []struct {
		name  string
		pairs map[string]string
		field string
	}{
		{"no auth_date", map[string]string{"query_id": "Q"}, "auth_date"},
		{"auth_date not integer", map[string]string{"auth_date": "1.5"}, "auth_date"},
		{"can_send_after not integer", map[string]string{"auth_date": "1770000000", "can_send_after": "x"}, "can_send_after"},
		{"user not json", map[string]string{"auth_date": "1770000000", "user": "not-json"}, "user"},
		{"receiver not json", map[string]string{"auth_date": "1770000000", "receiver": "{"}, "receiver"},
		{"chat not json", map[string]string{"auth_date": "1770000000", "chat": `{"id":"x"}`}, "chat"},
		{"all bad", map[string]string{"auth_date": "x", "can_send_after": "x", "user": "x", "receiver": "x", "chat": "x"}, "auth_date"},
		{"from can_send_after", map[string]string{"auth_date": "1", "can_send_after": "x", "user": "x", "receiver": "x", "chat": "x"}, "can_send_after"},
		{"from user", map[string]string{"auth_date": "1", "user": "x", "receiver": "x", "chat": "x"}, "user"},
		{"from receiver", map[string]string{"auth_date": "1", "receiver": "x", "chat": "x"}, "receiver"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := ValidateWebAppInitData(signWebAppQuery(testWebAppToken, tt.pairs), testWebAppToken)
			requireWebAppError(t, data, err, ErrorWebAppInitDataMalformed)
			if !strings.Contains(err.Error(), "field "+tt.field+":") {
				t.Fatalf("err = %v, want field %s", err, tt.field)
			}
		})
	}
}

func TestValidateWebAppInitData_Null(t *testing.T) {
	pairs := map[string]string{
		"auth_date": "1770000000",
		"user":      "null",
		"receiver":  "null",
		"chat":      "null",
	}
	data, err := ValidateWebAppInitData(signWebAppQuery(testWebAppToken, pairs), testWebAppToken)
	if err != nil {
		t.Fatal(err)
	}
	if data.User != nil || data.Receiver != nil || data.Chat != nil {
		t.Fatalf("data = %+v", data)
	}
}

func TestValidateWebAppInitData_PrefixNotStripped(t *testing.T) {
	raw := signWebAppQuery(testWebAppToken, map[string]string{"auth_date": "1770000000", "query_id": "Q"})

	for name, input := range map[string]string{
		"question mark":         "?" + raw,
		"fragment":              "#tgWebAppData=" + raw,
		"fragment, escaped":     "#tgWebAppData=" + url.QueryEscape(raw) + "&tgWebAppVersion=8.0",
		"tgWebAppData, escaped": "tgWebAppData=" + url.QueryEscape(raw),
	} {
		t.Run(name, func(t *testing.T) {
			data, err := ValidateWebAppInitData(input, testWebAppToken)
			if !errors.Is(err, ErrorWebAppInitDataInvalid) && !errors.Is(err, ErrorWebAppInitDataUnsigned) {
				t.Fatalf("err = %v", err)
			}
			if data != nil {
				t.Fatalf("data = %+v", data)
			}
		})
	}
}

func TestValidateWebAppInitDataSignature(t *testing.T) {
	const botID = 7342037359
	priv := testWebAppKey()
	pub := priv.Public().(ed25519.PublicKey)

	pairs := fullWebAppPairs()
	delete(pairs, "signature")
	raw := signEd25519WebAppQuery(priv, botID, pairs)

	t.Run("valid", func(t *testing.T) {
		data, err := ValidateWebAppInitDataSignature(raw, botID, pub)
		if err != nil {
			t.Fatal(err)
		}
		if data.User == nil || data.User.ID != fullWebAppUserID || data.Chat == nil || data.Chat.Title != "Group" ||
			data.AuthDate != 1770000000 || data.Hash != webAppHash(testWebAppToken, pairs) ||
			data.Signature != webAppTestSignature(priv, botID, pairs) {
			t.Fatalf("data = %+v", data)
		}
	})

	t.Run("minimal", func(t *testing.T) {
		pairs := map[string]string{"auth_date": "1770000000"}
		raw := reencode(t, signEd25519WebAppQuery(priv, botID, pairs), func(v url.Values) { v.Del("hash") })
		data, err := ValidateWebAppInitDataSignature(raw, botID, pub)
		if err != nil {
			t.Fatal(err)
		}
		if data.AuthDate != 1770000000 || data.User != nil || data.Receiver != nil || data.Chat != nil {
			t.Fatalf("data = %+v", data)
		}
	})

	t.Run("padding", func(t *testing.T) {
		padded := reencode(t, raw, func(v url.Values) {
			sig, err := base64.RawURLEncoding.DecodeString(v.Get("signature"))
			if err != nil {
				t.Fatal(err)
			}
			v.Set("signature", base64.URLEncoding.EncodeToString(sig))
		})
		if !strings.Contains(padded, "%3D") {
			t.Fatalf("fixture has no padding: %s", padded)
		}
		if _, err := ValidateWebAppInitDataSignature(padded, botID, pub); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("hash not required", func(t *testing.T) {
		noHash := reencode(t, raw, func(v url.Values) { v.Del("hash") })
		if _, err := ValidateWebAppInitDataSignature(noHash, botID, pub); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("test public key does not match", func(t *testing.T) {
		key, err := hex.DecodeString(WebAppTestPublicKeyHex)
		if err != nil {
			t.Fatal(err)
		}
		data, err := ValidateWebAppInitDataSignature(raw, botID, key)
		requireWebAppError(t, data, err, ErrorWebAppInitDataInvalid)
	})

	t.Run("other bot id", func(t *testing.T) {
		data, err := ValidateWebAppInitDataSignature(raw, botID+1, pub)
		requireWebAppError(t, data, err, ErrorWebAppInitDataInvalid)
	})

	invalid := []struct {
		name   string
		mutate func(url.Values)
	}{
		{"change user", func(v url.Values) { v.Set("user", `{"id":1}`) }},
		{"add field", func(v url.Values) { v.Set("extra", "1") }},
		{"remove field", func(v url.Values) { v.Del("start_param") }},
		{"not base64url", func(v url.Values) { v.Set("signature", "!!!") }},
		{"std base64", func(v url.Values) { v.Set("signature", strings.Repeat("+/", 43)) }},
		{"short signature", func(v url.Values) { v.Set("signature", base64.RawURLEncoding.EncodeToString([]byte("short"))) }},
	}
	for _, tt := range invalid {
		t.Run(tt.name, func(t *testing.T) {
			data, err := ValidateWebAppInitDataSignature(reencode(t, raw, tt.mutate), botID, pub)
			requireWebAppError(t, data, err, ErrorWebAppInitDataInvalid)
		})
	}

	unsigned := map[string]string{
		"empty":           "",
		"no signature":    reencode(t, raw, func(v url.Values) { v.Del("signature") }),
		"empty signature": reencode(t, raw, func(v url.Values) { v.Set("signature", "") }),
	}
	for name, input := range unsigned {
		t.Run(name, func(t *testing.T) {
			data, err := ValidateWebAppInitDataSignature(input, botID, pub)
			requireWebAppError(t, data, err, ErrorWebAppInitDataUnsigned)
		})
	}

	malformed := map[string]string{
		"repeated key":       raw + "&auth_date=1",
		"unsigned duplicate": "auth_date=1&auth_date=2",
		"semicolon":          strings.Replace(raw, "&", ";", 1),
		"no auth_date":       signEd25519WebAppQuery(priv, botID, map[string]string{"query_id": "Q"}),
		"bad auth_date":      signEd25519WebAppQuery(priv, botID, map[string]string{"auth_date": "x"}),
		"user not json":      signEd25519WebAppQuery(priv, botID, map[string]string{"auth_date": "1", "user": "x"}),
		"chat not json":      signEd25519WebAppQuery(priv, botID, map[string]string{"auth_date": "1", "chat": "x"}),
		"bad can_send":       signEd25519WebAppQuery(priv, botID, map[string]string{"auth_date": "1", "can_send_after": "x"}),
		"bad percent":        "auth_date=%zz&signature=" + url.QueryEscape(webAppTestSignature(priv, botID, nil)),
		"receiver error":     signEd25519WebAppQuery(priv, botID, map[string]string{"auth_date": "1", "receiver": "x"}),
	}
	for name, input := range malformed {
		t.Run(name, func(t *testing.T) {
			data, err := ValidateWebAppInitDataSignature(input, botID, pub)
			requireWebAppError(t, data, err, ErrorWebAppInitDataMalformed)
		})
	}
}

func TestDecodeWebAppSignature_Base64URL(t *testing.T) {
	// This byte sequence exercises the URL alphabet and trailing bits without a valid signature.
	sig := make([]byte, ed25519.SignatureSize)
	copy(sig, []byte{0xfb, 0xff, 0xff})
	padded := base64.URLEncoding.EncodeToString(sig)
	unpadded := base64.RawURLEncoding.EncodeToString(sig)

	for name, input := range map[string]string{"padded": padded, "unpadded": unpadded} {
		t.Run(name, func(t *testing.T) {
			got, err := decodeWebAppSignature(input)
			if err != nil || !reflect.DeepEqual(got, sig) {
				t.Fatalf("decoded = %x, err = %v, want %x", got, err, sig)
			}
		})
	}

	for name, input := range map[string]string{
		"partial padding":   unpadded + "=",
		"extra padding":     padded + "=",
		"carriage return":   unpadded[:4] + "\r" + unpadded[4:],
		"line feed":         unpadded[:4] + "\n" + unpadded[4:],
		"nonzero pad bits":  unpadded[:len(unpadded)-1] + "B",
		"standard alphabet": base64.RawStdEncoding.EncodeToString(sig),
	} {
		t.Run(name, func(t *testing.T) {
			if _, err := decodeWebAppSignature(input); err == nil {
				t.Fatal("malformed base64url decoded successfully")
			}
		})
	}
}

func TestValidateWebAppInitDataSignature_CallerErrors(t *testing.T) {
	pub := testWebAppKey().Public().(ed25519.PublicKey)

	tests := []struct {
		name  string
		botID int64
		key   ed25519.PublicKey
	}{
		{"nil key", 1, nil},
		{"short key", 1, pub[:31]},
		{"hex string as key", 1, ed25519.PublicKey(WebAppPublicKeyHex)},
		{"zero bot id", 0, pub},
		{"negative bot id", -1, pub},
		{"both wrong", 0, nil},
	}
	inputs := []string{"", "a=1;b=2", "auth_date=1&auth_date=2", "signature=!!!"}

	for _, tt := range tests {
		for _, input := range inputs {
			t.Run(tt.name+"/"+input, func(t *testing.T) {
				data, err := ValidateWebAppInitDataSignature(input, tt.botID, tt.key)
				requireWebAppConfigurationError(t, data, err)
			})
		}
	}
}

func TestWebAppPublicKeys(t *testing.T) {
	// https://core.telegram.org/bots/webapps#validating-data-for-third-party-use
	tests := []struct {
		name string
		key  string
		want string
	}{
		{"production", WebAppPublicKeyHex, "e7bf03a2fa4602af4580703d88dda5bb59f32ed8b02a56c187fe7d34caed242d"},
		{"test", WebAppTestPublicKeyHex, "40055058a4ee38156a06562e52eece92a771bcd8346a8c4615cb7376eddf72ec"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key, err := hex.DecodeString(tt.key)
			if err != nil {
				t.Fatal(err)
			}
			if got := hex.EncodeToString(key); got != tt.want {
				t.Fatalf("public key = %s, want %s", got, tt.want)
			}
		})
	}
}

func TestWebAppInitData_CheckAuthDate(t *testing.T) {
	now := time.Unix(1770000000, 0)

	tests := []struct {
		name   string
		offset time.Duration
		maxAge time.Duration
		want   error
	}{
		{"expired", -2 * time.Hour, time.Hour, ErrorWebAppInitDataExpired},
		{"fresh", -time.Minute, time.Hour, nil},
		{"future", time.Hour, time.Minute, nil},
		{"zero max age", time.Hour, 0, ErrorWebAppInitDataExpired},
		{"negative max age", time.Hour, -time.Hour, ErrorWebAppInitDataExpired},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			raw := signWebAppQuery(testWebAppToken, map[string]string{
				"auth_date": strconv.FormatInt(now.Add(tt.offset).Unix(), 10),
			})
			data, err := ValidateWebAppInitData(raw, testWebAppToken)
			if err != nil {
				t.Fatal(err)
			}
			if err := data.checkAuthDate(now, tt.maxAge); err != tt.want {
				t.Fatalf("err = %v, want %v", err, tt.want)
			}
		})
	}
}

func TestWebAppInitData_CheckAuthDate_Public(t *testing.T) {
	raw := signWebAppQuery(testWebAppToken, map[string]string{"auth_date": "0"})
	data, err := ValidateWebAppInitData(raw, testWebAppToken)
	if err != nil {
		t.Fatal(err)
	}
	if err := data.CheckAuthDate(time.Hour); !errors.Is(err, ErrorWebAppInitDataExpired) {
		t.Fatalf("err = %v, want %v", err, ErrorWebAppInitDataExpired)
	}
	const maxAge = time.Duration(1<<63 - 1)
	if err := data.CheckAuthDate(maxAge); err != nil {
		t.Fatalf("within max age: %v", err)
	}
}

func TestWebAppInitData_CheckAuthDateCutoff(t *testing.T) {
	const maxAge = time.Hour
	authTime := time.Unix(1770000000, 0)
	data := WebAppInitData{AuthDate: int(authTime.Unix())}
	deadline := authTime.Add(maxAge)

	if err := data.checkAuthDate(deadline, maxAge); err != nil {
		t.Fatalf("at deadline: %v", err)
	}
	if err := data.checkAuthDate(deadline.Add(time.Nanosecond), maxAge); !errors.Is(err, ErrorWebAppInitDataExpired) {
		t.Fatalf("after deadline: err = %v, want %v", err, ErrorWebAppInitDataExpired)
	}
}

func TestValidateWebappRequest_EmptyToken(t *testing.T) {
	values := url.Values{"auth_date": {"1770000000"}}
	if fields := prepareWebAppRequest(values, ""); fields != nil {
		t.Fatalf("legacy signed fields = %v, want nil", fields)
	}
	if user, ok := ValidateWebappRequest(values, ""); ok || user != nil {
		t.Fatalf("ok=%v user=%+v", ok, user)
	}
}

func TestWebAppInitData_CheckAuthDateExtremes(t *testing.T) {
	now := time.Unix(1770000000, 0)
	maxTimestamp := int(^uint(0) >> 1)
	future := WebAppInitData{AuthDate: maxTimestamp}
	if err := future.checkAuthDate(now, time.Hour); err != nil {
		t.Fatalf("max timestamp: %v", err)
	}

	past := WebAppInitData{AuthDate: -maxTimestamp - 1}
	if err := past.checkAuthDate(now, time.Hour); !errors.Is(err, ErrorWebAppInitDataExpired) {
		t.Fatalf("min timestamp: err = %v, want %v", err, ErrorWebAppInitDataExpired)
	}

	if strconv.IntSize == 64 {
		oldTimestamp := time.Date(1600, time.January, 1, 0, 0, 0, 0, time.UTC).Unix()
		old := WebAppInitData{AuthDate: int(oldTimestamp)}
		if err := old.checkAuthDate(now, time.Duration(1<<63-1)); !errors.Is(err, ErrorWebAppInitDataExpired) {
			t.Fatalf("age exceeds max duration: err = %v, want %v", err, ErrorWebAppInitDataExpired)
		}
	}
}

func TestValidateWebappRequest_UserNull(t *testing.T) {
	values, err := url.ParseQuery(signWebAppQuery(testWebAppToken, map[string]string{
		"auth_date": "1770000000",
		"user":      "null",
	}))
	if err != nil {
		t.Fatal(err)
	}
	user, ok := ValidateWebappRequest(values, testWebAppToken)
	if !ok || user == nil || *user != (WebAppUser{}) {
		t.Fatalf("ok=%v user=%+v", ok, user)
	}
}

func TestValidateWebappRequest_RepeatedValues(t *testing.T) {
	values, err := url.ParseQuery(signWebAppQuery(testWebAppToken, fullWebAppPairs()))
	if err != nil {
		t.Fatal(err)
	}
	values.Add("user", `{"id":42,"first_name":"Bob"}`)
	values.Add("hash", "unused")
	before := url.Values{}
	for k, v := range values {
		before[k] = append([]string(nil), v...)
	}

	user, ok := ValidateWebappRequest(values, testWebAppToken)
	if !ok || user == nil || user.ID != fullWebAppUserID {
		t.Fatalf("ok=%v user=%+v", ok, user)
	}
	if !reflect.DeepEqual(values, before) {
		t.Fatalf("values modified: %v", values)
	}
}

func TestValidateWebappRequest_FullInitData(t *testing.T) {
	values, err := url.ParseQuery(signWebAppQuery(testWebAppToken, fullWebAppPairs()))
	if err != nil {
		t.Fatal(err)
	}
	user, ok := ValidateWebappRequest(values, testWebAppToken)
	if !ok || user == nil || user.ID != fullWebAppUserID || user.FirstName != "Alice" {
		t.Fatalf("ok=%v user=%+v", ok, user)
	}
	noUser := fullWebAppPairs()
	delete(noUser, "user")
	values, err = url.ParseQuery(signWebAppQuery(testWebAppToken, noUser))
	if err != nil {
		t.Fatal(err)
	}
	user, ok = ValidateWebappRequest(values, testWebAppToken)
	if !ok || user == nil || *user != (WebAppUser{}) {
		t.Fatalf("no user: ok=%v user=%+v", ok, user)
	}

	values["unsent_field"] = []string{}
	if user, ok := ValidateWebappRequest(values, testWebAppToken); ok || user != nil {
		t.Fatalf("empty value slice: ok=%v user=%+v", ok, user)
	}
}
