package models

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// InputRichMessage https://core.telegram.org/bots/api#inputrichmessage
//
// Describes a rich message to be sent. Exactly one of the fields HTML or
// Markdown must be used.
type InputRichMessage struct {
	HTML                string `json:"html,omitempty"`
	Markdown            string `json:"markdown,omitempty"`
	IsRTL               bool   `json:"is_rtl,omitempty"`
	SkipEntityDetection bool   `json:"skip_entity_detection,omitempty"`
}

// RichMessage https://core.telegram.org/bots/api#richmessage
//
// Rich formatted message.
type RichMessage struct {
	Blocks []RichBlock `json:"blocks"`
	IsRTL  bool        `json:"is_rtl,omitempty"`
}

// RichTextType https://core.telegram.org/bots/api#richtext
type RichTextType string

const (
	RichTextTypeBold                   RichTextType = "bold"
	RichTextTypeItalic                 RichTextType = "italic"
	RichTextTypeUnderline              RichTextType = "underline"
	RichTextTypeStrikethrough          RichTextType = "strikethrough"
	RichTextTypeSpoiler                RichTextType = "spoiler"
	RichTextTypeDateTime               RichTextType = "date_time"
	RichTextTypeTextMention            RichTextType = "text_mention"
	RichTextTypeSubscript              RichTextType = "subscript"
	RichTextTypeSuperscript            RichTextType = "superscript"
	RichTextTypeMarked                 RichTextType = "marked"
	RichTextTypeCode                   RichTextType = "code"
	RichTextTypeCustomEmoji            RichTextType = "custom_emoji"
	RichTextTypeMathematicalExpression RichTextType = "mathematical_expression"
	RichTextTypeURL                    RichTextType = "url"
	RichTextTypeEmailAddress           RichTextType = "email_address"
	RichTextTypePhoneNumber            RichTextType = "phone_number"
	RichTextTypeBankCardNumber         RichTextType = "bank_card_number"
	RichTextTypeMention                RichTextType = "mention"
	RichTextTypeHashtag                RichTextType = "hashtag"
	RichTextTypeCashtag                RichTextType = "cashtag"
	RichTextTypeBotCommand             RichTextType = "bot_command"
	RichTextTypeAnchor                 RichTextType = "anchor"
	RichTextTypeAnchorLink             RichTextType = "anchor_link"
	RichTextTypeReference              RichTextType = "reference"
	RichTextTypeReferenceLink          RichTextType = "reference_link"
)

// RichText https://core.telegram.org/bots/api#richtext
//
// Represents a rich formatted text. It can be a plain string (Plain), an array
// of RichText (Array), or one of the typed entities below.
type RichText struct {
	// Plain is set when the rich text is a plain string.
	Plain *string
	// Array is set when the rich text is an array of RichText.
	Array []RichText
	// Type is the entity type when the rich text is a typed object.
	Type RichTextType

	Bold                   *RichTextBold
	Italic                 *RichTextItalic
	Underline              *RichTextUnderline
	Strikethrough          *RichTextStrikethrough
	Spoiler                *RichTextSpoiler
	DateTime               *RichTextDateTime
	TextMention            *RichTextTextMention
	Subscript              *RichTextSubscript
	Superscript            *RichTextSuperscript
	Marked                 *RichTextMarked
	Code                   *RichTextCode
	CustomEmoji            *RichTextCustomEmoji
	MathematicalExpression *RichTextMathematicalExpression
	URL                    *RichTextURL
	EmailAddress           *RichTextEmailAddress
	PhoneNumber            *RichTextPhoneNumber
	BankCardNumber         *RichTextBankCardNumber
	Mention                *RichTextMention
	Hashtag                *RichTextHashtag
	Cashtag                *RichTextCashtag
	BotCommand             *RichTextBotCommand
	Anchor                 *RichTextAnchor
	AnchorLink             *RichTextAnchorLink
	Reference              *RichTextReference
	ReferenceLink          *RichTextReferenceLink
}

