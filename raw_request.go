package bot

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
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

// rawRequest materialises the multipart body up front so net/http can set
// Request.ContentLength and Request.GetBody. Without them http2.Transport cannot
// replay a POST after the server sends GOAWAY, and every call on a draining
// connection fails. The price is that an upload is held in memory for the
// duration of the request. A method with no fields is sent without a body and
// without a Content-Type, as local telegram-bot-api servers reject an empty
// multipart body.
func (b *Bot) rawRequest(ctx context.Context, method string, params any, dest any) error {
	var bodyBuf bytes.Buffer
	form := multipart.NewWriter(&bodyBuf)

	var fieldsCount int
	if params != nil && !reflect.ValueOf(params).IsNil() {
		var errFormData error
		fieldsCount, errFormData = buildRequestForm(form, params)
		if errFormData != nil {
			return fmt.Errorf("error build request form for method %s, %w", method, errFormData)
		}
	}

	var requestBody io.Reader = http.NoBody
	var contentType string
	if fieldsCount > 0 {
		if errFormClose := form.Close(); errFormClose != nil {
			return fmt.Errorf("error form close for method %s, %w", method, errFormClose)
		}
		requestBody = bytes.NewReader(bodyBuf.Bytes())
		contentType = form.FormDataContentType()
	}

	u := b.url + "/bot" + b.token + "/"
	if b.testEnvironment {
		u += "test/"
	}
	u += method

	if b.isDebug && strings.ToLower(method) != "getupdates" {
		requestDebugData, _ := json.Marshal(params)
		b.debugHandler("request url: %s, payload: %s", strings.Replace(u, b.token, "***", 1), requestDebugData)
	}

	req, errRequest := http.NewRequestWithContext(ctx, http.MethodPost, u, requestBody)
	if errRequest != nil {
		return fmt.Errorf("error create request for method %s, %w", method, errRequest)
	}

	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}

	resp, errDo := b.client.Do(req)
	if errDo != nil {
		var netErr *url.Error
		if errors.As(errDo, &netErr) {
			netErr.URL = strings.Replace(netErr.URL, b.token, "***", -1)
		}

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
		case http.StatusForbidden:
			return fmt.Errorf("%w, %s", ErrorForbidden, r.Description)
		case http.StatusBadRequest:
			if r.Parameters.MigrateToChatID != 0 {
				err := &MigrateError{
					Message:         fmt.Sprintf("%s: %s", ErrorBadRequest, r.Description),
					MigrateToChatID: r.Parameters.MigrateToChatID,
				}

				return err
			}
			return fmt.Errorf("%w, %s", ErrorBadRequest, r.Description)
		case http.StatusUnauthorized:
			return fmt.Errorf("%w, %s", ErrorUnauthorized, r.Description)
		case http.StatusNotFound:
			return fmt.Errorf("%w, %s", ErrorNotFound, r.Description)
		case http.StatusConflict:
			return fmt.Errorf("%w, %s", ErrorConflict, r.Description)
		case http.StatusTooManyRequests:
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
			b.debugHandler("response from '%s' with payload '%s'", strings.Replace(u, b.token, "***", 1), body)
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
