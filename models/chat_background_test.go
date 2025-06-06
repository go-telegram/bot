package models

import (
	"encoding/json"
	"testing"
)

func TestChatBackground_UnmarshalJSON_fill(t *testing.T) {
	src := `{"type":{"type":"fill","fill":{"type":"solid","color":123},"dark_theme_dimming":2}}`

	var cb ChatBackground
	err := json.Unmarshal([]byte(src), &cb)
	if err != nil {
		t.Fatal(err)
	}

	if cb.Type.Type != ChatBackgroundTypeFill {
		t.Fatal("invalid type")
	}

	if cb.Type.Fill == nil {
		t.Fatal("invalid Fill")
	}

	if cb.Type.Fill.Fill.Type != BackgroundFillTypeSolid {
		t.Fatal("invalid fill type")
	}

	if cb.Type.Fill.Fill.Solid.Color != 123 {
		t.Fatalf("invalid color %d, expect 123", cb.Type.Fill.Fill.Solid.Color)
	}

	if cb.Type.Fill.DarkThemeDimming != 2 {
		t.Fatalf("invalid dark theme dimming %d, expect 2", cb.Type.Fill.DarkThemeDimming)
	}
}

func TestChatBackground_UnmarshalJSON_wallpaper(t *testing.T) {
	src := `{"type":{"type":"wallpaper","document":{"file_name":"BG_1.jpg","mime_type":"image/jpeg",
	"thumbnail":{"file_id":"test","file_unique_id":"test","file_size":4260,"width":156,"height":320},
	"thumb":{"file_id":"test","file_unique_id":"test","file_size":4260,"width":156,"height":320},
	"file_id":"test","file_unique_id":"test","file_size":299202},"dark_theme_dimming":0}}`

	var cb ChatBackground
	err := json.Unmarshal([]byte(src), &cb)
	if err != nil {
		t.Fatal(err)
	}

	if cb.Type.Type != ChatBackgroundTypeWallpaper {
		t.Fatal("invalid type")
	}

	if cb.Type.Wallpaper == nil {
		t.Fatal("invalid Wallpaper")
	}

	if cb.Type.Wallpaper.Document.FileID != "test" {
		t.Fatal("invalid document file id")
	}
}

func TestChatBackground_UnmarshalJSON_pattern(t *testing.T) {
	src := `{"type":{"type":"pattern","document":{"file_id":"test","file_unique_id":"test","file_size":123,"file_path":"test"},"fill":{"type":"solid","solid":{"color":123}},"intensity":1,"is_inverted":true,"is_moving":true}}`

	var cb ChatBackground
	err := json.Unmarshal([]byte(src), &cb)
	if err != nil {
		t.Fatal(err)
	}

	if cb.Type.Type != ChatBackgroundTypePattern {
		t.Fatal("invalid type")
	}

	if cb.Type.Pattern == nil {
		t.Fatal("invalid Pattern")
	}

	if cb.Type.Pattern.Document.FileID != "test" {
		t.Fatal("invalid document file id")
	}
}

func TestChatBackground_UnmarshalJSON_chat_theme(t *testing.T) {
	src := `{"type":{"type":"chat_theme","theme_name":"test"}}`

	var cb ChatBackground
	err := json.Unmarshal([]byte(src), &cb)
	if err != nil {
		t.Fatal(err)
	}

	if cb.Type.Type != ChatBackgroundTypeChatTheme {
		t.Fatal("invalid type")
	}

	if cb.Type.Theme == nil {
		t.Fatal("invalid Theme")
	}

	if cb.Type.Theme.ThemeName != "test" {
		t.Fatal("invalid theme name")
	}
}
