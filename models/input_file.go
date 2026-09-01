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

// MarshalJSON encodes the upload as its attach:// reference, for the places the
// value is marshalled instead of becoming a form field of its own, e.g. the
// thumbnail of an InputMedia. Filename doubles as the name of the form part.
func (i *InputFileUpload) MarshalJSON() ([]byte, error) {
	return json.Marshal("attach://" + i.Filename)
}

type InputFileString struct {
	Data string
}

func (*InputFileString) inputFileTag() {}

func (i *InputFileString) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.Data)
}

func (i *InputFileString) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &i.Data)
}

// normalizeInputFile turns a typed nil pointer into an untyped nil, so an omitempty
// field carrying one is omitted instead of encoded as a null the Bot API rejects.
func normalizeInputFile(f InputFile) InputFile {
	switch v := f.(type) {
	case *InputFileUpload:
		if v == nil {
			return nil
		}
	case *InputFileString:
		if v == nil {
			return nil
		}
	}
	return f
}
