package methods

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type GetFileParams struct {
	FileID string `json:"file_id" rules:"required"`
}

// GetFile https://core.telegram.org/bots/api#getfile
func GetFile(ctx context.Context, b *bot.Bot, params *GetFileParams) (*models.File, error) {
	result := &models.File{}

	err := bot.RawRequest(ctx, b, "getFile", params, result)

	return result, err
}
