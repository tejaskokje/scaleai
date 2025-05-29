package utils

import (
    "bytes"
    "encoding/json"
    "io"
    "net/http"
    "time"
)

// HTTPClient is a wrapper for making HTTP requests
type HTTPClient struct {
    client *http.Client
}

// NewHTTPClient creates a new HTTP client with default settings
func NewHTTPClient() *HTTPClient {
    return &HTTPClient{
        client: &http.Client{
            Timeout: time.Second * 30,
        },
    }
}

// Get performs an HTTP GET request
func (c *HTTPClient) Get(url string) (*http.Response, error) {
    return c.client.Get(url)
}

// Post performs an HTTP POST request with JSON body
func (c *HTTPClient) Post(url string, data interface{}) (*http.Response, error) {
    jsonData, err := json.Marshal(data)
    if err != nil {
        return nil, err
    }

    return c.client.Post(url, "application/json", bytes.NewBuffer(jsonData))
}

// GetJSON performs a GET request and unmarshals the response into the provided interface
func (c *HTTPClient) GetJSON(url string, v interface{}) error {
    resp, err := c.Get(url)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return err
    }

    return json.Unmarshal(body, v)
}