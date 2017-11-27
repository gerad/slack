package slack

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
)

var (
	// ErrMissingToken is returned when Exec is called without first setting an
	// API token
	ErrMissingToken = errors.New("Slack API token missing")
)

// Exec executes a slack rpc command using the default client
func Exec(cmd string, args map[string]interface{}, v interface{}) error {
	return DefaultClient.Exec(cmd, args, v)
}

// Exec executes a slack rpc command
func (client *Client) Exec(cmd string, args map[string]interface{}, v interface{}) error {
	if client.Token == "" {
		return ErrMissingToken
	}

	params := client.params(args)

	res, err := client.HTTPClient.Get("https://slack.com/api/" + cmd + "?" + params.Encode())

	if err != nil {
		return err
	}

	if v == nil {
		return nil
	}

	return json.NewDecoder(res.Body).Decode(v)
}

// params converts args to url.Values for subsequent encoding by Exec
func (client *Client) params(args map[string]interface{}) url.Values {
	params := url.Values{}
	if args != nil {
		for k, v := range args {
			params.Add(k, fmt.Sprintf("%s", v))
		}
	}
	params.Add("token", client.Token)
	return params
}
