package bot

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"reflect"
	"strings"

	"github.com/go-telegram/bot/models"
)

type MarshalInputMedia interface {
	MarshalInputMedia() ([]byte, error)
}

//var marshalInputMediaType = reflect.TypeOf(new(MarshalInputMedia))

// buildRequestForm builds form-data for request
// if params contains InputFile of type InputFileUpload, it will be added to form-data ad upload file. Also, for InputMedia attachments
// It returns content-type and form-data. And error if occurred
func buildRequestForm(params any) (string, io.Reader, error) {
	if params == nil {
		return "application/json", http.NoBody, nil
	}

	buf := bytes.NewBuffer(nil)
	form := multipart.NewWriter(buf)
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

		var w io.Writer
		var errCreateField error
		var data io.Reader

		switch vv := v.Field(i).Interface().(type) {
		case string:
			data = strings.NewReader(vv)
			w, errCreateField = form.CreateFormField(fieldName)
		case *models.InputFileUpload:
			data = vv.Data
			w, errCreateField = form.CreateFormFile(fieldName, vv.Filename)
		case *models.InputFileString:
			data = strings.NewReader(vv.Data)
			w, errCreateField = form.CreateFormField(fieldName)
		case *models.InputMediaPhoto, *models.InputMediaVideo, *models.InputMediaAnimation, *models.InputMediaAudio, *models.InputMediaDocument:
			media := vv.(models.InputMedia)
			// add fields with media attachment if needed
			if strings.HasPrefix(media.GetMedia(), "attach://") {
				filename := strings.TrimPrefix(media.GetMedia(), "attach://")
				mediaAttachmentField, errCreateMediaAttachmentField := form.CreateFormFile(filename, filename)
				if errCreateMediaAttachmentField != nil {
					return "", nil, errCreateMediaAttachmentField
				}
				_, errCopy := io.Copy(mediaAttachmentField, media.Attachment())
				if errCopy != nil {
					return "", nil, errCopy
				}
				fieldsCount++
			}
			line, errEncode := media.MarshalInputMedia()
			if errEncode != nil {
				return "", nil, errEncode
			}

			data = bytes.NewReader(line)
			w, errCreateField = form.CreateFormField(fieldName)
		case []models.InputMedia:
			var lines []string
			for _, media := range vv {
				// add fields with media attachment if needed
				if strings.HasPrefix(media.GetMedia(), "attach://") {
					filename := strings.TrimPrefix(media.GetMedia(), "attach://")
					mediaAttachmentField, errCreateMediaAttachmentField := form.CreateFormFile(filename, filename)
					if errCreateMediaAttachmentField != nil {
						return "", nil, errCreateMediaAttachmentField
					}
					_, errCopy := io.Copy(mediaAttachmentField, media.Attachment())
					if errCopy != nil {
						return "", nil, errCopy
					}
					fieldsCount++
				}

				line, errEncode := media.MarshalInputMedia()
				if errEncode != nil {
					return "", nil, errEncode
				}
				lines = append(lines, string(line))
			}

			data = strings.NewReader("[" + strings.Join(lines, ",") + "]")
			w, errCreateField = form.CreateFormField(fieldName)
		default:
			d, errMarshal := json.Marshal(v.Field(i).Interface())
			if errMarshal != nil {
				return "", nil, errMarshal
			}
			d = bytes.Trim(d, "\"")
			data = bytes.NewReader(d)
			w, errCreateField = form.CreateFormField(fieldName)
		}

		if errCreateField != nil {
			return "", nil, errCreateField
		}

		_, errCopy := io.Copy(w, data)
		if errCopy != nil {
			return "", nil, errCopy
		}
		fieldsCount++
	}

	errClose := form.Close()
	if errClose != nil {
		return "", nil, errClose
	}

	if fieldsCount == 0 {
		return "application/json", http.NoBody, nil
	}

	return form.FormDataContentType(), buf, nil
}
