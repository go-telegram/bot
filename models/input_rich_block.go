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

	InputRichBlockParagraph              *InputRichBlockParagraph
	InputRichBlockSectionHeading         *InputRichBlockSectionHeading
	InputRichBlockPreformatted           *InputRichBlockPreformatted
	InputRichBlockFooter                 *InputRichBlockFooter
	InputRichBlockDivider                *InputRichBlockDivider
	InputRichBlockMathematicalExpression *InputRichBlockMathematicalExpression
	InputRichBlockAnchor                 *InputRichBlockAnchor
	InputRichBlockList                   *InputRichBlockList
	InputRichBlockBlockQuotation         *InputRichBlockBlockQuotation
	InputRichBlockPullQuotation          *InputRichBlockPullQuotation
	InputRichBlockCollage                *InputRichBlockCollage
	InputRichBlockSlideshow              *InputRichBlockSlideshow
	InputRichBlockTable                  *InputRichBlockTable
	InputRichBlockDetails                *InputRichBlockDetails
	InputRichBlockMap                    *InputRichBlockMap
	InputRichBlockAnimation              *InputRichBlockAnimation
	InputRichBlockAudio                  *InputRichBlockAudio
	InputRichBlockPhoto                  *InputRichBlockPhoto
	InputRichBlockVideo                  *InputRichBlockVideo
	InputRichBlockVoiceNote              *InputRichBlockVoiceNote
	InputRichBlockThinking               *InputRichBlockThinking
}

// variantIsNil reports whether the variant pointer matching Type is unset. An
// InputRichBlock is built by the caller, so a Type without its variant is a
// plausible mistake that must not reach json.Marshal as a nil dereference.
func (rb InputRichBlock) variantIsNil() bool {
	switch rb.Type {
	case RichBlockTypeParagraph:
		return rb.InputRichBlockParagraph == nil
	case RichBlockTypeSectionHeading:
		return rb.InputRichBlockSectionHeading == nil
	case RichBlockTypePreformatted:
		return rb.InputRichBlockPreformatted == nil
	case RichBlockTypeFooter:
		return rb.InputRichBlockFooter == nil
	case RichBlockTypeDivider:
		return rb.InputRichBlockDivider == nil
	case RichBlockTypeMathematicalExpression:
		return rb.InputRichBlockMathematicalExpression == nil
	case RichBlockTypeAnchor:
		return rb.InputRichBlockAnchor == nil
	case RichBlockTypeList:
		return rb.InputRichBlockList == nil
	case RichBlockTypeBlockQuotation:
		return rb.InputRichBlockBlockQuotation == nil
	case RichBlockTypePullQuotation:
		return rb.InputRichBlockPullQuotation == nil
	case RichBlockTypeCollage:
		return rb.InputRichBlockCollage == nil
	case RichBlockTypeSlideshow:
		return rb.InputRichBlockSlideshow == nil
	case RichBlockTypeTable:
		return rb.InputRichBlockTable == nil
	case RichBlockTypeDetails:
		return rb.InputRichBlockDetails == nil
	case RichBlockTypeMap:
		return rb.InputRichBlockMap == nil
	case RichBlockTypeAnimation:
		return rb.InputRichBlockAnimation == nil
	case RichBlockTypeAudio:
		return rb.InputRichBlockAudio == nil
	case RichBlockTypePhoto:
		return rb.InputRichBlockPhoto == nil
	case RichBlockTypeVideo:
		return rb.InputRichBlockVideo == nil
	case RichBlockTypeVoiceNote:
		return rb.InputRichBlockVoiceNote == nil
	case RichBlockTypeThinking:
		return rb.InputRichBlockThinking == nil
	}
	return true
}

