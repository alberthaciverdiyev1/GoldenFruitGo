package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	BaseURL string
	Token   string
	HTTP    *http.Client
}

func NewClient(baseURL string) *Client {
	return &Client{
		BaseURL: baseURL,
		HTTP:    &http.Client{Timeout: 10 * time.Second},
	}
}

func (c *Client) doRequest(method, endpoint string, data interface{}) (*http.Response, error) {
	var body *bytes.Buffer
	if data != nil {
		jsonData, err := json.Marshal(data)
		if err != nil {
			return nil, fmt.Errorf("json marshal xətası: %v", err)
		}
		body = bytes.NewBuffer(jsonData)
	} else {
		body = bytes.NewBuffer([]byte{})
	}

	req, err := http.NewRequest(method, c.BaseURL+endpoint, body)
	if err != nil {
		return nil, fmt.Errorf("request yaradıla bilmədi: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	if c.Token != "" {
		req.Header.Set("Authorization", "Bearer "+c.Token)
	}

	return c.HTTP.Do(req)
}

func (c *Client) Get(endpoint string) (*http.Response, error) {
	return c.doRequest("GET", endpoint, nil)
}

func (c *Client) Post(endpoint string, data interface{}) (*http.Response, error) {
	return c.doRequest("POST", endpoint, data)
}

func (c *Client) Put(endpoint string, data interface{}) (*http.Response, error) {
	return c.doRequest("PUT", endpoint, data)
}

func (c *Client) Delete(endpoint string) (*http.Response, error) {
	return c.doRequest("DELETE", endpoint, nil)
}
