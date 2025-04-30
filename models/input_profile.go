package models

import (
	"encoding/json"
	"io"
)

// InputProfilePhoto https://core.telegram.org/bots/api#inputprofilephoto
type InputProfilePhoto interface {
	inputProfilePhotoTag()
}

// InputProfilePhotoStatic https://core.telegram.org/bots/api#inputprofilephotostatic
type InputProfilePhotoStatic struct {
	Photo string `json:"photo"`

	MediaAttachment io.Reader `json:"-"`
}

func (i InputProfilePhotoStatic) GetMedia() string      { return i.Photo }
func (i InputProfilePhotoStatic) Attachment() io.Reader { return i.MediaAttachment }
func (InputProfilePhotoStatic) inputProfilePhotoTag()   {}

func (i InputProfilePhotoStatic) MarshalInputMedia() ([]byte, error) {
	ret := struct {
		Type string `json:"type"`
		InputProfilePhotoStatic
	}{
		Type:                    "static",
		InputProfilePhotoStatic: i,
	}

	return json.Marshal(&ret)
}

// InputProfilePhotoAnimated https://core.telegram.org/bots/api#inputprofilephotoanimated
type InputProfilePhotoAnimated struct {
	Animation          string  `json:"animation"`
	MainFrameTimestamp float64 `json:"main_frame_timestamp,omitempty"`

	MediaAttachment io.Reader `json:"-"`
}

func (i InputProfilePhotoAnimated) GetMedia() string      { return i.Animation }
func (i InputProfilePhotoAnimated) Attachment() io.Reader { return i.MediaAttachment }
func (InputProfilePhotoAnimated) inputProfilePhotoTag()   {}

func (i InputProfilePhotoAnimated) MarshalInputMedia() ([]byte, error) {
	ret := struct {
		Type string `json:"type"`
		InputProfilePhotoAnimated
	}{
		Type:                      "animated",
		InputProfilePhotoAnimated: i,
	}

	return json.Marshal(&ret)
}
