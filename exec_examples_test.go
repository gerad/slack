package slack_test

import (
	"fmt"

	"github.com/gerad/slack"
)

func ExampleExec() {
	var res interface{}
	args := slack.Args{"foo": "bar"}

	slack.Exec("api.test", args, &res)

	fmt.Println(res)
}
