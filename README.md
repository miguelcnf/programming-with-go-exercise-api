# Programming with Go API exercise

Implement the `/clock/current` http handler as described in the `internal/clock/handler.go` file:

```go
// CurrentHandler returns a 200 response with the current time as a Time JSON representation.
// Callers might request a response in the timezone of their location by setting the header WS-Location to a
// location name corresponding to a file in the IANA Time Zone database, such as "America/New_York".
// If an unknown location is received it serves a 400 response.
func CurrentHandler(w http.ResponseWriter, r *http.Request) {
	// TODO
}}
```

References:

* [net/http](https://golang.org/pkg/net/http/)
* [time](https://golang.org/pkg/time/)
* [encoding/json](https://golang.org/pkg/encoding/json/)