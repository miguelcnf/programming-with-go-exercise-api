package clock

import (
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
	// TODO
}
