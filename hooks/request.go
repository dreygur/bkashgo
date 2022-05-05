package hooks

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func DoRequest(data []byte, username, password, url string, authorized bool) ([]byte, error) {
	client := &http.Client{}
	r, err := http.NewRequest("POST", url, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Content-Length", strconv.Itoa(len(data)))
	if authorized {
		r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", username))
		r.Header.Add("X-APP-Key", password)
	} else {
		r.Header.Add("username", username)
		r.Header.Add("password", password)
	}

	response, err := client.Do(r)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
