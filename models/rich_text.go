package models

import (
	"bytes"
	"encoding/json"
	"fmt"
)

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
// RichText is polymorphic: a rich text value is encoded by Telegram as a bare
// JSON string, a JSON array of RichText (a sequence), or a tagged object with a
// "type" discriminator. The form is recorded in Type: it is the empty string for
// the plain-string and array forms, set in PlainText and Array respectively, and
// otherwise holds the discriminator with the matching variant pointer populated.
type RichText struct {
	Type      RichTextType
	PlainText string
	Array     []RichText

	RichTextBold                   *RichTextBold
	RichTextItalic                 *RichTextItalic
	RichTextUnderline              *RichTextUnderline
	RichTextStrikethrough          *RichTextStrikethrough
	RichTextSpoiler                *RichTextSpoiler
	RichTextDateTime               *RichTextDateTime
	RichTextTextMention            *RichTextTextMention
	RichTextSubscript              *RichTextSubscript
	RichTextSuperscript            *RichTextSuperscript
	RichTextMarked                 *RichTextMarked
	RichTextCode                   *RichTextCode
	RichTextCustomEmoji            *RichTextCustomEmoji
	RichTextMathematicalExpression *RichTextMathematicalExpression
	RichTextURL                    *RichTextURL
	RichTextEmailAddress           *RichTextEmailAddress
	RichTextPhoneNumber            *RichTextPhoneNumber
	RichTextBankCardNumber         *RichTextBankCardNumber
	RichTextMention                *RichTextMention
	RichTextHashtag                *RichTextHashtag
	RichTextCashtag                *RichTextCashtag
	RichTextBotCommand             *RichTextBotCommand
	RichTextAnchor                 *RichTextAnchor
	RichTextAnchorLink             *RichTextAnchorLink
	RichTextReference              *RichTextReference
	RichTextReferenceLink          *RichTextReferenceLink
}

// MarshalJSON implements json.Marshaler. The value receiver ensures the encoding
// is applied even when a RichText is reached as a (non-pointer) struct field.
func (rt RichText) MarshalJSON() ([]byte, error) {
	if rt.Array != nil {
		return json.Marshal(rt.Array)
	}

	switch rt.Type {
	case "":
		return json.Marshal(rt.PlainText)
	case RichTextTypeBold:
		rt.RichTextBold.Type = RichTextTypeBold
		return json.Marshal(rt.RichTextBold)
	case RichTextTypeItalic:
		rt.RichTextItalic.Type = RichTextTypeItalic
		return json.Marshal(rt.RichTextItalic)
	case RichTextTypeUnderline:
		rt.RichTextUnderline.Type = RichTextTypeUnderline
		return json.Marshal(rt.RichTextUnderline)
	case RichTextTypeStrikethrough:
		rt.RichTextStrikethrough.Type = RichTextTypeStrikethrough
		return json.Marshal(rt.RichTextStrikethrough)
	case RichTextTypeSpoiler:
		rt.RichTextSpoiler.Type = RichTextTypeSpoiler
		return json.Marshal(rt.RichTextSpoiler)
	case RichTextTypeDateTime:
		rt.RichTextDateTime.Type = RichTextTypeDateTime
		return json.Marshal(rt.RichTextDateTime)
	case RichTextTypeTextMention:
		rt.RichTextTextMention.Type = RichTextTypeTextMention
		return json.Marshal(rt.RichTextTextMention)
	case RichTextTypeSubscript:
		rt.RichTextSubscript.Type = RichTextTypeSubscript
		return json.Marshal(rt.RichTextSubscript)
	case RichTextTypeSuperscript:
		rt.RichTextSuperscript.Type = RichTextTypeSuperscript
		return json.Marshal(rt.RichTextSuperscript)
	case RichTextTypeMarked:
		rt.RichTextMarked.Type = RichTextTypeMarked
		return json.Marshal(rt.RichTextMarked)
	case RichTextTypeCode:
		rt.RichTextCode.Type = RichTextTypeCode
		return json.Marshal(rt.RichTextCode)
	case RichTextTypeCustomEmoji:
		rt.RichTextCustomEmoji.Type = RichTextTypeCustomEmoji
		return json.Marshal(rt.RichTextCustomEmoji)
	case RichTextTypeMathematicalExpression:
		rt.RichTextMathematicalExpression.Type = RichTextTypeMathematicalExpression
		return json.Marshal(rt.RichTextMathematicalExpression)
	case RichTextTypeURL:
		rt.RichTextURL.Type = RichTextTypeURL
		return json.Marshal(rt.RichTextURL)
	case RichTextTypeEmailAddress:
		rt.RichTextEmailAddress.Type = RichTextTypeEmailAddress
		return json.Marshal(rt.RichTextEmailAddress)
	case RichTextTypePhoneNumber:
		rt.RichTextPhoneNumber.Type = RichTextTypePhoneNumber
		return json.Marshal(rt.RichTextPhoneNumber)
	case RichTextTypeBankCardNumber:
		rt.RichTextBankCardNumber.Type = RichTextTypeBankCardNumber
		return json.Marshal(rt.RichTextBankCardNumber)
	case RichTextTypeMention:
		rt.RichTextMention.Type = RichTextTypeMention
		return json.Marshal(rt.RichTextMention)
	case RichTextTypeHashtag:
		rt.RichTextHashtag.Type = RichTextTypeHashtag
		return json.Marshal(rt.RichTextHashtag)
	case RichTextTypeCashtag:
		rt.RichTextCashtag.Type = RichTextTypeCashtag
		return json.Marshal(rt.RichTextCashtag)
	case RichTextTypeBotCommand:
		rt.RichTextBotCommand.Type = RichTextTypeBotCommand
		return json.Marshal(rt.RichTextBotCommand)
	case RichTextTypeAnchor:
		rt.RichTextAnchor.Type = RichTextTypeAnchor
		return json.Marshal(rt.RichTextAnchor)
	case RichTextTypeAnchorLink:
		rt.RichTextAnchorLink.Type = RichTextTypeAnchorLink
		return json.Marshal(rt.RichTextAnchorLink)
	case RichTextTypeReference:
		rt.RichTextReference.Type = RichTextTypeReference
		return json.Marshal(rt.RichTextReference)
	case RichTextTypeReferenceLink:
		rt.RichTextReferenceLink.Type = RichTextTypeReferenceLink
		return json.Marshal(rt.RichTextReferenceLink)
	}

	return nil, fmt.Errorf("unsupported RichText type %q", rt.Type)
}

