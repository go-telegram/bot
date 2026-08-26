package models

import (
	"strings"
	"testing"
)

func TestInputMedia_GetThumbnail(t *testing.T) {
	thumb := &InputFileUpload{Filename: "thumb.jpg", Data: strings.NewReader("thumb")}

	media := []interface {
		GetThumbnail() InputFile
	}{
		&InputMediaVideo{Thumbnail: thumb},
		&InputMediaAnimation{Thumbnail: thumb},
		&InputMediaAudio{Thumbnail: thumb},
		&InputMediaDocument{Thumbnail: thumb},
		&InputPaidMediaVideo{Thumbnail: thumb},
	}

	for _, m := range media {
		if m.GetThumbnail() != InputFile(thumb) {
			t.Fatalf("%T returned another thumbnail", m)
		}
	}
}

func TestInputMedia_GetThumbnailNotSet(t *testing.T) {
	if (&InputMediaVideo{}).GetThumbnail() != nil {
		t.Fatal("expected no thumbnail")
	}
}
