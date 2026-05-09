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

// rawRequest sends one Telegram Bot API call as a multipart/form-data POST.
//
// The body is built into a bytes.Buffer up front, then handed to
// http.NewRequestWithContext as a *bytes.Reader. That lets net/http
// auto-populate Request.GetBody and Request.ContentLength, which the
// http2.Transport needs in order to retry the request when Telegram's
// server sends a GOAWAY frame mid-flight (a routine part of HTTP/2
// connection draining). The earlier io.Pipe-based body could not be
// replayed, so every POST issued on a draining connection failed with
// "cannot retry err [http2: ...] after Request.Body was written;
// define Request.GetBody to avoid this error". Long-poll getUpdates
// (a GET, no body) was unaffected, so a bot would keep receiving
// updates while silently failing every reply.
//
// Trade-off: the entire request body is held in memory until the
// request is sent. For the vast majority of methods (sendMessage,
// editMessageText, callback answers, …) the body is a few KB. For
// file uploads the body is bounded by Telegram's bot-API per-request
// limit (~50 MB) and only held for the duration of the upload.
func (b *Bot) rawRequest(ctx context.Context, method string, params any, dest any) error {
	var bodyBuf bytes.Buffer
	form := multipart.NewWriter(&bodyBuf)

	if params != nil && !reflect.ValueOf(params).IsNil() {
		if _, errFormData := buildRequestForm(form, params); errFormData != nil {
			return fmt.Errorf("error build request form for method %s, %w", method, errFormData)
		}
		if errFormClose := form.Close(); errFormClose != nil {
			return fmt.Errorf("error form close for method %s, %w", method, errFormClose)
		}
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

	req, errRequest := http.NewRequestWithContext(ctx, http.MethodPost, u, bytes.NewReader(bodyBuf.Bytes()))
	if errRequest != nil {
		return fmt.Errorf("error create request for method %s, %w", method, errRequest)
	}

	req.Header.Add("Content-Type", form.FormDataContentType())

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
