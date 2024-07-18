package models

import "testing"

func TestPaidMedia_UnmarshalJSON_Preview(t *testing.T) {
	src := `{"type":"preview","width":42,"height":43,"duration":44}`

	p := &PaidMedia{}
	err := p.UnmarshalJSON([]byte(src))
	if err != nil {
		t.Fatal(err)
	}

	if p.Type != PaidMediaTypePreview {
		t.Fatal("wrong type")
	}

	if p.Preview == nil {
		t.Fatal("Preview is nil")
	}

	if p.Preview.Width != 42 {
		t.Fatal("wrong width")
	}

	if p.Preview.Height != 43 {
		t.Fatal("wrong height")
	}

	if p.Preview.Duration != 44 {
		t.Fatal("wrong duration")
	}
}

func TestPaidMedia_UnmarshalJSON_Photo(t *testing.T) {
	src := `{"type":"photo","photo":[{"width":42,"height":43}]}`

	p := &PaidMedia{}
	err := p.UnmarshalJSON([]byte(src))
	if err != nil {
		t.Fatal(err)
	}

	if p.Type != PaidMediaTypePhoto {
		t.Fatal("wrong type")
	}

	if p.Photo == nil {
		t.Fatal("Photo is nil")
	}

	if len(p.Photo.Photo) != 1 {
		t.Fatal("wrong photo length")
	}

	if p.Photo.Photo[0].Width != 42 {
		t.Fatal("wrong width")
	}
	if p.Photo.Photo[0].Height != 43 {
		t.Fatal("wrong height")
	}
}

func TestPaidMedia_UnmarshalJSON_Video(t *testing.T) {
	src := `{"type":"video","video":{"width":42,"height":43,"duration":44}}`

	p := &PaidMedia{}
	err := p.UnmarshalJSON([]byte(src))
	if err != nil {
		t.Fatal(err)
	}

	if p.Type != PaidMediaTypeVideo {
		t.Fatal("wrong type")
	}

	if p.Video == nil {
		t.Fatal("Video is nil")
	}

	if p.Video.Video.Width != 42 {
		t.Fatal("wrong width")
	}
	if p.Video.Video.Height != 43 {
		t.Fatal("wrong height")
	}
	if p.Video.Video.Duration != 44 {
		t.Fatal("wrong duration")
	}
}
