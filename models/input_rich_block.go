package models

import (
	"encoding/json"
	"fmt"
)

// InputRichBlock https://core.telegram.org/bots/api#inputrichblock
//
// InputRichBlock is a tagged union describing a single block of an outgoing rich
// message. Type holds the discriminator (reusing the RichBlockType values) and
// the matching variant pointer is populated.
type InputRichBlock struct {
	Type RichBlockType

	InputRichBlockParagraph                *InputRichBlockParagraph
	InputRichBlockSectionHeading           *InputRichBlockSectionHeading
	InputRichBlockPreformatted             *InputRichBlockPreformatted
	InputRichBlockFooter                   *InputRichBlockFooter
	InputRichBlockDivider                  *InputRichBlockDivider
	InputRichBlockMathematicalExpression   *InputRichBlockMathematicalExpression
	InputRichBlockAnchor                   *InputRichBlockAnchor
	InputRichBlockList                     *InputRichBlockList
	InputRichBlockBlockQuotation           *InputRichBlockBlockQuotation
	InputRichBlockExpandableBlockQuotation *InputRichBlockExpandableBlockQuotation
	InputRichBlockPullQuotation            *InputRichBlockPullQuotation
	InputRichBlockCollage                  *InputRichBlockCollage
	InputRichBlockSlideshow                *InputRichBlockSlideshow
	InputRichBlockTable                    *InputRichBlockTable
	InputRichBlockDetails                  *InputRichBlockDetails
	InputRichBlockMap                      *InputRichBlockMap
	InputRichBlockButtons                  *InputRichBlockButtons
	InputRichBlockAnimation                *InputRichBlockAnimation
	InputRichBlockAudio                    *InputRichBlockAudio
	InputRichBlockDocument                 *InputRichBlockDocument
	InputRichBlockPhoto                    *InputRichBlockPhoto
	InputRichBlockVideo                    *InputRichBlockVideo
	InputRichBlockVoiceNote                *InputRichBlockVoiceNote
	InputRichBlockThinking                 *InputRichBlockThinking
}