func (rt *RichText) UnmarshalJSON(data []byte) error {
	trimmed := bytes.TrimSpace(data)
	if len(trimmed) == 0 || bytes.Equal(trimmed, []byte("null")) {
		return nil
	}

	switch trimmed[0] {
	case '"':
		rt.Type = ""
		return json.Unmarshal(trimmed, &rt.PlainText)
	case '[':
		rt.Type = ""
		return json.Unmarshal(trimmed, &rt.Array)
	case '{':
		// fall through to the tagged-object dispatch below
	default:
		return fmt.Errorf("unexpected RichText JSON: %s", trimmed)
	}

	v := struct {
		Type RichTextType `json:"type"`
	}{}
	if err := json.Unmarshal(trimmed, &v); err != nil {
		return err
	}

	switch v.Type {
	case RichTextTypeBold:
		rt.Type = RichTextTypeBold
		rt.RichTextBold = &RichTextBold{}
		return json.Unmarshal(trimmed, rt.RichTextBold)
	case RichTextTypeItalic:
		rt.Type = RichTextTypeItalic
		rt.RichTextItalic = &RichTextItalic{}
		return json.Unmarshal(trimmed, rt.RichTextItalic)
	case RichTextTypeUnderline:
		rt.Type = RichTextTypeUnderline
		rt.RichTextUnderline = &RichTextUnderline{}
		return json.Unmarshal(trimmed, rt.RichTextUnderline)
	case RichTextTypeStrikethrough:
		rt.Type = RichTextTypeStrikethrough
		rt.RichTextStrikethrough = &RichTextStrikethrough{}
		return json.Unmarshal(trimmed, rt.RichTextStrikethrough)
	case RichTextTypeSpoiler:
		rt.Type = RichTextTypeSpoiler
		rt.RichTextSpoiler = &RichTextSpoiler{}
		return json.Unmarshal(trimmed, rt.RichTextSpoiler)
	case RichTextTypeDateTime:
		rt.Type = RichTextTypeDateTime
		rt.RichTextDateTime = &RichTextDateTime{}
		return json.Unmarshal(trimmed, rt.RichTextDateTime)
	case RichTextTypeTextMention:
		rt.Type = RichTextTypeTextMention
		rt.RichTextTextMention = &RichTextTextMention{}
		return json.Unmarshal(trimmed, rt.RichTextTextMention)
	case RichTextTypeSubscript:
		rt.Type = RichTextTypeSubscript
		rt.RichTextSubscript = &RichTextSubscript{}
		return json.Unmarshal(trimmed, rt.RichTextSubscript)
	case RichTextTypeSuperscript:
		rt.Type = RichTextTypeSuperscript
		rt.RichTextSuperscript = &RichTextSuperscript{}
		return json.Unmarshal(trimmed, rt.RichTextSuperscript)
	case RichTextTypeMarked:
		rt.Type = RichTextTypeMarked
		rt.RichTextMarked = &RichTextMarked{}
		return json.Unmarshal(trimmed, rt.RichTextMarked)
	case RichTextTypeCode:
		rt.Type = RichTextTypeCode
		rt.RichTextCode = &RichTextCode{}
		return json.Unmarshal(trimmed, rt.RichTextCode)
	case RichTextTypeCustomEmoji:
		rt.Type = RichTextTypeCustomEmoji
		rt.RichTextCustomEmoji = &RichTextCustomEmoji{}
		return json.Unmarshal(trimmed, rt.RichTextCustomEmoji)
	case RichTextTypeMathematicalExpression:
		rt.Type = RichTextTypeMathematicalExpression
		rt.RichTextMathematicalExpression = &RichTextMathematicalExpression{}
		return json.Unmarshal(trimmed, rt.RichTextMathematicalExpression)
	case RichTextTypeURL:
		rt.Type = RichTextTypeURL
		rt.RichTextURL = &RichTextURL{}
		return json.Unmarshal(trimmed, rt.RichTextURL)
	case RichTextTypeEmailAddress:
		rt.Type = RichTextTypeEmailAddress
		rt.RichTextEmailAddress = &RichTextEmailAddress{}
		return json.Unmarshal(trimmed, rt.RichTextEmailAddress)
	case RichTextTypePhoneNumber:
		rt.Type = RichTextTypePhoneNumber
		rt.RichTextPhoneNumber = &RichTextPhoneNumber{}
		return json.Unmarshal(trimmed, rt.RichTextPhoneNumber)
	case RichTextTypeBankCardNumber:
		rt.Type = RichTextTypeBankCardNumber
		rt.RichTextBankCardNumber = &RichTextBankCardNumber{}
		return json.Unmarshal(trimmed, rt.RichTextBankCardNumber)
	case RichTextTypeMention:
		rt.Type = RichTextTypeMention
		rt.RichTextMention = &RichTextMention{}
		return json.Unmarshal(trimmed, rt.RichTextMention)
	case RichTextTypeHashtag:
		rt.Type = RichTextTypeHashtag
		rt.RichTextHashtag = &RichTextHashtag{}
		return json.Unmarshal(trimmed, rt.RichTextHashtag)
	case RichTextTypeCashtag:
		rt.Type = RichTextTypeCashtag
		rt.RichTextCashtag = &RichTextCashtag{}
		return json.Unmarshal(trimmed, rt.RichTextCashtag)
	case RichTextTypeBotCommand:
		rt.Type = RichTextTypeBotCommand
		rt.RichTextBotCommand = &RichTextBotCommand{}
		return json.Unmarshal(trimmed, rt.RichTextBotCommand)
	case RichTextTypeAnchor:
		rt.Type = RichTextTypeAnchor
		rt.RichTextAnchor = &RichTextAnchor{}
		return json.Unmarshal(trimmed, rt.RichTextAnchor)
	case RichTextTypeAnchorLink:
		rt.Type = RichTextTypeAnchorLink
		rt.RichTextAnchorLink = &RichTextAnchorLink{}
		return json.Unmarshal(trimmed, rt.RichTextAnchorLink)
	case RichTextTypeReference:
		rt.Type = RichTextTypeReference
		rt.RichTextReference = &RichTextReference{}
		return json.Unmarshal(trimmed, rt.RichTextReference)
	case RichTextTypeReferenceLink:
		rt.Type = RichTextTypeReferenceLink
		rt.RichTextReferenceLink = &RichTextReferenceLink{}
		return json.Unmarshal(trimmed, rt.RichTextReferenceLink)
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
	UnixTime       int          `json:"unix_time"`
	DateTimeFormat string       `json:"date_time_format"`
}

// RichTextTextMention https://core.telegram.org/bots/api#richtexttextmention
type RichTextTextMention struct {
	Type RichTextType `json:"type"`
	Text RichText     `json:"text"`
	User *User        `json:"user"`
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
