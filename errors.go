package bot

import "errors"

var (
	ErrEmptyChatID     = errors.New("empty chat id")
	ErrEmptyText       = errors.New("empty text")
	ErrEmptyFromChatID = errors.New("empty from chat id")
	ErrEmptyMessageID  = errors.New("empty message id")
	ErrEmptyPhoto      = errors.New("empty photo")
	ErrEmptyAudio      = errors.New("empty audio")
	ErrEmptyDocument   = errors.New("empty document")
	ErrEmptyVideo      = errors.New("empty video")
	ErrEmptyAnimation  = errors.New("empty animation")
	ErrEmptyVoice      = errors.New("empty voice")
	ErrEmptyVideoNote  = errors.New("empty video note")
	ErrEmptyMedia      = errors.New("empty media")
	ErrEmptyLatitude   = errors.New("empty latitude")
	ErrEmptyLongitude  = errors.New("empty longitude")
)
