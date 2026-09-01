package bot

import (
	"bytes"
	"errors"
	"io"
	"mime/multipart"
	"strings"
	"testing"

	"github.com/go-telegram/bot/models"
)

type structReader struct{}

func (structReader) Read(p []byte) (int, error) {
	return 0, io.EOF
}

func Test_addFormFieldInputFileUpload_structReaderDoesNotPanic(t *testing.T) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	form := newRequestForm(writer)
	err := addFormFieldInputFileUpload(form, "file", &models.InputFileUpload{
		Filename: "x.bin",
		Data:     structReader{},
	})
	if err != nil {
		t.Fatal(err)
	}
	if err := writer.Close(); err != nil {
		t.Fatal(err)
	}
}

func Test_readerIsNil(t *testing.T) {
	if !readerIsNil(nil) {
		t.Fatal("nil reader")
	}
	var p *strings.Reader
	var iface io.Reader = p
	if !readerIsNil(iface) {
		t.Fatal("typed nil pointer in interface")
	}
	if readerIsNil(strings.NewReader("a")) {
		t.Fatal("non-nil reader")
	}
	if readerIsNil(structReader{}) {
		t.Fatal("struct value reader is usable, not nil")
	}
}

func assertFormData(t *testing.T, data, expect string) {
	if !strings.Contains(expect, "\r\n") {
		expect = strings.ReplaceAll(expect, "\n", "\r\n")
	}

	if data != expect {
		t.Fatalf("\nExpect:\n%s\nGot:\n%s\n", expect, data)
	}
}

func Test_buildRequestForm(t *testing.T) {
	params := struct {
		String                     string                     `json:"string"`
		InputFileUpload            *models.InputFileUpload    `json:"input_file_upload"`
		InputFileString            *models.InputFileString    `json:"input_file_string"`
		InputMediaSlice            []models.InputMedia        `json:"input_media_slice"`
		InlineQueryResultSlice     []models.InlineQueryResult `json:"inline_query_result_slice"`
		DefaultInt                 int                        `json:"default_int"`
		InputMediaInterface        models.InputMedia          `json:"input_media_interface"`
		InlineQueryResultInterface models.InlineQueryResult   `json:"inline_query_result_interface"`
		InputStickerSlice          []models.InputSticker      `json:"input_sticker_slice"`
		NoJSONTag1                 string
		NoJSONTag2                 string `json:"-"`
		OmitEmptyString            string `json:"omit_empty_string,omitempty"`
	}{
		String:          "foo",
		InputFileUpload: &models.InputFileUpload{Filename: "file.png", Data: strings.NewReader("content input file upload")},
		InputFileString: &models.InputFileString{Data: "content input file string"},
		InputMediaSlice: []models.InputMedia{
			&models.InputMediaPhoto{Media: "attach://foo.png", Caption: "bar", MediaAttachment: strings.NewReader("media attachment")},
			&models.InputMediaPhoto{Media: "foobar", Caption: "bar"},
		},
		InlineQueryResultSlice: []models.InlineQueryResult{
			&models.InlineQueryResultArticle{Title: "foo", Description: "bar", InputMessageContent: &models.InputTextMessageContent{MessageText: "foo"}},
		},
		DefaultInt:                 42,
		InputMediaInterface:        &models.InputMediaPhoto{Media: "foo", Caption: "bar", ParseMode: "baz"},
		InlineQueryResultInterface: &models.InlineQueryResultArticle{Title: "foo", Description: "bar", InputMessageContent: &models.InputTextMessageContent{MessageText: "foo"}},
		InputStickerSlice: []models.InputSticker{
			{
				Sticker:           "attach://sticker.png",
				Format:            "foo",
				EmojiList:         []string{"bar"},
				StickerAttachment: strings.NewReader("sticker file"),
			},
			{
				Sticker:   "foo",
				Format:    "bar",
				EmojiList: []string{"baz"},
			},
		},
		NoJSONTag1:      "foo",
		NoJSONTag2:      "bar",
		OmitEmptyString: "",
	}

	buf := bytes.NewBuffer(nil)
	form := multipart.NewWriter(buf)
	form.SetBoundary("XXX") //nolint

	fieldsCount, errBuild := buildRequestForm(form, &params)
	if errBuild != nil {
		t.Error(errBuild)
		return
	}
	if err := form.Close(); err != nil {
		t.Errorf("failed to close form: %v", err)
	}

	expect := `--XXX
Content-Disposition: form-data; name="string"

foo
--XXX
Content-Disposition: form-data; name="input_file_upload"; filename="file.png"
Content-Type: application/octet-stream

content input file upload
--XXX
Content-Disposition: form-data; name="input_file_string"

content input file string
--XXX
Content-Disposition: form-data; name="foo.png"; filename="foo.png"
Content-Type: application/octet-stream

media attachment
--XXX
Content-Disposition: form-data; name="input_media_slice"

[{"type":"photo","media":"attach://foo.png","caption":"bar"},{"type":"photo","media":"foobar","caption":"bar"}]
--XXX
Content-Disposition: form-data; name="inline_query_result_slice"

[{"type":"article","id":"","title":"foo","input_message_content":{"message_text":"foo"},"description":"bar"}]
--XXX
Content-Disposition: form-data; name="default_int"

42
--XXX
Content-Disposition: form-data; name="input_media_interface"

{"type":"photo","media":"foo","caption":"bar","parse_mode":"baz"}
--XXX
Content-Disposition: form-data; name="inline_query_result_interface"

{"type":"article","id":"","title":"foo","input_message_content":{"message_text":"foo"},"description":"bar"}
--XXX
Content-Disposition: form-data; name="sticker.png"; filename="sticker.png"
Content-Type: application/octet-stream

sticker file
--XXX
Content-Disposition: form-data; name="input_sticker_slice"

[{"sticker":"attach://sticker.png","format":"foo","emoji_list":["bar"]},{"sticker":"foo","format":"bar","emoji_list":["baz"]}]
--XXX--
`
	assertEqualInt(t, fieldsCount, 7)
	assertFormData(t, buf.String(), expect)
}

