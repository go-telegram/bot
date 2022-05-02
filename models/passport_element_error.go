package models

import (
	"encoding/json"
)

// PassportElementError https://core.telegram.org/bots/api#passportelementerror
type PassportElementError interface {
	passportElementErrorTag()

	MarshalCustom() ([]byte, error)
}

// PassportElementErrorDataField https://core.telegram.org/bots/api#passportelementerrordatafield
type PassportElementErrorDataField struct {
	Type      string `json:"type"`
	FieldName string `json:"field_name"`
	DataHash  string `json:"data_hash"`
	Message   string `json:"message"`
}

func (PassportElementErrorDataField) passportElementErrorTag() {}

func (m *PassportElementErrorDataField) MarshalCustom() ([]byte, error) {
	ret := struct {
		Type string `json:"type"`
		*PassportElementErrorDataField
	}{
		Type:                          "data",
		PassportElementErrorDataField: m,
	}

	return json.Marshal(&ret)
}

// PassportElementErrorFrontSide https://core.telegram.org/bots/api#passportelementerrorfrontside
type PassportElementErrorFrontSide struct {
	Type     string `json:"type"`
	FileHash string `json:"file_hash"`
	Message  string `json:"message"`
}

func (PassportElementErrorFrontSide) passportElementErrorTag() {}

func (m *PassportElementErrorFrontSide) MarshalCustom() ([]byte, error) {
	ret := struct {
		Type string `json:"type"`
		*PassportElementErrorFrontSide
	}{
		Type:                          "front_side",
		PassportElementErrorFrontSide: m,
	}

	return json.Marshal(&ret)
}

// PassportElementErrorReverseSide https://core.telegram.org/bots/api#passportelementerrorreverseside
type PassportElementErrorReverseSide struct {
	Type     string `json:"type"`
	FileHash string `json:"file_hash"`
	Message  string `json:"message"`
}

func (PassportElementErrorReverseSide) passportElementErrorTag() {}

func (m *PassportElementErrorReverseSide) MarshalCustom() ([]byte, error) {
	ret := struct {
		Type string `json:"type"`
		*PassportElementErrorReverseSide
	}{
		Type:                            "reverse_side",
		PassportElementErrorReverseSide: m,
	}

	return json.Marshal(&ret)
}

// PassportElementErrorSelfie https://core.telegram.org/bots/api#passportelementerrorselfie
type PassportElementErrorSelfie struct {
	Type     string `json:"type"`
	FileHash string `json:"file_hash"`
	Message  string `json:"message"`
}

func (PassportElementErrorSelfie) passportElementErrorTag() {}

func (m *PassportElementErrorSelfie) MarshalCustom() ([]byte, error) {
	ret := struct {
		Type string `json:"type"`
		*PassportElementErrorSelfie
	}{
		Type:                       "selfie",
		PassportElementErrorSelfie: m,
	}

	return json.Marshal(&ret)
}

// PassportElementErrorFile https://core.telegram.org/bots/api#passportelementerrorfile
type PassportElementErrorFile struct {
	Type     string `json:"type"`
	FileHash string `json:"file_hash"`
	Message  string `json:"message"`
}

func (PassportElementErrorFile) passportElementErrorTag() {}

func (m *PassportElementErrorFile) MarshalCustom() ([]byte, error) {
	ret := struct {
		Type string `json:"type"`
		*PassportElementErrorFile
	}{
		Type:                     "file",
		PassportElementErrorFile: m,
	}

	return json.Marshal(&ret)
}

// PassportElementErrorFiles https://core.telegram.org/bots/api#passportelementerrorfiles
type PassportElementErrorFiles struct {
	Type       string   `json:"type"`
	FileHashes []string `json:"file_hashes"`
	Message    string   `json:"message"`
}

func (PassportElementErrorFiles) passportElementErrorTag() {}

func (m *PassportElementErrorFiles) MarshalCustom() ([]byte, error) {
	ret := struct {
		Type string `json:"type"`
		*PassportElementErrorFiles
	}{
		Type:                      "files",
		PassportElementErrorFiles: m,
	}

	return json.Marshal(&ret)
}

// PassportElementErrorTranslationFile https://core.telegram.org/bots/api#passportelementerrortranslationfile
type PassportElementErrorTranslationFile struct {
	Type     string `json:"type"`
	FileHash string `json:"file_hash"`
	Message  string `json:"message"`
}

func (PassportElementErrorTranslationFile) passportElementErrorTag() {}

func (m *PassportElementErrorTranslationFile) MarshalCustom() ([]byte, error) {
	ret := struct {
		Type string `json:"type"`
		*PassportElementErrorTranslationFile
	}{
		Type:                                "translation_file",
		PassportElementErrorTranslationFile: m,
	}

	return json.Marshal(&ret)
}

// PassportElementErrorTranslationFiles https://core.telegram.org/bots/api#passportelementerrortranslationfiles
type PassportElementErrorTranslationFiles struct {
	Type       string   `json:"type"`
	FileHashes []string `json:"file_hashes"`
	Message    string   `json:"message"`
}

func (PassportElementErrorTranslationFiles) passportElementErrorTag() {}

func (m *PassportElementErrorTranslationFiles) MarshalCustom() ([]byte, error) {
	ret := struct {
		Type string `json:"type"`
		*PassportElementErrorTranslationFiles
	}{
		Type:                                 "translation_files",
		PassportElementErrorTranslationFiles: m,
	}

	return json.Marshal(&ret)
}

// PassportElementErrorUnspecified https://core.telegram.org/bots/api#passportelementerrorunspecified
type PassportElementErrorUnspecified struct {
	Type        string `json:"type"`
	ElementHash string `json:"element_hash"`
	Message     string `json:"message"`
}

func (PassportElementErrorUnspecified) passportElementErrorTag() {}

func (m *PassportElementErrorUnspecified) MarshalCustom() ([]byte, error) {
	ret := struct {
		Type string `json:"type"`
		*PassportElementErrorUnspecified
	}{
		Type:                            "unspecified",
		PassportElementErrorUnspecified: m,
	}

	return json.Marshal(&ret)
}