func (rt RichText) MarshalJSON() ([]byte, error) {
	if rt.Plain != nil {
		return json.Marshal(*rt.Plain)
	}
	if rt.Array != nil {
		return json.Marshal(rt.Array)
	}

	switch rt.Type {
	case RichTextTypeBold:
		rt.Bold.Type = RichTextTypeBold
		return json.Marshal(rt.Bold)
	case RichTextTypeItalic:
		rt.Italic.Type = RichTextTypeItalic
		return json.Marshal(rt.Italic)
	case RichTextTypeUnderline:
		rt.Underline.Type = RichTextTypeUnderline
		return json.Marshal(rt.Underline)
	case RichTextTypeStrikethrough:
		rt.Strikethrough.Type = RichTextTypeStrikethrough
		return json.Marshal(rt.Strikethrough)
	case RichTextTypeSpoiler:
		rt.Spoiler.Type = RichTextTypeSpoiler
		return json.Marshal(rt.Spoiler)
	case RichTextTypeDateTime:
		rt.DateTime.Type = RichTextTypeDateTime
		return json.Marshal(rt.DateTime)
	case RichTextTypeTextMention:
		rt.TextMention.Type = RichTextTypeTextMention
		return json.Marshal(rt.TextMention)
	case RichTextTypeSubscript:
		rt.Subscript.Type = RichTextTypeSubscript
		return json.Marshal(rt.Subscript)
	case RichTextTypeSuperscript:
		rt.Superscript.Type = RichTextTypeSuperscript
		return json.Marshal(rt.Superscript)
	case RichTextTypeMarked:
		rt.Marked.Type = RichTextTypeMarked
		return json.Marshal(rt.Marked)
	case RichTextTypeCode:
		rt.Code.Type = RichTextTypeCode
		return json.Marshal(rt.Code)
	case RichTextTypeCustomEmoji:
		rt.CustomEmoji.Type = RichTextTypeCustomEmoji
		return json.Marshal(rt.CustomEmoji)
	case RichTextTypeMathematicalExpression:
		rt.MathematicalExpression.Type = RichTextTypeMathematicalExpression
		return json.Marshal(rt.MathematicalExpression)
	case RichTextTypeURL:
		rt.URL.Type = RichTextTypeURL
		return json.Marshal(rt.URL)
	case RichTextTypeEmailAddress:
		rt.EmailAddress.Type = RichTextTypeEmailAddress
		return json.Marshal(rt.EmailAddress)
	case RichTextTypePhoneNumber:
		rt.PhoneNumber.Type = RichTextTypePhoneNumber
		return json.Marshal(rt.PhoneNumber)
	case RichTextTypeBankCardNumber:
		rt.BankCardNumber.Type = RichTextTypeBankCardNumber
		return json.Marshal(rt.BankCardNumber)
	case RichTextTypeMention:
		rt.Mention.Type = RichTextTypeMention
		return json.Marshal(rt.Mention)
	case RichTextTypeHashtag:
		rt.Hashtag.Type = RichTextTypeHashtag
		return json.Marshal(rt.Hashtag)
	case RichTextTypeCashtag:
		rt.Cashtag.Type = RichTextTypeCashtag
		return json.Marshal(rt.Cashtag)
	case RichTextTypeBotCommand:
		rt.BotCommand.Type = RichTextTypeBotCommand
		return json.Marshal(rt.BotCommand)
	case RichTextTypeAnchor:
		rt.Anchor.Type = RichTextTypeAnchor
		return json.Marshal(rt.Anchor)
	case RichTextTypeAnchorLink:
		rt.AnchorLink.Type = RichTextTypeAnchorLink
		return json.Marshal(rt.AnchorLink)
	case RichTextTypeReference:
		rt.Reference.Type = RichTextTypeReference
		return json.Marshal(rt.Reference)
	case RichTextTypeReferenceLink:
		rt.ReferenceLink.Type = RichTextTypeReferenceLink
		return json.Marshal(rt.ReferenceLink)
	}

	return []byte("null"), nil
}