// MarshalJSON implements json.Marshaler. The value receiver ensures the encoding
// is applied even when an InputRichBlock is reached as a (non-pointer) struct field.
func (rb InputRichBlock) MarshalJSON() ([]byte, error) {
	switch rb.Type {
	case RichBlockTypeParagraph:
		return marshalVariant("InputRichBlock", rb.Type, rb.InputRichBlockParagraph, func(v *InputRichBlockParagraph) { v.Type = rb.Type })
	case RichBlockTypeSectionHeading:
		return marshalVariant("InputRichBlock", rb.Type, rb.InputRichBlockSectionHeading, func(v *InputRichBlockSectionHeading) { v.Type = rb.Type })
	case RichBlockTypePreformatted:
		return marshalVariant("InputRichBlock", rb.Type, rb.InputRichBlockPreformatted, func(v *InputRichBlockPreformatted) { v.Type = rb.Type })
	case RichBlockTypeFooter:
		return marshalVariant("InputRichBlock", rb.Type, rb.InputRichBlockFooter, func(v *InputRichBlockFooter) { v.Type = rb.Type })
	case RichBlockTypeDivider:
		return marshalVariant("InputRichBlock", rb.Type, rb.InputRichBlockDivider, func(v *InputRichBlockDivider) { v.Type = rb.Type })
	case RichBlockTypeMathematicalExpression:
		return marshalVariant("InputRichBlock", rb.Type, rb.InputRichBlockMathematicalExpression, func(v *InputRichBlockMathematicalExpression) { v.Type = rb.Type })
	case RichBlockTypeAnchor:
		return marshalVariant("InputRichBlock", rb.Type, rb.InputRichBlockAnchor, func(v *InputRichBlockAnchor) { v.Type = rb.Type })
	case RichBlockTypeList:
		return marshalVariant("InputRichBlock", rb.Type, rb.InputRichBlockList, func(v *InputRichBlockList) { v.Type = rb.Type })
	case RichBlockTypeBlockQuotation:
		return marshalVariant("InputRichBlock", rb.Type, rb.InputRichBlockBlockQuotation, func(v *InputRichBlockBlockQuotation) { v.Type = rb.Type })
	case RichBlockTypeExpandableBlockQuotation:
		return marshalVariant("InputRichBlock", rb.Type, rb.InputRichBlockExpandableBlockQuotation, func(v *InputRichBlockExpandableBlockQuotation) { v.Type = rb.Type })
	case RichBlockTypePullQuotation:
		return marshalVariant("InputRichBlock", rb.Type, rb.InputRichBlockPullQuotation, func(v *InputRichBlockPullQuotation) { v.Type = rb.Type })
	case RichBlockTypeCollage:
		return marshalVariant("InputRichBlock", rb.Type, rb.InputRichBlockCollage, func(v *InputRichBlockCollage) { v.Type = rb.Type })
	case RichBlockTypeSlideshow:
		return marshalVariant("InputRichBlock", rb.Type, rb.InputRichBlockSlideshow, func(v *InputRichBlockSlideshow) { v.Type = rb.Type })
	case RichBlockTypeTable:
		return marshalVariant("InputRichBlock", rb.Type, rb.InputRichBlockTable, func(v *InputRichBlockTable) { v.Type = rb.Type })
	case RichBlockTypeDetails:
		return marshalVariant("InputRichBlock", rb.Type, rb.InputRichBlockDetails, func(v *InputRichBlockDetails) { v.Type = rb.Type })
	case RichBlockTypeMap:
		return marshalVariant("InputRichBlock", rb.Type, rb.InputRichBlockMap, func(v *InputRichBlockMap) { v.Type = rb.Type })
	case RichBlockTypeButtons:
		return marshalVariant("InputRichBlock", rb.Type, rb.InputRichBlockButtons, func(v *InputRichBlockButtons) { v.Type = rb.Type })
	case RichBlockTypeAnimation:
		return marshalVariant("InputRichBlock", rb.Type, rb.InputRichBlockAnimation, func(v *InputRichBlockAnimation) { v.Type = rb.Type })
	case RichBlockTypeAudio:
		return marshalVariant("InputRichBlock", rb.Type, rb.InputRichBlockAudio, func(v *InputRichBlockAudio) { v.Type = rb.Type })
	case RichBlockTypeDocument:
		return marshalVariant("InputRichBlock", rb.Type, rb.InputRichBlockDocument, func(v *InputRichBlockDocument) { v.Type = rb.Type })
	case RichBlockTypePhoto:
		return marshalVariant("InputRichBlock", rb.Type, rb.InputRichBlockPhoto, func(v *InputRichBlockPhoto) { v.Type = rb.Type })
	case RichBlockTypeVideo:
		return marshalVariant("InputRichBlock", rb.Type, rb.InputRichBlockVideo, func(v *InputRichBlockVideo) { v.Type = rb.Type })
	case RichBlockTypeVoiceNote:
		return marshalVariant("InputRichBlock", rb.Type, rb.InputRichBlockVoiceNote, func(v *InputRichBlockVoiceNote) { v.Type = rb.Type })
	case RichBlockTypeThinking:
		return marshalVariant("InputRichBlock", rb.Type, rb.InputRichBlockThinking, func(v *InputRichBlockThinking) { v.Type = rb.Type })
	}

	return nil, fmt.Errorf("unsupported InputRichBlock type %q", rb.Type)
}

