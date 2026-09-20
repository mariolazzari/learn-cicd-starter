package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	t.Run("returns API key", func(t *testing.T) {
		headers := http.Header{}
		headers.Set("Authorization", "ApiKey test-key")

		req := &http.Request{Header: headers}

		got, err := GetAPIKey(req.Header)
		if err != nil {
			t.Fatalf("GetAPIKey() error = %v", err)
		}

		if got != "test-key" {
			t.Errorf("GetAPIKey() = %q, want %q", got, "test-key")
		}
	})

	t.Run("returns error when API key is missing", func(t *testing.T) {
		req := &http.Request{Header: http.Header{}}

		_, err := GetAPIKey(req.Header)
		if err == nil {
			t.Fatal("GetAPIKey() expected error, got nil")
		}
	})

	t.Run("returns error when API key is missing", func(t *testing.T) {

		t.Fatal("GetAPIKey() fake error")

	})
}
