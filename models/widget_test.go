package models

import (
	"encoding/json"
	"testing"
)

func TestWidgetQueryType_RoundTrip(t *testing.T) {
	cases := []string{
		`{"type":"data"}`,
		`{"type":"data","wallet_card":{"address":"EQA","amount":"10","stickers":[]}}`,
		`{"type":"sign_data","wallet":{"address":"EQA","amount":"10","public_key":"aa","state_init":"bb"},"sign_data":"{\"payload\":1}"}`,
		`{"type":"sent_transaction","wallet":{"address":"EQA","amount":"10"},"msg_hash":"abc"}`,
	}
	for _, src := range cases {
		t.Run(src, func(t *testing.T) {
			var q WidgetQueryType
			if err := json.Unmarshal([]byte(src), &q); err != nil {
				t.Fatalf("unmarshal: %v", err)
			}
			out, err := json.Marshal(q)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			if string(out) != src {
				t.Fatalf("round-trip mismatch:\n got %s\nwant %s", out, src)
			}
		})
	}
}

func TestWidgetQueryType_Variants(t *testing.T) {
	var data WidgetQueryType
	if err := json.Unmarshal([]byte(`{"type":"data","wallet_card":{"address":"EQA","amount":"10","stickers":[]}}`), &data); err != nil {
		t.Fatal(err)
	}
	if data.WidgetQueryTypeData == nil || data.WidgetQueryTypeData.WalletCard == nil || data.WidgetQueryTypeData.WalletCard.Amount != "10" {
		t.Fatalf("data variant: %+v", data.WidgetQueryTypeData)
	}

	var signed WidgetQueryType
	if err := json.Unmarshal([]byte(`{"type":"sign_data","wallet":{"address":"EQA","amount":"10"},"sign_data":"raw"}`), &signed); err != nil {
		t.Fatal(err)
	}
	if signed.WidgetQueryTypeSignData == nil || signed.WidgetQueryTypeSignData.SignData != "raw" || signed.WidgetQueryTypeSignData.Wallet.Address != "EQA" {
		t.Fatalf("sign variant: %+v", signed.WidgetQueryTypeSignData)
	}

	var sent WidgetQueryType
	if err := json.Unmarshal([]byte(`{"type":"sent_transaction","wallet":{"address":"EQA","amount":"10"},"msg_hash":"abc"}`), &sent); err != nil {
		t.Fatal(err)
	}
	if sent.WidgetQueryTypeSentTransaction == nil || sent.WidgetQueryTypeSentTransaction.MsgHash != "abc" {
		t.Fatalf("sent variant: %+v", sent.WidgetQueryTypeSentTransaction)
	}
}

func TestSharedWidgetQuery_Decode(t *testing.T) {
	var q SharedWidgetQuery
	src := `{"id":"1","from":{"id":9},"type":{"type":"data"},"data":"ping","widget_id":"w"}`
	if err := json.Unmarshal([]byte(src), &q); err != nil {
		t.Fatal(err)
	}
	if q.From.ID != 9 || q.WidgetID != "w" || q.Data != "ping" || q.Type.WidgetQueryTypeData == nil {
		t.Fatalf("%+v", q)
	}
}

func TestMessage_WidgetAndServiceFields(t *testing.T) {
	var msg Message
	src := `{"message_id":1,"date":1,"chat":{"id":2,"type":"private"},"widget":{"html":"<b>hi</b>","id":"w"},"paid_messages_waived":{},"ton_wallet_transfer":{"transaction_id":"","peer_address":"EQ","amount":"-5"}}`
	if err := json.Unmarshal([]byte(src), &msg); err != nil {
		t.Fatal(err)
	}
	if msg.Widget == nil || msg.Widget.HTML != "<b>hi</b>" || msg.Widget.ID != "w" {
		t.Fatalf("widget: %+v", msg.Widget)
	}
	if msg.PaidMessagesWaived == nil {
		t.Fatal("paid messages waived missing")
	}
	if msg.TonWalletTransfer == nil || msg.TonWalletTransfer.TransactionID != "" || msg.TonWalletTransfer.Amount != "-5" {
		t.Fatalf("transfer: %+v", msg.TonWalletTransfer)
	}
}

func TestInputWidgetMessageContent_Marshal(t *testing.T) {
	article := &InlineQueryResultArticle{
		ID:    "1",
		Title: "widget",
		InputMessageContent: &InputWidgetMessageContent{
			HTML:         "<b>x</b>",
			AllowQueries: true,
			IsShared:     true,
		},
	}
	out, err := article.MarshalCustom()
	if err != nil {
		t.Fatal(err)
	}
	var decoded map[string]any
	if err := json.Unmarshal(out, &decoded); err != nil {
		t.Fatal(err)
	}
	if decoded["type"] != "article" {
		t.Fatalf("type %v", decoded["type"])
	}
	content, ok := decoded["input_message_content"].(map[string]any)
	if !ok {
		t.Fatalf("content %#v", decoded["input_message_content"])
	}
	if content["html"] != "<b>x</b>" || content["allow_queries"] != true || content["is_shared"] != true {
		t.Fatalf("content %#v", content)
	}

	shared := &InlineQueryResultArticle{
		ID:                  "2",
		Title:               "shared",
		InputMessageContent: &InputSharedWidgetMessageContent{WidgetID: "w", ChatData: "c"},
	}
	out, err = shared.MarshalCustom()
	if err != nil {
		t.Fatal(err)
	}
	decoded = nil
	if err := json.Unmarshal(out, &decoded); err != nil {
		t.Fatal(err)
	}
	content, ok = decoded["input_message_content"].(map[string]any)
	if !ok || content["widget_id"] != "w" || content["chat_data"] != "c" {
		t.Fatalf("shared content %#v", decoded["input_message_content"])
	}
}
