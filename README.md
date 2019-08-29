# Programming with Go API exercise

Implement the `/clock/current` http handler as described in the `internal/clock/handler.go` file:

```go
// CurrentHandler returns the current time as a Time JSON representation.
//
// Callers might request a response in the timezone of their location by by setting the header WS-Local to a
// location name corresponding to a file in the IANA Time Zone database, such as "America/New_York".
func CurrentHandler(w http.ResponseWriter, r *http.Request) {
    // TODO
}
```