func (rt *RichText) UnmarshalJSON(data []byte) error {
	trimmed := bytes.TrimSpace(data)
	if len(trimmed) == 0 || string(trimmed) == "null" {
		return nil
	}

	switch trimmed[0] {
	case '"':
		var s string
		if err := json.Unmarshal(data, &s); err != nil {
			return err
		}
		rt.Plain = &s
		return nil
	case '[':
		var arr []RichText
		if err := json.Unmarshal(data, &arr); err != nil {
			return err
		}
		rt.Array = arr
		return nil
	}

	v := struct {
		Type RichTextType `json:"type"`
	}{}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	rt.Type = v.Type

	switch v.Type {
	case RichTextTypeBold:
		rt.Bold = &RichTextBold{}
		return json.Unmarshal(data, rt.Bold)
	case RichTextTypeItalic:
		rt.Italic = &RichTextItalic{}
		return json.Unmarshal(data, rt.Italic)
	case RichTextTypeUnderline:
		rt.Underline = &RichTextUnderline{}
		return json.Unmarshal(data, rt.Underline)
	case RichTextTypeStrikethrough:
		rt.Strikethrough = &RichTextStrikethrough{}
		return json.Unmarshal(data, rt.Strikethrough)
	case RichTextTypeSpoiler:
		rt.Spoiler = &RichTextSpoiler{}
		return json.Unmarshal(data, rt.Spoiler)
	case RichTextTypeDateTime:
		rt.DateTime = &RichTextDateTime{}
		return json.Unmarshal(data, rt.DateTime)
	case RichTextTypeTextMention:
		rt.TextMention = &RichTextTextMention{}
		return json.Unmarshal(data, rt.TextMention)
	case RichTextTypeSubscript:
		rt.Subscript = &RichTextSubscript{}
		return json.Unmarshal(data, rt.Subscript)
	case RichTextTypeSuperscript:
		rt.Superscript = &RichTextSuperscript{}
		return json.Unmarshal(data, rt.Superscript)
	case RichTextTypeMarked:
		rt.Marked = &RichTextMarked{}
		return json.Unmarshal(data, rt.Marked)
	case RichTextTypeCode:
		rt.Code = &RichTextCode{}
		return json.Unmarshal(data, rt.Code)
	case RichTextTypeCustomEmoji:
		rt.CustomEmoji = &RichTextCustomEmoji{}
		return json.Unmarshal(data, rt.CustomEmoji)
	case RichTextTypeMathematicalExpression:
		rt.MathematicalExpression = &RichTextMathematicalExpression{}
		return json.Unmarshal(data, rt.MathematicalExpression)
	case RichTextTypeURL:
		rt.URL = &RichTextURL{}
		return json.Unmarshal(data, rt.URL)
	case RichTextTypeEmailAddress:
		rt.EmailAddress = &RichTextEmailAddress{}
		return json.Unmarshal(data, rt.EmailAddress)
	case RichTextTypePhoneNumber:
		rt.PhoneNumber = &RichTextPhoneNumber{}
		return json.Unmarshal(data, rt.PhoneNumber)
	case RichTextTypeBankCardNumber:
		rt.BankCardNumber = &RichTextBankCardNumber{}
		return json.Unmarshal(data, rt.BankCardNumber)
	case RichTextTypeMention:
		rt.Mention = &RichTextMention{}
		return json.Unmarshal(data, rt.Mention)
	case RichTextTypeHashtag:
		rt.Hashtag = &RichTextHashtag{}
		return json.Unmarshal(data, rt.Hashtag)
	case RichTextTypeCashtag:
		rt.Cashtag = &RichTextCashtag{}
		return json.Unmarshal(data, rt.Cashtag)
	case RichTextTypeBotCommand:
		rt.BotCommand = &RichTextBotCommand{}
		return json.Unmarshal(data, rt.BotCommand)
	case RichTextTypeAnchor:
		rt.Anchor = &RichTextAnchor{}
		return json.Unmarshal(data, rt.Anchor)
	case RichTextTypeAnchorLink:
		rt.AnchorLink = &RichTextAnchorLink{}
		return json.Unmarshal(data, rt.AnchorLink)
	case RichTextTypeReference:
		rt.Reference = &RichTextReference{}
		return json.Unmarshal(data, rt.Reference)
	case RichTextTypeReferenceLink:
		rt.ReferenceLink = &RichTextReferenceLink{}
		return json.Unmarshal(data, rt.ReferenceLink)
	}

	return fmt.Errorf("unsupported RichText type %q", v.Type)
}

// RichTextBold https://core.telegram.org/bots/api#richtextbold
type RichTextBold struct {
	Type RichTextType `json:"type"`
	Text RichText     `json:"text"`
}

// RichTextItalic https://core.telegram.org/bots/api#richtextitalic
type RichTextItalic struct {
	Type RichTextType `json:"type"`
	Text RichText     `json:"text"`
}