// MarshalJSON implements json.Marshaler. The value receiver ensures the encoding
// is applied even when an InputRichBlock is reached as a (non-pointer) struct field.
func (rb InputRichBlock) MarshalJSON() ([]byte, error) {
	if rb.variantIsNil() {
		return nil, fmt.Errorf("nil variant for InputRichBlock type %q", rb.Type)
	}

	switch rb.Type {
	case RichBlockTypeParagraph:
		rb.InputRichBlockParagraph.Type = RichBlockTypeParagraph
		return json.Marshal(rb.InputRichBlockParagraph)
	case RichBlockTypeSectionHeading:
		rb.InputRichBlockSectionHeading.Type = RichBlockTypeSectionHeading
		return json.Marshal(rb.InputRichBlockSectionHeading)
	case RichBlockTypePreformatted:
		rb.InputRichBlockPreformatted.Type = RichBlockTypePreformatted
		return json.Marshal(rb.InputRichBlockPreformatted)
	case RichBlockTypeFooter:
		rb.InputRichBlockFooter.Type = RichBlockTypeFooter
		return json.Marshal(rb.InputRichBlockFooter)
	case RichBlockTypeDivider:
		rb.InputRichBlockDivider.Type = RichBlockTypeDivider
		return json.Marshal(rb.InputRichBlockDivider)
	case RichBlockTypeMathematicalExpression:
		rb.InputRichBlockMathematicalExpression.Type = RichBlockTypeMathematicalExpression
		return json.Marshal(rb.InputRichBlockMathematicalExpression)
	case RichBlockTypeAnchor:
		rb.InputRichBlockAnchor.Type = RichBlockTypeAnchor
		return json.Marshal(rb.InputRichBlockAnchor)
	case RichBlockTypeList:
		rb.InputRichBlockList.Type = RichBlockTypeList
		return json.Marshal(rb.InputRichBlockList)
	case RichBlockTypeBlockQuotation:
		rb.InputRichBlockBlockQuotation.Type = RichBlockTypeBlockQuotation
		return json.Marshal(rb.InputRichBlockBlockQuotation)
	case RichBlockTypePullQuotation:
		rb.InputRichBlockPullQuotation.Type = RichBlockTypePullQuotation
		return json.Marshal(rb.InputRichBlockPullQuotation)
	case RichBlockTypeCollage:
		rb.InputRichBlockCollage.Type = RichBlockTypeCollage
		return json.Marshal(rb.InputRichBlockCollage)
	case RichBlockTypeSlideshow:
		rb.InputRichBlockSlideshow.Type = RichBlockTypeSlideshow
		return json.Marshal(rb.InputRichBlockSlideshow)
	case RichBlockTypeTable:
		rb.InputRichBlockTable.Type = RichBlockTypeTable
		return json.Marshal(rb.InputRichBlockTable)
	case RichBlockTypeDetails:
		rb.InputRichBlockDetails.Type = RichBlockTypeDetails
		return json.Marshal(rb.InputRichBlockDetails)
	case RichBlockTypeMap:
		rb.InputRichBlockMap.Type = RichBlockTypeMap
		return json.Marshal(rb.InputRichBlockMap)
	case RichBlockTypeAnimation:
		rb.InputRichBlockAnimation.Type = RichBlockTypeAnimation
		return json.Marshal(rb.InputRichBlockAnimation)
	case RichBlockTypeAudio:
		rb.InputRichBlockAudio.Type = RichBlockTypeAudio
		return json.Marshal(rb.InputRichBlockAudio)
	case RichBlockTypePhoto:
		rb.InputRichBlockPhoto.Type = RichBlockTypePhoto
		return json.Marshal(rb.InputRichBlockPhoto)
	case RichBlockTypeVideo:
		rb.InputRichBlockVideo.Type = RichBlockTypeVideo
		return json.Marshal(rb.InputRichBlockVideo)
	case RichBlockTypeVoiceNote:
		rb.InputRichBlockVoiceNote.Type = RichBlockTypeVoiceNote
		return json.Marshal(rb.InputRichBlockVoiceNote)
	case RichBlockTypeThinking:
		rb.InputRichBlockThinking.Type = RichBlockTypeThinking
		return json.Marshal(rb.InputRichBlockThinking)
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
	case RichBlockTypeAnimation:
		rb.Type = RichBlockTypeAnimation
		rb.InputRichBlockAnimation = &InputRichBlockAnimation{}
		return json.Unmarshal(data, rb.InputRichBlockAnimation)
	case RichBlockTypeAudio:
		rb.Type = RichBlockTypeAudio
		rb.InputRichBlockAudio = &InputRichBlockAudio{}
		return json.Unmarshal(data, rb.InputRichBlockAudio)
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
