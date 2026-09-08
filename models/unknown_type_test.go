package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

const unknownType = "type_added_in_future_api"

// Telegram adds new variants with each Bot API release. A polymorphic model must keep
// the unknown discriminator and decode without error, otherwise the whole update is lost.
func TestPolymorphicUnmarshal_UnknownDiscriminator(t *testing.T) {
	var (
		chatMember     ChatMember
		reaction       ReactionType
		boostSource    ChatBoostSource
		ownedGift      OwnedGift
		menuButton     MenuButton
		messageOrigin  MessageOrigin
		storyArea      StoryAreaType
		partner        TransactionPartner
		withdrawal     RevenueWithdrawalState
		backgroundType BackgroundType
		backgroundFill BackgroundFill
		richBlock      RichBlock
		richText       RichText
		paidMedia      PaidMedia
	)

	cases := []struct {
		name string
		src  string
		dst  any
		typ  func() string
	}{
		{"ChatMember", `{"status":"` + unknownType + `","user":{"id":1}}`, &chatMember, func() string { return string(chatMember.Type) }},
		{"ReactionType", `{"type":"` + unknownType + `"}`, &reaction, func() string { return string(reaction.Type) }},
		{"ChatBoostSource", `{"source":"` + unknownType + `"}`, &boostSource, func() string { return string(boostSource.Source) }},
		{"OwnedGift", `{"type":"` + unknownType + `"}`, &ownedGift, func() string { return string(ownedGift.Type) }},
		{"MenuButton", `{"type":"` + unknownType + `"}`, &menuButton, func() string { return string(menuButton.Type) }},
		{"MessageOrigin", `{"type":"` + unknownType + `"}`, &messageOrigin, func() string { return string(messageOrigin.Type) }},
		{"StoryAreaType", `{"type":"` + unknownType + `"}`, &storyArea, func() string { return string(storyArea.Type) }},
		{"TransactionPartner", `{"type":"` + unknownType + `"}`, &partner, func() string { return string(partner.Type) }},
		{"RevenueWithdrawalState", `{"type":"` + unknownType + `"}`, &withdrawal, func() string { return string(withdrawal.Type) }},
		{"BackgroundType", `{"type":"` + unknownType + `"}`, &backgroundType, func() string { return string(backgroundType.Type) }},
		{"BackgroundFill", `{"type":"` + unknownType + `"}`, &backgroundFill, func() string { return string(backgroundFill.Type) }},
		{"RichBlock", `{"type":"` + unknownType + `"}`, &richBlock, func() string { return string(richBlock.Type) }},
		{"RichText", `{"type":"` + unknownType + `"}`, &richText, func() string { return string(richText.Type) }},
		{"PaidMedia", `{"type":"` + unknownType + `"}`, &paidMedia, func() string { return string(paidMedia.Type) }},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if unmarshalErr := json.Unmarshal([]byte(c.src), c.dst); unmarshalErr != nil {
				t.Fatalf("unknown discriminator must not fail: %v", unmarshalErr)
			}
			if got := c.typ(); got != unknownType {
				t.Fatalf("discriminator lost: got %q", got)
			}
			assertNoVariantSet(t, c.dst)
		})
	}
}

// assertNoVariantSet checks every variant pointer of a polymorphic struct stays nil.
func assertNoVariantSet(t *testing.T, dst any) {
	t.Helper()
	v := reflect.ValueOf(dst).Elem()
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		if f.Kind() == reflect.Pointer && !f.IsNil() {
			t.Fatalf("variant %s set for unknown type", v.Type().Field(i).Name)
		}
	}
}

// A value decoded from a newer Bot API release must survive a round trip. Anyone who
// logs, persists or queues updates as JSON, or echoes a value back into a request,
// would otherwise break on the day Telegram ships a new variant.
func TestPolymorphicMarshal_UnknownDiscriminator(t *testing.T) {
	cases := []struct {
		name  string
		value json.Unmarshaler
	}{
		{"ChatMember", &ChatMember{}},
		{"ReactionType", &ReactionType{}},
		{"ChatBoostSource", &ChatBoostSource{}},
		{"MenuButton", &MenuButton{}},
		{"MessageOrigin", &MessageOrigin{}},
		{"BackgroundType", &BackgroundType{}},
		{"BackgroundFill", &BackgroundFill{}},
		{"RichBlock", &RichBlock{}},
		{"RichText", &RichText{}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			field := "type"
			switch c.name {
			case "ChatMember":
				field = "status"
			case "ChatBoostSource":
				field = "source"
			}
			src := `{"` + field + `":"` + unknownType + `"}`

			if unmarshalErr := json.Unmarshal([]byte(src), c.value); unmarshalErr != nil {
				t.Fatalf("unknown discriminator must not fail to decode: %v", unmarshalErr)
			}

			out, marshalErr := json.Marshal(c.value)
			if marshalErr != nil {
				t.Fatalf("unknown discriminator must not fail to encode: %v", marshalErr)
			}
			if string(out) != src {
				t.Fatalf("round-trip mismatch:\n got %s\nwant %s", out, src)
			}
		})
	}
}