// RichTextUnderline https://core.telegram.org/bots/api#richtextunderline
type RichTextUnderline struct {
	Type RichTextType `json:"type"`
	Text RichText     `json:"text"`
}

// RichTextStrikethrough https://core.telegram.org/bots/api#richtextstrikethrough
type RichTextStrikethrough struct {
	Type RichTextType `json:"type"`
	Text RichText     `json:"text"`
}

// RichTextSpoiler https://core.telegram.org/bots/api#richtextspoiler
type RichTextSpoiler struct {
	Type RichTextType `json:"type"`
	Text RichText     `json:"text"`
}

// RichTextDateTime https://core.telegram.org/bots/api#richtextdatetime
type RichTextDateTime struct {
	Type           RichTextType `json:"type"`
	Text           RichText     `json:"text"`
	UnixTime       int64        `json:"unix_time"`
	DateTimeFormat string       `json:"date_time_format"`
}

// RichTextTextMention https://core.telegram.org/bots/api#richtexttextmention
type RichTextTextMention struct {
	Type RichTextType `json:"type"`
	Text RichText     `json:"text"`
	User User         `json:"user"`
}

// RichTextSubscript https://core.telegram.org/bots/api#richtextsubscript
type RichTextSubscript struct {
	Type RichTextType `json:"type"`
	Text RichText     `json:"text"`
}

// RichTextSuperscript https://core.telegram.org/bots/api#richtextsuperscript
type RichTextSuperscript struct {
	Type RichTextType `json:"type"`
	Text RichText     `json:"text"`
}

// RichTextMarked https://core.telegram.org/bots/api#richtextmarked
type RichTextMarked struct {
	Type RichTextType `json:"type"`
	Text RichText     `json:"text"`
}

// RichTextCode https://core.telegram.org/bots/api#richtextcode
type RichTextCode struct {
	Type RichTextType `json:"type"`
	Text RichText     `json:"text"`
}

// RichTextCustomEmoji https://core.telegram.org/bots/api#richtextcustomemoji
type RichTextCustomEmoji struct {
	Type            RichTextType `json:"type"`
	CustomEmojiID   string       `json:"custom_emoji_id"`
	AlternativeText string       `json:"alternative_text"`
}

// RichTextMathematicalExpression https://core.telegram.org/bots/api#richtextmathematicalexpression
type RichTextMathematicalExpression struct {
	Type       RichTextType `json:"type"`
	Expression string       `json:"expression"`
}

// RichTextURL https://core.telegram.org/bots/api#richtexturl
type RichTextURL struct {
	Type RichTextType `json:"type"`
	Text RichText     `json:"text"`
	URL  string       `json:"url"`
}

// RichTextEmailAddress https://core.telegram.org/bots/api#richtextemailaddress
type RichTextEmailAddress struct {
	Type         RichTextType `json:"type"`
	Text         RichText     `json:"text"`
	EmailAddress string       `json:"email_address"`
}

// RichTextPhoneNumber https://core.telegram.org/bots/api#richtextphonenumber
type RichTextPhoneNumber struct {
	Type        RichTextType `json:"type"`
	Text        RichText     `json:"text"`
	PhoneNumber string       `json:"phone_number"`
}

// RichTextBankCardNumber https://core.telegram.org/bots/api#richtextbankcardnumber
type RichTextBankCardNumber struct {
	Type           RichTextType `json:"type"`
	Text           RichText     `json:"text"`
	BankCardNumber string       `json:"bank_card_number"`
}

// RichTextMention https://core.telegram.org/bots/api#richtextmention
type RichTextMention struct {
	Type     RichTextType `json:"type"`
	Text     RichText     `json:"text"`
	Username string       `json:"username"`
}

// RichTextHashtag https://core.telegram.org/bots/api#richtexthashtag
type RichTextHashtag struct {
	Type    RichTextType `json:"type"`
	Text    RichText     `json:"text"`
	Hashtag string       `json:"hashtag"`
}

// RichTextCashtag https://core.telegram.org/bots/api#richtextcashtag
type RichTextCashtag struct {
	Type    RichTextType `json:"type"`
	Text    RichText     `json:"text"`
	Cashtag string       `json:"cashtag"`
}

