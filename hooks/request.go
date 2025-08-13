package hooks

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

type Request struct {
	Debug      bool
	Payload    interface{}
	Username   string
	Password   string
	Url        string
	Authorized bool
}

func DoRequest(request *Request) ([]byte, error) {
	data, err := json.Marshal(request.Payload)
	if err != nil {
		return nil, err
	}

	// Request: Debug Console
	if request.Debug {
		var formatted bytes.Buffer
		json.Indent(&formatted, data, "", "  ")
		fmt.Println(">>> REQUEST JSON:")
		fmt.Println(formatted.String())
	}

	client := &http.Client{}
	r, err := http.NewRequest("POST", request.Url, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(r.Context(), time.Second*30)
	defer cancel()

	r = r.WithContext(ctx)

	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Content-Length", strconv.Itoa(len(data)))
	if request.Authorized {
		r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", request.Username))
		r.Header.Add("X-APP-Key", request.Password)
	} else {
		r.Header.Add("username", request.Username)
		r.Header.Add("password", request.Password)
	}

	response, err := client.Do(r)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// Response: Debug Console
	if request.Debug {
		var formattedResp bytes.Buffer
		json.Indent(&formattedResp, body, "", "  ")
		fmt.Println("<<< RESPONSE JSON:")
		fmt.Println(formattedResp.String())
	}

	return body, nil
}
