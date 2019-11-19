package clock

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

const (
	CurrentEndpoint = "/clock/current"
)

// Time represents a time response with both time and timezone information.
type Time struct {
	Time     time.Time `json:"Time"`
	Timezone string    `json:"Zone"`
}

// CurrentHandler returns a 200 response with the current time in UTC as a Time JSON representation.
// Callers might request a response in the timezone of their location by setting the header WS-Location to a
// location name corresponding to a file in the IANA Time Zone database, such as "America/New_York".
// If an unknown location is received it serves a 400 response.
func CurrentHandler(w http.ResponseWriter, r *http.Request) {
	var (
		now      = time.Now().In(time.UTC)
		timezone = time.UTC.String()
	)

	location := r.Header.Get("WS-Location")
	if location != "" {
		l, err := time.LoadLocation(location)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		now = now.In(l)
		timezone, _ = now.Zone()
	}

	t := Time{
		Time:     now,
		Timezone: timezone,
	}

	resp, err := json.Marshal(t)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(resp)
	if err != nil {
		log.Printf("error writing http response: %v", err)
	}
}
