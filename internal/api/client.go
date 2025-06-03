package api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

// Client communicates with an AI chat API.
//
// It sends chat messages and returns the AI's reply.
// Example usage:
//
//	cli := api.NewClient()
//	reply, err := cli.SendMessage(ctx, "hello")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(reply)
//
// Output:
// ai response
//
// The example demonstrates typical usage of the Client type.
// To actually run it, replace the URL and provide an API key.
type Client struct {
	httpClient *http.Client
	baseURL    string
	apiKey     string
}

// Option configures a Client.
type Option func(*Client)

// WithHTTPClient sets a custom HTTP client.
func WithHTTPClient(hc *http.Client) Option {
	return func(c *Client) { c.httpClient = hc }
}

// WithBaseURL sets the API base URL.
func WithBaseURL(url string) Option {
	return func(c *Client) { c.baseURL = url }
}

// WithAPIKey sets the API key used for authentication.
func WithAPIKey(key string) Option {
	return func(c *Client) { c.apiKey = key }
}

// NewClient returns a configured Client instance.
func NewClient(opts ...Option) *Client {
	c := &Client{httpClient: http.DefaultClient, baseURL: "https://api.example.com"}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

// SendMessage sends a chat message and returns the reply.
func (c *Client) SendMessage(ctx context.Context, msg string) (reply string, err error) {
	payload := map[string]string{"message": msg}
	buf, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("marshal payload: %w", err)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL+"/chat", bytes.NewReader(buf))
	if err != nil {
		return "", fmt.Errorf("new request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("do request: %w", err)
	}
	defer func() {
		if cerr := resp.Body.Close(); cerr != nil && err == nil {
			err = fmt.Errorf("close body: %w", cerr)
		}
	}()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status: %s", resp.Status)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("read body: %w", err)
	}
	var out struct{ Reply string }
	if err := json.Unmarshal(body, &out); err != nil {
		return "", fmt.Errorf("decode body: %w", err)
	}
	if out.Reply == "" {
		return "", errors.New("empty reply")
	}
	return out.Reply, nil
}
