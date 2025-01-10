package models

import (
	"encoding/json"
)

// InlineQuery https://core.telegram.org/bots/api#inlinequery
type InlineQuery struct {
	ID       string    `json:"id"`
	From     *User     `json:"from"`
	Query    string    `json:"query"`
	Offset   string    `json:"offset"`
	ChatType string    `json:"chat_type,omitempty"`
	Location *Location `json:"location,omitempty"`
}

// AnswerInlineQuery https://core.telegram.org/bots/api#answerinlinequery
type AnswerInlineQuery struct {
	InlineQueryID     string              `json:"inline_query_id"`
	Results           []InlineQueryResult `json:"results"`
	CacheTime         int                 `json:"cache_time,omitempty"`
	IsPersonal        bool                `json:"is_personal,omitempty"`
	NextOffset        string              `json:"next_offset,omitempty"`
	SwitchPmText      string              `json:"switch_pm_text,omitempty"`
	SwitchPmParameter string              `json:"switch_pm_parameter,omitempty"`
}

// InlineQueryResultsButton https://core.telegram.org/bots/api#inlinequeryresultsbutton
type InlineQueryResultsButton struct {
	Text           string      `json:"text"`
	WebApp         *WebAppInfo `json:"web_app"`
	StartParameter string      `json:"start_parameter,omitempty"`
}

// InlineQueryResult https://core.telegram.org/bots/api#inlinequeryresult
type InlineQueryResult interface {
	inlineQueryResultTag()

	MarshalCustom() ([]byte, error)
}

// InlineQueryResultArticle https://core.telegram.org/bots/api#inlinequeryresultarticle
type InlineQueryResultArticle struct {
	ID                  string              `json:"id"`
	Title               string              `json:"title,omitempty"`
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
	ReplyMarkup         ReplyMarkup         `json:"reply_markup,omitempty"`
	URL                 string              `json:"url,omitempty"`
	Description         string              `json:"description,omitempty"`
	ThumbnailURL        string              `json:"thumbnail_url,omitempty"`
	ThumbnailWidth      int                 `json:"thumbnail_width,omitempty"`
	ThumbnailHeight     int                 `json:"thumbnail_height,omitempty"`
}

func (*InlineQueryResultArticle) inlineQueryResultTag() {}

func (m *InlineQueryResultArticle) MarshalCustom() ([]byte, error) {
	ret := struct {
		Type string `json:"type"`
		*InlineQueryResultArticle
	}{
		Type:                     "article",
		InlineQueryResultArticle: m,
	}

	return json.Marshal(&ret)
}

// InlineQueryResultPhoto https://core.telegram.org/bots/api#inlinequeryresultphoto
type InlineQueryResultPhoto struct {
	ID                    string              `json:"id"`
	PhotoURL              string              `json:"photo_url"`
	ThumbnailURL          string              `json:"thumbnail_url"`
	PhotoWidth            int                 `json:"photo_width,omitempty"`
	PhotoHeight           int                 `json:"photo_height,omitempty"`
	Title                 string              `json:"title,omitempty"`
	Description           string              `json:"description,omitempty"`
	Caption               string              `json:"caption,omitempty"`
	ParseMode             ParseMode           `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity     `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia bool                `json:"show_caption_above_media,omitempty"`
	ReplyMarkup           ReplyMarkup         `json:"reply_markup,omitempty"`
	InputMessageContent   InputMessageContent `json:"input_message_content,omitempty"`
}

func (*InlineQueryResultPhoto) inlineQueryResultTag() {}

func (m *InlineQueryResultPhoto) MarshalCustom() ([]byte, error) {
	ret := struct {
		Type string `json:"type"`
		*InlineQueryResultPhoto
	}{
		Type:                   "photo",
		InlineQueryResultPhoto: m,
	}

	return json.Marshal(&ret)
}