// RichTextBotCommand https://core.telegram.org/bots/api#richtextbotcommand
type RichTextBotCommand struct {
	Type       RichTextType `json:"type"`
	Text       RichText     `json:"text"`
	BotCommand string       `json:"bot_command"`
}

// RichTextAnchor https://core.telegram.org/bots/api#richtextanchor
type RichTextAnchor struct {
	Type RichTextType `json:"type"`
	Name string       `json:"name"`
}

// RichTextAnchorLink https://core.telegram.org/bots/api#richtextanchorlink
type RichTextAnchorLink struct {
	Type       RichTextType `json:"type"`
	Text       RichText     `json:"text"`
	AnchorName string       `json:"anchor_name"`
}

// RichTextReference https://core.telegram.org/bots/api#richtextreference
type RichTextReference struct {
	Type RichTextType `json:"type"`
	Text RichText     `json:"text"`
	Name string       `json:"name"`
}

// RichTextReferenceLink https://core.telegram.org/bots/api#richtextreferencelink
type RichTextReferenceLink struct {
	Type          RichTextType `json:"type"`
	Text          RichText     `json:"text"`
	ReferenceName string       `json:"reference_name"`
}

// RichBlockCaption https://core.telegram.org/bots/api#richblockcaption
type RichBlockCaption struct {
	Text   RichText  `json:"text"`
	Credit *RichText `json:"credit,omitempty"`
}

// RichBlockTableCell https://core.telegram.org/bots/api#richblocktablecell
type RichBlockTableCell struct {
	Text     *RichText `json:"text,omitempty"`
	IsHeader bool      `json:"is_header,omitempty"`
	Colspan  int       `json:"colspan,omitempty"`
	Rowspan  int       `json:"rowspan,omitempty"`
	Align    string    `json:"align"`
	VAlign   string    `json:"valign"`
}

// RichBlockListItem https://core.telegram.org/bots/api#richblocklistitem
type RichBlockListItem struct {
	Label       string      `json:"label"`
	Blocks      []RichBlock `json:"blocks"`
	HasCheckbox bool        `json:"has_checkbox,omitempty"`
	IsChecked   bool        `json:"is_checked,omitempty"`
	Value       int         `json:"value,omitempty"`
	Type        string      `json:"type,omitempty"`
}

// RichBlockType https://core.telegram.org/bots/api#richblock
type RichBlockType string

const (
	RichBlockTypeParagraph              RichBlockType = "paragraph"
	RichBlockTypeHeading                RichBlockType = "heading"
	RichBlockTypePre                    RichBlockType = "pre"
	RichBlockTypeFooter                 RichBlockType = "footer"
	RichBlockTypeDivider                RichBlockType = "divider"
	RichBlockTypeMathematicalExpression RichBlockType = "mathematical_expression"
	RichBlockTypeAnchor                 RichBlockType = "anchor"
	RichBlockTypeList                   RichBlockType = "list"
	RichBlockTypeBlockquote             RichBlockType = "blockquote"
	RichBlockTypePullquote              RichBlockType = "pullquote"
	RichBlockTypeCollage                RichBlockType = "collage"
	RichBlockTypeSlideshow              RichBlockType = "slideshow"
	RichBlockTypeTable                  RichBlockType = "table"
	RichBlockTypeDetails                RichBlockType = "details"
	RichBlockTypeMap                    RichBlockType = "map"
	RichBlockTypeAnimation              RichBlockType = "animation"
	RichBlockTypeAudio                  RichBlockType = "audio"
	RichBlockTypePhoto                  RichBlockType = "photo"
	RichBlockTypeVideo                  RichBlockType = "video"
	RichBlockTypeVoiceNote              RichBlockType = "voice_note"
	RichBlockTypeThinking               RichBlockType = "thinking"
)