func (rb *InputRichBlock) UnmarshalJSON(data []byte) error {
	v := struct {
		Type RichBlockType `json:"type"`
	}{}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	switch v.Type {
	case RichBlockTypeParagraph:
		rb.Type = RichBlockTypeParagraph
		rb.InputRichBlockParagraph = &InputRichBlockParagraph{}
		return json.Unmarshal(data, rb.InputRichBlockParagraph)
	case RichBlockTypeSectionHeading:
		rb.Type = RichBlockTypeSectionHeading
		rb.InputRichBlockSectionHeading = &InputRichBlockSectionHeading{}
		return json.Unmarshal(data, rb.InputRichBlockSectionHeading)
	case RichBlockTypePreformatted:
		rb.Type = RichBlockTypePreformatted
		rb.InputRichBlockPreformatted = &InputRichBlockPreformatted{}
		return json.Unmarshal(data, rb.InputRichBlockPreformatted)
	case RichBlockTypeFooter:
		rb.Type = RichBlockTypeFooter
		rb.InputRichBlockFooter = &InputRichBlockFooter{}
		return json.Unmarshal(data, rb.InputRichBlockFooter)
	case RichBlockTypeDivider:
		rb.Type = RichBlockTypeDivider
		rb.InputRichBlockDivider = &InputRichBlockDivider{}
		return json.Unmarshal(data, rb.InputRichBlockDivider)
	case RichBlockTypeMathematicalExpression:
		rb.Type = RichBlockTypeMathematicalExpression
		rb.InputRichBlockMathematicalExpression = &InputRichBlockMathematicalExpression{}
		return json.Unmarshal(data, rb.InputRichBlockMathematicalExpression)
	case RichBlockTypeAnchor:
		rb.Type = RichBlockTypeAnchor
		rb.InputRichBlockAnchor = &InputRichBlockAnchor{}
		return json.Unmarshal(data, rb.InputRichBlockAnchor)
	case RichBlockTypeList:
		rb.Type = RichBlockTypeList
		rb.InputRichBlockList = &InputRichBlockList{}
		return json.Unmarshal(data, rb.InputRichBlockList)
	case RichBlockTypeBlockQuotation:
		rb.Type = RichBlockTypeBlockQuotation
		rb.InputRichBlockBlockQuotation = &InputRichBlockBlockQuotation{}
		return json.Unmarshal(data, rb.InputRichBlockBlockQuotation)
	case RichBlockTypeExpandableBlockQuotation:
		rb.Type = RichBlockTypeExpandableBlockQuotation
		rb.InputRichBlockExpandableBlockQuotation = &InputRichBlockExpandableBlockQuotation{}
		return json.Unmarshal(data, rb.InputRichBlockExpandableBlockQuotation)
	case RichBlockTypePullQuotation:
		rb.Type = RichBlockTypePullQuotation
		rb.InputRichBlockPullQuotation = &InputRichBlockPullQuotation{}
		return json.Unmarshal(data, rb.InputRichBlockPullQuotation)
	case RichBlockTypeCollage:
		rb.Type = RichBlockTypeCollage
		rb.InputRichBlockCollage = &InputRichBlockCollage{}
		return json.Unmarshal(data, rb.InputRichBlockCollage)
	case RichBlockTypeSlideshow:
		rb.Type = RichBlockTypeSlideshow
		rb.InputRichBlockSlideshow = &InputRichBlockSlideshow{}
		return json.Unmarshal(data, rb.InputRichBlockSlideshow)
	case RichBlockTypeTable:
		rb.Type = RichBlockTypeTable
		rb.InputRichBlockTable = &InputRichBlockTable{}
		return json.Unmarshal(data, rb.InputRichBlockTable)
	case RichBlockTypeDetails:
		rb.Type = RichBlockTypeDetails
		rb.InputRichBlockDetails = &InputRichBlockDetails{}
		return json.Unmarshal(data, rb.InputRichBlockDetails)
	case RichBlockTypeMap:
		rb.Type = RichBlockTypeMap
		rb.InputRichBlockMap = &InputRichBlockMap{}
		return json.Unmarshal(data, rb.InputRichBlockMap)
	case RichBlockTypeButtons:
		rb.Type = RichBlockTypeButtons
		rb.InputRichBlockButtons = &InputRichBlockButtons{}
		return json.Unmarshal(data, rb.InputRichBlockButtons)
	case RichBlockTypeAnimation:
		rb.Type = RichBlockTypeAnimation
		rb.InputRichBlockAnimation = &InputRichBlockAnimation{}
		return json.Unmarshal(data, rb.InputRichBlockAnimation)
	case RichBlockTypeAudio:
		rb.Type = RichBlockTypeAudio
		rb.InputRichBlockAudio = &InputRichBlockAudio{}
		return json.Unmarshal(data, rb.InputRichBlockAudio)
	case RichBlockTypeDocument:
		rb.Type = RichBlockTypeDocument
		rb.InputRichBlockDocument = &InputRichBlockDocument{}
		return json.Unmarshal(data, rb.InputRichBlockDocument)
	case RichBlockTypePhoto:
		rb.Type = RichBlockTypePhoto
		rb.InputRichBlockPhoto = &InputRichBlockPhoto{}
		return json.Unmarshal(data, rb.InputRichBlockPhoto)
	case RichBlockTypeVideo:
		rb.Type = RichBlockTypeVideo
		rb.InputRichBlockVideo = &InputRichBlockVideo{}
		return json.Unmarshal(data, rb.InputRichBlockVideo)
	case RichBlockTypeVoiceNote:
		rb.Type = RichBlockTypeVoiceNote
		rb.InputRichBlockVoiceNote = &InputRichBlockVoiceNote{}
		return json.Unmarshal(data, rb.InputRichBlockVoiceNote)
	case RichBlockTypeThinking:
		rb.Type = RichBlockTypeThinking
		rb.InputRichBlockThinking = &InputRichBlockThinking{}
		return json.Unmarshal(data, rb.InputRichBlockThinking)
	}

	return fmt.Errorf("unsupported InputRichBlock type %q", v.Type)
}

