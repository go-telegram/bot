package models

import (
	"encoding/json"
	"testing"
)

func TestChatBackground_UnmarshalJSON_fill(t *testing.T) {
	src := `{"type":"fill","fill":{"type":"solid","color":123},"dark_theme_dimming":2}`

	var cb ChatBackground
	err := json.Unmarshal([]byte(src), &cb)
	if err != nil {
		t.Fatal(err)
	}

	if cb.Type != ChatBackgroundTypeFill {
		t.Fatal("invalid type")
	}

	if cb.Fill == nil {
		t.Fatal("invalid Fill")
	}

	if cb.Fill.Fill.Type != BackgroundFillTypeSolid {
		t.Fatal("invalid fill type")
	}

	if cb.Fill.Fill.Solid.Color != 123 {
		t.Fatalf("invalid color %d, expect 123", cb.Fill.Fill.Solid.Color)
	}

	if cb.Fill.DarkThemeDimming != 2 {
		t.Fatalf("invalid dark theme dimming %d, expect 2", cb.Fill.DarkThemeDimming)
	}
}

func TestChatBackground_UnmarshalJSON_wallpaper(t *testing.T) {
	src := `{"type":"wallpaper","document":{"file_id":"test","file_unique_id":"test2"},"dark_theme_dimming":2,"is_blurred":true,"is_moving":true}`

	var cb ChatBackground
	err := json.Unmarshal([]byte(src), &cb)
	if err != nil {
		t.Fatal(err)
	}

	if cb.Type != ChatBackgroundTypeWallpaper {
		t.Fatal("invalid type")
	}

	if cb.Wallpaper == nil {
		t.Fatal("invalid Wallpaper")
	}

	if cb.Wallpaper.Document.FileID != "test" {
		t.Fatal("invalid document file id")
	}
}

func TestChatBackground_UnmarshalJSON_pattern(t *testing.T) {
	src := `{"type":"pattern","document":{"file_id":"test","file_unique_id":"test","file_size":123,"file_path":"test"},"fill":{"type":"solid","solid":{"color":123}},"intensity":1,"is_inverted":true,"is_moving":true}`

	var cb ChatBackground
	err := json.Unmarshal([]byte(src), &cb)
	if err != nil {
		t.Fatal(err)
	}

	if cb.Type != ChatBackgroundTypePattern {
		t.Fatal("invalid type")
	}

	if cb.Pattern == nil {
		t.Fatal("invalid Pattern")
	}

	if cb.Pattern.Document.FileID != "test" {
		t.Fatal("invalid document file id")
	}
}

func TestChatBackground_UnmarshalJSON_chat_theme(t *testing.T) {
	src := `{"type":"chat_theme","theme_name":"test"}`

	var cb ChatBackground
	err := json.Unmarshal([]byte(src), &cb)
	if err != nil {
		t.Fatal(err)
	}

	if cb.Type != ChatBackgroundTypeChatTheme {
		t.Fatal("invalid type")
	}

	if cb.Theme == nil {
		t.Fatal("invalid Theme")
	}

	if cb.Theme.ThemeName != "test" {
		t.Fatal("invalid theme name")
	}
}