// An empty discriminator is an unset value, not a variant from a future release:
// there is no payload to encode, so it stays an error.
func TestPolymorphicMarshal_EmptyDiscriminator(t *testing.T) {
	// RichText is absent on purpose: an empty Type is its plain-string form.
	cases := []struct {
		name  string
		value json.Marshaler
	}{
		{"ChatMember", &ChatMember{}},
		{"ReactionType", &ReactionType{}},
		{"ChatBoostSource", &ChatBoostSource{}},
		{"MenuButton", &MenuButton{}},
		{"MessageOrigin", &MessageOrigin{}},
		{"BackgroundType", &BackgroundType{}},
		{"BackgroundFill", &BackgroundFill{}},
		{"RichBlock", &RichBlock{}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if _, marshalErr := c.value.MarshalJSON(); marshalErr == nil {
				t.Fatal("expected an error for an empty discriminator")
			}
		})
	}
}

// ReactionTypePaid is a known variant since Bot API 7.6, but MarshalJSON never had a
// case for it, so a paid reaction read from an update could not be sent back.
func TestReactionType_PaidRoundTrip(t *testing.T) {
	src := `{"type":"paid"}`

	rt := &ReactionType{}
	if unmarshalErr := json.Unmarshal([]byte(src), rt); unmarshalErr != nil {
		t.Fatalf("decode: %v", unmarshalErr)
	}
	if rt.ReactionTypePaid == nil {
		t.Fatal("paid variant not populated")
	}

	out, marshalErr := json.Marshal(rt)
	if marshalErr != nil {
		t.Fatalf("encode: %v", marshalErr)
	}
	if string(out) != src {
		t.Fatalf("round-trip mismatch:\n got %s\nwant %s", out, src)
	}
}

// An empty Type is RichText's plain-string form, so a tagged object without a "type"
// must not decode into it: that would silently alias a malformed value to "".
func TestRichText_ObjectWithoutType(t *testing.T) {
	var rt RichText
	if unmarshalErr := json.Unmarshal([]byte(`{"text":"hi"}`), &rt); unmarshalErr == nil {
		t.Fatal("expected an error for a tagged object without a type")
	}
}

// A Type set without its variant pointer is a caller mistake. It must come back as an
// error from MarshalJSON, never as a nil-pointer panic inside encoding/json.
func TestReactionType_MarshalNilVariant(t *testing.T) {
	for _, typ := range []ReactionTypeType{ReactionTypeTypeEmoji, ReactionTypeTypeCustomEmoji, ReactionTypeTypePaid} {
		t.Run(string(typ), func(t *testing.T) {
			if _, marshalErr := json.Marshal(&ReactionType{Type: typ}); marshalErr == nil {
				t.Fatal("expected an error for a Type without its variant")
			}
		})
	}
}

// A tagged object without its discriminator is malformed, not a variant from a future
// release. Decoding must reject it, the same way MarshalJSON rejects an empty Type,
// so a value that decodes can always be encoded again.
func TestPolymorphicUnmarshal_EmptyDiscriminator(t *testing.T) {
	cases := []struct {
		name string
		dst  json.Unmarshaler
	}{
		{"ChatMember", &ChatMember{}},
		{"ReactionType", &ReactionType{}},
		{"ChatBoostSource", &ChatBoostSource{}},
		{"OwnedGift", &OwnedGift{}},
		{"MenuButton", &MenuButton{}},
		{"MessageOrigin", &MessageOrigin{}},
		{"StoryAreaType", &StoryAreaType{}},
		{"TransactionPartner", &TransactionPartner{}},
		{"RevenueWithdrawalState", &RevenueWithdrawalState{}},
		{"BackgroundType", &BackgroundType{}},
		{"BackgroundFill", &BackgroundFill{}},
		{"RichBlock", &RichBlock{}},
		{"RichText", &RichText{}},
		{"PaidMedia", &PaidMedia{}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if unmarshalErr := json.Unmarshal([]byte(`{}`), c.dst); unmarshalErr == nil {
				t.Fatal("expected an error for an object without a discriminator")
			}
		})
	}
}