// InputRichBlockListItem https://core.telegram.org/bots/api#inputrichblocklistitem
//
// An item in a list to be sent.
type InputRichBlockListItem struct {
	Blocks      []InputRichBlock `json:"blocks"`
	HasCheckbox bool             `json:"has_checkbox,omitempty"`
	IsChecked   bool             `json:"is_checked,omitempty"`
	Value       int              `json:"value,omitempty"`
	Type        string           `json:"type,omitempty"`
}

// InputRichBlockParagraph https://core.telegram.org/bots/api#inputrichblockparagraph
type InputRichBlockParagraph struct {
	Type RichBlockType `json:"type"` // always "paragraph"
	Text RichText      `json:"text"`
}

// InputRichBlockSectionHeading https://core.telegram.org/bots/api#inputrichblocksectionheading
type InputRichBlockSectionHeading struct {
	Type RichBlockType `json:"type"` // always "heading"
	Text RichText      `json:"text"`
	Size int           `json:"size"`
}

// InputRichBlockPreformatted https://core.telegram.org/bots/api#inputrichblockpreformatted
type InputRichBlockPreformatted struct {
	Type     RichBlockType `json:"type"` // always "pre"
	Text     RichText      `json:"text"`
	Language string        `json:"language,omitempty"`
}

// InputRichBlockFooter https://core.telegram.org/bots/api#inputrichblockfooter
type InputRichBlockFooter struct {
	Type RichBlockType `json:"type"` // always "footer"
	Text RichText      `json:"text"`
}

// InputRichBlockDivider https://core.telegram.org/bots/api#inputrichblockdivider
type InputRichBlockDivider struct {
	Type RichBlockType `json:"type"` // always "divider"
}

// InputRichBlockMathematicalExpression https://core.telegram.org/bots/api#inputrichblockmathematicalexpression
type InputRichBlockMathematicalExpression struct {
	Type       RichBlockType `json:"type"` // always "mathematical_expression"
	Expression string        `json:"expression"`
}

// InputRichBlockAnchor https://core.telegram.org/bots/api#inputrichblockanchor
type InputRichBlockAnchor struct {
	Type RichBlockType `json:"type"` // always "anchor"
	Name string        `json:"name"`
}

// InputRichBlockList https://core.telegram.org/bots/api#inputrichblocklist
type InputRichBlockList struct {
	Type  RichBlockType            `json:"type"` // always "list"
	Items []InputRichBlockListItem `json:"items"`
}

// InputRichBlockBlockQuotation https://core.telegram.org/bots/api#inputrichblockblockquotation
type InputRichBlockBlockQuotation struct {
	Type   RichBlockType    `json:"type"` // always "blockquote"
	Blocks []InputRichBlock `json:"blocks"`
	Credit *RichText        `json:"credit,omitempty"`
}

// InputRichBlockExpandableBlockQuotation https://core.telegram.org/bots/api#inputrichblockexpandableblockquotation
type InputRichBlockExpandableBlockQuotation struct {
	Type   RichBlockType `json:"type"` // always "expandable_blockquote"
	Text   RichText      `json:"text"`
	Credit *RichText     `json:"credit,omitempty"`
}

// InputRichBlockPullQuotation https://core.telegram.org/bots/api#inputrichblockpullquotation
type InputRichBlockPullQuotation struct {
	Type   RichBlockType `json:"type"` // always "pullquote"
	Text   RichText      `json:"text"`
	Credit *RichText     `json:"credit,omitempty"`
}

// InputRichBlockCollage https://core.telegram.org/bots/api#inputrichblockcollage
type InputRichBlockCollage struct {
	Type    RichBlockType     `json:"type"` // always "collage"
	Blocks  []InputRichBlock  `json:"blocks"`
	Caption *RichBlockCaption `json:"caption,omitempty"`
}

