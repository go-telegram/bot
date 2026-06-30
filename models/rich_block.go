package models

import (
	"encoding/json"
	"fmt"
)

// RichBlockType https://core.telegram.org/bots/api#richblock
type RichBlockType string

const (
	RichBlockTypeParagraph              RichBlockType = "paragraph"
	RichBlockTypeSectionHeading         RichBlockType = "heading"
	RichBlockTypePreformatted           RichBlockType = "pre"
	RichBlockTypeFooter                 RichBlockType = "footer"
	RichBlockTypeDivider                RichBlockType = "divider"
	RichBlockTypeMathematicalExpression RichBlockType = "mathematical_expression"
	RichBlockTypeAnchor                 RichBlockType = "anchor"
	RichBlockTypeList                   RichBlockType = "list"
	RichBlockTypeBlockQuotation         RichBlockType = "blockquote"
	RichBlockTypePullQuotation          RichBlockType = "pullquote"
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
// RichBlock is a tagged union; Type holds the discriminator and the matching
// variant pointer is populated.
type RichBlock struct {
	Type RichBlockType

	RichBlockParagraph              *RichBlockParagraph
	RichBlockSectionHeading         *RichBlockSectionHeading
	RichBlockPreformatted           *RichBlockPreformatted
	RichBlockFooter                 *RichBlockFooter
	RichBlockDivider                *RichBlockDivider
	RichBlockMathematicalExpression *RichBlockMathematicalExpression
	RichBlockAnchor                 *RichBlockAnchor
	RichBlockList                   *RichBlockList
	RichBlockBlockQuotation         *RichBlockBlockQuotation
	RichBlockPullQuotation          *RichBlockPullQuotation
	RichBlockCollage                *RichBlockCollage
	RichBlockSlideshow              *RichBlockSlideshow
	RichBlockTable                  *RichBlockTable
	RichBlockDetails                *RichBlockDetails
	RichBlockMap                    *RichBlockMap
	RichBlockAnimation              *RichBlockAnimation
	RichBlockAudio                  *RichBlockAudio
	RichBlockPhoto                  *RichBlockPhoto
	RichBlockVideo                  *RichBlockVideo
	RichBlockVoiceNote              *RichBlockVoiceNote
	RichBlockThinking               *RichBlockThinking
}

// MarshalJSON implements json.Marshaler. The value receiver ensures the encoding
// is applied even when a RichBlock is reached as a (non-pointer) struct field.
func (rb RichBlock) MarshalJSON() ([]byte, error) {
	switch rb.Type {
	case RichBlockTypeParagraph:
		rb.RichBlockParagraph.Type = RichBlockTypeParagraph
		return json.Marshal(rb.RichBlockParagraph)
	case RichBlockTypeSectionHeading:
		rb.RichBlockSectionHeading.Type = RichBlockTypeSectionHeading
		return json.Marshal(rb.RichBlockSectionHeading)
	case RichBlockTypePreformatted:
		rb.RichBlockPreformatted.Type = RichBlockTypePreformatted
		return json.Marshal(rb.RichBlockPreformatted)
	case RichBlockTypeFooter:
		rb.RichBlockFooter.Type = RichBlockTypeFooter
		return json.Marshal(rb.RichBlockFooter)
	case RichBlockTypeDivider:
		rb.RichBlockDivider.Type = RichBlockTypeDivider
		return json.Marshal(rb.RichBlockDivider)
	case RichBlockTypeMathematicalExpression:
		rb.RichBlockMathematicalExpression.Type = RichBlockTypeMathematicalExpression
		return json.Marshal(rb.RichBlockMathematicalExpression)
	case RichBlockTypeAnchor:
		rb.RichBlockAnchor.Type = RichBlockTypeAnchor
		return json.Marshal(rb.RichBlockAnchor)
	case RichBlockTypeList:
		rb.RichBlockList.Type = RichBlockTypeList
		return json.Marshal(rb.RichBlockList)
	case RichBlockTypeBlockQuotation:
		rb.RichBlockBlockQuotation.Type = RichBlockTypeBlockQuotation
		return json.Marshal(rb.RichBlockBlockQuotation)
	case RichBlockTypePullQuotation:
		rb.RichBlockPullQuotation.Type = RichBlockTypePullQuotation
		return json.Marshal(rb.RichBlockPullQuotation)
	case RichBlockTypeCollage:
		rb.RichBlockCollage.Type = RichBlockTypeCollage
		return json.Marshal(rb.RichBlockCollage)
	case RichBlockTypeSlideshow:
		rb.RichBlockSlideshow.Type = RichBlockTypeSlideshow
		return json.Marshal(rb.RichBlockSlideshow)
	case RichBlockTypeTable:
		rb.RichBlockTable.Type = RichBlockTypeTable
		return json.Marshal(rb.RichBlockTable)
	case RichBlockTypeDetails:
		rb.RichBlockDetails.Type = RichBlockTypeDetails
		return json.Marshal(rb.RichBlockDetails)
	case RichBlockTypeMap:
		rb.RichBlockMap.Type = RichBlockTypeMap
		return json.Marshal(rb.RichBlockMap)
	case RichBlockTypeAnimation:
		rb.RichBlockAnimation.Type = RichBlockTypeAnimation
		return json.Marshal(rb.RichBlockAnimation)
	case RichBlockTypeAudio:
		rb.RichBlockAudio.Type = RichBlockTypeAudio
		return json.Marshal(rb.RichBlockAudio)
	case RichBlockTypePhoto:
		rb.RichBlockPhoto.Type = RichBlockTypePhoto
		return json.Marshal(rb.RichBlockPhoto)
	case RichBlockTypeVideo:
		rb.RichBlockVideo.Type = RichBlockTypeVideo
		return json.Marshal(rb.RichBlockVideo)
	case RichBlockTypeVoiceNote:
		rb.RichBlockVoiceNote.Type = RichBlockTypeVoiceNote
		return json.Marshal(rb.RichBlockVoiceNote)
	case RichBlockTypeThinking:
		rb.RichBlockThinking.Type = RichBlockTypeThinking
		return json.Marshal(rb.RichBlockThinking)
	}

	return nil, fmt.Errorf("unsupported RichBlock type %q", rb.Type)
}

func (rb *RichBlock) UnmarshalJSON(data []byte) error {
	v := struct {
		Type RichBlockType `json:"type"`
	}{}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	switch v.Type {
	case RichBlockTypeParagraph:
		rb.Type = RichBlockTypeParagraph
		rb.RichBlockParagraph = &RichBlockParagraph{}
		return json.Unmarshal(data, rb.RichBlockParagraph)
	case RichBlockTypeSectionHeading:
		rb.Type = RichBlockTypeSectionHeading
		rb.RichBlockSectionHeading = &RichBlockSectionHeading{}
		return json.Unmarshal(data, rb.RichBlockSectionHeading)
	case RichBlockTypePreformatted:
		rb.Type = RichBlockTypePreformatted
		rb.RichBlockPreformatted = &RichBlockPreformatted{}
		return json.Unmarshal(data, rb.RichBlockPreformatted)
	case RichBlockTypeFooter:
		rb.Type = RichBlockTypeFooter
		rb.RichBlockFooter = &RichBlockFooter{}
		return json.Unmarshal(data, rb.RichBlockFooter)
	case RichBlockTypeDivider:
		rb.Type = RichBlockTypeDivider
		rb.RichBlockDivider = &RichBlockDivider{}
		return json.Unmarshal(data, rb.RichBlockDivider)
	case RichBlockTypeMathematicalExpression:
		rb.Type = RichBlockTypeMathematicalExpression
		rb.RichBlockMathematicalExpression = &RichBlockMathematicalExpression{}
		return json.Unmarshal(data, rb.RichBlockMathematicalExpression)
	case RichBlockTypeAnchor:
		rb.Type = RichBlockTypeAnchor
		rb.RichBlockAnchor = &RichBlockAnchor{}
		return json.Unmarshal(data, rb.RichBlockAnchor)
	case RichBlockTypeList:
		rb.Type = RichBlockTypeList
		rb.RichBlockList = &RichBlockList{}
		return json.Unmarshal(data, rb.RichBlockList)
	case RichBlockTypeBlockQuotation:
		rb.Type = RichBlockTypeBlockQuotation
		rb.RichBlockBlockQuotation = &RichBlockBlockQuotation{}
		return json.Unmarshal(data, rb.RichBlockBlockQuotation)
	case RichBlockTypePullQuotation:
		rb.Type = RichBlockTypePullQuotation
		rb.RichBlockPullQuotation = &RichBlockPullQuotation{}
		return json.Unmarshal(data, rb.RichBlockPullQuotation)
	case RichBlockTypeCollage:
		rb.Type = RichBlockTypeCollage
		rb.RichBlockCollage = &RichBlockCollage{}
		return json.Unmarshal(data, rb.RichBlockCollage)
	case RichBlockTypeSlideshow:
		rb.Type = RichBlockTypeSlideshow
		rb.RichBlockSlideshow = &RichBlockSlideshow{}
		return json.Unmarshal(data, rb.RichBlockSlideshow)
	case RichBlockTypeTable:
		rb.Type = RichBlockTypeTable
		rb.RichBlockTable = &RichBlockTable{}
		return json.Unmarshal(data, rb.RichBlockTable)
	case RichBlockTypeDetails:
		rb.Type = RichBlockTypeDetails
		rb.RichBlockDetails = &RichBlockDetails{}
		return json.Unmarshal(data, rb.RichBlockDetails)
	case RichBlockTypeMap:
		rb.Type = RichBlockTypeMap
		rb.RichBlockMap = &RichBlockMap{}
		return json.Unmarshal(data, rb.RichBlockMap)
	case RichBlockTypeAnimation:
		rb.Type = RichBlockTypeAnimation
		rb.RichBlockAnimation = &RichBlockAnimation{}
		return json.Unmarshal(data, rb.RichBlockAnimation)
	case RichBlockTypeAudio:
		rb.Type = RichBlockTypeAudio
		rb.RichBlockAudio = &RichBlockAudio{}
		return json.Unmarshal(data, rb.RichBlockAudio)
	case RichBlockTypePhoto:
		rb.Type = RichBlockTypePhoto
		rb.RichBlockPhoto = &RichBlockPhoto{}
		return json.Unmarshal(data, rb.RichBlockPhoto)
	case RichBlockTypeVideo:
		rb.Type = RichBlockTypeVideo
		rb.RichBlockVideo = &RichBlockVideo{}
		return json.Unmarshal(data, rb.RichBlockVideo)
	case RichBlockTypeVoiceNote:
		rb.Type = RichBlockTypeVoiceNote
		rb.RichBlockVoiceNote = &RichBlockVoiceNote{}
		return json.Unmarshal(data, rb.RichBlockVoiceNote)
	case RichBlockTypeThinking:
		rb.Type = RichBlockTypeThinking
		rb.RichBlockThinking = &RichBlockThinking{}
		return json.Unmarshal(data, rb.RichBlockThinking)
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
//
// A block with a "Thinking..." placeholder, corresponding to the custom HTML tag
// <tg-thinking>.
type RichBlockThinking struct {
	Type RichBlockType `json:"type"`
	Text RichText      `json:"text"`
}