// InlineQueryResultGif https://core.telegram.org/bots/api#inlinequeryresultgif
type InlineQueryResultGif struct {
	ID                    string              `json:"id"`
	GifURL                string              `json:"gif_url"`
	GifWidth              int                 `json:"gif_width,omitempty"`
	GifHeight             int                 `json:"gif_height,omitempty"`
	GifDuration           int                 `json:"gif_duration,omitempty"`
	ThumbnailURL          string              `json:"thumbnail_url"`
	ThumbnailMimeType     string              `json:"thumbnail_mime_type,omitempty"`
	Title                 string              `json:"title,omitempty"`
	Caption               string              `json:"caption,omitempty"`
	ParseMode             ParseMode           `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity     `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia bool                `json:"show_caption_above_media,omitempty"`
	ReplyMarkup           ReplyMarkup         `json:"reply_markup,omitempty"`
	InputMessageContent   InputMessageContent `json:"input_message_content,omitempty"`
}

func (*InlineQueryResultGif) inlineQueryResultTag() {}

func (m *InlineQueryResultGif) MarshalCustom() ([]byte, error) {
	ret := struct {
		Type string `json:"type"`
		*InlineQueryResultGif
	}{
		Type:                 "gif",
		InlineQueryResultGif: m,
	}

	return json.Marshal(&ret)
}

// InlineQueryResultMpeg4Gif https://core.telegram.org/bots/api#inlinequeryresultmpeg4gif
type InlineQueryResultMpeg4Gif struct {
	ID                    string              `json:"id"`
	Mpeg4URL              string              `json:"mpeg4_url"`
	Mpeg4Width            int                 `json:"mpeg4_width,omitempty"`
	Mpeg4Height           int                 `json:"mpeg4_height,omitempty"`
	Mpeg4Duration         int                 `json:"mpeg4_duration,omitempty"`
	ThumbnailURL          string              `json:"thumbnail_url"`
	ThumbnailMimeType     string              `json:"thumbnail_mime_type,omitempty"`
	Title                 string              `json:"title,omitempty"`
	Caption               string              `json:"caption,omitempty"`
	ParseMode             ParseMode           `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity     `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia bool                `json:"show_caption_above_media,omitempty"`
	ReplyMarkup           ReplyMarkup         `json:"reply_markup,omitempty"`
	InputMessageContent   InputMessageContent `json:"input_message_content,omitempty"`
}

func (*InlineQueryResultMpeg4Gif) inlineQueryResultTag() {}

func (m *InlineQueryResultMpeg4Gif) MarshalCustom() ([]byte, error) {
	ret := struct {
		Type string `json:"type"`
		*InlineQueryResultMpeg4Gif
	}{
		Type:                      "mpeg4_gif",
		InlineQueryResultMpeg4Gif: m,
	}

	return json.Marshal(&ret)
}

// InlineQueryResultVideo https://core.telegram.org/bots/api#inlinequeryresultvideo
type InlineQueryResultVideo struct {
	ID                    string              `json:"id"`
	VideoURL              string              `json:"video_url"`
	MimeType              string              `json:"mime_type,omitempty"`
	ThumbnailURL          string              `json:"thumbnail_url"`
	Title                 string              `json:"title,omitempty"`
	Caption               string              `json:"caption,omitempty"`
	ParseMode             ParseMode           `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity     `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia bool                `json:"show_caption_above_media,omitempty"`
	VideoWidth            int                 `json:"video_width,omitempty"`
	VideoHeight           int                 `json:"video_height,omitempty"`
	VideoDuration         int                 `json:"video_duration,omitempty"`
	Description           string              `json:"description,omitempty"`
	ReplyMarkup           ReplyMarkup         `json:"reply_markup,omitempty"`
	InputMessageContent   InputMessageContent `json:"input_message_content,omitempty"`
}

func (*InlineQueryResultVideo) inlineQueryResultTag() {}

func (m *InlineQueryResultVideo) MarshalCustom() ([]byte, error) {
	ret := struct {
		Type string `json:"type"`
		*InlineQueryResultVideo
	}{
		Type:                   "video",
		InlineQueryResultVideo: m,
	}

	return json.Marshal(&ret)
}