// InputRichBlockSlideshow https://core.telegram.org/bots/api#inputrichblockslideshow
type InputRichBlockSlideshow struct {
	Type    RichBlockType     `json:"type"` // always "slideshow"
	Blocks  []InputRichBlock  `json:"blocks"`
	Caption *RichBlockCaption `json:"caption,omitempty"`
}

// InputRichBlockTable https://core.telegram.org/bots/api#inputrichblocktable
type InputRichBlockTable struct {
	Type       RichBlockType          `json:"type"` // always "table"
	Cells      [][]RichBlockTableCell `json:"cells"`
	IsBordered bool                   `json:"is_bordered,omitempty"`
	IsStriped  bool                   `json:"is_striped,omitempty"`
	IsCompact  bool                   `json:"is_compact,omitempty"`
	Caption    *RichText              `json:"caption,omitempty"`
}

// InputRichBlockDetails https://core.telegram.org/bots/api#inputrichblockdetails
type InputRichBlockDetails struct {
	Type    RichBlockType    `json:"type"` // always "details"
	Summary RichText         `json:"summary"`
	Blocks  []InputRichBlock `json:"blocks"`
	IsOpen  bool             `json:"is_open,omitempty"`
}

// InputRichBlockMap https://core.telegram.org/bots/api#inputrichblockmap
type InputRichBlockMap struct {
	Type     RichBlockType     `json:"type"` // always "map"
	Location Location          `json:"location"`
	Zoom     int               `json:"zoom"`
	Width    int               `json:"width"`
	Height   int               `json:"height"`
	Caption  *RichBlockCaption `json:"caption,omitempty"`
}

// InputRichBlockButtons https://core.telegram.org/bots/api#inputrichblockbuttons
type InputRichBlockButtons struct {
	Type    RichBlockType       `json:"type"` // always "buttons"
	Buttons []RichMessageButton `json:"buttons"`
	Align   string              `json:"align,omitempty"`
}

// InputRichBlockAnimation https://core.telegram.org/bots/api#inputrichblockanimation
type InputRichBlockAnimation struct {
	Type      RichBlockType       `json:"type"` // always "animation"
	Animation InputMediaAnimation `json:"animation"`
	Caption   *RichBlockCaption   `json:"caption,omitempty"`
}

// InputRichBlockAudio https://core.telegram.org/bots/api#inputrichblockaudio
type InputRichBlockAudio struct {
	Type    RichBlockType     `json:"type"` // always "audio"
	Audio   InputMediaAudio   `json:"audio"`
	Caption *RichBlockCaption `json:"caption,omitempty"`
}

// InputRichBlockDocument https://core.telegram.org/bots/api#inputrichblockdocument
type InputRichBlockDocument struct {
	Type     RichBlockType      `json:"type"` // always "document"
	Document InputMediaDocument `json:"document"`
	Caption  *RichBlockCaption  `json:"caption,omitempty"`
}

// InputRichBlockPhoto https://core.telegram.org/bots/api#inputrichblockphoto
type InputRichBlockPhoto struct {
	Type    RichBlockType     `json:"type"` // always "photo"
	Photo   InputMediaPhoto   `json:"photo"`
	Caption *RichBlockCaption `json:"caption,omitempty"`
}

// InputRichBlockVideo https://core.telegram.org/bots/api#inputrichblockvideo
type InputRichBlockVideo struct {
	Type    RichBlockType     `json:"type"` // always "video"
	Video   InputMediaVideo   `json:"video"`
	Caption *RichBlockCaption `json:"caption,omitempty"`
}

// InputRichBlockVoiceNote https://core.telegram.org/bots/api#inputrichblockvoicenote
type InputRichBlockVoiceNote struct {
	Type      RichBlockType       `json:"type"` // always "voice_note"
	VoiceNote InputMediaVoiceNote `json:"voice_note"`
	Caption   *RichBlockCaption   `json:"caption,omitempty"`
}

// InputRichBlockThinking https://core.telegram.org/bots/api#inputrichblockthinking
type InputRichBlockThinking struct {
	Type RichBlockType `json:"type"` // always "thinking"
	Text RichText      `json:"text"`
}
