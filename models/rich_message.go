package models

// RichMessage https://core.telegram.org/bots/api#richmessage
//
// Describes a rich formatted message.
type RichMessage struct {
	Blocks []RichBlock `json:"blocks"`
	IsRTL  bool        `json:"is_rtl,omitempty"`
}

// InputRichMessage https://core.telegram.org/bots/api#inputrichmessage
//
// Describes a rich formatted message to be sent. The thinking block is expressed
// by the custom HTML tag <tg-thinking> within HTML.
type InputRichMessage struct {
	HTML                string `json:"html,omitempty"`
	Markdown            string `json:"markdown,omitempty"`
	IsRTL               bool   `json:"is_rtl,omitempty"`
	SkipEntityDetection bool   `json:"skip_entity_detection,omitempty"`
}

// InputRichMessageContent https://core.telegram.org/bots/api#inputrichmessagecontent
//
// Represents the content of a rich message to be sent as the result of an inline
// query.
type InputRichMessageContent struct {
	RichMessage InputRichMessage `json:"rich_message"`
}

// RichBlockCaption https://core.telegram.org/bots/api#richblockcaption
//
// Caption of a rich formatted block.
type RichBlockCaption struct {
	Text   RichText  `json:"text"`
	Credit *RichText `json:"credit,omitempty"`
}

// RichBlockListItem https://core.telegram.org/bots/api#richblocklistitem
//
// An item of a list.
type RichBlockListItem struct {
	Label       string      `json:"label"`
	Blocks      []RichBlock `json:"blocks"`
	HasCheckbox bool        `json:"has_checkbox,omitempty"`
	IsChecked   bool        `json:"is_checked,omitempty"`
	Value       int         `json:"value,omitempty"`
	Type        string      `json:"type,omitempty"`
}

// RichBlockTableCell https://core.telegram.org/bots/api#richblocktablecell
//
// Cell in a table.
type RichBlockTableCell struct {
	Text     *RichText `json:"text,omitempty"`
	IsHeader bool      `json:"is_header,omitempty"`
	Colspan  int       `json:"colspan,omitempty"`
	Rowspan  int       `json:"rowspan,omitempty"`
	Align    string    `json:"align"`
	Valign   string    `json:"valign"`
}
