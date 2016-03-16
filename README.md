# Slack

A minimal go Slack API client.

## About

This is an extraction from a side-project to learn and play with go.

Most people will be better served by <https://github.com/nlopes/slack>.

## Usage

The core api is extremely minimal, just one method `Exec` exposes all of the
[Slack Web API](https://api.slack.com/methods):

```go
import (
  "fmt"
  "github.com/gerad/slack"
)

var response interface{}
var err = slack.Exec("api.test", nil, &response)
fmt.Println(response)
```

Optionally, for convenience, types are defined for API arguments and responses.

```go
import (
  "fmt"
  "github.com/gerad/slack"
)

var resp = slack.APITestResp{}
var args = slack.APITestArgs{}
var err = slack.Exec("api.test", &args, &resp)
fmt.Println(resp)
```

## License

ISC
