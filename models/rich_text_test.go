package models

import (
	"encoding/json"
	"testing"
)

// richTextRoundTrip unmarshals src into a RichText then marshals it back and
// asserts the JSON is byte-identical. Our struct field order (type, then text,
// then extras) matches Telegram's encoding, so a string compare is meaningful.
func richTextRoundTrip(t *testing.T, src string) RichText {
	t.Helper()

	var rt RichText
	if err := json.Unmarshal([]byte(src), &rt); err != nil {
		t.Fatalf("unmarshal %s: %v", src, err)
	}

	out, err := json.Marshal(rt)
	if err != nil {
		t.Fatalf("marshal %s: %v", src, err)
	}

	if string(out) != src {
		t.Fatalf("round-trip mismatch:\n got %s\nwant %s", out, src)
	}

	return rt
}

func TestRichText_PlainString(t *testing.T) {
	rt := richTextRoundTrip(t, `"hello world"`)
	if rt.Type != "" {
		t.Fatalf("plain string should have empty Type, got %q", rt.Type)
	}
	if rt.PlainText != "hello world" {
		t.Fatalf("wrong PlainText %q", rt.PlainText)
	}
}

func TestRichText_Array(t *testing.T) {
	rt := richTextRoundTrip(t, `["a","b","c"]`)
	if len(rt.Array) != 3 {
		t.Fatalf("expected 3 array elements, got %d", len(rt.Array))
	}
	if rt.Array[1].PlainText != "b" {
		t.Fatalf("wrong array element %q", rt.Array[1].PlainText)
	}
}

func TestRichText_TaggedBold(t *testing.T) {
	rt := richTextRoundTrip(t, `{"type":"bold","text":"hi"}`)
	if rt.Type != RichTextTypeBold {
		t.Fatalf("wrong type %q", rt.Type)
	}
	if rt.RichTextBold == nil || rt.RichTextBold.Text.PlainText != "hi" {
		t.Fatal("bold text not decoded")
	}
}

func TestRichText_Recursive(t *testing.T) {
	// bold wrapping italic wrapping a plain string — exercises the recursive
	// Text RichText field through both directions.
	rt := richTextRoundTrip(t, `{"type":"bold","text":{"type":"italic","text":"x"}}`)
	if rt.RichTextBold == nil {
		t.Fatal("outer bold nil")
	}
	inner := rt.RichTextBold.Text
	if inner.Type != RichTextTypeItalic || inner.RichTextItalic == nil {
		t.Fatal("inner italic not decoded")
	}
	if inner.RichTextItalic.Text.PlainText != "x" {
		t.Fatal("innermost plain text wrong")
	}
}

func TestRichText_ArrayOfMixed(t *testing.T) {
	// a sequence mixing a plain string and a tagged object, like a paragraph body.
	rt := richTextRoundTrip(t, `[{"type":"bold","text":"a"},"b"]`)
	if len(rt.Array) != 2 {
		t.Fatalf("expected 2 elements, got %d", len(rt.Array))
	}
	if rt.Array[0].Type != RichTextTypeBold {
		t.Fatal("first element should be bold")
	}
	if rt.Array[1].PlainText != "b" {
		t.Fatal("second element should be plain 'b'")
	}
}

func TestRichText_VariantsWithExtraFields(t *testing.T) {
	for _, src := range []string{
		`{"type":"url","text":"link","url":"https://example.com"}`,
		`{"type":"custom_emoji","custom_emoji_id":"123","alternative_text":"🙂"}`,
		`{"type":"mathematical_expression","expression":"x^2"}`,
		`{"type":"anchor","name":"top"}`,
		`{"type":"reference_link","text":"see","reference_name":"r1"}`,
	} {
		richTextRoundTrip(t, src)
	}
}

func TestRichText_Null(t *testing.T) {
	var rt RichText
	if err := rt.UnmarshalJSON([]byte("null")); err != nil {
		t.Fatalf("null should not error: %v", err)
	}
	if rt.Type != "" || rt.PlainText != "" || rt.Array != nil {
		t.Fatal("null should leave RichText zero")
	}
}

func TestRichText_UnknownType(t *testing.T) {
	var rt RichText
	err := rt.UnmarshalJSON([]byte(`{"type":"definitely_not_a_real_type"}`))
	if err == nil {
		t.Fatal("expected error for unknown RichText type")
	}
}
