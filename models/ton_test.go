package models

import (
	"encoding/json"
	"testing"
)

// The 10.4 rename keeps the old name usable.
var _ StoryAreaPosition = AreaPosition{}

func TestChatFullInfo_TonWalletCard(t *testing.T) {
	src := `{"id":1,"type":"private","max_reaction_count":0,"bio":"","join_to_send_messages":false,"join_by_request":false,"accepted_gift_types":{"unlimited_gifts":false,"limited_gifts":false,"unique_gifts":false,"premium_subscription":false},` +
		`"ton_wallet_card":{"address":"EQabc","amount":"123456","stickers":[{"position":{"x_percentage":50,"y_percentage":50,"width_percentage":10,"height_percentage":10,"rotation_angle":0,"corner_radius_percentage":0},"nft":{"address":"EQnft","collection_address":"EQcol","index":"7","name":"Cap","description":"d","attributes":[{"trait_type":"Model","value":"Red"}]}}],"background_nft_address":"EQnft"}}`
	var c ChatFullInfo
	if err := json.Unmarshal([]byte(src), &c); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	card := c.TonWalletCard
	if card == nil || card.Address != "EQabc" || card.Amount != "123456" || card.BackgroundNftAddress != "EQnft" {
		t.Fatalf("ton_wallet_card not decoded: %+v", card)
	}
	if len(card.Stickers) != 1 || card.Stickers[0].Position.XPercentage != 50 || card.Stickers[0].Nft.Index != "7" {
		t.Fatalf("stickers not decoded: %+v", card.Stickers)
	}
	if attrs := card.Stickers[0].Nft.Attributes; len(attrs) != 1 || attrs[0].TraitType != "Model" || attrs[0].Value != "Red" {
		t.Fatalf("nft attributes not decoded: %+v", attrs)
	}
}

func TestMessage_TonWalletTransfer(t *testing.T) {
	src := `{"message_id":1,"date":1,"chat":{"id":1,"type":"private"},"ton_wallet_transfer":{"transaction_id":"tx","peer_address":"EQpeer","amount":"-500","comment":"thanks"}}`
	var m Message
	if err := json.Unmarshal([]byte(src), &m); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	tr := m.TonWalletTransfer
	if tr == nil || tr.TransactionID != "tx" || tr.PeerAddress != "EQpeer" || tr.Amount != "-500" || tr.Comment != "thanks" {
		t.Fatalf("ton_wallet_transfer not decoded: %+v", tr)
	}
}

func TestMessage_PaidMessagesWaived(t *testing.T) {
	src := `{"message_id":1,"date":1,"chat":{"id":1,"type":"private"},"paid_messages_waived":{}}`
	var m Message
	if err := json.Unmarshal([]byte(src), &m); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if m.PaidMessagesWaived == nil {
		t.Fatal("paid_messages_waived not decoded")
	}
}

func TestMessageEntity_TonAddressAndBankCard(t *testing.T) {
	src := `[{"type":"ton_address","offset":0,"length":48},{"type":"bank_card_number","offset":49,"length":16}]`
	var es []MessageEntity
	if err := json.Unmarshal([]byte(src), &es); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if es[0].Type != MessageEntityTypeTonAddress || es[1].Type != MessageEntityTypeBankCardNumber {
		t.Fatalf("entity types: %+v", es)
	}
}
