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
			if err := json.Unmarshal([]byte(c.src), c.dst); err != nil {
				t.Fatalf("unknown discriminator must not fail: %v", err)
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
