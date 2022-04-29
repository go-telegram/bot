package bot

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type apiResponse struct {
	OK          bool            `json:"ok"`
	Result      json.RawMessage `json:"result,omitempty"`
	Description string          `json:"description,omitempty"`
}

func RawRequest(ctx context.Context, b *Bot, method string, params any, dest any) error {
	errValidate := paramsValidate(params)
	if errValidate != nil {
		return fmt.Errorf("error params validate for method %s, %w", method, errValidate)
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
