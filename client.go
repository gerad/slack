package slack

import "net/http"

// Client is a slack api client
type Client struct {
	// Token is the slack api token
	// Tokens can be generated at: https://api.slack.com/docs/oauth-test-tokens
	Token string

	// HTTPClient is the http.Client used to connect to slack.
	// If none is provided, http.DefaultClient will be used.
	HTTPClient *http.Client
}

// New creates a new slack api client with the provided slack api token
func New(token string) *Client {
	return &Client{Token: token}
}

// DefaultClient is the default Client that is used by Exec
var DefaultClient = &Client{}
