package models

import (
	"encoding/json"
	"fmt"
)

// BackgroundType https://core.telegram.org/bots/api#backgroundtype
type BackgroundType string

const (
	ChatBackgroundTypeFill      BackgroundType = "fill"
	ChatBackgroundTypeWallpaper BackgroundType = "wallpaper"
	ChatBackgroundTypePattern   BackgroundType = "pattern"
	ChatBackgroundTypeChatTheme BackgroundType = "chat_theme"
)

// ChatBackground https://core.telegram.org/bots/api#chatbackground
type ChatBackground struct {
	Type      BackgroundType
	Fill      *BackgroundTypeFill
	Wallpaper *BackgroundTypeWallpaper
	Pattern   *BackgroundTypePattern
	Theme     *BackgroundTypeChatTheme
}

func (cb *ChatBackground) UnmarshalJSON(data []byte) error {
	v := struct {
		Type string `json:"type"`
	}{}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	switch v.Type {
	case "fill":
		cb.Type = ChatBackgroundTypeFill
		cb.Fill = &BackgroundTypeFill{}
		return json.Unmarshal(data, cb.Fill)
	case "wallpaper":
		cb.Type = ChatBackgroundTypeWallpaper
		cb.Wallpaper = &BackgroundTypeWallpaper{}
		return json.Unmarshal(data, cb.Wallpaper)
	case "pattern":
		cb.Type = ChatBackgroundTypePattern
		cb.Pattern = &BackgroundTypePattern{}
		return json.Unmarshal(data, cb.Pattern)
	case "chat_theme":
		cb.Type = ChatBackgroundTypeChatTheme
		cb.Theme = &BackgroundTypeChatTheme{}
		return json.Unmarshal(data, cb.Theme)
	}

	return fmt.Errorf("unsupported ChatBackground type")
}

// BackgroundTypeFill https://core.telegram.org/bots/api#backgroundtypefill
type BackgroundTypeFill struct {
	Fill             BackgroundFill `json:"fill"`
	DarkThemeDimming int            `json:"dark_theme_dimming"`
}

// BackgroundTypeWallpaper https://core.telegram.org/bots/api#backgroundtypewallpaper
type BackgroundTypeWallpaper struct {
	Document         Document `json:"document"`
	DarkThemeDimming int      `json:"dark_theme_dimming"`
	IsBlurred        bool     `json:"is_blurred,omitempty"`
	IsMoving         bool     `json:"is_moving,omitempty"`
}

// BackgroundTypePattern https://core.telegram.org/bots/api#backgroundtypepattern
type BackgroundTypePattern struct {
	Document   Document       `json:"document"`
	Fill       BackgroundFill `json:"fill"`
	Intensity  int            `json:"intensity"`
	IsInverted bool           `json:"is_inverted,omitempty"`
	IsMoving   bool           `json:"is_moving,omitempty"`
}

// BackgroundTypeChatTheme https://core.telegram.org/bots/api#backgroundtypechattheme
type BackgroundTypeChatTheme struct {
	ThemeName string `json:"theme_name"`
}

type BackgroundFillType string

const (
	BackgroundFillTypeSolid            BackgroundFillType = "solid"
	BackgroundFillTypeGradient         BackgroundFillType = "gradient"
	BackgroundFillTypeFreeformGradient BackgroundFillType = "freeform_gradient"
)

type BackgroundFill struct {
	Type             BackgroundFillType `json:"type"`
	Solid            *BackgroundFillSolid
	Gradient         *BackgroundFillGradient
	FreeformGradient *BackgroundFillFreeformGradient
}

func (bf *BackgroundFill) UnmarshalJSON(data []byte) error {
	v := struct {
		Type string `json:"type"`
	}{}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	switch v.Type {
	case "solid":
		bf.Type = BackgroundFillTypeSolid
		bf.Solid = &BackgroundFillSolid{}
		return json.Unmarshal(data, bf.Solid)
	case "gradient":
		bf.Type = BackgroundFillTypeGradient
		bf.Gradient = &BackgroundFillGradient{}
		return json.Unmarshal(data, bf.Gradient)
	case "freeform_gradient":
		bf.Type = BackgroundFillTypeFreeformGradient
		bf.FreeformGradient = &BackgroundFillFreeformGradient{}
		return json.Unmarshal(data, bf.FreeformGradient)
	}

	return fmt.Errorf("unsupported BackgroundFill type")
}

// BackgroundFillSolid https://core.telegram.org/bots/api#backgroundfillsolid
type BackgroundFillSolid struct {
	Color int `json:"color"`
}

// BackgroundFillGradient https://core.telegram.org/bots/api#backgroundfillgradient
type BackgroundFillGradient struct {
	TopColor      int `json:"top_color"`
	BottomColor   int `json:"bottom_color"`
	RotationAngle int `json:"rotation_angle"`
}

// BackgroundFillFreeformGradient https://core.telegram.org/bots/api#backgroundfillfreeformgradient
type BackgroundFillFreeformGradient struct {
	Colors []int `json:"colors"`
}
