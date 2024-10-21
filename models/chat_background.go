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
		Type struct {
			Type BackgroundType `json:"type"`
		} `json:"type"`
	}{}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	switch v.Type.Type {
	case ChatBackgroundTypeFill:
		cb.Type = ChatBackgroundTypeFill
		fillContainer := &struct {
			Type BackgroundTypeFill `json:"type"`
		}{}
		err := json.Unmarshal(data, fillContainer)
		if err != nil {
			return err
		}
		cb.Fill = &fillContainer.Type
		return nil
	case ChatBackgroundTypeWallpaper:
		cb.Type = ChatBackgroundTypeWallpaper
		wallpaperContainer := &struct {
			Type BackgroundTypeWallpaper `json:"type"`
		}{}
		err := json.Unmarshal(data, wallpaperContainer)
		if err != nil {
			return err
		}
		cb.Wallpaper = &wallpaperContainer.Type
		return nil
	case ChatBackgroundTypePattern:
		cb.Type = ChatBackgroundTypePattern
		patternContainer := &struct {
			Type BackgroundTypePattern `json:"type"`
		}{}
		err := json.Unmarshal(data, patternContainer)
		if err != nil {
			return err
		}
		cb.Pattern = &patternContainer.Type
		return nil
	case ChatBackgroundTypeChatTheme:
		cb.Type = ChatBackgroundTypeChatTheme
		chatThemeContainer := &struct {
			Type BackgroundTypeChatTheme `json:"type"`
		}{}
		err := json.Unmarshal(data, chatThemeContainer)
		if err != nil {
			return err
		}
		cb.Theme = &chatThemeContainer.Type
		return nil
	}

	return fmt.Errorf("unsupported ChatBackground type")
}

func (cb *ChatBackground) MarshalJSON() ([]byte, error) {
	switch cb.Type {
	case ChatBackgroundTypeFill:
		cb.Fill.Type = ChatBackgroundTypeFill
		return json.Marshal(cb.Fill)
	case ChatBackgroundTypeWallpaper:
		cb.Wallpaper.Type = ChatBackgroundTypeWallpaper
		return json.Marshal(cb.Wallpaper)
	case ChatBackgroundTypePattern:
		cb.Pattern.Type = ChatBackgroundTypePattern
		return json.Marshal(cb.Pattern)
	case ChatBackgroundTypeChatTheme:
		cb.Theme.Type = ChatBackgroundTypeChatTheme
		return json.Marshal(cb.Theme)
	}

	return nil, fmt.Errorf("unsupported ChatBackground type")
}

// BackgroundTypeFill https://core.telegram.org/bots/api#backgroundtypefill
type BackgroundTypeFill struct {
	Type             BackgroundType `json:"type"`
	Fill             BackgroundFill `json:"fill"`
	DarkThemeDimming int            `json:"dark_theme_dimming"`
}

// BackgroundTypeWallpaper https://core.telegram.org/bots/api#backgroundtypewallpaper
type BackgroundTypeWallpaper struct {
	Type             BackgroundType `json:"type"`
	Document         Document       `json:"document"`
	DarkThemeDimming int            `json:"dark_theme_dimming"`
	IsBlurred        bool           `json:"is_blurred,omitempty"`
	IsMoving         bool           `json:"is_moving,omitempty"`
}

// BackgroundTypePattern https://core.telegram.org/bots/api#backgroundtypepattern
type BackgroundTypePattern struct {
	Type       BackgroundType `json:"type"`
	Document   Document       `json:"document"`
	Fill       BackgroundFill `json:"fill"`
	Intensity  int            `json:"intensity"`
	IsInverted bool           `json:"is_inverted,omitempty"`
	IsMoving   bool           `json:"is_moving,omitempty"`
}

// BackgroundTypeChatTheme https://core.telegram.org/bots/api#backgroundtypechattheme
type BackgroundTypeChatTheme struct {
	Type      BackgroundType `json:"type"`
	ThemeName string         `json:"theme_name"`
}

type BackgroundFillType string

const (
	BackgroundFillTypeSolid            BackgroundFillType = "solid"
	BackgroundFillTypeGradient         BackgroundFillType = "gradient"
	BackgroundFillTypeFreeformGradient BackgroundFillType = "freeform_gradient"
)

type BackgroundFill struct {
	Type             BackgroundFillType
	Solid            *BackgroundFillSolid
	Gradient         *BackgroundFillGradient
	FreeformGradient *BackgroundFillFreeformGradient
}

func (bf *BackgroundFill) UnmarshalJSON(data []byte) error {
	v := struct {
		Type BackgroundFillType `json:"type"`
	}{}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	switch v.Type {
	case BackgroundFillTypeSolid:
		bf.Type = BackgroundFillTypeSolid
		bf.Solid = &BackgroundFillSolid{}
		return json.Unmarshal(data, bf.Solid)
	case BackgroundFillTypeGradient:
		bf.Type = BackgroundFillTypeGradient
		bf.Gradient = &BackgroundFillGradient{}
		return json.Unmarshal(data, bf.Gradient)
	case BackgroundFillTypeFreeformGradient:
		bf.Type = BackgroundFillTypeFreeformGradient
		bf.FreeformGradient = &BackgroundFillFreeformGradient{}
		return json.Unmarshal(data, bf.FreeformGradient)
	}

	return fmt.Errorf("unsupported BackgroundFill type")
}

func (bf *BackgroundFill) MarshalJSON() ([]byte, error) {
	switch bf.Type {
	case BackgroundFillTypeSolid:
		bf.Solid.Type = BackgroundFillTypeSolid
		return json.Marshal(bf.Solid)
	case BackgroundFillTypeGradient:
		bf.Gradient.Type = BackgroundFillTypeGradient
		return json.Marshal(bf.Gradient)
	case BackgroundFillTypeFreeformGradient:
		bf.FreeformGradient.Type = BackgroundFillTypeFreeformGradient
		return json.Marshal(bf.FreeformGradient)
	}

	return nil, fmt.Errorf("unsupported BackgroundFill type")
}

// BackgroundFillSolid https://core.telegram.org/bots/api#backgroundfillsolid
type BackgroundFillSolid struct {
	Type  BackgroundFillType `json:"type"`
	Color int                `json:"color"`
}

// BackgroundFillGradient https://core.telegram.org/bots/api#backgroundfillgradient
type BackgroundFillGradient struct {
	Type          BackgroundFillType `json:"type"`
	TopColor      int                `json:"top_color"`
	BottomColor   int                `json:"bottom_color"`
	RotationAngle int                `json:"rotation_angle"`
}

// BackgroundFillFreeformGradient https://core.telegram.org/bots/api#backgroundfillfreeformgradient
type BackgroundFillFreeformGradient struct {
	Type   BackgroundFillType `json:"type"`
	Colors []int              `json:"colors"`
}
