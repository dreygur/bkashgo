package hooks

import (
	"net/http"
	"strconv"
)

func Getheaders(r *http.Request, data []byte, username, password string) {
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Content-Length", strconv.Itoa(len(data)))
	r.Header.Add("username", username)
	r.Header.Add("password", password)
}
