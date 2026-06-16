package models

import (
	"encoding/json"
	"testing"
)

// InputRichMessageContent must be usable as InlineQueryResultArticle.InputMessageContent.
func TestInputRichMessageContent_implementsInputMessageContent(t *testing.T) {
	var _ InputMessageContent = InputRichMessageContent{}
	var _ InputMessageContent = &InputRichMessageContent{}

	article := &InlineQueryResultArticle{
		ID:           "share_1",
		Title:        "Title",
		Description:  "Description",
		ThumbnailURL: "https://cdn.example.com/thumb.jpg",
		InputMessageContent: &InputRichMessageContent{
			RichMessage: InputRichMessage{
				Markdown: "# Title\n\nDescription",
			},
		},
	}

	if article.InputMessageContent == nil {
		t.Fatal("InputMessageContent is nil")
	}
}

func TestInputRichMessageContent_marshalJSON(t *testing.T) {
	c := InputRichMessageContent{
		RichMessage: InputRichMessage{
			Markdown:            "# Title\n\nDescription\n\n![](https://cdn.example.com/photo.jpg)",
			IsRTL:               true,
			SkipEntityDetection: true,
		},
	}

	data, err := json.Marshal(c)
	if err != nil {
		t.Fatal(err)
	}

	got := string(data)
	expect := `{"rich_message":{"markdown":"# Title\n\nDescription\n\n![](https://cdn.example.com/photo.jpg)","is_rtl":true,"skip_entity_detection":true}}`
	if got != expect {
		t.Fatalf("unexpected json\n got: %s\nwant: %s", got, expect)
	}
}

func TestInputRichMessage_marshalOmitEmpty(t *testing.T) {
	data, err := json.Marshal(InputRichMessage{Markdown: "**hi**"})
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != `{"markdown":"**hi**"}` {
		t.Fatalf("unexpected json: %s", string(data))
	}
}

// Message.rich_message must deserialize, including typed rich text and rich blocks.
func TestMessage_unmarshalRichMessage(t *testing.T) {
	src := `{
		"message_id": 1,
		"date": 0,
		"chat": {"id": 42},
		"rich_message": {
			"is_rtl": false,
			"blocks": [
				{"type": "heading", "size": 1, "text": "Report"},
				{"type": "paragraph", "text": [
					"intro ",
					{"type": "bold", "text": "bold"},
					{"type": "url", "text": "link", "url": "https://t.me/"}
				]},
				{"type": "divider"},
				{"type": "photo", "photo": [{"file_id": "abc", "file_unique_id": "u", "width": 100, "height": 200}], "caption": {"text": "cap"}}
			]
		}
	}`

	msg := Message{}
	if err := json.Unmarshal([]byte(src), &msg); err != nil {
		t.Fatal(err)
	}

	if msg.RichMessage == nil {
		t.Fatal("RichMessage is nil")
	}
	if len(msg.RichMessage.Blocks) != 4 {
		t.Fatalf("expected 4 blocks, got %d", len(msg.RichMessage.Blocks))
	}

	heading := msg.RichMessage.Blocks[0]
	if heading.Type != RichBlockTypeHeading || heading.SectionHeading == nil {
		t.Fatal("first block is not a heading")
	}
	if heading.SectionHeading.Size != 1 {
		t.Fatalf("unexpected heading size: %d", heading.SectionHeading.Size)
	}
	if heading.SectionHeading.Text.Plain == nil || *heading.SectionHeading.Text.Plain != "Report" {
		t.Fatal("unexpected heading text")
	}

	paragraph := msg.RichMessage.Blocks[1]
	if paragraph.Type != RichBlockTypeParagraph || paragraph.Paragraph == nil {
		t.Fatal("second block is not a paragraph")
	}
	arr := paragraph.Paragraph.Text.Array
	if len(arr) != 3 {
		t.Fatalf("expected 3 rich text parts, got %d", len(arr))
	}
	if arr[1].Type != RichTextTypeBold || arr[1].Bold == nil {
		t.Fatal("second part is not bold")
	}
	if arr[2].Type != RichTextTypeURL || arr[2].URL == nil || arr[2].URL.URL != "https://t.me/" {
		t.Fatal("third part is not a url")
	}

	if msg.RichMessage.Blocks[2].Type != RichBlockTypeDivider {
		t.Fatal("third block is not a divider")
	}

	photo := msg.RichMessage.Blocks[3]
	if photo.Type != RichBlockTypePhoto || photo.Photo == nil {
		t.Fatal("fourth block is not a photo")
	}
	if len(photo.Photo.Photo) != 1 || photo.Photo.Photo[0].FileID != "abc" {
		t.Fatal("unexpected photo content")
	}
	if photo.Photo.Caption == nil || photo.Photo.Caption.Text.Plain == nil || *photo.Photo.Caption.Text.Plain != "cap" {
		t.Fatal("unexpected photo caption")
	}
}

func TestRichBlock_roundTrip(t *testing.T) {
	src := `{"type":"paragraph","text":[{"type":"bold","text":"x"},{"type":"italic","text":"y"}]}`

	block := RichBlock{}
	if err := json.Unmarshal([]byte(src), &block); err != nil {
		t.Fatal(err)
	}

	data, err := json.Marshal(block)
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != src {
		t.Fatalf("round trip mismatch\n got: %s\nwant: %s", string(data), src)
	}
}

// Existing input message content types must keep working.
func TestInputTextMessageContent_stillWorks(t *testing.T) {
	var _ InputMessageContent = &InputTextMessageContent{}

	data, err := json.Marshal(&InputTextMessageContent{MessageText: "hello"})
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != `{"message_text":"hello"}` {
		t.Fatalf("unexpected json: %s", string(data))
	}
}
