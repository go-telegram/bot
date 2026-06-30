package models

import (
	"encoding/json"
	"testing"
)

func richBlockRoundTrip(t *testing.T, src string) RichBlock {
	t.Helper()

	var rb RichBlock
	if err := json.Unmarshal([]byte(src), &rb); err != nil {
		t.Fatalf("unmarshal %s: %v", src, err)
	}

	out, err := json.Marshal(rb)
	if err != nil {
		t.Fatalf("marshal %s: %v", src, err)
	}

	if string(out) != src {
		t.Fatalf("round-trip mismatch:\n got %s\nwant %s", out, src)
	}

	return rb
}

func TestRichBlock_Thinking(t *testing.T) {
	rb := richBlockRoundTrip(t, `{"type":"thinking","text":"weighing the options"}`)
	if rb.Type != RichBlockTypeThinking {
		t.Fatalf("wrong type %q", rb.Type)
	}
	if rb.RichBlockThinking == nil || rb.RichBlockThinking.Text.PlainText != "weighing the options" {
		t.Fatal("thinking text not decoded")
	}
}

func TestRichBlock_ParagraphWithFormatting(t *testing.T) {
	richBlockRoundTrip(t, `{"type":"paragraph","text":{"type":"bold","text":"hi"}}`)
}

func TestRichBlock_Divider(t *testing.T) {
	richBlockRoundTrip(t, `{"type":"divider"}`)
}

func TestRichBlock_NestedList(t *testing.T) {
	rb := richBlockRoundTrip(t, `{"type":"list","items":[{"label":"1","blocks":[{"type":"paragraph","text":"first"}]}]}`)
	if rb.RichBlockList == nil || len(rb.RichBlockList.Items) != 1 {
		t.Fatal("list items not decoded")
	}
	if rb.RichBlockList.Items[0].Blocks[0].Type != RichBlockTypeParagraph {
		t.Fatal("nested paragraph block not decoded")
	}
}

func TestRichBlock_Table(t *testing.T) {
	richBlockRoundTrip(t, `{"type":"table","cells":[[{"align":"left","valign":"top"},{"text":"h","align":"center","valign":"middle"}]]}`)
}

func TestRichBlock_BlockQuotationWithBlocks(t *testing.T) {
	richBlockRoundTrip(t, `{"type":"blockquote","blocks":[{"type":"paragraph","text":"quoted"}]}`)
}

func TestRichBlock_UnknownType(t *testing.T) {
	var rb RichBlock
	if err := json.Unmarshal([]byte(`{"type":"not_a_block"}`), &rb); err == nil {
		t.Fatal("expected error for unknown RichBlock type")
	}
}

func TestRichMessage_RoundTrip(t *testing.T) {
	src := `{"blocks":[{"type":"thinking","text":"hmm"},{"type":"paragraph","text":"answer"}]}`

	var rm RichMessage
	if err := json.Unmarshal([]byte(src), &rm); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if len(rm.Blocks) != 2 {
		t.Fatalf("expected 2 blocks, got %d", len(rm.Blocks))
	}
	if rm.Blocks[0].Type != RichBlockTypeThinking {
		t.Fatal("first block should be thinking")
	}

	out, err := json.Marshal(rm)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	if string(out) != src {
		t.Fatalf("round-trip mismatch:\n got %s\nwant %s", out, src)
	}
}
