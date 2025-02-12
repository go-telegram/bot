package models

import (
	"encoding/json"
	"io"
)

// InputMedia https://core.telegram.org/bots/api#inputmedia
type InputMedia interface {
	inputMediaTag()

	MarshalInputMedia() ([]byte, error)
	Attachment() io.Reader
	GetMedia() string
}

// InputMediaPhoto https://core.telegram.org/bots/api#inputmediaphoto
type InputMediaPhoto struct {
	Media                 string          `json:"media"`
	Caption               string          `json:"caption,omitempty"`
	ParseMode             ParseMode       `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia bool            `json:"show_caption_above_media,omitempty"`
	HasSpoiler            bool            `json:"has_spoiler,omitempty"`

	MediaAttachment io.Reader `json:"-"`
}

func (m *InputMediaPhoto) Attachment() io.Reader {
	return m.MediaAttachment
}

func (m *InputMediaPhoto) GetMedia() string {
	return m.Media
}

func (m *InputMediaPhoto) MarshalInputMedia() ([]byte, error) {
	ret := struct {
		Type string `json:"type"`
		*InputMediaPhoto
	}{
		Type:            "photo",
		InputMediaPhoto: m,
	}

	return json.Marshal(&ret)
}

func (InputMediaPhoto) inputMediaTag() {}

// InputMediaVideo https://core.telegram.org/bots/api#inputmediavideo
type InputMediaVideo struct {
	Media                 string          `json:"media"`
	Thumbnail             InputFile       `json:"thumbnail,omitempty"`
	Cover                 string          `json:"cover,omitempty"`
	StartTimestamp        int             `json:"start_timestamp,omitempty"`
	Caption               string          `json:"caption,omitempty"`
	ParseMode             ParseMode       `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia bool            `json:"show_caption_above_media,omitempty"`
	Width                 int             `json:"width,omitempty"`
	Height                int             `json:"height,omitempty"`
	Duration              int             `json:"duration,omitempty"`
	SupportsStreaming     bool            `json:"supports_streaming,omitempty"`
	HasSpoiler            bool            `json:"has_spoiler,omitempty"`

	MediaAttachment io.Reader `json:"-"`
}

func (m *InputMediaVideo) Attachment() io.Reader {
	return m.MediaAttachment
}

func (m *InputMediaVideo) GetMedia() string {
	return m.Media
}

func (m InputMediaVideo) MarshalInputMedia() ([]byte, error) {
	return json.Marshal(&struct {
		Type string `json:"type"`
		InputMediaVideo
	}{
		Type:            "video",
		InputMediaVideo: m,
	})
}

func (InputMediaVideo) inputMediaTag() {}

// InputMediaAnimation https://core.telegram.org/bots/api#inputmediaanimation
type InputMediaAnimation struct {
	Media                 string          `json:"media"`
	Thumbnail             InputFile       `json:"thumbnail,omitempty"`
	Caption               string          `json:"caption,omitempty"`
	ParseMode             ParseMode       `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia bool            `json:"show_caption_above_media,omitempty"`
	Width                 int             `json:"width,omitempty"`
	Height                int             `json:"height,omitempty"`
	Duration              int             `json:"duration,omitempty"`
	HasSpoiler            bool            `json:"has_spoiler,omitempty"`

	MediaAttachment io.Reader `json:"-"`
}

func (m *InputMediaAnimation) Attachment() io.Reader {
	return m.MediaAttachment
}

func (m *InputMediaAnimation) GetMedia() string {
	return m.Media
}

func (m InputMediaAnimation) MarshalInputMedia() ([]byte, error) {
	return json.Marshal(&struct {
		Type string `json:"type"`
		InputMediaAnimation
	}{
		Type:                "animation",
		InputMediaAnimation: m,
	})
}

func (InputMediaAnimation) inputMediaTag() {}

// InputMediaAudio https://core.telegram.org/bots/api#inputmediaaudio
type InputMediaAudio struct {
	Media           string          `json:"media"`
	Thumbnail       InputFile       `json:"thumbnail,omitempty"`
	Caption         string          `json:"caption,omitempty"`
	ParseMode       ParseMode       `json:"parse_mode,omitempty"`
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
	Duration        int             `json:"duration,omitempty"`
	Performer       string          `json:"performer,omitempty"`
	Title           string          `json:"title,omitempty"`

	MediaAttachment io.Reader `json:"-"`
}

func (m *InputMediaAudio) Attachment() io.Reader {
	return m.MediaAttachment
}

func (m *InputMediaAudio) GetMedia() string {
	return m.Media
}

func (m InputMediaAudio) MarshalInputMedia() ([]byte, error) {
	return json.Marshal(&struct {
		Type string `json:"type"`
		InputMediaAudio
	}{
		Type:            "audio",
		InputMediaAudio: m,
	})
}

func (InputMediaAudio) inputMediaTag() {}

// InputMediaDocument https://core.telegram.org/bots/api#inputmediadocument
type InputMediaDocument struct {
	Media                       string          `json:"media"`
	Thumbnail                   InputFile       `json:"thumbnail,omitempty"`
	Caption                     string          `json:"caption,omitempty"`
	ParseMode                   ParseMode       `json:"parse_mode,omitempty"`
	CaptionEntities             []MessageEntity `json:"caption_entities,omitempty"`
	DisableContentTypeDetection bool            `json:"disable_content_type_detection,omitempty"`

	MediaAttachment io.Reader `json:"-"`
}

func (m *InputMediaDocument) Attachment() io.Reader {
	return m.MediaAttachment
}

func (m *InputMediaDocument) GetMedia() string {
	return m.Media
}

func (m InputMediaDocument) MarshalInputMedia() ([]byte, error) {
	return json.Marshal(&struct {
		Type string `json:"type"`
		InputMediaDocument
	}{
		Type:               "document",
		InputMediaDocument: m,
	})
}

func (InputMediaDocument) inputMediaTag() {}