// RichBlock https://core.telegram.org/bots/api#richblock
//
// Represents a block in a rich formatted message.
type RichBlock struct {
	Type RichBlockType

	Paragraph              *RichBlockParagraph
	SectionHeading         *RichBlockSectionHeading
	Preformatted           *RichBlockPreformatted
	Footer                 *RichBlockFooter
	Divider                *RichBlockDivider
	MathematicalExpression *RichBlockMathematicalExpression
	Anchor                 *RichBlockAnchor
	List                   *RichBlockList
	BlockQuotation         *RichBlockBlockQuotation
	PullQuotation          *RichBlockPullQuotation
	Collage                *RichBlockCollage
	Slideshow              *RichBlockSlideshow
	Table                  *RichBlockTable
	Details                *RichBlockDetails
	Map                    *RichBlockMap
	Animation              *RichBlockAnimation
	Audio                  *RichBlockAudio
	Photo                  *RichBlockPhoto
	Video                  *RichBlockVideo
	VoiceNote              *RichBlockVoiceNote
	Thinking               *RichBlockThinking
}

func (b RichBlock) MarshalJSON() ([]byte, error) {
	switch b.Type {
	case RichBlockTypeParagraph:
		b.Paragraph.Type = RichBlockTypeParagraph
		return json.Marshal(b.Paragraph)
	case RichBlockTypeHeading:
		b.SectionHeading.Type = RichBlockTypeHeading
		return json.Marshal(b.SectionHeading)
	case RichBlockTypePre:
		b.Preformatted.Type = RichBlockTypePre
		return json.Marshal(b.Preformatted)
	case RichBlockTypeFooter:
		b.Footer.Type = RichBlockTypeFooter
		return json.Marshal(b.Footer)
	case RichBlockTypeDivider:
		b.Divider.Type = RichBlockTypeDivider
		return json.Marshal(b.Divider)
	case RichBlockTypeMathematicalExpression:
		b.MathematicalExpression.Type = RichBlockTypeMathematicalExpression
		return json.Marshal(b.MathematicalExpression)
	case RichBlockTypeAnchor:
		b.Anchor.Type = RichBlockTypeAnchor
		return json.Marshal(b.Anchor)
	case RichBlockTypeList:
		b.List.Type = RichBlockTypeList
		return json.Marshal(b.List)
	case RichBlockTypeBlockquote:
		b.BlockQuotation.Type = RichBlockTypeBlockquote
		return json.Marshal(b.BlockQuotation)
	case RichBlockTypePullquote:
		b.PullQuotation.Type = RichBlockTypePullquote
		return json.Marshal(b.PullQuotation)
	case RichBlockTypeCollage:
		b.Collage.Type = RichBlockTypeCollage
		return json.Marshal(b.Collage)
	case RichBlockTypeSlideshow:
		b.Slideshow.Type = RichBlockTypeSlideshow
		return json.Marshal(b.Slideshow)
	case RichBlockTypeTable:
		b.Table.Type = RichBlockTypeTable
		return json.Marshal(b.Table)
	case RichBlockTypeDetails:
		b.Details.Type = RichBlockTypeDetails
		return json.Marshal(b.Details)
	case RichBlockTypeMap:
		b.Map.Type = RichBlockTypeMap
		return json.Marshal(b.Map)
	case RichBlockTypeAnimation:
		b.Animation.Type = RichBlockTypeAnimation
		return json.Marshal(b.Animation)
	case RichBlockTypeAudio:
		b.Audio.Type = RichBlockTypeAudio
		return json.Marshal(b.Audio)
	case RichBlockTypePhoto:
		b.Photo.Type = RichBlockTypePhoto
		return json.Marshal(b.Photo)
	case RichBlockTypeVideo:
		b.Video.Type = RichBlockTypeVideo
		return json.Marshal(b.Video)
	case RichBlockTypeVoiceNote:
		b.VoiceNote.Type = RichBlockTypeVoiceNote
		return json.Marshal(b.VoiceNote)
	case RichBlockTypeThinking:
		b.Thinking.Type = RichBlockTypeThinking
		return json.Marshal(b.Thinking)
	}

	return nil, fmt.Errorf("unsupported RichBlock type %q", b.Type)
}

