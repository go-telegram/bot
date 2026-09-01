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

// A thumbnail left as a typed nil pointer is not a thumbnail: omitempty cannot see it
// through the interface, so it has to be normalised or the request carries a null
// where the Bot API expects a string.
func TestInputMedia_MarshalTypedNilThumbnailOmitted(t *testing.T) {
	var upload *InputFileUpload

	tests := map[string]interface {
		MarshalInputMedia() ([]byte, error)
	}{
		"video":      &InputMediaVideo{Media: "file_id", Thumbnail: upload},
		"animation":  &InputMediaAnimation{Media: "file_id", Thumbnail: upload},
		"audio":      &InputMediaAudio{Media: "file_id", Thumbnail: upload},
		"document":   &InputMediaDocument{Media: "file_id", Thumbnail: upload},
		"paid video": &InputPaidMediaVideo{Media: "file_id", Thumbnail: upload},
	}

	for name, media := range tests {
		t.Run(name, func(t *testing.T) {
			data, err := media.MarshalInputMedia()
			if err != nil {
				t.Fatal(err)
			}
			if strings.Contains(string(data), "thumbnail") {
				t.Fatalf("thumbnail not omitted: %s", data)
			}
		})
	}
}

func TestInputMedia_MarshalTypedNilInputFileStringThumbnailOmitted(t *testing.T) {
	var file *InputFileString

	data, err := (&InputMediaVideo{Media: "file_id", Thumbnail: file}).MarshalInputMedia()
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(data), "thumbnail") {
		t.Fatalf("thumbnail not omitted: %s", data)
	}
}
