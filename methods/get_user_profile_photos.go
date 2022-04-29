package methods

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type GetUserProfilePhotosParams struct {
	UserID int `json:"user_id" rules:"required"`
	Offset int `json:"offset,omitempty"`
	Limit  int `json:"limit,omitempty"`
}

// GetUserProfilePhotos https://core.telegram.org/bots/api#getuserprofilephotos
func GetUserProfilePhotos(ctx context.Context, b *bot.Bot, params *GetUserProfilePhotosParams) (*models.UserProfilePhotos, error) {
	result := &models.UserProfilePhotos{}

	err := bot.RawRequest(ctx, b, "getUserProfilePhotos", params, result)

	return result, err
}
