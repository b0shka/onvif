package networking

import (
	"bytes"
	"context"
	"net/http"
)

// SendSoap send soap message
func SendSoap(httpClient *http.Client, endpoint, message string, headers map[string]string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewBufferString(message))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/soap+xml; charset=utf-8")
	for name, value := range headers {
		req.Header.Set(name, value)
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// SendSoapWithCtx send soap message with context
func SendSoapWithCtx(ctx context.Context, httpClient *http.Client, endpoint, message string, headers map[string]string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, bytes.NewBufferString(message))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/soap+xml; charset=utf-8")
	for name, value := range headers {
		req.Header.Set(name, value)
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
