package payment

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"go.uber.org/zap"
)

// IHTTPRequester interface for http requests.
type IHTTPRequester interface {
	Post(ctx context.Context, path string, body interface{}) (IResponser, error)
}

type headers map[string]string

// HTTPRequester wrapper for http requests.
type HTTPRequester struct {
	l       *zap.Logger
	baseURL string
	headers headers
	timeout time.Duration
}

// NewHTTPRequester HTTPRequester's constructor.
func NewHTTPRequester(l *zap.Logger, baseURL string, headers headers, timeout time.Duration) *HTTPRequester {
	return &HTTPRequester{
		l:       l,
		baseURL: baseURL,
		headers: headers,
		timeout: timeout,
	}
}

// Post performs a post request.
func (r *HTTPRequester) Post(ctx context.Context, path string, body interface{}) (IResponser, error) {
	return r.do("POST", path, body)
}

func (r *HTTPRequester) do(method, path string, body interface{}) (IResponser, error) {
	c := http.Client{
		Timeout: r.timeout,
	}

	bodyEncoded := new(bytes.Buffer)
	if body != nil {
		if eErr := json.NewEncoder(bodyEncoded).Encode(body); eErr != nil {
			return nil, NewInternalServerError(eErr.Error())
		}
	}

	req, rErr := http.NewRequest(method, r.baseURL+path, bodyEncoded)
	if rErr != nil {
		return nil, NewInternalServerError(rErr.Error())
	}

	req.Header.Add("content-type", "application/json")
	for k, v := range r.headers {
		req.Header.Add(k, v)
	}

	r.l.Info(
		"logging http request",
		zap.String("baseUrl", r.baseURL),
		zap.String("path", path),
		zap.String("method", method),
		zap.String("body", bodyEncoded.String()),
	)

	res, dErr := c.Do(req)
	if dErr != nil {
		return nil, NewInternalServerError(dErr.Error())
	}

	bodyDecoded, iErr := ioutil.ReadAll(res.Body)
	if iErr != nil {
		return nil, NewInternalServerError(iErr.Error())
	}

	bodyCompacted := new(bytes.Buffer)
	json.Compact(bodyCompacted, bodyDecoded)

	r.l.Info(
		"logging http response",
		zap.String("baseUrl", r.baseURL),
		zap.String("path", path),
		zap.String("method", method),
		zap.String("body", bodyCompacted.String()),
		zap.Int("statusCode", res.StatusCode),
	)

	return NewResponse(res.StatusCode, bodyDecoded), nil
}
