package network

import (
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	IsStatusSuccess = map[int]bool{
		http.StatusOK:      true,
		http.StatusCreated: true,
	}
)

// CallHTTPRequest ...
func CallHTTPRequest(ctx context.Context, method, url string, body io.Reader, headers ...map[string]string) (res []byte, code int, err error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	// iterate optional data of headers
	for _, header := range headers {
		for key, value := range header {
			req.Header.Set(key, value)
		}
	}

	client := &http.Client{Timeout: time.Minute}
	r, err := client.Do(req)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	defer r.Body.Close()

	res, err = ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	return res, r.StatusCode, nil
}
