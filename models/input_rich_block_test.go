package models

import (
	"encoding/json"
	"strings"
	"testing"
)

// TestInputRichBlock_Marshal verifies each block variant marshals with the
// correct "type" discriminator injected from the union tag.
func TestInputRichBlock_Marshal(t *testing.T) {
	cases := []struct {
		name     string
		block    InputRichBlock
		wantType string
	}{
		{"paragraph", InputRichBlock{Type: RichBlockTypeParagraph, InputRichBlockParagraph: &InputRichBlockParagraph{Text: RichText{PlainText: "hi"}}}, "paragraph"},
		{"heading", InputRichBlock{Type: RichBlockTypeSectionHeading, InputRichBlockSectionHeading: &InputRichBlockSectionHeading{Text: RichText{PlainText: "h"}, Size: 2}}, "heading"},
		{"divider", InputRichBlock{Type: RichBlockTypeDivider, InputRichBlockDivider: &InputRichBlockDivider{}}, "divider"},
		{"math", InputRichBlock{Type: RichBlockTypeMathematicalExpression, InputRichBlockMathematicalExpression: &InputRichBlockMathematicalExpression{Expression: "x^2"}}, "mathematical_expression"},
		{"thinking", InputRichBlock{Type: RichBlockTypeThinking, InputRichBlockThinking: &InputRichBlockThinking{Text: RichText{PlainText: "reasoning"}}}, "thinking"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out, err := json.Marshal(c.block)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			if !strings.Contains(string(out), `"type":"`+c.wantType+`"`) {
				t.Fatalf("missing type %q in %s", c.wantType, out)
			}
		})
	}
}

// TestInputRichBlockList_Nested verifies list items carry nested blocks.
func TestInputRichBlockList_Nested(t *testing.T) {
	block := InputRichBlock{
		Type: RichBlockTypeList,
		InputRichBlockList: &InputRichBlockList{
			Items: []InputRichBlockListItem{
				{Blocks: []InputRichBlock{{Type: RichBlockTypeParagraph, InputRichBlockParagraph: &InputRichBlockParagraph{Text: RichText{PlainText: "item"}}}}, HasCheckbox: true, IsChecked: true},
			},
		},
	}
	out, err := json.Marshal(block)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	s := string(out)
	if !strings.Contains(s, `"type":"list"`) || !strings.Contains(s, `"has_checkbox":true`) {
		t.Fatalf("unexpected list encoding: %s", s)
	}
}

// TestInputRichMessage_BlocksAndMedia verifies the 10.2 fields on InputRichMessage.
func TestInputRichMessage_BlocksAndMedia(t *testing.T) {
	in := InputRichMessage{
		Markdown: "![img](tg://photo?id=abc)",
		Media: []InputRichMessageMedia{
			{ID: "abc", Media: &InputMediaPhoto{Media: "file_id_123"}},
		},
		Blocks: []InputRichBlock{
			{Type: RichBlockTypeParagraph, InputRichBlockParagraph: &InputRichBlockParagraph{Text: RichText{PlainText: "p"}}},
		},
	}
	out, err := json.Marshal(in)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	s := string(out)
	if !strings.Contains(s, `"media"`) || !strings.Contains(s, `"blocks"`) {
		t.Fatalf("expected media and blocks in %s", s)
	}
}

// TestInputMediaVoiceNote verifies the new voice note input media marshals with type.
func TestInputMediaVoiceNote(t *testing.T) {
	m := InputMediaVoiceNote{Media: "attach://voice", Duration: 5}
	out, err := m.MarshalInputMedia()
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	if !strings.Contains(string(out), `"type":"voice_note"`) {
		t.Fatalf("expected voice_note type: %s", out)
	}
}

