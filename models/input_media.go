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

// inputMediaPhotoAlias strips the methods of InputMediaPhoto so the embedded value below is
// encoded as plain fields instead of recursing back into MarshalJSON.
type inputMediaPhotoAlias InputMediaPhoto

func (m InputMediaPhoto) MarshalInputMedia() ([]byte, error) {
	return json.Marshal(struct {
		Type string `json:"type"`
		inputMediaPhotoAlias
	}{
		Type:                 "photo",
		inputMediaPhotoAlias: inputMediaPhotoAlias(m),
	})
}

// MarshalJSON keeps the required "type" discriminator when the value is reached
// through a plain json.Marshal, e.g. nested inside a rich message.
func (m InputMediaPhoto) MarshalJSON() ([]byte, error) {
	return m.MarshalInputMedia()
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

func (m *InputMediaVideo) GetThumbnail() InputFile {
	return m.Thumbnail
}

// inputMediaVideoAlias strips the methods of InputMediaVideo so the embedded value below is
// encoded as plain fields instead of recursing back into MarshalJSON.
type inputMediaVideoAlias InputMediaVideo

func (m InputMediaVideo) MarshalInputMedia() ([]byte, error) {
	m.Thumbnail = normalizeInputFile(m.Thumbnail)
	return json.Marshal(struct {
		Type string `json:"type"`
		inputMediaVideoAlias
	}{
		Type:                 "video",
		inputMediaVideoAlias: inputMediaVideoAlias(m),
	})
}

// MarshalJSON keeps the required "type" discriminator when the value is reached
// through a plain json.Marshal, e.g. nested inside a rich message.
func (m InputMediaVideo) MarshalJSON() ([]byte, error) {
	return m.MarshalInputMedia()
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

func (m *InputMediaAnimation) GetThumbnail() InputFile {
	return m.Thumbnail
}

// inputMediaAnimationAlias strips the methods of InputMediaAnimation so the embedded value below is
// encoded as plain fields instead of recursing back into MarshalJSON.
type inputMediaAnimationAlias InputMediaAnimation

func (m InputMediaAnimation) MarshalInputMedia() ([]byte, error) {
	m.Thumbnail = normalizeInputFile(m.Thumbnail)
	return json.Marshal(struct {
		Type string `json:"type"`
		inputMediaAnimationAlias
	}{
		Type:                     "animation",
		inputMediaAnimationAlias: inputMediaAnimationAlias(m),
	})
}

// MarshalJSON keeps the required "type" discriminator when the value is reached
// through a plain json.Marshal, e.g. nested inside a rich message.
func (m InputMediaAnimation) MarshalJSON() ([]byte, error) {
	return m.MarshalInputMedia()
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

func (m *InputMediaAudio) GetThumbnail() InputFile {
	return m.Thumbnail
}

// inputMediaAudioAlias strips the methods of InputMediaAudio so the embedded value below is
// encoded as plain fields instead of recursing back into MarshalJSON.
type inputMediaAudioAlias InputMediaAudio

func (m InputMediaAudio) MarshalInputMedia() ([]byte, error) {
	m.Thumbnail = normalizeInputFile(m.Thumbnail)
	return json.Marshal(struct {
		Type string `json:"type"`
		inputMediaAudioAlias
	}{
		Type:                 "audio",
		inputMediaAudioAlias: inputMediaAudioAlias(m),
	})
}

// MarshalJSON keeps the required "type" discriminator when the value is reached
// through a plain json.Marshal, e.g. nested inside a rich message.
func (m InputMediaAudio) MarshalJSON() ([]byte, error) {
	return m.MarshalInputMedia()
}

func (InputMediaAudio) inputMediaTag() {}

// InputMediaVoiceNote https://core.telegram.org/bots/api#inputmediavoicenote
type InputMediaVoiceNote struct {
	Media           string          `json:"media"`
	Caption         string          `json:"caption,omitempty"`
	ParseMode       ParseMode       `json:"parse_mode,omitempty"`
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
	Duration        int             `json:"duration,omitempty"`

	MediaAttachment io.Reader `json:"-"`
}

func (m *InputMediaVoiceNote) Attachment() io.Reader {
	return m.MediaAttachment
}

func (m *InputMediaVoiceNote) GetMedia() string {
	return m.Media
}

// inputMediaVoiceNoteAlias strips the methods of InputMediaVoiceNote so the embedded value below is
// encoded as plain fields instead of recursing back into MarshalJSON.
type inputMediaVoiceNoteAlias InputMediaVoiceNote

func (m InputMediaVoiceNote) MarshalInputMedia() ([]byte, error) {
	return json.Marshal(struct {
		Type string `json:"type"`
		inputMediaVoiceNoteAlias
	}{
		Type:                     "voice_note",
		inputMediaVoiceNoteAlias: inputMediaVoiceNoteAlias(m),
	})
}

// MarshalJSON keeps the required "type" discriminator when the value is reached
// through a plain json.Marshal, e.g. nested inside a rich message.
func (m InputMediaVoiceNote) MarshalJSON() ([]byte, error) {
	return m.MarshalInputMedia()
}

func (InputMediaVoiceNote) inputMediaTag() {}

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

func (m *InputMediaDocument) GetThumbnail() InputFile {
	return m.Thumbnail
}

// inputMediaDocumentAlias strips the methods of InputMediaDocument so the embedded value below is
// encoded as plain fields instead of recursing back into MarshalJSON.
type inputMediaDocumentAlias InputMediaDocument

func (m InputMediaDocument) MarshalInputMedia() ([]byte, error) {
	m.Thumbnail = normalizeInputFile(m.Thumbnail)
	return json.Marshal(struct {
		Type string `json:"type"`
		inputMediaDocumentAlias
	}{
		Type:                    "document",
		inputMediaDocumentAlias: inputMediaDocumentAlias(m),
	})
}

// MarshalJSON keeps the required "type" discriminator when the value is reached
// through a plain json.Marshal, e.g. nested inside a rich message.
func (m InputMediaDocument) MarshalJSON() ([]byte, error) {
	return m.MarshalInputMedia()
}

func (InputMediaDocument) inputMediaTag() {}

// InputMediaLivePhoto https://core.telegram.org/bots/api#inputmedialivephoto
type InputMediaLivePhoto struct {
	Media                 string          `json:"media"`
	Photo                 string          `json:"photo"`
	Caption               string          `json:"caption,omitempty"`
	ParseMode             ParseMode       `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia bool            `json:"show_caption_above_media,omitempty"`
	HasSpoiler            bool            `json:"has_spoiler,omitempty"`
	MediaAttachment       io.Reader       `json:"-"`
	PhotoAttachment       io.Reader       `json:"-"`
}

func (m *InputMediaLivePhoto) Attachment() io.Reader {
	return m.MediaAttachment
}

func (m *InputMediaLivePhoto) GetMedia() string {
	return m.Media
}

// inputMediaLivePhotoAlias strips the methods of InputMediaLivePhoto so the embedded value below is
// encoded as plain fields instead of recursing back into MarshalJSON.
type inputMediaLivePhotoAlias InputMediaLivePhoto

func (m InputMediaLivePhoto) MarshalInputMedia() ([]byte, error) {
	return json.Marshal(struct {
		Type string `json:"type"`
		inputMediaLivePhotoAlias
	}{
		Type:                     "live_photo",
		inputMediaLivePhotoAlias: inputMediaLivePhotoAlias(m),
	})
}

// MarshalJSON keeps the required "type" discriminator when the value is reached
// through a plain json.Marshal, e.g. nested inside a rich message.
func (m InputMediaLivePhoto) MarshalJSON() ([]byte, error) {
	return m.MarshalInputMedia()
}

func (InputMediaLivePhoto) inputMediaTag() {}

// InputMediaLocation https://core.telegram.org/bots/api#inputmedialocation
type InputMediaLocation struct {
	Latitude           float64 `json:"latitude"`
	Longitude          float64 `json:"longitude"`
	HorizontalAccuracy float64 `json:"horizontal_accuracy,omitempty"`
}

func (m *InputMediaLocation) Attachment() io.Reader {
	return nil
}

func (m *InputMediaLocation) GetMedia() string {
	return ""
}

// inputMediaLocationAlias strips the methods of InputMediaLocation so the embedded value below is
// encoded as plain fields instead of recursing back into MarshalJSON.
type inputMediaLocationAlias InputMediaLocation

func (m InputMediaLocation) MarshalInputMedia() ([]byte, error) {
	return json.Marshal(struct {
		Type string `json:"type"`
		inputMediaLocationAlias
	}{
		Type:                    "location",
		inputMediaLocationAlias: inputMediaLocationAlias(m),
	})
}

// MarshalJSON keeps the required "type" discriminator when the value is reached
// through a plain json.Marshal, e.g. nested inside a rich message.
func (m InputMediaLocation) MarshalJSON() ([]byte, error) {
	return m.MarshalInputMedia()
}

func (InputMediaLocation) inputMediaTag() {}

// InputMediaSticker https://core.telegram.org/bots/api#inputmediasticker
type InputMediaSticker struct {
	Media           string    `json:"media"`
	Emoji           string    `json:"emoji,omitempty"`
	MediaAttachment io.Reader `json:"-"`
}

func (m *InputMediaSticker) Attachment() io.Reader {
	return m.MediaAttachment
}

func (m *InputMediaSticker) GetMedia() string {
	return m.Media
}

// inputMediaStickerAlias strips the methods of InputMediaSticker so the embedded value below is
// encoded as plain fields instead of recursing back into MarshalJSON.
type inputMediaStickerAlias InputMediaSticker

func (m InputMediaSticker) MarshalInputMedia() ([]byte, error) {
	return json.Marshal(struct {
		Type string `json:"type"`
		inputMediaStickerAlias
	}{
		Type:                   "sticker",
		inputMediaStickerAlias: inputMediaStickerAlias(m),
	})
}

// MarshalJSON keeps the required "type" discriminator when the value is reached
// through a plain json.Marshal, e.g. nested inside a rich message.
func (m InputMediaSticker) MarshalJSON() ([]byte, error) {
	return m.MarshalInputMedia()
}

func (InputMediaSticker) inputMediaTag() {}

// InputMediaVenue https://core.telegram.org/bots/api#inputmediavenue
type InputMediaVenue struct {
	Latitude        float64 `json:"latitude"`
	Longitude       float64 `json:"longitude"`
	Title           string  `json:"title"`
	Address         string  `json:"address"`
	FoursquareID    string  `json:"foursquare_id,omitempty"`
	FoursquareType  string  `json:"foursquare_type,omitempty"`
	GooglePlaceID   string  `json:"google_place_id,omitempty"`
	GooglePlaceType string  `json:"google_place_type,omitempty"`
}

func (m *InputMediaVenue) Attachment() io.Reader {
	return nil
}

func (m *InputMediaVenue) GetMedia() string {
	return ""
}

// inputMediaVenueAlias strips the methods of InputMediaVenue so the embedded value below is
// encoded as plain fields instead of recursing back into MarshalJSON.
type inputMediaVenueAlias InputMediaVenue

func (m InputMediaVenue) MarshalInputMedia() ([]byte, error) {
	return json.Marshal(struct {
		Type string `json:"type"`
		inputMediaVenueAlias
	}{
		Type:                 "venue",
		inputMediaVenueAlias: inputMediaVenueAlias(m),
	})
}

// MarshalJSON keeps the required "type" discriminator when the value is reached
// through a plain json.Marshal, e.g. nested inside a rich message.
func (m InputMediaVenue) MarshalJSON() ([]byte, error) {
	return m.MarshalInputMedia()
}

func (InputMediaVenue) inputMediaTag() {}
