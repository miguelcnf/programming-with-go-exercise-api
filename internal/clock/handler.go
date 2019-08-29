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

// CurrentHandler returns the current time as a Time JSON representation.
//
// Callers might request a response in the timezone of their location by by setting the header WS-Local to a
// location name corresponding to a file in the IANA Time Zone database, such as "America/New_York".
func CurrentHandler(w http.ResponseWriter, r *http.Request) {
	// TODO
}
