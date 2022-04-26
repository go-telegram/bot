package bot

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"reflect"
	"strings"

	"github.com/go-telegram/bot/models"
)

type RequestParams interface {
	Validate() error
}

type apiResponse struct {
	OK          bool            `json:"ok"`
	Result      json.RawMessage `json:"result,omitempty"`
	Description string          `json:"description,omitempty"`
}

func RawRequest(ctx context.Context, b *Bot, method string, params RequestParams, dest any) error {
	if params != nil {
		errValidate := params.Validate()
		if errValidate != nil {
			return fmt.Errorf("error params validate for method %s, %w", method, errValidate)
		}
	}

	contentType, formData, errFormData := buildRequestForm(params)
	if errFormData != nil {
		return fmt.Errorf("error build request form for method %s, %w", method, errFormData)
	}

	u := b.url + "/bot" + b.token + "/" + method

	if b.isDebug && strings.ToLower(method) != "getupdates" {
		requestDebugData, _ := json.Marshal(params)
		b.debug("request url: %s, payload: %s", u, requestDebugData)
	}

	req, errRequest := http.NewRequestWithContext(ctx, http.MethodPost, u, formData)
	if errRequest != nil {
		return fmt.Errorf("error create request for method %s, %w", method, errRequest)
	}

	req.Header.Add("Content-Type", contentType)

	resp, errDo := b.client.Do(req)
	if errDo != nil {
		return fmt.Errorf("error do request for method %s, %w", method, errDo)
	}
	defer resp.Body.Close()

	body, errReadBody := io.ReadAll(resp.Body)
	if errReadBody != nil {
		return fmt.Errorf("error read response body for method %s, %w", method, errReadBody)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected response statusCode %d for method %s, %s", resp.StatusCode, method, body)
	}

	r := apiResponse{}

	errDecode := json.Unmarshal(body, &r)
	if errDecode != nil {
		return fmt.Errorf("error decode response body for method %s, %s, %w", method, body, errDecode)
	}

	if !r.OK {
		return fmt.Errorf("error response from telegram for method %s, %s", method, r.Description)
	}

	if !bytes.Equal(r.Result, []byte("[]")) {
		b.debug("response from '%s' with payload '%s'", u, body)
	}

	if dest != nil {
		errDecodeDest := json.Unmarshal(r.Result, dest)
		if errDecodeDest != nil {
			return fmt.Errorf("error decode response result for method %s, %w", method, errDecodeDest)
		}
	}

	return nil
}

// buildRequestForm builds form-data for request
// if params contains InputFile of type InputFileUpload, it will be added to form-data ad upload file
// It returns content-type and form-data. And error if occurred
func buildRequestForm(params RequestParams) (string, io.Reader, error) {
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
