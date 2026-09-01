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

// mediaThumbnail is implemented by the media types carrying a thumbnail, which is
// uploaded as its own attach:// part instead of a form field of its own.
type mediaThumbnail interface {
	GetThumbnail() models.InputFile
}

var customMarshalInterface = reflect.TypeOf(new(customMarshal)).Elem()
var inputMediaInterface = reflect.TypeOf(new(inputMedia)).Elem()

// formWriter is the part of multipart.Writer used to build a request, narrowed so the
// created part names can be tracked, see requestForm.
type formWriter interface {
	CreateFormField(fieldName string) (io.Writer, error)
	CreateFormFile(fieldName, filename string) (io.Writer, error)
}

// requestForm tracks the names of the parts written to one request. A part name is
// what an attach:// reference resolves against, so two different files under one name
// are ambiguous and rejected, while one file referenced twice reuses the part already
// written. A file part and a form field sharing a name are the same ambiguity.
type requestForm struct {
	w          formWriter
	fileParts  map[string]io.Reader
	fieldNames map[string]struct{}
}

func newRequestForm(w formWriter) *requestForm {
	return &requestForm{
		w:          w,
		fileParts:  map[string]io.Reader{},
		fieldNames: map[string]struct{}{},
	}
}

// hasFilePart reports whether a file part of that name is written, i.e. an attach://
// reference to it already resolves and needs no attachment of its own.
func (f *requestForm) hasFilePart(partName string) bool {
	_, ok := f.fileParts[partName]
	return ok
}

// addFilePart writes data as the file part partName. Repeating a name with the same
// reader, or with none, references the file already uploaded under it.
func (f *requestForm) addFilePart(partName, filename string, data io.Reader) error {
	if written, ok := f.fileParts[partName]; ok {
		if readerIsNil(data) || sameReader(written, data) {
			return nil
		}
		return fmt.Errorf("duplicate form part name %q", partName)
	}
	if _, ok := f.fieldNames[partName]; ok {
		return fmt.Errorf("duplicate form part name %q", partName)
	}
	w, errCreatePart := f.w.CreateFormFile(partName, filename)
	if errCreatePart != nil {
		return errCreatePart
	}
	f.fileParts[partName] = data
	_, errCopy := io.Copy(w, data)
	return errCopy
}

// addFieldPart writes value as the form field fieldName.
func (f *requestForm) addFieldPart(fieldName string, value io.Reader) error {
	if _, ok := f.fieldNames[fieldName]; ok {
		return fmt.Errorf("duplicate form part name %q", fieldName)
	}
	if _, ok := f.fileParts[fieldName]; ok {
		return fmt.Errorf("duplicate form part name %q", fieldName)
	}
	w, errCreateField := f.w.CreateFormField(fieldName)
	if errCreateField != nil {
		return errCreateField
	}
	f.fieldNames[fieldName] = struct{}{}
	_, errCopy := io.Copy(w, value)
	return errCopy
}

// buildRequestForm builds form-data for request
// if params contains InputFile of type InputFileUpload, it will be added to form-data ad upload file. Also, for InputMedia attachments
func buildRequestForm(w *multipart.Writer, params any) (int, error) {
	form := newRequestForm(w)
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
	return isNilValue(r)
}

// isNilValue reports whether a value is nil, including a typed nil pointer carried by
// a non-nil interface, which an == nil check does not catch.
func isNilValue(value any) bool {
	if value == nil {
		return true
	}
	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.Pointer, reflect.Interface, reflect.Map, reflect.Slice, reflect.Chan, reflect.Func:
		return v.IsNil()
	default:
		return false
	}
}

// sameReader reports whether two attachments are one reader, i.e. one file referenced
// twice rather than two files claiming a single part name.
func sameReader(a, b io.Reader) bool {
	ta := reflect.TypeOf(a)
	if ta == nil || ta != reflect.TypeOf(b) || !ta.Comparable() {
		return false
	}
	return a == b
}

func addFormFieldInputFileUpload(form *requestForm, fieldName string, value *models.InputFileUpload) error {
	if isNilValue(value) {
		return fmt.Errorf("nil value for field %s", fieldName)
	}
	if readerIsNil(value.Data) {
		return fmt.Errorf("nil data for field %s", fieldName)
	}
	return form.addFilePart(fieldName, value.Filename, value.Data)
}

func addFormFieldInputMediaItem(form *requestForm, value inputMedia) ([]byte, error) {
	if err := addInputMediaAttachment(form, value); err != nil {
		return nil, err
	}
	return value.MarshalInputMedia()
}

// addInputMediaAttachment adds the attach:// upload parts referenced by a media item, if any.
func addInputMediaAttachment(form *requestForm, value inputMedia) error {
	if strings.HasPrefix(value.GetMedia(), "attach://") {
		filename := strings.TrimPrefix(value.GetMedia(), "attach://")
		if !form.hasFilePart(filename) && readerIsNil(value.Attachment()) {
			return fmt.Errorf("nil attachment for attach://%s", filename)
		}
		if err := form.addFilePart(filename, filename, value.Attachment()); err != nil {
			return err
		}
	}
	if thumbnailed, ok := value.(mediaThumbnail); ok {
		if err := addInputFileAttachment(form, thumbnailed.GetThumbnail()); err != nil {
			return err
		}
	}
	if live, ok := value.(*models.InputMediaLivePhoto); ok && strings.HasPrefix(live.Photo, "attach://") {
		filename := strings.TrimPrefix(live.Photo, "attach://")
		if !form.hasFilePart(filename) && readerIsNil(live.PhotoAttachment) {
			return fmt.Errorf("nil PhotoAttachment for attach://%s", filename)
		}
		if err := form.addFilePart(filename, filename, live.PhotoAttachment); err != nil {
			return err
		}
	}
	return nil
}

