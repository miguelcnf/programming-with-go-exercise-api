package clock

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCurrentHandler(t *testing.T) {
	t.Run("should return bad request for invalid location", func(t *testing.T) {
		// TODO
	})

	t.Run("should return time in UTC by default", func(t *testing.T) {
		recorder := httptest.NewRecorder()

		req, err := http.NewRequest(http.MethodGet, CurrentEndpoint, nil)
		if err != nil {
			t.Fatalf("expected error to be nil but was: %v", err)
		}

		CurrentHandler(recorder, req)

		// TODO
	})

	t.Run("should return time in specific timezone", func(t *testing.T) {
		// TODO
	})
}
