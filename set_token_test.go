package slack_test

import (
	"testing"

	"github.com/gerad/slack"
)

func TestSetToken(t *testing.T) {
	token := "slack-api-token"
	client := &slack.Client{}
	client.SetToken(token)
	if client.Token != token {
		t.Error(token)
	}
}
