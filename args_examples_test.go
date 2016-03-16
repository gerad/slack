package slack_test

import "github.com/gerad/slack"

func ExampleArgs() {
	slack.Exec("api.test", slack.Args{
		"error": "my_error",
		"foo":   "bar",
	}, nil)
}