func Test_addFormFieldInputMedia_nilAttachment(t *testing.T) {
	writer := multipart.NewWriter(bytes.NewBuffer(nil))
	form := newRequestForm(writer)
	err := addFormFieldInputMedia(form, "media", &models.InputMediaPhoto{
		Media: "attach://photo.png",
	})
	if err == nil {
		t.Fatal("expected error for attach:// with nil MediaAttachment")
	}
	if !strings.Contains(err.Error(), "nil attachment") {
		t.Fatalf("unexpected error: %v", err)
	}
}

func Test_addFormFieldInputStickerSlice_nilAttachment(t *testing.T) {
	writer := multipart.NewWriter(bytes.NewBuffer(nil))
	form := newRequestForm(writer)
	err := addFormFieldInputStickerSlice(form, "stickers", []models.InputSticker{{
		Sticker: "attach://sticker.png",
		Format:  "static",
	}})
	if err == nil {
		t.Fatal("expected error for attach:// with nil StickerAttachment")
	}
	if !strings.Contains(err.Error(), "nil StickerAttachment") {
		t.Fatalf("unexpected error: %v", err)
	}
}

func Test_buildRequestForm_inputRichMessageMedia(t *testing.T) {
	params := struct {
		RichMessage models.InputRichMessage `json:"rich_message"`
	}{
		RichMessage: models.InputRichMessage{
			HTML: "tg://document?id=doc1",
			Media: []models.InputRichMessageMedia{
				{
					ID: "doc1",
					Media: &models.InputMediaDocument{
						Media:           "attach://doc1.pdf",
						MediaAttachment: strings.NewReader("document content"),
					},
				},
				{
					ID:    "doc2",
					Media: &models.InputMediaDocument{Media: "file_id"},
				},
			},
		},
	}

	buf := bytes.NewBuffer(nil)
	form := multipart.NewWriter(buf)
	form.SetBoundary("XXX") //nolint

	fieldsCount, errBuild := buildRequestForm(form, &params)
	if errBuild != nil {
		t.Error(errBuild)
		return
	}
	if err := form.Close(); err != nil {
		t.Errorf("failed to close form: %v", err)
	}

	expect := `--XXX
Content-Disposition: form-data; name="doc1.pdf"; filename="doc1.pdf"
Content-Type: application/octet-stream

document content
--XXX
Content-Disposition: form-data; name="rich_message"

{"html":"tg://document?id=doc1","media":[{"id":"doc1","media":{"type":"document","media":"attach://doc1.pdf"}},{"id":"doc2","media":{"type":"document","media":"file_id"}}]}
--XXX--
`
	assertEqualInt(t, fieldsCount, 1)
	assertFormData(t, buf.String(), expect)
}