// InlineQueryResultAudio https://core.telegram.org/bots/api#inlinequeryresultaudio
type InlineQueryResultAudio struct {
	ID                  string              `json:"id"`
	AudioURL            string              `json:"audio_url"`
	Title               string              `json:"title,omitempty"`
	Caption             string              `json:"caption,omitempty"`
	ParseMode           ParseMode           `json:"parse_mode,omitempty"`
	CaptionEntities     []MessageEntity     `json:"caption_entities,omitempty"`
	Performer           string              `json:"performer,omitempty"`
	AudioDuration       int                 `json:"audio_duration,omitempty"`
	ReplyMarkup         ReplyMarkup         `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

func (*InlineQueryResultAudio) inlineQueryResultTag() {}

func (m *InlineQueryResultAudio) MarshalCustom() ([]byte, error) {
	ret := struct {
		Type string `json:"type"`
		*InlineQueryResultAudio
	}{
		Type:                   "audio",
		InlineQueryResultAudio: m,
	}

	return json.Marshal(&ret)
}

// InlineQueryResultVoice https://core.telegram.org/bots/api#inlinequeryresultvoice
type InlineQueryResultVoice struct {
	ID                  string              `json:"id"`
	VoiceURL            string              `json:"voice_url"`
	Title               string              `json:"title,omitempty"`
	Caption             string              `json:"caption,omitempty"`
	ParseMode           ParseMode           `json:"parse_mode,omitempty"`
	CaptionEntities     []MessageEntity     `json:"caption_entities,omitempty"`
	VoiceDuration       int                 `json:"voice_duration,omitempty"`
	ReplyMarkup         ReplyMarkup         `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

func (*InlineQueryResultVoice) inlineQueryResultTag() {}

func (m *InlineQueryResultVoice) MarshalCustom() ([]byte, error) {
	ret := struct {
		Type string `json:"type"`
		*InlineQueryResultVoice
	}{
		Type:                   "voice",
		InlineQueryResultVoice: m,
	}

	return json.Marshal(&ret)
}

// InlineQueryResultDocument https://core.telegram.org/bots/api#inlinequeryresultdocument
type InlineQueryResultDocument struct {
	ID                  string              `json:"id"`
	Title               string              `json:"title,omitempty"`
	Caption             string              `json:"caption,omitempty"`
	ParseMode           ParseMode           `json:"parse_mode,omitempty"`
	CaptionEntities     []MessageEntity     `json:"caption_entities,omitempty"`
	DocumentURL         string              `json:"document_url"`
	MimeType            string              `json:"mime_type"`
	Description         string              `json:"description,omitempty"`
	ReplyMarkup         ReplyMarkup         `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
	ThumbnailURL        string              `json:"thumbnail_url,omitempty"`
	ThumbnailWidth      int                 `json:"thumbnail_width,omitempty"`
	ThumbnailHeight     int                 `json:"thumbnail_height,omitempty"`
}

func (*InlineQueryResultDocument) inlineQueryResultTag() {}

func (m *InlineQueryResultDocument) MarshalCustom() ([]byte, error) {
	ret := struct {
		Type string `json:"type"`
		*InlineQueryResultDocument
	}{
		Type:                      "document",
		InlineQueryResultDocument: m,
	}

	return json.Marshal(&ret)
}

// InlineQueryResultLocation https://core.telegram.org/bots/api#inlinequeryresultlocation
type InlineQueryResultLocation struct {
	ID                   string              `json:"id"`
	Latitude             float64             `json:"latitude"`
	Longitude            float64             `json:"longitude"`
	Title                string              `json:"title,omitempty"`
	HorizontalAccuracy   float64             `json:"horizontal_accuracy,omitempty"`
	LivePeriod           int                 `json:"live_period,omitempty"`
	Heading              int                 `json:"heading,omitempty"`
	ProximityAlertRadius int                 `json:"proximity_alert_radius,omitempty"`
	ReplyMarkup          ReplyMarkup         `json:"reply_markup,omitempty"`
	InputMessageContent  InputMessageContent `json:"input_message_content,omitempty"`
	ThumbnailURL         string              `json:"thumbnail_url,omitempty"`
	ThumbnailWidth       int                 `json:"thumbnail_width,omitempty"`
	ThumbnailHeight      int                 `json:"thumbnail_height,omitempty"`
}

func (*InlineQueryResultLocation) inlineQueryResultTag() {}

func (m *InlineQueryResultLocation) MarshalCustom() ([]byte, error) {
	ret := struct {
		Type string `json:"type"`
		*InlineQueryResultLocation
	}{
		Type:                      "location",
		InlineQueryResultLocation: m,
	}

	return json.Marshal(&ret)
}

// InlineQueryResultVenue https://core.telegram.org/bots/api#inlinequeryresultvenue
type InlineQueryResultVenue struct {
	ID                  string              `json:"id"`
	Latitude            float64             `json:"latitude"`
	Longitude           float64             `json:"longitude"`
	Title               string              `json:"title,omitempty"`
	Address             string              `json:"address"`
	FoursquareID        string              `json:"foursquare_id,omitempty"`
	FoursquareType      string              `json:"foursquare_type,omitempty"`
	GooglePlaceID       string              `json:"google_place_id,omitempty"`
	GooglePlaceType     string              `json:"google_place_type,omitempty"`
	ReplyMarkup         ReplyMarkup         `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
	ThumbnailURL        string              `json:"thumbnail_url,omitempty"`
	ThumbnailWidth      int                 `json:"thumbnail_width,omitempty"`
	ThumbnailHeight     int                 `json:"thumbnail_height,omitempty"`
}

func (*InlineQueryResultVenue) inlineQueryResultTag() {}

func (m *InlineQueryResultVenue) MarshalCustom() ([]byte, error) {
	ret := struct {
		Type string `json:"type"`
		*InlineQueryResultVenue
	}{
		Type:                   "venue",
		InlineQueryResultVenue: m,
	}

	return json.Marshal(&ret)
}

// InlineQueryResultContact https://core.telegram.org/bots/api#inlinequeryresultcontact
type InlineQueryResultContact struct {
	ID                  string              `json:"id"`
	PhoneNumber         string              `json:"phone_number"`
	FirstName           string              `json:"first_name"`
	LastName            string              `json:"last_name,omitempty"`
	VCard               string              `json:"vcard,omitempty"`
	ReplyMarkup         ReplyMarkup         `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
	ThumbnailURL        string              `json:"thumbnail_url,omitempty"`
	ThumbnailWidth      int                 `json:"thumbnail_width,omitempty"`
	ThumbnailHeight     int                 `json:"thumbnail_height,omitempty"`
}

func (*InlineQueryResultContact) inlineQueryResultTag() {}

func (m *InlineQueryResultContact) MarshalCustom() ([]byte, error) {
	ret := struct {
		Type string `json:"type"`
		*InlineQueryResultContact
	}{
		Type:                     "contact",
		InlineQueryResultContact: m,
	}

	return json.Marshal(&ret)
}

// InlineQueryResultGame https://core.telegram.org/bots/api#inlinequeryresultgame
type InlineQueryResultGame struct {
	ID            string      `json:"id"`
	GameShortName string      `json:"game_short_name"`
	ReplyMarkup   ReplyMarkup `json:"reply_markup,omitempty"`
}

func (*InlineQueryResultGame) inlineQueryResultTag() {}

func (m *InlineQueryResultGame) MarshalCustom() ([]byte, error) {
	ret := struct {
		Type string `json:"type"`
		*InlineQueryResultGame
	}{
		Type:                  "game",
		InlineQueryResultGame: m,
	}

	return json.Marshal(&ret)
}

// InlineQueryResultCachedPhoto https://core.telegram.org/bots/api#inlinequeryresultcachedphoto
type InlineQueryResultCachedPhoto struct {
	ID                    string              `json:"id"`
	PhotoFileID           string              `json:"photo_file_id"`
	Title                 string              `json:"title,omitempty"`
	Description           string              `json:"description,omitempty"`
	Caption               string              `json:"caption,omitempty"`
	ParseMode             ParseMode           `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity     `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia bool                `json:"show_caption_above_media,omitempty"`
	ReplyMarkup           ReplyMarkup         `json:"reply_markup,omitempty"`
	InputMessageContent   InputMessageContent `json:"input_message_content,omitempty"`
}

func (*InlineQueryResultCachedPhoto) inlineQueryResultTag() {}

func (m *InlineQueryResultCachedPhoto) MarshalCustom() ([]byte, error) {
	ret := struct {
		Type string `json:"type"`
		*InlineQueryResultCachedPhoto
	}{
		Type:                         "photo",
		InlineQueryResultCachedPhoto: m,
	}

	return json.Marshal(&ret)
}

// InlineQueryResultCachedGif https://core.telegram.org/bots/api#inlinequeryresultcachedgif
type InlineQueryResultCachedGif struct {
	ID                    string              `json:"id"`
	GifFileID             string              `json:"gif_file_id"`
	Title                 string              `json:"title,omitempty"`
	Caption               string              `json:"caption,omitempty"`
	ParseMode             ParseMode           `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity     `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia bool                `json:"show_caption_above_media,omitempty"`
	ReplyMarkup           ReplyMarkup         `json:"reply_markup,omitempty"`
	InputMessageContent   InputMessageContent `json:"input_message_content,omitempty"`
}

func (*InlineQueryResultCachedGif) inlineQueryResultTag() {}

func (m *InlineQueryResultCachedGif) MarshalCustom() ([]byte, error) {
	ret := struct {
		Type string `json:"type"`
		*InlineQueryResultCachedGif
	}{
		Type:                       "gif",
		InlineQueryResultCachedGif: m,
	}

	return json.Marshal(&ret)
}

// InlineQueryResultCachedMpeg4Gif https://core.telegram.org/bots/api#inlinequeryresultcachedmpeg4gif
type InlineQueryResultCachedMpeg4Gif struct {
	ID                    string              `json:"id"`
	Mpeg4FileID           string              `json:"mpeg4_file_id"`
	Title                 string              `json:"title,omitempty"`
	Caption               string              `json:"caption,omitempty"`
	ParseMode             ParseMode           `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity     `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia bool                `json:"show_caption_above_media,omitempty"`
	ReplyMarkup           ReplyMarkup         `json:"reply_markup,omitempty"`
	InputMessageContent   InputMessageContent `json:"input_message_content,omitempty"`
}

func (*InlineQueryResultCachedMpeg4Gif) inlineQueryResultTag() {}

func (m *InlineQueryResultCachedMpeg4Gif) MarshalCustom() ([]byte, error) {
	ret := struct {
		Type string `json:"type"`
		*InlineQueryResultCachedMpeg4Gif
	}{
		Type:                            "mpeg4_gif",
		InlineQueryResultCachedMpeg4Gif: m,
	}

	return json.Marshal(&ret)
}

// InlineQueryResultCachedSticker https://core.telegram.org/bots/api#inlinequeryresultcachedsticker
type InlineQueryResultCachedSticker struct {
	ID                  string              `json:"id"`
	StickerFileID       string              `json:"sticker_file_id"`
	ReplyMarkup         ReplyMarkup         `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

func (*InlineQueryResultCachedSticker) inlineQueryResultTag() {}

func (m *InlineQueryResultCachedSticker) MarshalCustom() ([]byte, error) {
	ret := struct {
		Type string `json:"type"`
		*InlineQueryResultCachedSticker
	}{
		Type:                           "sticker",
		InlineQueryResultCachedSticker: m,
	}

	return json.Marshal(&ret)
}

// InlineQueryResultCachedDocument https://core.telegram.org/bots/api#inlinequeryresultcacheddocument
type InlineQueryResultCachedDocument struct {
	ID                  string              `json:"id"`
	Title               string              `json:"title,omitempty"`
	DocumentFileID      string              `json:"document_file_id"`
	Description         string              `json:"description,omitempty"`
	Caption             string              `json:"caption,omitempty"`
	ParseMode           ParseMode           `json:"parse_mode,omitempty"`
	CaptionEntities     []MessageEntity     `json:"caption_entities,omitempty"`
	ReplyMarkup         ReplyMarkup         `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

func (*InlineQueryResultCachedDocument) inlineQueryResultTag() {}

func (m *InlineQueryResultCachedDocument) MarshalCustom() ([]byte, error) {
	ret := struct {
		Type string `json:"type"`
		*InlineQueryResultCachedDocument
	}{
		Type:                            "document",
		InlineQueryResultCachedDocument: m,
	}

	return json.Marshal(&ret)
}

// InlineQueryResultCachedVideo https://core.telegram.org/bots/api#inlinequeryresultcachedvideo
type InlineQueryResultCachedVideo struct {
	ID                    string              `json:"id"`
	VideoFileID           string              `json:"video_file_id"`
	Title                 string              `json:"title,omitempty"`
	Description           string              `json:"description,omitempty"`
	Caption               string              `json:"caption,omitempty"`
	ParseMode             ParseMode           `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity     `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia bool                `json:"show_caption_above_media,omitempty"`
	ReplyMarkup           ReplyMarkup         `json:"reply_markup,omitempty"`
	InputMessageContent   InputMessageContent `json:"input_message_content,omitempty"`
}

func (*InlineQueryResultCachedVideo) inlineQueryResultTag() {}

func (m *InlineQueryResultCachedVideo) MarshalCustom() ([]byte, error) {
	ret := struct {
		Type string `json:"type"`
		*InlineQueryResultCachedVideo
	}{
		Type:                         "video",
		InlineQueryResultCachedVideo: m,
	}

	return json.Marshal(&ret)
}

// InlineQueryResultCachedVoice https://core.telegram.org/bots/api#inlinequeryresultcachedvoice
type InlineQueryResultCachedVoice struct {
	ID                  string              `json:"id"`
	VoiceFileID         string              `json:"voice_file_id"`
	Title               string              `json:"title,omitempty"`
	Caption             string              `json:"caption,omitempty"`
	ParseMode           ParseMode           `json:"parse_mode,omitempty"`
	CaptionEntities     []MessageEntity     `json:"caption_entities,omitempty"`
	ReplyMarkup         ReplyMarkup         `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

func (*InlineQueryResultCachedVoice) inlineQueryResultTag() {}

func (m *InlineQueryResultCachedVoice) MarshalCustom() ([]byte, error) {
	ret := struct {
		Type string `json:"type"`
		*InlineQueryResultCachedVoice
	}{
		Type:                         "voice",
		InlineQueryResultCachedVoice: m,
	}

	return json.Marshal(&ret)
}

// InlineQueryResultCachedAudio https://core.telegram.org/bots/api#inlinequeryresultcachedaudio
type InlineQueryResultCachedAudio struct {
	ID                  string              `json:"id"`
	AudioFileID         string              `json:"audio_file_id"`
	Caption             string              `json:"caption,omitempty"`
	ParseMode           ParseMode           `json:"parse_mode,omitempty"`
	CaptionEntities     []MessageEntity     `json:"caption_entities,omitempty"`
	ReplyMarkup         ReplyMarkup         `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

func (*InlineQueryResultCachedAudio) inlineQueryResultTag() {}

func (m *InlineQueryResultCachedAudio) MarshalCustom() ([]byte, error) {
	ret := struct {
		Type string `json:"type"`
		*InlineQueryResultCachedAudio
	}{
		Type:                         "audio",
		InlineQueryResultCachedAudio: m,
	}

	return json.Marshal(&ret)
}

// InputMessageContent https://core.telegram.org/bots/api#inputmessagecontent
type InputMessageContent interface {
	inputMessageContentTag()
}

// InputTextMessageContent https://core.telegram.org/bots/api#inputtextmessagecontent
type InputTextMessageContent struct {
	MessageText        string              `json:"message_text"`
	ParseMode          ParseMode           `json:"parse_mode,omitempty"`
	Entities           []MessageEntity     `json:"entities,omitempty"`
	LinkPreviewOptions *LinkPreviewOptions `json:"link_preview_options,omitempty"`
}

func (InputTextMessageContent) inputMessageContentTag() {}

// InputLocationMessageContent https://core.telegram.org/bots/api#inputlocationmessagecontent
type InputLocationMessageContent struct {
	Latitude             float64 `json:"latitude"`
	Longitude            float64 `json:"longitude"`
	HorizontalAccuracy   float64 `json:"horizontal_accuracy,omitempty"`
	LivePeriod           int     `json:"live_period,omitempty"`
	Heading              int     `json:"heading,omitempty"`
	ProximityAlertRadius int     `json:"proximity_alert_radius,omitempty"`
}

func (InputLocationMessageContent) inputMessageContentTag() {}

// InputVenueMessageContent https://core.telegram.org/bots/api#inputvenuemessagecontent
type InputVenueMessageContent struct {
	Latitude        float64 `json:"latitude"`
	Longitude       float64 `json:"longitude"`
	Title           string  `json:"title"`
	Address         string  `json:"address"`
	FoursquareID    string  `json:"foursquare_id,omitempty"`
	FoursquareType  string  `json:"foursquare_type,omitempty"`
	GooglePlaceID   string  `json:"google_place_id,omitempty"`
	GooglePlaceType string  `json:"google_place_type,omitempty"`
}

func (InputVenueMessageContent) inputMessageContentTag() {}

// InputContactMessageContent https://core.telegram.org/bots/api#inputcontactmessagecontent
type InputContactMessageContent struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name,omitempty"`
	VCard       string `json:"vcard,omitempty"`
}

func (InputContactMessageContent) inputMessageContentTag() {}

// LabeledPrice https://core.telegram.org/bots/api#labeledprice
type LabeledPrice struct {
	Label  string `json:"label"`
	Amount int    `json:"amount"`
}

// InputInvoiceMessageContent https://core.telegram.org/bots/api#inputinvoicemessagecontent
type InputInvoiceMessageContent struct {
	Title                     string         `json:"title"`
	Description               string         `json:"description"`
	Payload                   string         `json:"payload"`
	ProviderToken             string         `json:"provider_token,omitempty"`
	Currency                  string         `json:"currency"`
	Prices                    []LabeledPrice `json:"prices"`
	MaxTipAmount              int            `json:"max_tip_amount,omitempty"`
	SuggestedTipAmounts       []int          `json:"suggested_tip_amounts,omitempty"`
	ProviderData              string         `json:"provider_data,omitempty"`
	PhotoURL                  string         `json:"photo_url,omitempty"`
	PhotoSize                 int            `json:"photo_size,omitempty"`
	PhotoWidth                int            `json:"photo_width,omitempty"`
	PhotoHeight               int            `json:"photo_height,omitempty"`
	NeedName                  bool           `json:"need_name,omitempty"`
	NeedPhoneNumber           bool           `json:"need_phone_number,omitempty"`
	NeedEmail                 bool           `json:"need_email,omitempty"`
	NeedShippingAddress       bool           `json:"need_shipping_address,omitempty"`
	SendPhoneNumberToProvider bool           `json:"send_phone_number_to_provider,omitempty"`
	SendEmailToProvider       bool           `json:"send_email_to_provider,omitempty"`
	IsFlexible                bool           `json:"is_flexible,omitempty"`
}

func (InputInvoiceMessageContent) inputMessageContentTag() {}
