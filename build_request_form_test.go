package bot

import (
	"bytes"
	"mime/multipart"
	"strings"
	"testing"

	"github.com/go-telegram/bot/models"
)

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
		NoJSONTag1:                 "foo",
		NoJSONTag2:                 "bar",
		OmitEmptyString:            "",
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
--XXX--
`
	assertEqualInt(t, fieldsCount, 6)
	assertFormData(t, buf.String(), expect)
}
