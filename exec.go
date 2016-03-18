package slack

import (
	"encoding/json"
	"errors"
	"net/url"
)

var (
	// ErrMissingToken is returned when Exec is called without first setting an
	// API token
	ErrMissingToken = errors.New("Slack API token missing")
)

// Argsable allows Exec to take any args that fulfil the interface
// of having an Args() function that returns map[string]string of
// key -> value
type Argsable interface {
	Args() map[string]string
}

// Exec executes a slack rpc command using the default client
func Exec(cmd string, args Argsable, v interface{}) error {
	return DefaultClient.Exec(cmd, args, v)
}

// Exec executes a slack rpc command
func (client *Client) Exec(cmd string, args Argsable, v interface{}) error {
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
func (client *Client) params(args Argsable) url.Values {
	params := url.Values{}
	if args != nil {
		for k, v := range args.Args() {
			params.Add(k, v)
		}
	}
	params.Add("token", client.Token)
	return params
}
