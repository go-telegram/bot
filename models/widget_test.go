package models

import (
	"encoding/json"
	"testing"
)

// Both widget contents are usable where an inline query result takes its content.
var (
	_ InputMessageContent = InputWidgetMessageContent{}
	_ InputMessageContent = InputSharedWidgetMessageContent{}
)

// widgetQueryTypeRoundTrip unmarshals src into a WidgetQueryType, marshals it back
// and asserts the JSON is byte-identical.
func widgetQueryTypeRoundTrip(t *testing.T, src string) WidgetQueryType {
	t.Helper()

	var w WidgetQueryType
	if err := json.Unmarshal([]byte(src), &w); err != nil {
		t.Fatalf("unmarshal %s: %v", src, err)
	}
	out, err := json.Marshal(w)
	if err != nil {
		t.Fatalf("marshal %s: %v", src, err)
	}
	if string(out) != src {
		t.Fatalf("round-trip mismatch:\n got %s\nwant %s", out, src)
	}
	return w
}

func TestWidgetQueryType_Data(t *testing.T) {
	w := widgetQueryTypeRoundTrip(t, `{"type":"data","wallet_card":{"address":"EQabc","amount":"1000","stickers":[]}}`)
	if w.Type != WidgetQueryTypeTypeData {
		t.Fatalf("wrong type %q", w.Type)
	}
	if w.Data == nil || w.Data.WalletCard == nil || w.Data.WalletCard.Address != "EQabc" {
		t.Fatalf("wallet card not decoded: %+v", w.Data)
	}
}

func TestWidgetQueryType_DataWithoutWallet(t *testing.T) {
	w := widgetQueryTypeRoundTrip(t, `{"type":"data"}`)
	if w.Data == nil || w.Data.WalletCard != nil {
		t.Fatalf("expected data variant without wallet card: %+v", w.Data)
	}
}

func TestWidgetQueryType_SignData(t *testing.T) {
	w := widgetQueryTypeRoundTrip(t, `{"type":"sign_data","wallet":{"address":"EQabc","amount":"5","public_key":"aa","state_init":"bb"},"sign_data":"{\"x\":1}"}`)
	if w.SignData == nil || w.SignData.Wallet.PublicKey != "aa" || w.SignData.SignData != `{"x":1}` {
		t.Fatalf("sign_data not decoded: %+v", w.SignData)
	}
}

func TestWidgetQueryType_SentTransaction(t *testing.T) {
	w := widgetQueryTypeRoundTrip(t, `{"type":"sent_transaction","wallet":{"address":"EQabc","amount":"5"},"msg_hash":"deadbeef"}`)
	if w.SentTransaction == nil || w.SentTransaction.MsgHash != "deadbeef" {
		t.Fatalf("sent_transaction not decoded: %+v", w.SentTransaction)
	}
}

// An unknown discriminator from a later Bot API release keeps its value in Type,
// leaves every variant nil and still encodes, so a batch of updates is not lost.
func TestWidgetQueryType_Unknown(t *testing.T) {
	w := widgetQueryTypeRoundTrip(t, `{"type":"future"}`)
	if w.Type != "future" || w.Data != nil || w.SignData != nil || w.SentTransaction != nil {
		t.Fatalf("unknown variant not tolerated: %+v", w)
	}
}

func TestWidgetQueryType_MissingDiscriminator(t *testing.T) {
	var w WidgetQueryType
	if err := json.Unmarshal([]byte(`{}`), &w); err == nil {
		t.Fatal("expected an error for a query type without a discriminator")
	}
	if _, err := json.Marshal(WidgetQueryType{Type: WidgetQueryTypeTypeData}); err == nil {
		t.Fatal("expected an error for a set Type without its variant")
	}
}

func TestUpdate_WidgetQuery(t *testing.T) {
	src := `{"update_id":1,"widget_query":{"id":"q1","from":{"id":5,"is_bot":false,"first_name":"A"},"chat":{"id":7,"type":"private"},"message_id":42,"chat_instance":"ci","type":{"type":"data"},"data":"clicked","widget_id":"w1"}}`
	var u Update
	if err := json.Unmarshal([]byte(src), &u); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	q := u.WidgetQuery
	if q == nil || q.ID != "q1" || q.From.ID != 5 || q.Chat == nil || q.Chat.ID != 7 || q.MessageID != 42 {
		t.Fatalf("widget_query not decoded: %+v", q)
	}
	if q.Type.Type != WidgetQueryTypeTypeData || q.Data != "clicked" || q.WidgetID != "w1" {
		t.Fatalf("widget_query fields: %+v", q)
	}
}

func TestUpdate_SharedWidgetQuery(t *testing.T) {
	src := `{"update_id":1,"shared_widget_query":{"id":"q2","from":{"id":5,"is_bot":false,"first_name":"A"},"type":{"type":"sent_transaction","wallet":{"address":"EQabc","amount":"1"},"msg_hash":"h"},"data":"d","widget_id":"w1"}}`
	var u Update
	if err := json.Unmarshal([]byte(src), &u); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	q := u.SharedWidgetQuery
	if q == nil || q.ID != "q2" || q.WidgetID != "w1" || q.Type.SentTransaction == nil || q.Type.SentTransaction.MsgHash != "h" {
		t.Fatalf("shared_widget_query not decoded: %+v", q)
	}
}

func TestMessage_Widget(t *testing.T) {
	src := `{"message_id":1,"date":1,"chat":{"id":1,"type":"private"},"widget":{"id":"w1","width":300,"height":120,"html":"<b>hi</b>","public_data":"p","chat_data":"c"}}`
	var m Message
	if err := json.Unmarshal([]byte(src), &m); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	w := m.Widget
	if w == nil || w.ID != "w1" || w.Width != 300 || w.Height != 120 || w.HTML != "<b>hi</b>" || w.PublicData != "p" || w.ChatData != "c" {
		t.Fatalf("widget not decoded: %+v", w)
	}
}

func TestExternalReplyInfo_Widget(t *testing.T) {
	src := `{"origin":{"type":"hidden_user","date":1,"sender_user_name":"x"},"widget":{"html":"<i>w</i>"}}`
	var e ExternalReplyInfo
	if err := json.Unmarshal([]byte(src), &e); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if e.Widget == nil || e.Widget.HTML != "<i>w</i>" {
		t.Fatalf("widget not decoded: %+v", e.Widget)
	}
}

func TestInputWidgetMessageContent_Marshal(t *testing.T) {
	out, err := json.Marshal(InputWidgetMessageContent{HTML: "x", Width: 200, AllowQueries: true, IsShared: true})
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	if want := `{"html":"x","width":200,"allow_queries":true,"is_shared":true}`; string(out) != want {
		t.Fatalf("got %s, want %s", out, want)
	}

	out, err = json.Marshal(InputSharedWidgetMessageContent{WidgetID: "w1", ChatData: "c"})
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	if want := `{"widget_id":"w1","chat_data":"c"}`; string(out) != want {
		t.Fatalf("got %s, want %s", out, want)
	}
}
