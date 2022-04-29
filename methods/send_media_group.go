package methods

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// InputMedia https://core.telegram.org/bots/api#inputmedia
type InputMedia interface {
	inputMediaTag()
}

type SendMediaGroupParams struct {
	ChatID                   any          `json:"chat_id" rules:"required,type:string|int"`
	Media                    []InputMedia `json:"media" rules:"required,min:1"`
	DisableNotification      bool         `json:"disable_notification,omitempty"`
	ProtectContent           bool         `json:"protect_content,omitempty"`
	ReplyToMessageID         int          `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool         `json:"allow_sending_without_reply,omitempty"`
}

// SendMediaGroup https://core.telegram.org/bots/api#sendmediagroup
func SendMediaGroup(ctx context.Context, b *bot.Bot, params *SendMediaGroupParams) ([]*models.Message, error) {
	var res []*models.Message

	err := bot.RawRequest(ctx, b, "sendMediaGroup", params, &res)

	return res, err
}

// InputMediaPhoto https://core.telegram.org/bots/api#inputmediaphoto
type InputMediaPhoto struct {
	Type            string                 `json:"type" rules:"required,equals:photo"`
	Media           string                 `json:"media" rules:"required"`
	Caption         string                 `json:"caption,omitempty"`
	ParseMode       models.ParseMode       `json:"parse_mode,omitempty"`
	CaptionEntities []models.MessageEntity `json:"caption_entities,omitempty"`
}

func (InputMediaPhoto) inputMediaTag() {}

// InputMediaVideo https://core.telegram.org/bots/api#inputmediavideo
type InputMediaVideo struct {
	Type              string                 `json:"type" rules:"required,equals:video"`
	Media             string                 `json:"media" rules:"required"`
	Thumb             string                 `json:"thumb,omitempty"`
	Caption           string                 `json:"caption,omitempty"`
	ParseMode         models.ParseMode       `json:"parse_mode,omitempty"`
	CaptionEntities   []models.MessageEntity `json:"caption_entities,omitempty"`
	Width             int                    `json:"width,omitempty"`
	Height            int                    `json:"height,omitempty"`
	Duration          int                    `json:"duration,omitempty"`
	SupportsStreaming bool                   `json:"supports_streaming,omitempty"`
}

func (InputMediaVideo) inputMediaTag() {}

// InputMediaAnimation https://core.telegram.org/bots/api#inputmediaanimation
type InputMediaAnimation struct {
	Type            string                 `json:"type" rules:"required,equals:animation"`
	Media           string                 `json:"media" rules:"required"`
	Thumb           string                 `json:"thumb,omitempty"`
	Caption         string                 `json:"caption,omitempty"`
	ParseMode       models.ParseMode       `json:"parse_mode,omitempty"`
	CaptionEntities []models.MessageEntity `json:"caption_entities,omitempty"`
	Width           int                    `json:"width,omitempty"`
	Height          int                    `json:"height,omitempty"`
	Duration        int                    `json:"duration,omitempty"`
}

func (InputMediaAnimation) inputMediaTag() {}

// InputMediaAudio https://core.telegram.org/bots/api#inputmediaaudio
type InputMediaAudio struct {
	Type            string                 `json:"type" rules:"required,equals:audio"`
	Media           string                 `json:"media" rules:"required"`
	Thumb           string                 `json:"thumb,omitempty"`
	Caption         string                 `json:"caption,omitempty"`
	ParseMode       models.ParseMode       `json:"parse_mode,omitempty"`
	CaptionEntities []models.MessageEntity `json:"caption_entities,omitempty"`
	Duration        int                    `json:"duration,omitempty"`
	Performer       string                 `json:"performer,omitempty"`
	Title           string                 `json:"title,omitempty"`
}

func (InputMediaAudio) inputMediaTag() {}

// InputMediaDocument https://core.telegram.org/bots/api#inputmediadocument
type InputMediaDocument struct {
	Type                        string                 `json:"type" rules:"required,equals:document"`
	Media                       string                 `json:"media" rules:"required"`
	Thumb                       string                 `json:"thumb,omitempty"`
	Caption                     string                 `json:"caption,omitempty"`
	ParseMode                   models.ParseMode       `json:"parse_mode,omitempty"`
	CaptionEntities             []models.MessageEntity `json:"caption_entities,omitempty"`
	DisableContentTypeDetection bool                   `json:"disable_content_type_detection,omitempty"`
}

func (InputMediaDocument) inputMediaTag() {}
