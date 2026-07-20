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
