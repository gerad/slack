package slack_test

import (
	"testing"

	"github.com/gerad/slack"
)

func TestNew(t *testing.T) {
	token := "slack-api-token"
	client := slack.New(token)
	if client.Token != token {
		t.Error(client.Token)
	}
}
