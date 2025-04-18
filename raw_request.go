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
	ErrorCode   int             `json:"error_code,omitempty"`
	Parameters  struct {
		RetryAfter      int `json:"retry_after,omitempty"`
		MigrateToChatID int `json:"migrate_to_chat_id,omitempty"`
	} `json:"parameters,omitempty"`
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

	u := b.url + "/bot" + b.token + "/"
	if b.testEnvironment {
		u += "test/"
	}
	u += method

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
	defer func() {
		if err := resp.Body.Close(); err != nil {
			b.errorsHandler(fmt.Errorf("failed to close response body: %w", err))
		}
	}()

	body, errReadBody := io.ReadAll(resp.Body)
	if errReadBody != nil {
		return fmt.Errorf("error read response body for method %s, %w", method, errReadBody)
	}

	r := apiResponse{}

	errDecode := json.Unmarshal(body, &r)
	if errDecode != nil {
		return fmt.Errorf("error decode response body for method %s, %s, %w", method, body, errDecode)
	}

	if !r.OK {
		switch r.ErrorCode {
		case 403:
			return fmt.Errorf("%w, %s", ErrorForbidden, r.Description)
		case 400:
			if r.Parameters.MigrateToChatID != 0 {
				err := &MigrateError{
					Message:         fmt.Sprintf("%s: %s", ErrorBadRequest, r.Description),
					MigrateToChatID: r.Parameters.MigrateToChatID,
				}

				return err
			}
			return fmt.Errorf("%w, %s", ErrorBadRequest, r.Description)
		case 401:
			return fmt.Errorf("%w, %s", ErrorUnauthorized, r.Description)
		case 404:
			return fmt.Errorf("%w, %s", ErrorNotFound, r.Description)
		case 409:
			return fmt.Errorf("%w, %s", ErrorConflict, r.Description)
		case 429:
			err := &TooManyRequestsError{
				Message:    fmt.Sprintf("%s, %s", ErrorTooManyRequests, r.Description),
				RetryAfter: r.Parameters.RetryAfter,
			}
			return err
		default:
			return fmt.Errorf("error response from telegram for method %s, %d %s", method, r.ErrorCode, r.Description)
		}
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