func (b *RichBlock) UnmarshalJSON(data []byte) error {
	v := struct {
		Type RichBlockType `json:"type"`
	}{}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	b.Type = v.Type

	switch v.Type {
	case RichBlockTypeParagraph:
		b.Paragraph = &RichBlockParagraph{}
		return json.Unmarshal(data, b.Paragraph)
	case RichBlockTypeHeading:
		b.SectionHeading = &RichBlockSectionHeading{}
		return json.Unmarshal(data, b.SectionHeading)
	case RichBlockTypePre:
		b.Preformatted = &RichBlockPreformatted{}
		return json.Unmarshal(data, b.Preformatted)
	case RichBlockTypeFooter:
		b.Footer = &RichBlockFooter{}
		return json.Unmarshal(data, b.Footer)
	case RichBlockTypeDivider:
		b.Divider = &RichBlockDivider{}
		return json.Unmarshal(data, b.Divider)
	case RichBlockTypeMathematicalExpression:
		b.MathematicalExpression = &RichBlockMathematicalExpression{}
		return json.Unmarshal(data, b.MathematicalExpression)
	case RichBlockTypeAnchor:
		b.Anchor = &RichBlockAnchor{}
		return json.Unmarshal(data, b.Anchor)
	case RichBlockTypeList:
		b.List = &RichBlockList{}
		return json.Unmarshal(data, b.List)
	case RichBlockTypeBlockquote:
		b.BlockQuotation = &RichBlockBlockQuotation{}
		return json.Unmarshal(data, b.BlockQuotation)
	case RichBlockTypePullquote:
		b.PullQuotation = &RichBlockPullQuotation{}
		return json.Unmarshal(data, b.PullQuotation)
	case RichBlockTypeCollage:
		b.Collage = &RichBlockCollage{}
		return json.Unmarshal(data, b.Collage)
	case RichBlockTypeSlideshow:
		b.Slideshow = &RichBlockSlideshow{}
		return json.Unmarshal(data, b.Slideshow)
	case RichBlockTypeTable:
		b.Table = &RichBlockTable{}
		return json.Unmarshal(data, b.Table)
	case RichBlockTypeDetails:
		b.Details = &RichBlockDetails{}
		return json.Unmarshal(data, b.Details)
	case RichBlockTypeMap:
		b.Map = &RichBlockMap{}
		return json.Unmarshal(data, b.Map)
	case RichBlockTypeAnimation:
		b.Animation = &RichBlockAnimation{}
		return json.Unmarshal(data, b.Animation)
	case RichBlockTypeAudio:
		b.Audio = &RichBlockAudio{}
		return json.Unmarshal(data, b.Audio)
	case RichBlockTypePhoto:
		b.Photo = &RichBlockPhoto{}
		return json.Unmarshal(data, b.Photo)
	case RichBlockTypeVideo:
		b.Video = &RichBlockVideo{}
		return json.Unmarshal(data, b.Video)
	case RichBlockTypeVoiceNote:
		b.VoiceNote = &RichBlockVoiceNote{}
		return json.Unmarshal(data, b.VoiceNote)
	case RichBlockTypeThinking:
		b.Thinking = &RichBlockThinking{}
		return json.Unmarshal(data, b.Thinking)
	}

	return fmt.Errorf("unsupported RichBlock type %q", v.Type)
}

// RichBlockParagraph https://core.telegram.org/bots/api#richblockparagraph
type RichBlockParagraph struct {
	Type RichBlockType `json:"type"`
	Text RichText      `json:"text"`
}

// RichBlockSectionHeading https://core.telegram.org/bots/api#richblocksectionheading
type RichBlockSectionHeading struct {
	Type RichBlockType `json:"type"`
	Text RichText      `json:"text"`
	Size int           `json:"size"`
}

// RichBlockPreformatted https://core.telegram.org/bots/api#richblockpreformatted
type RichBlockPreformatted struct {
	Type     RichBlockType `json:"type"`
	Text     RichText      `json:"text"`
	Language string        `json:"language,omitempty"`
}

// RichBlockFooter https://core.telegram.org/bots/api#richblockfooter
type RichBlockFooter struct {
	Type RichBlockType `json:"type"`
	Text RichText      `json:"text"`
}

// RichBlockDivider https://core.telegram.org/bots/api#richblockdivider
type RichBlockDivider struct {
	Type RichBlockType `json:"type"`
}

// RichBlockMathematicalExpression https://core.telegram.org/bots/api#richblockmathematicalexpression
type RichBlockMathematicalExpression struct {
	Type       RichBlockType `json:"type"`
	Expression string        `json:"expression"`
}

// RichBlockAnchor https://core.telegram.org/bots/api#richblockanchor
type RichBlockAnchor struct {
	Type RichBlockType `json:"type"`
	Name string        `json:"name"`
}

