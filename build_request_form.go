package bot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"reflect"
	"strings"

	"github.com/go-telegram/bot/models"
)

type inputMedia interface {
	MarshalInputMedia() ([]byte, error)
	Attachment() io.Reader
	GetMedia() string
}

type customMarshal interface {
	MarshalCustom() ([]byte, error)
}

var customMarshalInterface = reflect.TypeOf(new(customMarshal)).Elem()
var inputMediaInterface = reflect.TypeOf(new(inputMedia)).Elem()

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

		// check fields by interface (declared type)
		if v.Field(i).Type().Implements(customMarshalInterface) {
			err := addFormFieldCustomMarshal(form, fieldName, v.Field(i).Interface().(customMarshal))
			if err != nil {
				return 0, err
			}
			continue
		}
		if v.Field(i).Type().Implements(inputMediaInterface) {
			err := addFormFieldInputMedia(form, fieldName, v.Field(i).Interface().(inputMedia))
			if err != nil {
				return 0, err
			}
			continue
		}

		// check fields by interface (runtime concrete type behind interface fields)
		if v.Field(i).Kind() == reflect.Interface && !v.Field(i).IsNil() {
			if v.Field(i).Elem().Type().Implements(customMarshalInterface) {
				err := addFormFieldCustomMarshal(form, fieldName, v.Field(i).Interface().(customMarshal))
				if err != nil {
					return 0, err
				}
				continue
			}
			if v.Field(i).Elem().Type().Implements(inputMediaInterface) {
				err := addFormFieldInputMedia(form, fieldName, v.Field(i).Interface().(inputMedia))
				if err != nil {
					return 0, err
				}
				continue
			}
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
		case *models.InputMediaLivePhoto:
			err = addFormFieldInputMedia(form, fieldName, vv)
		case []models.InputMedia:
			var ss []inputMedia
			for _, m := range vv {
				ss = append(ss, m)
			}
			err = addFormFieldInputMediaSlice(form, fieldName, ss)
		case []models.InputPaidMedia:
			var ss []inputMedia
			for _, m := range vv {
				ss = append(ss, m)
			}
			err = addFormFieldInputMediaSlice(form, fieldName, ss)
		case []models.InlineQueryResult:
			err = addFormFieldInlineQueryResultSlice(form, fieldName, vv)
		case []models.InputSticker:
			err = addFormFieldInputStickerSlice(form, fieldName, vv)
		case models.InputRichMessage:
			err = addFormFieldInputRichMessage(form, fieldName, &vv)
		case *models.InputRichMessage:
			err = addFormFieldInputRichMessage(form, fieldName, vv)
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

func readerIsNil(r io.Reader) bool {
	if r == nil {
		return true
	}
	v := reflect.ValueOf(r)
	switch v.Kind() {
	case reflect.Pointer, reflect.Interface, reflect.Map, reflect.Slice, reflect.Chan, reflect.Func:
		return v.IsNil()
	default:
		return false
	}
}

func addFormFieldInputFileUpload(form *multipart.Writer, fieldName string, value *models.InputFileUpload) error {
	if readerIsNil(value.Data) {
		return fmt.Errorf("nil data for field %s", fieldName)
	}
	w, errCreateField := form.CreateFormFile(fieldName, value.Filename)
	if errCreateField != nil {
		return errCreateField
	}
	_, errCopy := io.Copy(w, value.Data)
	return errCopy
}

func addFormFieldInputMediaItem(form *multipart.Writer, value inputMedia) ([]byte, error) {
	if err := addInputMediaAttachment(form, value); err != nil {
		return nil, err
	}
	return value.MarshalInputMedia()
}

// addInputMediaAttachment adds the attach:// upload parts referenced by a media item, if any.
func addInputMediaAttachment(form *multipart.Writer, value inputMedia) error {
	if strings.HasPrefix(value.GetMedia(), "attach://") {
		filename := strings.TrimPrefix(value.GetMedia(), "attach://")
		if readerIsNil(value.Attachment()) {
			return fmt.Errorf("nil attachment for attach://%s", filename)
		}
		mediaAttachmentField, errCreateMediaAttachmentField := form.CreateFormFile(filename, filename)
		if errCreateMediaAttachmentField != nil {
			return errCreateMediaAttachmentField
		}
		_, errCopy := io.Copy(mediaAttachmentField, value.Attachment())
		if errCopy != nil {
			return errCopy
		}
	}
	if live, ok := value.(*models.InputMediaLivePhoto); ok && strings.HasPrefix(live.Photo, "attach://") {
		filename := strings.TrimPrefix(live.Photo, "attach://")
		if readerIsNil(live.PhotoAttachment) {
			return fmt.Errorf("nil PhotoAttachment for attach://%s", filename)
		}
		photoField, errCreate := form.CreateFormFile(filename, filename)
		if errCreate != nil {
			return errCreate
		}
		if _, err := io.Copy(photoField, live.PhotoAttachment); err != nil {
			return err
		}
	}
	return nil
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

func addFormFieldInputMedia(form *multipart.Writer, fieldName string, value inputMedia) error {
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

func addFormFieldInputMediaSlice(form *multipart.Writer, fieldName string, value []inputMedia) error {
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

// addFormFieldInputRichMessage adds the attach:// uploads referenced by a rich message
// before encoding it, so that nested media is uploaded like top level InputMedia is.
func addFormFieldInputRichMessage(form *multipart.Writer, fieldName string, value *models.InputRichMessage) error {
	if err := addInputRichBlockSliceAttachments(form, value.Blocks); err != nil {
		return err
	}
	for _, media := range value.Media {
		if media.Media == nil {
			continue
		}
		if err := addInputMediaAttachment(form, media.Media); err != nil {
			return err
		}
	}
	return addFormFieldDefault(form, fieldName, value)
}

func addInputRichBlockSliceAttachments(form *multipart.Writer, blocks []models.InputRichBlock) error {
	for _, block := range blocks {
		if err := addInputRichBlockAttachments(form, block); err != nil {
			return err
		}
	}
	return nil
}

// addInputRichBlockAttachments walks a single block: media blocks upload their attachment,
// container blocks recurse into their nested blocks.
func addInputRichBlockAttachments(form *multipart.Writer, block models.InputRichBlock) error {
	switch block.Type {
	case models.RichBlockTypeAnimation:
		if block.InputRichBlockAnimation != nil {
			return addInputMediaAttachment(form, &block.InputRichBlockAnimation.Animation)
		}
	case models.RichBlockTypeAudio:
		if block.InputRichBlockAudio != nil {
			return addInputMediaAttachment(form, &block.InputRichBlockAudio.Audio)
		}
	case models.RichBlockTypeDocument:
		if block.InputRichBlockDocument != nil {
			return addInputMediaAttachment(form, &block.InputRichBlockDocument.Document)
		}
	case models.RichBlockTypePhoto:
		if block.InputRichBlockPhoto != nil {
			return addInputMediaAttachment(form, &block.InputRichBlockPhoto.Photo)
		}
	case models.RichBlockTypeVideo:
		if block.InputRichBlockVideo != nil {
			return addInputMediaAttachment(form, &block.InputRichBlockVideo.Video)
		}
	case models.RichBlockTypeVoiceNote:
		if block.InputRichBlockVoiceNote != nil {
			return addInputMediaAttachment(form, &block.InputRichBlockVoiceNote.VoiceNote)
		}
	case models.RichBlockTypeList:
		if block.InputRichBlockList != nil {
			for _, item := range block.InputRichBlockList.Items {
				if err := addInputRichBlockSliceAttachments(form, item.Blocks); err != nil {
					return err
				}
			}
		}
	case models.RichBlockTypeBlockQuotation:
		if block.InputRichBlockBlockQuotation != nil {
			return addInputRichBlockSliceAttachments(form, block.InputRichBlockBlockQuotation.Blocks)
		}
	case models.RichBlockTypeCollage:
		if block.InputRichBlockCollage != nil {
			return addInputRichBlockSliceAttachments(form, block.InputRichBlockCollage.Blocks)
		}
	case models.RichBlockTypeSlideshow:
		if block.InputRichBlockSlideshow != nil {
			return addInputRichBlockSliceAttachments(form, block.InputRichBlockSlideshow.Blocks)
		}
	case models.RichBlockTypeDetails:
		if block.InputRichBlockDetails != nil {
			return addInputRichBlockSliceAttachments(form, block.InputRichBlockDetails.Blocks)
		}
	}
	return nil
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

func addFormFieldInputStickerSlice(form *multipart.Writer, fieldName string, value []models.InputSticker) error {
	var lines []string
	for _, sticker := range value {
		if strings.HasPrefix(sticker.Sticker, "attach://") {
			filename := strings.TrimPrefix(sticker.Sticker, "attach://")
			if readerIsNil(sticker.StickerAttachment) {
				return fmt.Errorf("nil StickerAttachment for attach://%s", filename)
			}
			attachmentField, errCreateAttachmentField := form.CreateFormFile(filename, filename)
			if errCreateAttachmentField != nil {
				return errCreateAttachmentField
			}
			_, errCopy := io.Copy(attachmentField, sticker.StickerAttachment)
			if errCopy != nil {
				return errCopy
			}
		}
		line, errEncode := json.Marshal(sticker)
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