func Test_buildRequestForm_inputRichMessageBlocks(t *testing.T) {
	params := struct {
		RichMessage *models.InputRichMessage `json:"rich_message"`
	}{
		RichMessage: &models.InputRichMessage{
			Blocks: []models.InputRichBlock{
				{
					Type: models.RichBlockTypePhoto,
					InputRichBlockPhoto: &models.InputRichBlockPhoto{
						Photo: models.InputMediaPhoto{
							Media:           "attach://photo.jpg",
							MediaAttachment: strings.NewReader("photo content"),
						},
					},
				},
				{
					Type: models.RichBlockTypeDetails,
					InputRichBlockDetails: &models.InputRichBlockDetails{
						Summary: models.RichText{PlainText: "sum"},
						Blocks: []models.InputRichBlock{
							{
								Type: models.RichBlockTypeVideo,
								InputRichBlockVideo: &models.InputRichBlockVideo{
									Video: models.InputMediaVideo{
										Media:           "attach://video.mp4",
										MediaAttachment: strings.NewReader("video content"),
									},
								},
							},
						},
					},
				},
				{
					Type: models.RichBlockTypeList,
					InputRichBlockList: &models.InputRichBlockList{
						Items: []models.InputRichBlockListItem{
							{
								Blocks: []models.InputRichBlock{
									{
										Type: models.RichBlockTypeAudio,
										InputRichBlockAudio: &models.InputRichBlockAudio{
											Audio: models.InputMediaAudio{
												Media:           "attach://audio.mp3",
												MediaAttachment: strings.NewReader("audio content"),
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	buf := bytes.NewBuffer(nil)
	form := multipart.NewWriter(buf)
	form.SetBoundary("XXX") //nolint

	fieldsCount, errBuild := buildRequestForm(form, &params)
	if errBuild != nil {
		t.Error(errBuild)
		return
	}
	if err := form.Close(); err != nil {
		t.Errorf("failed to close form: %v", err)
	}

	expect := `--XXX
Content-Disposition: form-data; name="photo.jpg"; filename="photo.jpg"
Content-Type: application/octet-stream

photo content
--XXX
Content-Disposition: form-data; name="video.mp4"; filename="video.mp4"
Content-Type: application/octet-stream

video content
--XXX
Content-Disposition: form-data; name="audio.mp3"; filename="audio.mp3"
Content-Type: application/octet-stream

audio content
--XXX
Content-Disposition: form-data; name="rich_message"

{"blocks":[{"type":"photo","photo":{"type":"photo","media":"attach://photo.jpg"}},{"type":"details","summary":"sum","blocks":[{"type":"video","video":{"type":"video","media":"attach://video.mp4"}}]},{"type":"list","items":[{"blocks":[{"type":"audio","audio":{"type":"audio","media":"attach://audio.mp3"}}]}]}]}
--XXX--
`
	assertEqualInt(t, fieldsCount, 1)
	assertFormData(t, buf.String(), expect)
}

func Test_addFormFieldInputRichMessage_nilAttachment(t *testing.T) {
	writer := multipart.NewWriter(bytes.NewBuffer(nil))
	form := newRequestForm(writer)
	err := addFormFieldInputRichMessage(form, "rich_message", &models.InputRichMessage{
		Media: []models.InputRichMessageMedia{
			{ID: "doc1", Media: &models.InputMediaDocument{Media: "attach://doc1.pdf"}},
		},
	})
	if err == nil {
		t.Fatal("expected error for attach:// with nil MediaAttachment")
	}
	if !strings.Contains(err.Error(), "nil attachment") {
		t.Fatalf("unexpected error: %v", err)
	}
}

func Test_buildRequestForm_inputMediaThumbnail(t *testing.T) {
	params := struct {
		Media     []models.InputMedia     `json:"media"`
		PaidMedia []models.InputPaidMedia `json:"paid_media"`
	}{
		Media: []models.InputMedia{
			&models.InputMediaVideo{
				Media:           "attach://video.mp4",
				MediaAttachment: strings.NewReader("video content"),
				Thumbnail:       &models.InputFileUpload{Filename: "thumb.jpg", Data: strings.NewReader("thumb content")},
			},
			&models.InputMediaAudio{
				Media:     "file_id",
				Thumbnail: &models.InputFileString{Data: "https://domain.com/thumb.jpg"},
			},
		},
		PaidMedia: []models.InputPaidMedia{
			&models.InputPaidMediaVideo{
				Media:     "file_id",
				Thumbnail: &models.InputFileUpload{Filename: "paid_thumb.jpg", Data: strings.NewReader("paid thumb content")},
			},
		},
	}

	buf := bytes.NewBuffer(nil)
	form := multipart.NewWriter(buf)
	form.SetBoundary("XXX") //nolint

	fieldsCount, errBuild := buildRequestForm(form, &params)
	if errBuild != nil {
		t.Error(errBuild)
		return
	}
	if err := form.Close(); err != nil {
		t.Errorf("failed to close form: %v", err)
	}

	expect := `--XXX
Content-Disposition: form-data; name="video.mp4"; filename="video.mp4"
Content-Type: application/octet-stream

video content
--XXX
Content-Disposition: form-data; name="thumb.jpg"; filename="thumb.jpg"
Content-Type: application/octet-stream

thumb content
--XXX
Content-Disposition: form-data; name="media"

[{"type":"video","media":"attach://video.mp4","thumbnail":"attach://thumb.jpg"},{"type":"audio","media":"file_id","thumbnail":"https://domain.com/thumb.jpg"}]
--XXX
Content-Disposition: form-data; name="paid_thumb.jpg"; filename="paid_thumb.jpg"
Content-Type: application/octet-stream

paid thumb content
--XXX
Content-Disposition: form-data; name="paid_media"

[{"type":"video","media":"file_id","thumbnail":"attach://paid_thumb.jpg"}]
--XXX--
`
	assertEqualInt(t, fieldsCount, 2)
	assertFormData(t, buf.String(), expect)
}

func Test_addFormFieldInputMedia_nilThumbnailData(t *testing.T) {
	writer := multipart.NewWriter(bytes.NewBuffer(nil))
	form := newRequestForm(writer)
	err := addFormFieldInputMedia(form, "media", &models.InputMediaVideo{
		Media:     "file_id",
		Thumbnail: &models.InputFileUpload{Filename: "thumb.jpg"},
	})
	if err == nil {
		t.Fatal("expected error for thumbnail with nil Data")
	}
	if !strings.Contains(err.Error(), "nil data for attach://thumb.jpg") {
		t.Fatalf("unexpected error: %v", err)
	}
}

type errWriter struct{}

func (errWriter) Write([]byte) (int, error) {
	return 0, errors.New("write failed")
}

// formFileNames returns the name of every file part of a built form, in order.
func formFileNames(t *testing.T, body, boundary string) string {
	t.Helper()

	var names []string
	r := multipart.NewReader(strings.NewReader(body), boundary)
	for {
		part, err := r.NextPart()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			t.Fatal(err)
		}
		if part.FileName() != "" {
			names = append(names, part.FormName())
		}
	}

	return strings.Join(names, ",")
}

func richDocumentBlock(name string) models.InputRichBlock {
	return models.InputRichBlock{
		Type: models.RichBlockTypeDocument,
		InputRichBlockDocument: &models.InputRichBlockDocument{
			Document: models.InputMediaDocument{
				Media:           "attach://" + name,
				MediaAttachment: strings.NewReader(name + " content"),
			},
		},
	}
}

func Test_addInputRichBlockAttachments_mediaBlocks(t *testing.T) {
	blocks := []models.InputRichBlock{
		{
			Type: models.RichBlockTypeAnimation,
			InputRichBlockAnimation: &models.InputRichBlockAnimation{
				Animation: models.InputMediaAnimation{Media: "attach://animation.gif", MediaAttachment: strings.NewReader("animation")},
			},
		},
		{
			Type: models.RichBlockTypeAudio,
			InputRichBlockAudio: &models.InputRichBlockAudio{
				Audio: models.InputMediaAudio{Media: "attach://audio.mp3", MediaAttachment: strings.NewReader("audio")},
			},
		},
		richDocumentBlock("document.pdf"),
		{
			Type: models.RichBlockTypePhoto,
			InputRichBlockPhoto: &models.InputRichBlockPhoto{
				Photo: models.InputMediaPhoto{Media: "attach://photo.jpg", MediaAttachment: strings.NewReader("photo")},
			},
		},
		{
			Type: models.RichBlockTypeVideo,
			InputRichBlockVideo: &models.InputRichBlockVideo{
				Video: models.InputMediaVideo{
					Media:           "attach://video.mp4",
					MediaAttachment: strings.NewReader("video"),
					Thumbnail:       &models.InputFileUpload{Filename: "video_thumb.jpg", Data: strings.NewReader("thumb")},
				},
			},
		},
		{
			Type: models.RichBlockTypeVoiceNote,
			InputRichBlockVoiceNote: &models.InputRichBlockVoiceNote{
				VoiceNote: models.InputMediaVoiceNote{Media: "attach://voice.ogg", MediaAttachment: strings.NewReader("voice")},
			},
		},
	}

	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	form := newRequestForm(writer)
	writer.SetBoundary("XXX") //nolint

	if err := addInputRichBlockSliceAttachments(form, blocks); err != nil {
		t.Fatal(err)
	}
	if err := writer.Close(); err != nil {
		t.Fatal(err)
	}

	assertEqualString(t,
		formFileNames(t, buf.String(), "XXX"),
		"animation.gif,audio.mp3,document.pdf,photo.jpg,video.mp4,video_thumb.jpg,voice.ogg",
	)
}

func Test_addInputRichBlockAttachments_containerBlocks(t *testing.T) {
	blocks := []models.InputRichBlock{
		{
			Type: models.RichBlockTypeList,
			InputRichBlockList: &models.InputRichBlockList{
				Items: []models.InputRichBlockListItem{{Blocks: []models.InputRichBlock{richDocumentBlock("list.pdf")}}},
			},
		},
		{
			Type: models.RichBlockTypeBlockQuotation,
			InputRichBlockBlockQuotation: &models.InputRichBlockBlockQuotation{
				Blocks: []models.InputRichBlock{richDocumentBlock("quote.pdf")},
			},
		},
		{
			Type: models.RichBlockTypeCollage,
			InputRichBlockCollage: &models.InputRichBlockCollage{
				Blocks: []models.InputRichBlock{richDocumentBlock("collage.pdf")},
			},
		},
		{
			Type: models.RichBlockTypeSlideshow,
			InputRichBlockSlideshow: &models.InputRichBlockSlideshow{
				Blocks: []models.InputRichBlock{richDocumentBlock("slideshow.pdf")},
			},
		},
		{
			Type: models.RichBlockTypeDetails,
			InputRichBlockDetails: &models.InputRichBlockDetails{
				Blocks: []models.InputRichBlock{richDocumentBlock("details.pdf")},
			},
		},
	}

	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	form := newRequestForm(writer)
	writer.SetBoundary("XXX") //nolint

	if err := addInputRichBlockSliceAttachments(form, blocks); err != nil {
		t.Fatal(err)
	}
	if err := writer.Close(); err != nil {
		t.Fatal(err)
	}

	assertEqualString(t,
		formFileNames(t, buf.String(), "XXX"),
		"list.pdf,quote.pdf,collage.pdf,slideshow.pdf,details.pdf",
	)
}

// A block whose variant pointer is not set carries no attachment. The encoder
// reports the broken union; the walk must not panic on the way there.
func Test_addInputRichBlockAttachments_missingVariant(t *testing.T) {
	types := []models.RichBlockType{
		models.RichBlockTypeAnimation,
		models.RichBlockTypeAudio,
		models.RichBlockTypeDocument,
		models.RichBlockTypePhoto,
		models.RichBlockTypeVideo,
		models.RichBlockTypeVoiceNote,
		models.RichBlockTypeList,
		models.RichBlockTypeBlockQuotation,
		models.RichBlockTypeCollage,
		models.RichBlockTypeSlideshow,
		models.RichBlockTypeDetails,
		models.RichBlockTypeParagraph,
	}

	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	form := newRequestForm(writer)
	writer.SetBoundary("XXX") //nolint

	for _, blockType := range types {
		if err := addInputRichBlockAttachments(form, models.InputRichBlock{Type: blockType}); err != nil {
			t.Fatalf("%s: %v", blockType, err)
		}
	}
	if err := writer.Close(); err != nil {
		t.Fatal(err)
	}

	assertEqualString(t, formFileNames(t, buf.String(), "XXX"), "")
}

func Test_addFormFieldInputRichMessage_nilAttachmentInNestedBlock(t *testing.T) {
	writer := multipart.NewWriter(bytes.NewBuffer(nil))
	form := newRequestForm(writer)
	err := addFormFieldInputRichMessage(form, "rich_message", &models.InputRichMessage{
		Blocks: []models.InputRichBlock{
			{
				Type: models.RichBlockTypeCollage,
				InputRichBlockCollage: &models.InputRichBlockCollage{
					Blocks: []models.InputRichBlock{
						{
							Type: models.RichBlockTypeDocument,
							InputRichBlockDocument: &models.InputRichBlockDocument{
								Document: models.InputMediaDocument{Media: "attach://doc1.pdf"},
							},
						},
					},
				},
			},
		},
	})
	if err == nil {
		t.Fatal("expected error for attach:// with nil MediaAttachment")
	}
	if !strings.Contains(err.Error(), "nil attachment for attach://doc1.pdf") {
		t.Fatalf("unexpected error: %v", err)
	}
}

func Test_addInputRichBlockAttachments_nilAttachmentInListItem(t *testing.T) {
	writer := multipart.NewWriter(bytes.NewBuffer(nil))
	form := newRequestForm(writer)
	err := addInputRichBlockAttachments(form, models.InputRichBlock{
		Type: models.RichBlockTypeList,
		InputRichBlockList: &models.InputRichBlockList{
			Items: []models.InputRichBlockListItem{
				{
					Blocks: []models.InputRichBlock{
						{
							Type: models.RichBlockTypeDocument,
							InputRichBlockDocument: &models.InputRichBlockDocument{
								Document: models.InputMediaDocument{Media: "attach://doc1.pdf"},
							},
						},
					},
				},
			},
		},
	})
	if err == nil {
		t.Fatal("expected error for attach:// with nil MediaAttachment")
	}
	if !strings.Contains(err.Error(), "nil attachment for attach://doc1.pdf") {
		t.Fatalf("unexpected error: %v", err)
	}
}

// The media of an InputRichMessageMedia is required, so a missing one cannot be
// encoded as a null and left to the server to reject.
func Test_addFormFieldInputRichMessage_nilMedia(t *testing.T) {
	form := newRequestForm(multipart.NewWriter(bytes.NewBuffer(nil)))

	err := addFormFieldInputRichMessage(form, "rich_message", &models.InputRichMessage{
		Media: []models.InputRichMessageMedia{{ID: "doc1"}},
	})
	if err == nil {
		t.Fatal("expected error for an entry with no media")
	}
	if !strings.Contains(err.Error(), "nil media for field rich_message at index 0") {
		t.Fatalf("unexpected error: %v", err)
	}
}

func Test_addFormFieldInputMedia_livePhoto(t *testing.T) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	form := newRequestForm(writer)
	writer.SetBoundary("XXX") //nolint

	err := addFormFieldInputMedia(form, "media", &models.InputMediaLivePhoto{
		Media:           "attach://live.jpg",
		MediaAttachment: strings.NewReader("still content"),
		Photo:           "attach://live.mov",
		PhotoAttachment: strings.NewReader("motion content"),
	})
	if err != nil {
		t.Fatal(err)
	}
	if err := writer.Close(); err != nil {
		t.Fatal(err)
	}

	assertEqualString(t, formFileNames(t, buf.String(), "XXX"), "live.jpg,live.mov")
}

func Test_addFormFieldInputMedia_livePhotoNilPhotoAttachment(t *testing.T) {
	writer := multipart.NewWriter(bytes.NewBuffer(nil))
	form := newRequestForm(writer)
	err := addFormFieldInputMedia(form, "media", &models.InputMediaLivePhoto{
		Media: "file_id",
		Photo: "attach://live.mov",
	})
	if err == nil {
		t.Fatal("expected error for attach:// with nil PhotoAttachment")
	}
	if !strings.Contains(err.Error(), "nil PhotoAttachment for attach://live.mov") {
		t.Fatalf("unexpected error: %v", err)
	}
}

func Test_addInputMediaAttachment_readErrors(t *testing.T) {
	t.Run("media", func(t *testing.T) {
		writer := multipart.NewWriter(bytes.NewBuffer(nil))
		form := newRequestForm(writer)
		err := addInputMediaAttachment(form, &models.InputMediaPhoto{
			Media:           "attach://photo.jpg",
			MediaAttachment: errReader(errors.New("read failed")),
		})
		if err == nil || !strings.Contains(err.Error(), "read failed") {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("live photo", func(t *testing.T) {
		writer := multipart.NewWriter(bytes.NewBuffer(nil))
		form := newRequestForm(writer)
		err := addInputMediaAttachment(form, &models.InputMediaLivePhoto{
			Media:           "file_id",
			Photo:           "attach://live.mov",
			PhotoAttachment: errReader(errors.New("read failed")),
		})
		if err == nil || !strings.Contains(err.Error(), "read failed") {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("thumbnail", func(t *testing.T) {
		writer := multipart.NewWriter(bytes.NewBuffer(nil))
		form := newRequestForm(writer)
		err := addInputMediaAttachment(form, &models.InputMediaVideo{
			Media:     "file_id",
			Thumbnail: &models.InputFileUpload{Filename: "thumb.jpg", Data: errReader(errors.New("read failed"))},
		})
		if err == nil || !strings.Contains(err.Error(), "read failed") {
			t.Fatalf("unexpected error: %v", err)
		}
	})
}

func Test_addInputMediaAttachment_writeErrors(t *testing.T) {
	t.Run("media", func(t *testing.T) {
		writer := multipart.NewWriter(errWriter{})
		form := newRequestForm(writer)
		err := addInputMediaAttachment(form, &models.InputMediaPhoto{
			Media:           "attach://photo.jpg",
			MediaAttachment: strings.NewReader("photo"),
		})
		if err == nil || !strings.Contains(err.Error(), "write failed") {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("live photo", func(t *testing.T) {
		writer := multipart.NewWriter(errWriter{})
		form := newRequestForm(writer)
		err := addInputMediaAttachment(form, &models.InputMediaLivePhoto{
			Media:           "file_id",
			Photo:           "attach://live.mov",
			PhotoAttachment: strings.NewReader("motion"),
		})
		if err == nil || !strings.Contains(err.Error(), "write failed") {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("thumbnail", func(t *testing.T) {
		writer := multipart.NewWriter(errWriter{})
		form := newRequestForm(writer)
		err := addInputMediaAttachment(form, &models.InputMediaVideo{
			Media:     "file_id",
			Thumbnail: &models.InputFileUpload{Filename: "thumb.jpg", Data: strings.NewReader("thumb")},
		})
		if err == nil || !strings.Contains(err.Error(), "write failed") {
			t.Fatalf("unexpected error: %v", err)
		}
	})
}

// A thumbnail left as a typed nil pointer is not a thumbnail: it carries no data and
// must not panic, and the key is omitted rather than sent as a null.
func Test_buildRequestForm_typedNilThumbnail(t *testing.T) {
	var thumb *models.InputFileUpload

	params := SendMediaGroupParams{
		ChatID: 1,
		Media: []models.InputMedia{
			&models.InputMediaVideo{
				Media:           "attach://video.mp4",
				MediaAttachment: strings.NewReader("video content"),
				Thumbnail:       thumb,
			},
		},
	}

	buf := bytes.NewBuffer(nil)
	form := multipart.NewWriter(buf)
	form.SetBoundary("XXX") //nolint

	if _, err := buildRequestForm(form, &params); err != nil {
		t.Fatal(err)
	}
	if err := form.Close(); err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(buf.String(), `[{"type":"video","media":"attach://video.mp4"}]`) {
		t.Fatalf("unexpected form data:\n%s", buf.String())
	}
}

// An InputRichMessageMedia holding a typed nil pointer carries no media either.
func Test_addFormFieldInputRichMessage_typedNilMedia(t *testing.T) {
	var doc *models.InputMediaDocument

	form := newRequestForm(multipart.NewWriter(bytes.NewBuffer(nil)))

	err := addFormFieldInputRichMessage(form, "rich_message", &models.InputRichMessage{
		Media: []models.InputRichMessageMedia{{ID: "doc1", Media: doc}},
	})
	if err == nil {
		t.Fatal("expected error for a typed nil media")
	}
	if !strings.Contains(err.Error(), "nil media for field rich_message at index 0") {
		t.Fatalf("unexpected error: %v", err)
	}
}

// Two attachments sharing a part name would silently resolve to one file.
func Test_buildRequestForm_duplicateThumbnailName(t *testing.T) {
	params := SendMediaGroupParams{
		ChatID: 1,
		Media: []models.InputMedia{
			&models.InputMediaVideo{
				Media:           "attach://a.mp4",
				MediaAttachment: strings.NewReader("a content"),
				Thumbnail:       &models.InputFileUpload{Filename: "thumb.jpg", Data: strings.NewReader("thumb a")},
			},
			&models.InputMediaVideo{
				Media:           "attach://b.mp4",
				MediaAttachment: strings.NewReader("b content"),
				Thumbnail:       &models.InputFileUpload{Filename: "thumb.jpg", Data: strings.NewReader("thumb b")},
			},
		},
	}

	form := multipart.NewWriter(bytes.NewBuffer(nil))

	_, err := buildRequestForm(form, &params)
	if err == nil {
		t.Fatal("expected error for two parts named thumb.jpg")
	}
	if !strings.Contains(err.Error(), `duplicate form part name "thumb.jpg"`) {
		t.Fatalf("unexpected error: %v", err)
	}
}

func Test_buildRequestForm_duplicateMediaAttachName(t *testing.T) {
	params := SendMediaGroupParams{
		ChatID: 1,
		Media: []models.InputMedia{
			&models.InputMediaPhoto{Media: "attach://photo.jpg", MediaAttachment: strings.NewReader("a")},
			&models.InputMediaPhoto{Media: "attach://photo.jpg", MediaAttachment: strings.NewReader("b")},
		},
	}

	form := multipart.NewWriter(bytes.NewBuffer(nil))

	_, err := buildRequestForm(form, &params)
	if err == nil {
		t.Fatal("expected error for two parts named photo.jpg")
	}
	if !strings.Contains(err.Error(), `duplicate form part name "photo.jpg"`) {
		t.Fatalf("unexpected error: %v", err)
	}
}

// A typed nil pointer in a top level InputFile field carries no data. It must be
// reported, not dereferenced: the form is built in a goroutine with no recover.
func Test_buildRequestForm_topLevelTypedNilInputFile(t *testing.T) {
	var thumb *models.InputFileUpload

	params := SendVideoParams{
		ChatID:    1,
		Video:     &models.InputFileString{Data: "file_id"},
		Thumbnail: thumb,
	}

	form := multipart.NewWriter(bytes.NewBuffer(nil))

	_, err := buildRequestForm(form, &params)
	if err == nil {
		t.Fatal("expected error for typed nil thumbnail")
	}
	if !strings.Contains(err.Error(), "nil value for field thumbnail") {
		t.Fatalf("unexpected error: %v", err)
	}
}

// Same for a typed nil element of a media slice.
func Test_buildRequestForm_typedNilMediaSliceItem(t *testing.T) {
	params := SendMediaGroupParams{
		ChatID: 1,
		Media: []models.InputMedia{
			&models.InputMediaPhoto{Media: "file_id"},
			(*models.InputMediaPhoto)(nil),
		},
	}

	form := multipart.NewWriter(bytes.NewBuffer(nil))

	_, err := buildRequestForm(form, &params)
	if err == nil {
		t.Fatal("expected error for typed nil media item")
	}
	if !strings.Contains(err.Error(), "nil value for field media at index 1") {
		t.Fatalf("unexpected error: %v", err)
	}
}

// A typed nil in a single InputMedia field is the same class of value.
func Test_buildRequestForm_typedNilMediaField(t *testing.T) {
	params := EditMessageMediaParams{
		Media: (*models.InputMediaPhoto)(nil),
	}

	form := multipart.NewWriter(bytes.NewBuffer(nil))

	_, err := buildRequestForm(form, &params)
	if err == nil {
		t.Fatal("expected error for typed nil media field")
	}
	if !strings.Contains(err.Error(), "nil value for field media") {
		t.Fatalf("unexpected error: %v", err)
	}
}

// One file referenced from two entries under a single attach:// name is a single
// part, not an ambiguity: the part already written is reused.
func Test_buildRequestForm_repeatedAttachNameSameReader(t *testing.T) {
	photo := strings.NewReader("photo content")

	params := SendMediaGroupParams{
		ChatID: 1,
		Media: []models.InputMedia{
			&models.InputMediaPhoto{Media: "attach://p.jpg", MediaAttachment: photo},
			&models.InputMediaPhoto{Media: "attach://p.jpg", MediaAttachment: photo},
		},
	}

	buf := bytes.NewBuffer(nil)
	form := multipart.NewWriter(buf)
	form.SetBoundary("XXX") //nolint

	if _, err := buildRequestForm(form, &params); err != nil {
		t.Fatal(err)
	}
	if err := form.Close(); err != nil {
		t.Fatal(err)
	}

	assertEqualString(t, formFileNames(t, buf.String(), "XXX"), "p.jpg")
}

// The second entry may leave the attachment out entirely and just carry the reference.
func Test_buildRequestForm_repeatedAttachNameWithoutReader(t *testing.T) {
	params := SendMediaGroupParams{
		ChatID: 1,
		Media: []models.InputMedia{
			&models.InputMediaPhoto{Media: "attach://p.jpg", MediaAttachment: strings.NewReader("photo content")},
			&models.InputMediaPhoto{Media: "attach://p.jpg"},
		},
	}

	buf := bytes.NewBuffer(nil)
	form := multipart.NewWriter(buf)
	form.SetBoundary("XXX") //nolint

	if _, err := buildRequestForm(form, &params); err != nil {
		t.Fatal(err)
	}
	if err := form.Close(); err != nil {
		t.Fatal(err)
	}

	assertEqualString(t, formFileNames(t, buf.String(), "XXX"), "p.jpg")
}

// One sticker image reused across two entries of a set.
func Test_buildRequestForm_repeatedStickerAttachName(t *testing.T) {
	sticker := strings.NewReader("sticker content")

	params := CreateNewStickerSetParams{
		Name:  "set",
		Title: "set",
		Stickers: []models.InputSticker{
			{Sticker: "attach://s.png", StickerAttachment: sticker, EmojiList: []string{"a"}},
			{Sticker: "attach://s.png", StickerAttachment: sticker, EmojiList: []string{"b"}},
		},
	}

	buf := bytes.NewBuffer(nil)
	form := multipart.NewWriter(buf)
	form.SetBoundary("XXX") //nolint

	if _, err := buildRequestForm(form, &params); err != nil {
		t.Fatal(err)
	}
	if err := form.Close(); err != nil {
		t.Fatal(err)
	}

	assertEqualString(t, formFileNames(t, buf.String(), "XXX"), "s.png")
}

// Two different files still cannot share one name.
func Test_buildRequestForm_repeatedAttachNameDifferentReader(t *testing.T) {
	params := SendMediaGroupParams{
		ChatID: 1,
		Media: []models.InputMedia{
			&models.InputMediaPhoto{Media: "attach://p.jpg", MediaAttachment: strings.NewReader("a")},
			&models.InputMediaPhoto{Media: "attach://p.jpg", MediaAttachment: strings.NewReader("b")},
		},
	}

	form := multipart.NewWriter(bytes.NewBuffer(nil))

	_, err := buildRequestForm(form, &params)
	if err == nil {
		t.Fatal("expected error for two files named p.jpg")
	}
	if !strings.Contains(err.Error(), `duplicate form part name "p.jpg"`) {
		t.Fatalf("unexpected error: %v", err)
	}
}

// A file part and a form field sharing a name is the ambiguity the check exists for.
func Test_buildRequestForm_filePartCollidesWithField(t *testing.T) {
	params := EditMessageMediaParams{
		Media: &models.InputMediaVideo{Media: "attach://media", MediaAttachment: strings.NewReader("video")},
	}

	form := multipart.NewWriter(bytes.NewBuffer(nil))

	_, err := buildRequestForm(form, &params)
	if err == nil {
		t.Fatal("expected error for a file part named like a form field")
	}
	if !strings.Contains(err.Error(), `duplicate form part name "media"`) {
		t.Fatalf("unexpected error: %v", err)
	}
}

// A nested upload with no Filename would be written as a part named "" and referenced
// as "attach://", which Telegram can only answer with an opaque Bad Request.
func Test_addInputFileAttachment_emptyFilename(t *testing.T) {
	form := newRequestForm(multipart.NewWriter(bytes.NewBuffer(nil)))

	err := addInputFileAttachment(form, &models.InputFileUpload{Data: strings.NewReader("thumb")})
	if err == nil {
		t.Fatal("expected error for an upload with no filename")
	}
	if !strings.Contains(err.Error(), "empty filename") {
		t.Fatalf("unexpected error: %v", err)
	}
}

// sliceReader is a reader that cannot be compared, so it can never be recognised as
// one file referenced twice.
type sliceReader []byte

func (sliceReader) Read([]byte) (int, error) {
	return 0, io.EOF
}

func Test_requestForm_addFilePart_incomparableReader(t *testing.T) {
	form := newRequestForm(multipart.NewWriter(bytes.NewBuffer(nil)))

	if err := form.addFilePart("p.jpg", "p.jpg", sliceReader("a")); err != nil {
		t.Fatal(err)
	}

	err := form.addFilePart("p.jpg", "p.jpg", sliceReader("a"))
	if err == nil {
		t.Fatal("expected error for a second reader under one part name")
	}
	if !strings.Contains(err.Error(), `duplicate form part name "p.jpg"`) {
		t.Fatalf("unexpected error: %v", err)
	}
}

func Test_requestForm_addFieldPart_duplicate(t *testing.T) {
	form := newRequestForm(multipart.NewWriter(bytes.NewBuffer(nil)))

	if err := form.addFieldPart("media", strings.NewReader("a")); err != nil {
		t.Fatal(err)
	}

	err := form.addFieldPart("media", strings.NewReader("b"))
	if err == nil {
		t.Fatal("expected error for two fields named media")
	}
	if !strings.Contains(err.Error(), `duplicate form part name "media"`) {
		t.Fatalf("unexpected error: %v", err)
	}
}

// A form field cannot take the name of a file part written before it either.
func Test_requestForm_addFilePart_collidesWithField(t *testing.T) {
	form := newRequestForm(multipart.NewWriter(bytes.NewBuffer(nil)))

	if err := form.addFieldPart("media", strings.NewReader("a")); err != nil {
		t.Fatal(err)
	}

	err := form.addFilePart("media", "media", strings.NewReader("b"))
	if err == nil {
		t.Fatal("expected error for a file part named like a form field")
	}
	if !strings.Contains(err.Error(), `duplicate form part name "media"`) {
		t.Fatalf("unexpected error: %v", err)
	}
}