// RichBlockList https://core.telegram.org/bots/api#richblocklist
type RichBlockList struct {
	Type  RichBlockType       `json:"type"`
	Items []RichBlockListItem `json:"items"`
}

// RichBlockBlockQuotation https://core.telegram.org/bots/api#richblockblockquotation
type RichBlockBlockQuotation struct {
	Type   RichBlockType `json:"type"`
	Blocks []RichBlock   `json:"blocks"`
	Credit *RichText     `json:"credit,omitempty"`
}

// RichBlockPullQuotation https://core.telegram.org/bots/api#richblockpullquotation
type RichBlockPullQuotation struct {
	Type   RichBlockType `json:"type"`
	Text   RichText      `json:"text"`
	Credit *RichText     `json:"credit,omitempty"`
}

// RichBlockCollage https://core.telegram.org/bots/api#richblockcollage
type RichBlockCollage struct {
	Type    RichBlockType     `json:"type"`
	Blocks  []RichBlock       `json:"blocks"`
	Caption *RichBlockCaption `json:"caption,omitempty"`
}

// RichBlockSlideshow https://core.telegram.org/bots/api#richblockslideshow
type RichBlockSlideshow struct {
	Type    RichBlockType     `json:"type"`
	Blocks  []RichBlock       `json:"blocks"`
	Caption *RichBlockCaption `json:"caption,omitempty"`
}

// RichBlockTable https://core.telegram.org/bots/api#richblocktable
type RichBlockTable struct {
	Type       RichBlockType          `json:"type"`
	Cells      [][]RichBlockTableCell `json:"cells"`
	IsBordered bool                   `json:"is_bordered,omitempty"`
	IsStriped  bool                   `json:"is_striped,omitempty"`
	Caption    *RichText              `json:"caption,omitempty"`
}

// RichBlockDetails https://core.telegram.org/bots/api#richblockdetails
type RichBlockDetails struct {
	Type    RichBlockType `json:"type"`
	Summary RichText      `json:"summary"`
	Blocks  []RichBlock   `json:"blocks"`
	IsOpen  bool          `json:"is_open,omitempty"`
}

// RichBlockMap https://core.telegram.org/bots/api#richblockmap
type RichBlockMap struct {
	Type     RichBlockType     `json:"type"`
	Location Location          `json:"location"`
	Zoom     int               `json:"zoom"`
	Width    int               `json:"width"`
	Height   int               `json:"height"`
	Caption  *RichBlockCaption `json:"caption,omitempty"`
}

// RichBlockAnimation https://core.telegram.org/bots/api#richblockanimation
type RichBlockAnimation struct {
	Type       RichBlockType     `json:"type"`
	Animation  Animation         `json:"animation"`
	HasSpoiler bool              `json:"has_spoiler,omitempty"`
	Caption    *RichBlockCaption `json:"caption,omitempty"`
}

// RichBlockAudio https://core.telegram.org/bots/api#richblockaudio
type RichBlockAudio struct {
	Type    RichBlockType     `json:"type"`
	Audio   Audio             `json:"audio"`
	Caption *RichBlockCaption `json:"caption,omitempty"`
}

// RichBlockPhoto https://core.telegram.org/bots/api#richblockphoto
type RichBlockPhoto struct {
	Type       RichBlockType     `json:"type"`
	Photo      []PhotoSize       `json:"photo"`
	HasSpoiler bool              `json:"has_spoiler,omitempty"`
	Caption    *RichBlockCaption `json:"caption,omitempty"`
}

// RichBlockVideo https://core.telegram.org/bots/api#richblockvideo
type RichBlockVideo struct {
	Type       RichBlockType     `json:"type"`
	Video      Video             `json:"video"`
	HasSpoiler bool              `json:"has_spoiler,omitempty"`
	Caption    *RichBlockCaption `json:"caption,omitempty"`
}

// RichBlockVoiceNote https://core.telegram.org/bots/api#richblockvoicenote
type RichBlockVoiceNote struct {
	Type      RichBlockType     `json:"type"`
	VoiceNote Voice             `json:"voice_note"`
	Caption   *RichBlockCaption `json:"caption,omitempty"`
}

// RichBlockThinking https://core.telegram.org/bots/api#richblockthinking
type RichBlockThinking struct {
	Type RichBlockType `json:"type"`
	Text RichText      `json:"text"`
}