// TestInputMedia_NestedTypeDiscriminator verifies InputMedia values keep their
// required "type" field when reached through a plain json.Marshal, which is how
// they are encoded once nested inside a rich message.
func TestInputMedia_NestedTypeDiscriminator(t *testing.T) {
	cases := []struct {
		name  string
		media any
		want  string
	}{
		{"photo", InputMediaPhoto{Media: "id"}, `"type":"photo"`},
		{"video", InputMediaVideo{Media: "id"}, `"type":"video"`},
		{"animation", InputMediaAnimation{Media: "id"}, `"type":"animation"`},
		{"audio", InputMediaAudio{Media: "id"}, `"type":"audio"`},
		{"voice_note", InputMediaVoiceNote{Media: "id"}, `"type":"voice_note"`},
		{"document", InputMediaDocument{Media: "id"}, `"type":"document"`},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out, err := json.Marshal(c.media)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			if !strings.Contains(string(out), c.want) {
				t.Fatalf("expected %s in %s", c.want, out)
			}
		})
	}
}

// TestInputRichMessage_NestedMediaType verifies the nested InputMedia inside a
// rich message carries its type through the json.Marshal path used when sending.
func TestInputRichMessage_NestedMediaType(t *testing.T) {
	in := InputRichMessage{
		Markdown: "![img](tg://photo?id=abc)",
		Media:    []InputRichMessageMedia{{ID: "abc", Media: &InputMediaPhoto{Media: "file_id_123"}}},
		Blocks: []InputRichBlock{
			{Type: RichBlockTypePhoto, InputRichBlockPhoto: &InputRichBlockPhoto{Photo: InputMediaPhoto{Media: "p"}}},
			{Type: RichBlockTypeAnimation, InputRichBlockAnimation: &InputRichBlockAnimation{Animation: InputMediaAnimation{Media: "a"}}},
			{Type: RichBlockTypeVoiceNote, InputRichBlockVoiceNote: &InputRichBlockVoiceNote{VoiceNote: InputMediaVoiceNote{Media: "v"}}},
		},
	}
	out, err := json.Marshal(in)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	s := string(out)
	for _, want := range []string{
		`"media":{"type":"photo"`,
		`"photo":{"type":"photo"`,
		`"animation":{"type":"animation"`,
		`"voice_note":{"type":"voice_note"`,
	} {
		if !strings.Contains(s, want) {
			t.Fatalf("expected %s in %s", want, s)
		}
	}
}

// TestInputRichBlock_NilVariant verifies a Type set without its matching variant
// pointer returns an error instead of panicking.
func TestInputRichBlock_NilVariant(t *testing.T) {
	if _, err := json.Marshal(InputRichBlock{Type: RichBlockTypeParagraph}); err == nil {
		t.Fatal("expected error for nil variant pointer")
	}
	if _, err := json.Marshal(InputRichBlock{}); err == nil {
		t.Fatal("expected error for empty InputRichBlock")
	}
}

// TestInputRichBlock_UnknownType verifies an unrecognised discriminator is
// reported as unsupported rather than as a missing variant pointer.
func TestInputRichBlock_UnknownType(t *testing.T) {
	_, err := json.Marshal(InputRichBlock{Type: "definitely_not_a_real_type"})
	if err == nil {
		t.Fatal("expected error for unknown block type")
	}
	if !strings.Contains(err.Error(), `unsupported InputRichBlock type "definitely_not_a_real_type"`) {
		t.Fatalf("unexpected error: %v", err)
	}
}

// TestInputRichBlock_MarshalDoesNotMutate verifies encoding stamps the union
// discriminator on a copy, leaving the caller's variant untouched. Without it
// two goroutines encoding the same block would race on the Type field.
func TestInputRichBlock_MarshalDoesNotMutate(t *testing.T) {
	variant := &InputRichBlockParagraph{Text: RichText{PlainText: "x"}}
	out, err := json.Marshal(InputRichBlock{Type: RichBlockTypeParagraph, InputRichBlockParagraph: variant})
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	if !strings.Contains(string(out), `"type":"paragraph"`) {
		t.Fatalf("missing discriminator in %s", out)
	}
	if variant.Type != "" {
		t.Fatalf("marshal mutated the caller's variant: Type = %q", variant.Type)
	}
}
