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
)

type apiResponse struct {
	OK          bool            `json:"ok"`
	Result      json.RawMessage `json:"result,omitempty"`
	Description string          `json:"description,omitempty"`
}

func (b *Bot) rawRequest(ctx context.Context, method string, params any, dest any) error {
	var httpBody io.Reader = http.NoBody
	var contentType string

	if params != nil && !reflect.ValueOf(params).IsNil() {
		buf := bytes.NewBuffer(nil)
		form := multipart.NewWriter(buf)

		fieldsCount, errFormData := buildRequestForm(form, params)
		if errFormData != nil {
			return fmt.Errorf("error build request form for method %s, %w", method, errFormData)
		}

		errFormClose := form.Close()
		if errFormClose != nil {
			return fmt.Errorf("error form close for method %s, %w", method, errFormClose)
		}

		if fieldsCount > 0 {
			httpBody = buf
			contentType = form.FormDataContentType()
		}
	}

	u := b.url + "/bot" + b.token + "/" + method

	if b.isDebug && strings.ToLower(method) != "getupdates" {
		requestDebugData, _ := json.Marshal(params)
		b.debugHandler("request url: %s, payload: %s", u, requestDebugData)
	}

	req, errRequest := http.NewRequestWithContext(ctx, http.MethodPost, u, httpBody)
	if errRequest != nil {
		return fmt.Errorf("error create request for method %s, %w", method, errRequest)
	}

	if contentType != "" {
		req.Header.Add("Content-Type", contentType)
	}

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
		if b.isDebug {
			b.debugHandler("response from '%s' with payload '%s'", u, body)
		}
	}

	if dest != nil {
		errDecodeDest := json.Unmarshal(r.Result, dest)
		if errDecodeDest != nil {
			return fmt.Errorf("error decode response result for method %s, %w", method, errDecodeDest)
		}
	}

	return nil
}
