package slack

// SetToken sets the slack api token of the default client
// Tokens can be generated at: https://api.slack.com/docs/oauth-test-tokens
func SetToken(token string) {
	DefaultClient.SetToken(token)
}

// SetToken sets the slack api token
// Tokens can be generated at: https://api.slack.com/docs/oauth-test-tokens
func (client *Client) SetToken(token string) {
	client.Token = token
}
