package models

import (
	"encoding/json"
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

// The reference is a JSON string: a filename carrying a quote or a backslash must be
// escaped, not concatenated into invalid JSON.
func TestInputFileUpload_MarshalJSONEscapes(t *testing.T) {
	data, err := json.Marshal(&InputFileUpload{Filename: `a"b\c.jpg`})
	if err != nil {
		t.Fatal(err)
	}

	var got string
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatal(err)
	}
	if got != `attach://a"b\c.jpg` {
		t.Fatalf("unexpected reference: %s", got)
	}
}

func TestInputFileString_MarshalJSONEscapes(t *testing.T) {
	data, err := json.Marshal(&InputFileString{Data: `file"id`})
	if err != nil {
		t.Fatal(err)
	}

	var got string
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatal(err)
	}
	if got != `file"id` {
		t.Fatalf("unexpected value: %s", got)
	}
}
