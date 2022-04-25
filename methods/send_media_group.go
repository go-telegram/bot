package methods

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// https://core.telegram.org/bots/api#sendmediagroup

type InputMedia interface {
	inputMediaTag()
}

type SendMediaGroupParams struct {
	ChatID                   string       `json:"chat_id"`
	Media                    []InputMedia `json:"media"`
	DisableNotification      bool         `json:"disable_notification,omitempty"`
	ProtectContent           bool         `json:"protect_content,omitempty"`
	ReplyToMessageID         int          `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool         `json:"allow_sending_without_reply,omitempty"`
}

func (p SendMediaGroupParams) Validate() error {
	if p.ChatID == "" {
		return bot.ErrEmptyChatID
	}
	if len(p.Media) == 0 {
		return bot.ErrEmptyMedia
	}
	return nil
}

func SendMediaGroup(ctx context.Context, b *bot.Bot, params *SendMediaGroupParams) ([]*models.Message, error) {
	var res []*models.Message

	err := bot.RawRequest(ctx, b, "sendMediaGroup", params, &res)

	return res, err
}

type InputMediaPhoto struct {
	Type            string                 `json:"type"`
	Media           string                 `json:"media"`
	Caption         string                 `json:"caption,omitempty"`
	ParseMode       models.ParseMode       `json:"parse_mode,omitempty"`
	CaptionEntities []models.MessageEntity `json:"caption_entities,omitempty"`
}

func (InputMediaPhoto) inputMediaTag() {}

type InputMediaVideo struct {
	Type              string                 `json:"type"`
	Media             string                 `json:"media"`
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

type InputMediaAnimation struct {
	Type            string                 `json:"type"`
	Media           string                 `json:"media"`
	Thumb           string                 `json:"thumb,omitempty"`
	Caption         string                 `json:"caption,omitempty"`
	ParseMode       models.ParseMode       `json:"parse_mode,omitempty"`
	CaptionEntities []models.MessageEntity `json:"caption_entities,omitempty"`
	Width           int                    `json:"width,omitempty"`
	Height          int                    `json:"height,omitempty"`
	Duration        int                    `json:"duration,omitempty"`
}

func (InputMediaAnimation) inputMediaTag() {}

type InputMediaAudio struct {
	Type            string                 `json:"type"`
	Media           string                 `json:"media"`
	Thumb           string                 `json:"thumb,omitempty"`
	Caption         string                 `json:"caption,omitempty"`
	ParseMode       models.ParseMode       `json:"parse_mode,omitempty"`
	CaptionEntities []models.MessageEntity `json:"caption_entities,omitempty"`
	Duration        int                    `json:"duration,omitempty"`
	Performer       string                 `json:"performer,omitempty"`
	Title           string                 `json:"title,omitempty"`
}

func (InputMediaAudio) inputMediaTag() {}

type InputMediaDocument struct {
	Type                        string                 `json:"type"`
	Media                       string                 `json:"media"`
	Thumb                       string                 `json:"thumb,omitempty"`
	Caption                     string                 `json:"caption,omitempty"`
	ParseMode                   models.ParseMode       `json:"parse_mode,omitempty"`
	CaptionEntities             []models.MessageEntity `json:"caption_entities,omitempty"`
	DisableContentTypeDetection bool                   `json:"disable_content_type_detection,omitempty"`
}

func (InputMediaDocument) inputMediaTag() {}