// addInputFileAttachment uploads a nested InputFile, which is marshalled as an
// attach:// reference instead of becoming a form field of its own. Anything but an
// upload (a file_id or an URL) carries no content and is left to the encoder, and so
// is a typed nil, which the encoder omits.
func addInputFileAttachment(form *requestForm, value models.InputFile) error {
	upload, ok := value.(*models.InputFileUpload)
	if !ok || upload == nil {
		return nil
	}
	if upload.Filename == "" {
		return fmt.Errorf("empty filename for nested upload, it is the attach:// reference")
	}
	if !form.hasFilePart(upload.Filename) && readerIsNil(upload.Data) {
		return fmt.Errorf("nil data for attach://%s", upload.Filename)
	}
	return form.addFilePart(upload.Filename, upload.Filename, upload.Data)
}

func addFormFieldCustomMarshal(form *requestForm, fieldName string, value customMarshal) error {
	line, errEncode := value.MarshalCustom()
	if errEncode != nil {
		return errEncode
	}
	return form.addFieldPart(fieldName, bytes.NewReader(line))
}

func addFormFieldInputMedia(form *requestForm, fieldName string, value inputMedia) error {
	if isNilValue(value) {
		return fmt.Errorf("nil value for field %s", fieldName)
	}

	line, err := addFormFieldInputMediaItem(form, value)
	if err != nil {
		return err
	}

	return form.addFieldPart(fieldName, bytes.NewReader(line))
}

func addFormFieldInputMediaSlice(form *requestForm, fieldName string, value []inputMedia) error {
	var lines []string
	for i, media := range value {
		if isNilValue(media) {
			return fmt.Errorf("nil value for field %s at index %d", fieldName, i)
		}
		line, err := addFormFieldInputMediaItem(form, media)
		if err != nil {
			return err
		}
		lines = append(lines, string(line))
	}

	return form.addFieldPart(fieldName, strings.NewReader("["+strings.Join(lines, ",")+"]"))
}

// addFormFieldInputRichMessage adds the attach:// uploads referenced by a rich message
// before encoding it, so that nested media is uploaded like top level InputMedia is.
func addFormFieldInputRichMessage(form *requestForm, fieldName string, value *models.InputRichMessage) error {
	if err := addInputRichBlockSliceAttachments(form, value.Blocks); err != nil {
		return err
	}
	for i, media := range value.Media {
		if isNilValue(media.Media) {
			return fmt.Errorf("nil media for field %s at index %d", fieldName, i)
		}
		if err := addInputMediaAttachment(form, media.Media); err != nil {
			return err
		}
	}
	return addFormFieldDefault(form, fieldName, value)
}

func addInputRichBlockSliceAttachments(form *requestForm, blocks []models.InputRichBlock) error {
	for _, block := range blocks {
		if err := addInputRichBlockAttachments(form, block); err != nil {
			return err
		}
	}
	return nil
}

// addInputRichBlockAttachments walks a single block: media blocks upload their attachment,
// container blocks recurse into their nested blocks.
func addInputRichBlockAttachments(form *requestForm, block models.InputRichBlock) error {
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

func addFormFieldInlineQueryResultSlice(form *requestForm, fieldName string, value []models.InlineQueryResult) error {
	var lines []string
	for _, media := range value {
		line, errEncode := media.MarshalCustom()
		if errEncode != nil {
			return errEncode
		}
		lines = append(lines, string(line))
	}

	return form.addFieldPart(fieldName, strings.NewReader("["+strings.Join(lines, ",")+"]"))
}

func addFormFieldInputStickerSlice(form *requestForm, fieldName string, value []models.InputSticker) error {
	var lines []string
	for _, sticker := range value {
		if strings.HasPrefix(sticker.Sticker, "attach://") {
			filename := strings.TrimPrefix(sticker.Sticker, "attach://")
			if !form.hasFilePart(filename) && readerIsNil(sticker.StickerAttachment) {
				return fmt.Errorf("nil StickerAttachment for attach://%s", filename)
			}
			if err := form.addFilePart(filename, filename, sticker.StickerAttachment); err != nil {
				return err
			}
		}
		line, errEncode := json.Marshal(sticker)
		if errEncode != nil {
			return errEncode
		}
		lines = append(lines, string(line))
	}

	return form.addFieldPart(fieldName, strings.NewReader("["+strings.Join(lines, ",")+"]"))
}

func addFormFieldDefault(form *requestForm, fieldName string, value any) error {
	d, errMarshal := json.Marshal(value)
	if errMarshal != nil {
		return errMarshal
	}
	d = bytes.Trim(d, "\"") // for strings values
	return form.addFieldPart(fieldName, bytes.NewReader(d))
}

func addFormFieldString(form *requestForm, fieldName string, value string) error {
	return form.addFieldPart(fieldName, strings.NewReader(value))
}
