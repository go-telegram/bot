package models

// InlineQuery https://core.telegram.org/bots/api#inlinequery
type InlineQuery struct {
	ID       string    `json:"id"`
	From     *User     `json:"from"`
	Query    string    `json:"query"`
	Offset   string    `json:"offset"`
	ChatType string    `json:"chat_type,omitempty"`
	Location *Location `json:"location,omitempty"`
}

// InlineQueryResult https://core.telegram.org/bots/api#inlinequeryresult
type InlineQueryResult interface {
	isInlineQueryResult()
}

type InputMessageContent interface {
	isInputMessageContent()
}

type InputTextMessageContent struct {
	MessageText           string          `json:"message_text"`
	ParseMode             ParseMode       `json:"parse_mode,omitempty"`
	Entities              []MessageEntity `json:"entities,omitempty"`
	DisableWebPagePreview bool            `json:"disable_web_page_preview,omitempty"`
}

func (InputTextMessageContent) isInputMessageContent() {}

type InlineQueryResultCachedPhoto struct {
	Type                string              `json:"type"`
	ID                  string              `json:"id"`
	PhotoFileID         string              `json:"photo_file_id"`
	Title               string              `json:"title,omitempty"`
	Description         string              `json:"description,omitempty"`
	Caption             string              `json:"caption,omitempty"`
	ParseMode           ParseMode           `json:"parse_mode,omitempty"`
	CaptionEntities     []MessageEntity     `json:"caption_entities,omitempty"`
	ReplyMarkup         ReplyMarkup         `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

func (InlineQueryResultCachedPhoto) isInlineQueryResult() {}

type InlineQueryResultCachedDocument struct {
	Type                string              `json:"type"`
	ID                  string              `json:"id"`
	Title               string              `json:"title"`
	DocumentFileID      string              `json:"document_file_id"`
	Description         string              `json:"description,omitempty"`
	Caption             string              `json:"caption,omitempty"`
	ParseMode           ParseMode           `json:"parse_mode,omitempty"`
	CaptionEntities     []MessageEntity     `json:"caption_entities,omitempty"`
	ReplyMarkup         ReplyMarkup         `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

func (InlineQueryResultCachedDocument) isInlineQueryResult() {}

type InlineQueryResultDocument struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	Title               string                `json:"title"`
	Caption             string                `json:"caption,omitempty"`
	ParseMode           ParseMode             `json:"parse_mode,omitempty"`
	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
	DocumentURL         string                `json:"document_url,omitempty"`
	MimeType            string                `json:"mime_type,omitempty"`
	Description         string                `json:"description,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
	ThumbURL            string                `json:"thumb_url,omitempty"`
	ThumbWidth          int                   `json:"thumb_width,omitempty"`
	ThumbHeight         int                   `json:"thumb_height,omitempty"`
}

func (InlineQueryResultDocument) isInlineQueryResult() {}

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

// todo:
