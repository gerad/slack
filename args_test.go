package slack_test

import (
	"reflect"
	"testing"

	"github.com/gerad/slack"
)

func TestArgs_Args(t *testing.T) {
	in := map[string]string{"foo": "bar"}
	out := slack.Args(in).Args()
	if !reflect.DeepEqual(in, out) {
		t.Error(out)
	}
}
