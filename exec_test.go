package slack_test

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/gerad/slack"
)

func TestClient_Exec(t *testing.T) {
	double := &RoundTripperDouble{}
	double.ParsedResponse = map[string]interface{}{
		"ok": true,
		"args": map[string]interface{}{
			"foo": "bar",
		},
	}

	var res interface{}
	client := slack.Client{"token", &http.Client{Transport: double}}
	args := slack.Args{"foo": "bar"}
	err := client.Exec("api.test", args, &res)

	if err != nil {
		t.Fail()
	}

	if !strings.HasPrefix(double.CapturedURL, "https://slack.com/api/api.test?") {
		t.Errorf("CapturedURL: %#v", double.CapturedURL)
	}

	if double.CapturedToken != "token" {
		t.Errorf("CapturedToken: %#v", double.CapturedToken)
	}

	if !reflect.DeepEqual(double.CapturedArgs, args.Args()) {
		t.Errorf("CapturedArgs: %#v", double.CapturedArgs)
	}

	if !reflect.DeepEqual(res, double.ParsedResponse) {
		t.Errorf("Parsed Response: %#v", res)
	}
}

func TestClient_Exec_nilArgs(t *testing.T) {
	double := &RoundTripperDouble{}

	var res interface{}
	client := slack.Client{"token", &http.Client{Transport: double}}
	err := client.Exec("api.test", nil, &res)

	if err != nil {
		t.Fail()
	}

	if double.CapturedToken != "token" {
		t.Errorf("CapturedToken: %#v", double.CapturedToken)
	}

	if !reflect.DeepEqual(double.CapturedArgs, map[string]string{}) {
		t.Errorf("CapturedArgs: %#v", double.CapturedArgs)
	}
}

func TestClient_Exec_nilValue(t *testing.T) {
	double := &RoundTripperDouble{}

	client := slack.Client{"token", &http.Client{Transport: double}}
	err := client.Exec("api.test", nil, nil)

	if err != nil {
		t.Fail()
	}
}

func TestClient_Exec_nilToken(t *testing.T) {
	double := &RoundTripperDouble{}

	client := slack.Client{HTTPClient: &http.Client{Transport: double}}
	err := client.Exec("api.test", nil, nil)

	if err != slack.ErrMissingToken {
		t.Error(err)
	}
}

func TestClient_Exec_httpError(t *testing.T) {
	double := &RoundTripperDouble{}
	double.ResponseError = errors.New("some http error")

	client := slack.Client{"token", &http.Client{Transport: double}}
	err := client.Exec("api.test", nil, nil)

	if !strings.Contains(err.Error(), double.ResponseError.Error()) {
		t.Errorf("%#v", err)
	}
}

type RoundTripperDouble struct {
	ParsedResponse interface{}
	ResponseError  error
	CapturedURL    string
	CapturedToken  string
	CapturedArgs   map[string]string
}

func (double *RoundTripperDouble) RoundTrip(req *http.Request) (*http.Response, error) {
	double.CapturedURL = req.URL.String()

	double.CapturedArgs = make(map[string]string)
	for k, vals := range req.URL.Query() {
		if k == "token" {
			double.CapturedToken = vals[0]
		} else {
			double.CapturedArgs[k] = vals[0]
		}
	}

	if double.ResponseError != nil {
		return nil, double.ResponseError
	}

	body, _ := json.Marshal(double.ParsedResponse)
	res := &http.Response{
		Body: ioutil.NopCloser(strings.NewReader(string(body))),
	}

	return res, nil
}
