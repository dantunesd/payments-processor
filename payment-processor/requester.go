package payment

import (
	"bytes"
	"encoding/json"
	"errors"
	"go.uber.org/zap"
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
	Logger  *zap.Logger
	BaseURL string
	Headers headers
	Timeout time.Duration
}

// NewHTTPRequester constructor.
func NewHTTPRequester(l *zap.Logger, baseURL string, headers headers, timeout time.Duration) IHTTPRequester {
	return &HTTPRequester{
		Logger:  l,
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

	r.Logger.Info(
		"logging http request",
		zap.String("path", path),
		zap.String("method", method),
		zap.String("body", bodyEncoded.String()),
	)

	res, dErr := c.Do(req)
	if dErr != nil {
		return dErr
	}

	resByte, rErr := ioutil.ReadAll(res.Body)
	if rErr != nil {
		return rErr
	}

	resBody := new(bytes.Buffer)
	json.Compact(resBody, resByte)

	r.Logger.Info(
		"logging http response",
		zap.String("path", path),
		zap.String("method", method),
		zap.Int("statusCode", res.StatusCode),
		zap.String("response", resBody.String()),
	)

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusMultipleChoices {
		return errors.New(resBody.String())
	}

	if uErr := json.Unmarshal(resBody.Bytes(), output); uErr != nil {
		return uErr
	}

	return nil
}
