package models

import (
	"strings"
	"testing"
)

func TestInputFileUpload_MarshalJSON(t *testing.T) {
	// Nested in an InputMedia the value is the attach:// reference Telegram resolves
	// against the uploaded part, not the part name itself.
	d, err := (&InputFileUpload{Filename: "thumb.jpg", Data: strings.NewReader("x")}).MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	if string(d) != `"attach://thumb.jpg"` {
		t.Fatalf("unexpected value %s", d)
	}
}
