package apize

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type rawResponse[T any] struct {
	Error *Error `json:"error,omitempty"`
	Data  *T     `json:"data,omitempty"`
}

func httpGet[T any](ctx context.Context, c *client, url string) (*T, error) {
	return httpRequest[T](ctx, c, http.MethodGet, url, nil)
}

func httpPost[T any](ctx context.Context, c *client, url string, req any) (*T, error) {
	return httpRequest[T](ctx, c, http.MethodPost, url, req)
}

func httpRequest[T any](ctx context.Context, c *client, method string, url string, req any) (*T, error) {
	// Format the full URL
	fullUrl := fmt.Sprintf("%s%s", c.baseurl, url)

	// Encode the body to JSON bytes
	var body io.Reader
	if method == http.MethodPost {
		if req == nil {
			body = strings.NewReader("{}")
		} else {
			// Encode the request to JSON bytes
			jsonBytes, err := json.Marshal(req)
			if err != nil {
				return nil, fmt.Errorf("encoding request to json: %s", err)
			}
			body = bytes.NewReader(jsonBytes)
		}
	}

	// Create the HTTP request
	httpReq, err := http.NewRequestWithContext(ctx, method, fullUrl, body)
	if err != nil {
		return nil, fmt.Errorf("creating http request: %s", err)
	}

	// Apply the token and content type headers
	if c.token != "" {
		httpReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))
	}
	if body != nil {
		httpReq.Header.Set("Content-Type", "application/json")
	}

	// Send the request using the configured client
	httpRes, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("sending http request: %s", err)
	}
	defer httpRes.Body.Close()

	// Read the response body into memory
	jsonBytes, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("reading http response body: %s", err)
	}

	// Try to decode the response into the appropriate type
	var rawResponse rawResponse[T]
	if err := json.Unmarshal(jsonBytes, &rawResponse); err != nil {
		return nil, fmt.Errorf("decoding response: %s", err)
	}

	// Return either the error or the response data
	if rawResponse.Error != nil {
		return nil, rawResponse.Error
	}
	return rawResponse.Data, nil
}
