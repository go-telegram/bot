package models

import (
	"encoding/json"
	"testing"
)

func TestChatBoostSource_premium(t *testing.T) {
	src := `{"source":"premium","user":{"id":123}}`

	var cbs ChatBoostSource
	err := json.Unmarshal([]byte(src), &cbs)
	if err != nil {
		t.Fatal(err)
	}

	if cbs.Source != ChatBoostSourceTypePremium {
		t.Fatal("invalid type")
	}

	if cbs.ChatBoostSourcePremium == nil {
		t.Fatal("invalid ChatBoostSourcePremium")
	}

	if cbs.ChatBoostSourcePremium.User.ID != 123 {
		t.Fatal("invalid user id")
	}
}

func TestChatBoostSource_gift_code(t *testing.T) {
	src := `{"source":"gift_code","user":{"id":123}}`

	var cbs ChatBoostSource
	err := json.Unmarshal([]byte(src), &cbs)
	if err != nil {
		t.Fatal(err)
	}

	if cbs.Source != ChatBoostSourceTypeGiftCode {
		t.Fatal("invalid type")
	}

	if cbs.ChatBoostSourceGiftCode == nil {
		t.Fatal("invalid ChatBoostSourceGiftCode")
	}

	if cbs.ChatBoostSourceGiftCode.User.ID != 123 {
		t.Fatal("invalid user id")
	}
}

func TestChatBoostSource_giveaway(t *testing.T) {
	src := `{"source":"giveaway","user":{"id":123}}`

	var cbs ChatBoostSource
	err := json.Unmarshal([]byte(src), &cbs)
	if err != nil {
		t.Fatal(err)
	}

	if cbs.Source != ChatBoostSourceTypeGiveaway {
		t.Fatal("invalid type")
	}

	if cbs.ChatBoostSourceGiveaway == nil {
		t.Fatal("invalid ChatBoostSourceGiveaway")
	}

	if cbs.ChatBoostSourceGiveaway.User.ID != 123 {
		t.Fatal("invalid user id")
	}
}
