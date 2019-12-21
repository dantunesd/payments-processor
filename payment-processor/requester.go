package payment

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

// IHTTPRequester interface for http requests
type IHTTPRequester interface {
	Post(path string, body, output interface{}) error
}

type headers map[string]string

// HTTPRequester wrapper for client http
type HTTPRequester struct {
	BaseURL string
	Headers headers
	Timeout time.Duration
}

// NewHTTPRequester constructor.
func NewHTTPRequester(baseURL string, headers headers, timeout time.Duration) *HTTPRequester {
	return &HTTPRequester{
		BaseURL: baseURL,
		Headers: headers,
		Timeout: timeout,
	}
}

// Post perform a post request
func (r *HTTPRequester) Post(path string, body, output interface{}) error {
	return r.do("POST", path, body, output)
}

func (r *HTTPRequester) do(method, path string, body, output interface{}) error {
	c := http.Client{
		Timeout: r.Timeout,
	}

	bodyEncoded := new(bytes.Buffer)
	if body != nil {
		if eErr := json.NewEncoder(bodyEncoded).Encode(body); eErr != nil {
			return eErr
		}
	}

	req, rErr := http.NewRequest(method, r.BaseURL+path, bodyEncoded)
	if rErr != nil {
		return rErr
	}

	for key, val := range r.Headers {
		req.Header.Add(key, val)
	}
	req.Header.Add("content-type", "application/json")

	res, dErr := c.Do(req)
	if dErr != nil {
		return dErr
	}

	resBody, rErr := ioutil.ReadAll(res.Body)
	if rErr != nil {
		return rErr
	}

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusMultipleChoices {
		return errors.New(string(resBody))
	}

	if uErr := json.Unmarshal(resBody, output); uErr != nil {
		return uErr
	}

	return nil
}
