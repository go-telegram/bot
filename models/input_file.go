package models

import (
	"encoding/json"
	"io"
)

type InputFileType int

// InputFile https://core.telegram.org/bots/api#inputfile
type InputFile interface {
	inputFileTag()
}

type InputFileUpload struct {
	Filename string
	Data     io.Reader
}

func (*InputFileUpload) inputFileTag() {}

func (i *InputFileUpload) MarshalJSON() ([]byte, error) {
	return []byte(`"@` + i.Filename + `"`), nil
}

type InputFileString struct {
	Data string
}

func (*InputFileString) inputFileTag() {}

func (i *InputFileString) MarshalJSON() ([]byte, error) {
	return []byte(`"` + i.Data + `"`), nil
}

func (i *InputFileString) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &i.Data)
}
