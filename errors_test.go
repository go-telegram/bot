package bot

import (
	"errors"
	"testing"
)

func TestTooManyRequestsError(t *testing.T) {
	err := &TooManyRequestsError{
		Message:    "rate limit exceeded",
		RetryAfter: 30,
	}

	expectedErrorMsg := "rate limit exceeded: retry_after 30"
	if err.Error() != expectedErrorMsg {
		t.Errorf("expected %s, got %s", expectedErrorMsg, err.Error())
	}

	if !IsTooManyRequestsError(err) {
		t.Errorf("expected IsTooManyRequestsError to return true")
	}

	var genericError error = err
	if !IsTooManyRequestsError(genericError) {
		t.Errorf("expected IsTooManyRequestsError to return true for generic error type")
	}
}

func TestMigrateError(t *testing.T) {
	err := &MigrateError{
		Message:         "chat migrated",
		MigrateToChatID: 12345,
	}

	expectedErrorMsg := "chat migrated: migrate_to_chat_id 12345"
	if err.Error() != expectedErrorMsg {
		t.Errorf("expected %s, got %s", expectedErrorMsg, err.Error())
	}

	if !IsMigrateError(err) {
		t.Errorf("expected IsMigrateError to return true")
	}

	var genericError error = err
	if !IsMigrateError(genericError) {
		t.Errorf("expected IsMigrateError to return true for generic error type")
	}
}

func TestStandardErrors(t *testing.T) {
	tests := []struct {
		err      error
		expected string
	}{
		{ErrorForbidden, "forbidden"},
		{ErrorBadRequest, "bad request"},
		{ErrorUnauthorized, "unauthorized"},
		{ErrorTooManyRequests, "too many requests"},
		{ErrorNotFound, "not found"},
		{ErrorConflict, "conflict"},
	}

	for _, tt := range tests {
		if tt.err.Error() != tt.expected {
			t.Errorf("expected %s, got %s", tt.expected, tt.err.Error())
		}
	}
}

func TestIsTooManyRequestsErrorFalse(t *testing.T) {
	err := errors.New("some other error")
	if IsTooManyRequestsError(err) {
		t.Errorf("expected IsTooManyRequestsError to return false")
	}
}

func TestIsMigrateErrorFalse(t *testing.T) {
	err := errors.New("some other error")
	if IsMigrateError(err) {
		t.Errorf("expected IsMigrateError to return false")
	}
}
