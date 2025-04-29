package models

import (
	"encoding/json"
	"fmt"
	"io"
)

type InputPaidMedia interface {
	inputPaidMediaTag()

	MarshalInputMedia() ([]byte, error)
	Attachment() io.Reader
	GetMedia() string
}

// InputPaidMediaPhoto https://core.telegram.org/bots/api#inputpaidmediaphoto
type InputPaidMediaPhoto struct {
	Media string `json:"media"`

	MediaAttachment io.Reader `json:"-"`
}

func (m *InputPaidMediaPhoto) inputPaidMediaTag() {}

func (m *InputPaidMediaPhoto) Attachment() io.Reader {
	return m.MediaAttachment
}

func (m *InputPaidMediaPhoto) GetMedia() string {
	return m.Media
}

func (m *InputPaidMediaPhoto) MarshalInputMedia() ([]byte, error) {
	ret := struct {
		Type string `json:"type"`
		*InputPaidMediaPhoto
	}{
		Type:                "photo",
		InputPaidMediaPhoto: m,
	}

	return json.Marshal(&ret)
}

// InputPaidMediaVideo https://core.telegram.org/bots/api#inputpaidmediavideo
type InputPaidMediaVideo struct {
	Media             string    `json:"media"`
	Thumbnail         InputFile `json:"thumbnail,omitempty"`
	Cover             string    `json:"cover,omitempty"`
	StartTimestamp    int       `json:"start_timestamp,omitempty"`
	Width             int       `json:"width,omitempty"`
	Height            int       `json:"height,omitempty"`
	Duration          int       `json:"duration,omitempty"`
	SupportsStreaming bool      `json:"supports_streaming,omitempty"`

	MediaAttachment io.Reader `json:"-"`
}

func (m *InputPaidMediaVideo) inputPaidMediaTag() {}

func (m *InputPaidMediaVideo) Attachment() io.Reader {
	return m.MediaAttachment
}

func (m *InputPaidMediaVideo) GetMedia() string {
	return m.Media
}

func (m *InputPaidMediaVideo) MarshalInputMedia() ([]byte, error) {
	ret := struct {
		Type string `json:"type"`
		*InputPaidMediaVideo
	}{
		Type:                "video",
		InputPaidMediaVideo: m,
	}

	return json.Marshal(&ret)
}

type PaidMediaType string

const (
	PaidMediaTypePreview PaidMediaType = "preview"
	PaidMediaTypePhoto   PaidMediaType = "photo"
	PaidMediaTypeVideo   PaidMediaType = "video"
)

// PaidMedia https://core.telegram.org/bots/api#paidmedia
type PaidMedia struct {
	Type PaidMediaType

	Preview *PaidMediaPreview
	Photo   *PaidMediaPhoto
	Video   *PaidMediaVideo
}

func (p *PaidMedia) UnmarshalJSON(data []byte) error {
	v := struct {
		Type PaidMediaType `json:"type"`
	}{}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}

	p.Type = v.Type

	switch v.Type {
	case PaidMediaTypePreview:
		p.Preview = &PaidMediaPreview{}
		return json.Unmarshal(data, p.Preview)
	case PaidMediaTypePhoto:
		p.Photo = &PaidMediaPhoto{}
		return json.Unmarshal(data, p.Photo)
	case PaidMediaTypeVideo:
		p.Video = &PaidMediaVideo{}
		return json.Unmarshal(data, p.Video)
	default:
		return fmt.Errorf("unsupported PaidMedia type, %v", v.Type)
	}
}

// PaidMediaPreview https://core.telegram.org/bots/api#paidmediapreview
type PaidMediaPreview struct {
	Type PaidMediaType

	Width    int `json:"width,omitempty"`
	Height   int `json:"height,omitempty"`
	Duration int `json:"duration,omitempty"`
}

// PaidMediaPhoto https://core.telegram.org/bots/api#paidmediaphoto
type PaidMediaPhoto struct {
	Type PaidMediaType

	Photo []PhotoSize `json:"photo"`
}

// PaidMediaVideo https://core.telegram.org/bots/api#paidmediavideo
type PaidMediaVideo struct {
	Type PaidMediaType

	Video Video `json:"video"`
}

// PaidMediaInfo https://core.telegram.org/bots/api#paidmediainfo
type PaidMediaInfo struct {
	StarCount int         `json:"star_count"`
	PaidMedia []PaidMedia `json:"paid_media"`
}

// PaidMediaPurchased https://core.telegram.org/bots/api#paidmediapurchased
type PaidMediaPurchased struct {
	From             User   `json:"from"`
	PaidMediaPayload string `json:"paid_media_payload"`
}

// PaidMessagePriceChanged https://core.telegram.org/bots/api#paidmessagepricechanged
type PaidMessagePriceChanged struct {
	PaidMessageStarCount int `json:"paid_message_star_count"`
}
