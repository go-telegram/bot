package models

import (
	"encoding/json"
	"fmt"
	"io"
)

type Story struct {
	Chat Chat `json:"chat"`
	ID   int  `json:"id"`
}

type InputStoryContent interface {
	inputStoryContentTag()
}

// InputStoryContentPhoto https://core.telegram.org/bots/api#inputstorycontentphoto
type InputStoryContentPhoto struct {
	Photo           string    `json:"photo"`
	MediaAttachment io.Reader `json:"-"`
}

func (i InputStoryContentPhoto) GetMedia() string      { return i.Photo }
func (i InputStoryContentPhoto) Attachment() io.Reader { return i.MediaAttachment }
func (InputStoryContentPhoto) inputStoryContentTag()   {}
func (i InputStoryContentPhoto) MarshalInputMedia() ([]byte, error) {
	ret := struct {
		Type string `json:"type"`
		InputStoryContentPhoto
	}{
		Type:                   "photo",
		InputStoryContentPhoto: i,
	}

	return json.Marshal(&ret)
}

// InputStoryContentVideo https://core.telegram.org/bots/api#inputstorycontentvideo
type InputStoryContentVideo struct {
	Video               string    `json:"video"`
	Duration            float64   `json:"duration,omitempty"`
	CoverFrameTimestamp float64   `json:"cover_frame_timestamp,omitempty"`
	IsAnimation         bool      `json:"is_animation,omitempty"`
	MediaAttachment     io.Reader `json:"-"`
}

func (i InputStoryContentVideo) GetMedia() string      { return i.Video }
func (i InputStoryContentVideo) Attachment() io.Reader { return i.MediaAttachment }
func (InputStoryContentVideo) inputStoryContentTag()   {}
func (i InputStoryContentVideo) MarshalInputMedia() ([]byte, error) {
	ret := struct {
		Type string `json:"type"`
		InputStoryContentVideo
	}{
		Type:                   "video",
		InputStoryContentVideo: i,
	}

	return json.Marshal(&ret)
}

// StoryArea https://core.telegram.org/bots/api#storyarea
type StoryArea struct {
	Position StoryAreaPosition `json:"position"`
	Type     StoryAreaType     `json:"type"`
}

// StoryAreaPosition https://core.telegram.org/bots/api#storyareaposition
type StoryAreaPosition struct {
	XPercentage            float64 `json:"x_percentage"`
	YPercentage            float64 `json:"y_percentage"`
	WidthPercentage        float64 `json:"width_percentage"`
	HeightPercentage       float64 `json:"height_percentage"`
	RotationAngle          float64 `json:"rotation_angle"`
	CornerRadiusPercentage float64 `json:"corner_radius_percentage"`
}

// LocationAddress https://core.telegram.org/bots/api#locationaddress
type LocationAddress struct {
	CountryCode string `json:"country_code"`
	State       string `json:"state,omitempty"`
	City        string `json:"city,omitempty"`
	Street      string `json:"street,omitempty"`
}

type StoryAreaTypeType string

var (
	StoryAreaTypeTypeLocation          StoryAreaTypeType = "location"
	StoryAreaTypeTypeSuggestedReaction StoryAreaTypeType = "suggested_reaction"
	StoryAreaTypeTypeLink              StoryAreaTypeType = "link"
	StoryAreaTypeTypeWeather           StoryAreaTypeType = "weather"
	StoryAreaTypeTypeUniqueGift        StoryAreaTypeType = "unique_gift"
)

// StoryAreaType https://core.telegram.org/bots/api#storyareatype
type StoryAreaType struct {
	Type StoryAreaTypeType

	StoryAreaTypeLocation          *StoryAreaTypeLocation
	StoryAreaTypeSuggestedReaction *StoryAreaTypeSuggestedReaction
	StoryAreaTypeLink              *StoryAreaTypeLink
	StoryAreaTypeWeather           *StoryAreaTypeWeather
	StoryAreaTypeUniqueGift        *StoryAreaTypeUniqueGift
}

func (s *StoryAreaType) UnmarshalJSON(data []byte) error {
	v := struct {
		Type StoryAreaTypeType `json:"type"`
	}{}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v.Type {
	case StoryAreaTypeTypeLocation:
		s.Type = StoryAreaTypeTypeLocation
		s.StoryAreaTypeLocation = &StoryAreaTypeLocation{
			Type: StoryAreaTypeTypeLocation,
		}
		return json.Unmarshal(data, s.StoryAreaTypeLocation)
	case StoryAreaTypeTypeSuggestedReaction:
		s.Type = StoryAreaTypeTypeSuggestedReaction
		s.StoryAreaTypeSuggestedReaction = &StoryAreaTypeSuggestedReaction{
			Type: StoryAreaTypeTypeSuggestedReaction,
		}
		return json.Unmarshal(data, s.StoryAreaTypeSuggestedReaction)
	case StoryAreaTypeTypeLink:
		s.Type = StoryAreaTypeTypeLink
		s.StoryAreaTypeLink = &StoryAreaTypeLink{
			Type: StoryAreaTypeTypeLink,
		}
		return json.Unmarshal(data, s.StoryAreaTypeLink)
	case StoryAreaTypeTypeWeather:
		s.Type = StoryAreaTypeTypeWeather
		s.StoryAreaTypeWeather = &StoryAreaTypeWeather{
			Type: StoryAreaTypeTypeWeather,
		}
		return json.Unmarshal(data, s.StoryAreaTypeWeather)
	case StoryAreaTypeTypeUniqueGift:
		s.Type = StoryAreaTypeTypeUniqueGift
		s.StoryAreaTypeUniqueGift = &StoryAreaTypeUniqueGift{
			Type: StoryAreaTypeTypeUniqueGift,
		}
		return json.Unmarshal(data, s.StoryAreaTypeUniqueGift)
	}

	return fmt.Errorf("unsupported StoryAreaType type")
}

// StoryAreaTypeLocation https://core.telegram.org/bots/api#storyareatypelocation
type StoryAreaTypeLocation struct {
	Type      StoryAreaTypeType `json:"type"`
	Latitude  float64           `json:"latitude"`
	Longitude float64           `json:"longitude"`
	Address   *LocationAddress  `json:"address,omitempty"`
}

// StoryAreaTypeSuggestedReaction https://core.telegram.org/bots/api#storyareatypesuggestedreaction
type StoryAreaTypeSuggestedReaction struct {
	Type         StoryAreaTypeType `json:"type"`
	ReactionType ReactionType      `json:"reaction_type"`
	IsDark       bool              `json:"is_dark,omitempty"`
	IsFlipped    bool              `json:"is_flipped,omitempty"`
}

// StoryAreaTypeLink https://core.telegram.org/bots/api#storyareatypelink
type StoryAreaTypeLink struct {
	Type StoryAreaTypeType `json:"type"`
	URL  string            `json:"url"`
}

// StoryAreaTypeWeather https://core.telegram.org/bots/api#storyareatypeweather
type StoryAreaTypeWeather struct {
	Type            StoryAreaTypeType `json:"type"`
	Temperature     float64           `json:"temperature"`
	Emoji           string            `json:"emoji"`
	BackgroundColor int               `json:"background_color"`
}

// StoryAreaTypeUniqueGift https://core.telegram.org/bots/api#storyareatypeuniquegift
type StoryAreaTypeUniqueGift struct {
	Type StoryAreaTypeType `json:"type"`
	Name string            `json:"name"`
}
