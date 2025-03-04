package middleware

import (
	"bytes"
	"io"
	"math"
	"net/http"
	"time"
)

const RetryCount = 3

type Retry struct {
	transport http.RoundTripper
}

func NewRetry(nextTransport http.RoundTripper) *Retry {
	return &Retry{transport: nextTransport}
}

func (r *Retry) RoundTrip(req *http.Request) (*http.Response, error) {
	// Clone the request body
	var bodyBytes []byte
	if req.Body != nil {
		bodyBytes, _ = io.ReadAll(req.Body)
		req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	}
	// Send the request
	resp, err := r.transport.RoundTrip(req)
	// Retry logic
	retries := 0
	for shouldRetry(err, resp) && retries < RetryCount {
		// Wait for the specified backoff period
		time.Sleep(backoff(retries))
		// We're going to retry, consume any response to reuse the connection.
		drainBody(resp)
		// Clone the request body again
		if req.Body != nil {
			req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		}
		// Retry the request
		resp, err = r.transport.RoundTrip(req)
		retries++
	}
	// Return the response
	return resp, err
}

func shouldRetry(err error, resp *http.Response) bool {
	if err != nil {
		return true
	}

	if resp.StatusCode == http.StatusTooManyRequests ||
		resp.StatusCode == 420 {
		return true
	}
	return false
}

func backoff(retries int) time.Duration {
	return time.Duration(math.Pow(2, float64(retries))) * time.Second
}

func drainBody(resp *http.Response) {
	if resp.Body != nil {
		io.Copy(io.Discard, resp.Body)
		resp.Body.Close()
	}
}
