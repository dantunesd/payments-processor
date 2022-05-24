package payment

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"go.uber.org/zap"
)

// IHTTPRequester interface for http requests.
type IHTTPRequester interface {
	Post(ctx context.Context, path string, body interface{}) (IResponser, error)
}

type headers map[string]string

// HTTPRequester wrapper for http requests.
type HTTPRequester struct {
	c       *http.Client
	l       *zap.Logger
	baseURL string
	headers headers
}

// NewHTTPRequester HTTPRequester's constructor.
func NewHTTPRequester(c *http.Client, l *zap.Logger, baseURL string, headers headers) *HTTPRequester {
	return &HTTPRequester{
		c:       c,
		l:       l,
		baseURL: baseURL,
		headers: headers,
	}
}

// Post performs a post request.
func (r *HTTPRequester) Post(ctx context.Context, path string, body interface{}) (IResponser, error) {
	return r.do("POST", path, body)
}

func (r *HTTPRequester) do(method, path string, body interface{}) (IResponser, error) {

	bodyEncoded := new(bytes.Buffer)
	if bErr := r.prepareBody(body, bodyEncoded); bErr != nil {
		return nil, NewInternalServerError(bErr.Error())
	}

	req, rErr := http.NewRequest(method, r.baseURL+path, bodyEncoded)
	if rErr != nil {
		return nil, NewInternalServerError(rErr.Error())
	}

	r.prepareHeaders(req)

	r.log("logging http request", r.baseURL, path, method, bodyEncoded.String(), nil)

	res, dErr := r.c.Do(req)
	if dErr != nil {
		return nil, NewInternalServerError(dErr.Error())
	}

	bodyDecoded, iErr := ioutil.ReadAll(res.Body)
	if iErr != nil {
		return nil, NewInternalServerError(iErr.Error())
	}

	bodyCompacted := new(bytes.Buffer)
	json.Compact(bodyCompacted, bodyDecoded)

	r.log("logging http response", r.baseURL, path, method, bodyCompacted.String(), res.StatusCode)

	return NewResponse(res.StatusCode, bodyDecoded), nil
}

func (r *HTTPRequester) prepareBody(body interface{}, bodyEncoded *bytes.Buffer) error {
	if body == nil {
		return nil
	}
	return json.NewEncoder(bodyEncoded).Encode(body)
}

func (r *HTTPRequester) prepareHeaders(req *http.Request) {
	req.Header.Add("content-type", "application/json")
	for k, v := range r.headers {
		req.Header.Add(k, v)
	}
}

func (r *HTTPRequester) log(message, baseURL, path, method, body string, statusCode interface{}) {
	r.l.Info(
		message,
		zap.String("baseUrl", baseURL),
		zap.String("path", path),
		zap.String("method", method),
		zap.String("body", body),
		zap.Any("statusCode", statusCode),
		zap.Any("headers", r.headers),
	)
}
