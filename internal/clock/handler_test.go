package clock

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCurrentHandler(t *testing.T) {
	t.Run("should return bad request for invalid location", func(t *testing.T) {
		recorder := httptest.NewRecorder()

		req, err := http.NewRequest(http.MethodGet, CurrentEndpoint, nil)
		if err != nil {
			t.Fatalf("expected error to be nil but was: %v", err)
		}
		req.Header.Set("WS-Location", "invalid")

		CurrentHandler(recorder, req)

		if recorder.Code != http.StatusBadRequest {
			t.Fatalf("expected http code to be 400 but was: %v", recorder.Code)
		}
	})

	t.Run("should return time in UTC by default", func(t *testing.T) {
		recorder := httptest.NewRecorder()

		req, err := http.NewRequest(http.MethodGet, CurrentEndpoint, nil)
		if err != nil {
			t.Fatalf("expected error to be nil but was: %v", err)
		}

		CurrentHandler(recorder, req)

		if recorder.Code != http.StatusOK {
			t.Fatalf("expected http code to be 200 but was: %v", recorder.Code)
		}

		body, err := ioutil.ReadAll(recorder.Result().Body)
		if err != nil {
			t.Fatalf("expected error to be nil but was: %v", err)
		}

		var resp Time
		err = json.Unmarshal(body, &resp)
		if err != nil {
			t.Fatalf("expected error to be nil but was: %v", err)
		}

		if resp.Timezone != "UTC" {
			t.Fatalf("expected timezone to be UTC but was: %v", resp.Timezone)
		}
	})

	t.Run("should return time in specific timezone", func(t *testing.T) {
		recorder := httptest.NewRecorder()

		req, err := http.NewRequest(http.MethodGet, CurrentEndpoint, nil)
		if err != nil {
			t.Fatalf("expected error to be nil but was: %v", err)
		}
		req.Header.Set("WS-Location", "America/New_York")

		CurrentHandler(recorder, req)

		if recorder.Code != http.StatusOK {
			t.Fatalf("expected http code to be 200 but was: %v", recorder.Code)
		}

		body, err := ioutil.ReadAll(recorder.Result().Body)
		if err != nil {
			t.Fatalf("expected error to be nil but was: %v", err)
		}

		var resp Time
		err = json.Unmarshal(body, &resp)
		if err != nil {
			t.Fatalf("expected error to be nil but was: %v", err)
		}

		if resp.Timezone != "EDT" {
			t.Fatalf("expected timezone to be EDT but was: %v", resp.Timezone)
		}
	})
}
