package models

import (
	"encoding/json"
	"testing"
)

func TestTonWalletCard_Decode(t *testing.T) {
	var card TonWalletCard
	src := `{
		"address":"EQA",
		"amount":"1000000000",
		"background_nft_address":"EQB",
		"stickers":[{
			"position":{"x_percentage":1,"y_percentage":2,"width_percentage":3,"height_percentage":4,"rotation_angle":5,"corner_radius_percentage":6},
			"nft":{"address":"EQN","index":"7","name":"Gift","description":"desc","attributes":[{"trait_type":"model","value":"gold"}]}
		}]
	}`
	if err := json.Unmarshal([]byte(src), &card); err != nil {
		t.Fatal(err)
	}
	if card.Amount != "1000000000" || card.BackgroundNftAddress != "EQB" || len(card.Stickers) != 1 {
		t.Fatalf("%+v", card)
	}
	sticker := card.Stickers[0]
	if sticker.Position.XPercentage != 1 || sticker.Nft.Index != "7" || sticker.Nft.Name != "Gift" {
		t.Fatalf("%+v", sticker)
	}
	if len(sticker.Nft.Attributes) != 1 || sticker.Nft.Attributes[0].Value != "gold" {
		t.Fatalf("attributes %+v", sticker.Nft.Attributes)
	}

	var legacy StoryAreaPosition = sticker.Position
	if legacy.WidthPercentage != 3 {
		t.Fatal("StoryAreaPosition alias did not keep the position")
	}
}

func TestTonWalletTransfer_EmptyTransactionID(t *testing.T) {
	src := `{"transaction_id":"","peer_address":"EQ","amount":"-5"}`
	var transfer TonWalletTransfer
	if err := json.Unmarshal([]byte(src), &transfer); err != nil {
		t.Fatal(err)
	}
	out, err := json.Marshal(transfer)
	if err != nil {
		t.Fatal(err)
	}
	if string(out) != src {
		t.Fatalf("got %s", out)
	}
}
