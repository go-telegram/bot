package bot

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"reflect"
	"strings"

	"github.com/go-telegram/bot/models"
)

type customMarshal interface {
	MarshalCustom() ([]byte, error)
}

var customMarshalInterface = reflect.TypeOf(new(customMarshal)).Elem()
var inputMediaInterface = reflect.TypeOf(new(models.InputMedia)).Elem()

// buildRequestForm builds form-data for request
// if params contains InputFile of type InputFileUpload, it will be added to form-data ad upload file. Also, for InputMedia attachments
func buildRequestForm(form *multipart.Writer, params any) (int, error) {
	v := reflect.ValueOf(params).Elem()

	var fieldsCount int

	for i := 0; i < v.NumField(); i++ {
		jsonTag := v.Type().Field(i).Tag.Get("json")
		if jsonTag == "-" || jsonTag == "" {
			continue
		}
		fieldName := strings.Split(jsonTag, ",")[0]
		omitempty := strings.Contains(jsonTag, ",omitempty")

		if omitempty && v.Field(i).IsZero() {
			continue
		}

		// check fields by interface
		if v.Field(i).Type().Implements(customMarshalInterface) {
			err := addFormFieldCustomMarshal(form, fieldName, v.Field(i).Interface().(customMarshal))
			if err != nil {
				return 0, err
			}
			continue
		}
		if v.Field(i).Type().Implements(inputMediaInterface) {
			err := addFormFieldInputMedia(form, fieldName, v.Field(i).Interface().(models.InputMedia))
			if err != nil {
				return 0, err
			}
			continue
		}

		var err error

		// check fields by type
		switch vv := v.Field(i).Interface().(type) {
		case string:
			err = addFormFieldString(form, fieldName, vv)
		case *models.InputFileUpload:
			err = addFormFieldInputFileUpload(form, fieldName, vv)
		case *models.InputFileString:
			err = addFormFieldString(form, fieldName, vv.Data)
		case []models.InputMedia:
			err = addFormFieldInputMediaSlice(form, fieldName, vv)
		case []models.InlineQueryResult:
			err = addFormFieldInlineQueryResultSlice(form, fieldName, vv)
		default:
			err = addFormFieldDefault(form, fieldName, v.Field(i).Interface())
		}
		if err != nil {
			return 0, err
		}

		fieldsCount++
	}

	return fieldsCount, nil
}

func addFormFieldInputFileUpload(form *multipart.Writer, fieldName string, value *models.InputFileUpload) error {
	w, errCreateField := form.CreateFormFile(fieldName, value.Filename)
	if errCreateField != nil {
		return errCreateField
	}
	_, errCopy := io.Copy(w, value.Data)
	return errCopy
}

func addFormFieldInputMediaItem(form *multipart.Writer, value models.InputMedia) ([]byte, error) {
	if strings.HasPrefix(value.GetMedia(), "attach://") {
		filename := strings.TrimPrefix(value.GetMedia(), "attach://")
		mediaAttachmentField, errCreateMediaAttachmentField := form.CreateFormFile(filename, filename)
		if errCreateMediaAttachmentField != nil {
			return nil, errCreateMediaAttachmentField
		}
		_, errCopy := io.Copy(mediaAttachmentField, value.Attachment())
		if errCopy != nil {
			return nil, errCopy
		}
	}
	return value.MarshalInputMedia()
}

func addFormFieldCustomMarshal(form *multipart.Writer, fieldName string, value customMarshal) error {
	line, errEncode := value.MarshalCustom()
	if errEncode != nil {
		return errEncode
	}
	w, errCreateField := form.CreateFormField(fieldName)
	if errCreateField != nil {
		return errCreateField
	}
	_, errCopy := io.Copy(w, bytes.NewReader(line))
	return errCopy
}

func addFormFieldInputMedia(form *multipart.Writer, fieldName string, value models.InputMedia) error {
	line, err := addFormFieldInputMediaItem(form, value)
	if err != nil {
		return err
	}

	w, errCreateField := form.CreateFormField(fieldName)
	if errCreateField != nil {
		return errCreateField
	}
	_, errCopy := io.Copy(w, bytes.NewReader(line))
	return errCopy
}

func addFormFieldInputMediaSlice(form *multipart.Writer, fieldName string, value []models.InputMedia) error {
	var lines []string
	for _, media := range value {
		line, err := addFormFieldInputMediaItem(form, media)
		if err != nil {
			return err
		}
		lines = append(lines, string(line))
	}

	w, errCreateField := form.CreateFormField(fieldName)
	if errCreateField != nil {
		return errCreateField
	}
	_, errCopy := io.Copy(w, strings.NewReader("["+strings.Join(lines, ",")+"]"))
	return errCopy
}

func addFormFieldInlineQueryResultSlice(form *multipart.Writer, fieldName string, value []models.InlineQueryResult) error {
	var lines []string
	for _, media := range value {
		line, errEncode := media.MarshalCustom()
		if errEncode != nil {
			return errEncode
		}
		lines = append(lines, string(line))
	}

	w, errCreateField := form.CreateFormField(fieldName)
	if errCreateField != nil {
		return errCreateField
	}
	_, errCopy := io.Copy(w, strings.NewReader("["+strings.Join(lines, ",")+"]"))
	return errCopy
}

func addFormFieldDefault(form *multipart.Writer, fieldName string, value any) error {
	d, errMarshal := json.Marshal(value)
	if errMarshal != nil {
		return errMarshal
	}
	d = bytes.Trim(d, "\"") // for strings values
	w, errCreateField := form.CreateFormField(fieldName)
	if errCreateField != nil {
		return errCreateField
	}
	_, errCopy := io.Copy(w, bytes.NewReader(d))
	return errCopy
}

func addFormFieldString(form *multipart.Writer, fieldName string, value string) error {
	w, errCreateField := form.CreateFormField(fieldName)
	if errCreateField != nil {
		return errCreateField
	}
	_, errCopy := io.Copy(w, strings.NewReader(value))
	return errCopy
}
