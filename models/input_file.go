package models

import (
	"io"
)

type InputFileType int

const (
	InputFileTypeUpload InputFileType = iota
	InputFileTypeString
)

type InputFile interface {
	inputFileTag()

	Type() InputFileType
}

type InputFileUpload struct {
	Filename string
	Data     io.Reader
}

func (InputFileUpload) inputFileTag() {}

func (i *InputFileUpload) Type() InputFileType {
	return InputFileTypeUpload
}

func (i *InputFileUpload) MarshalJSON() ([]byte, error) {
	return []byte(`"@` + i.Filename + `"`), nil
}

type InputFileString struct {
	Data string
}

func (InputFileString) inputFileTag() {}

func (i *InputFileString) Type() InputFileType {
	return InputFileTypeString
}

func (i *InputFileString) MarshalJSON() ([]byte, error) {
	return []byte(`"` + i.Data + `"`), nil
}
