package macromeasures

import (
	"testing"
)

var client *Client
var err error

var apiKey = "your-api-key"

func init() {
	client, err = NewClient(apiKey, DefaultRateLimit)
	if err != nil {
		panic(err.Error())
	}
}

func TestClient(t *testing.T) {
	client, err := NewClient(apiKey, DefaultRateLimit)
	if err != nil {
		t.Errorf("Macromeasures: Client: Error - %s", "an error should not have occured")
		t.Fail()
	}
	if client.Twitter == nil {
		t.Errorf("Macromeasures: Client: Error - %s", "twitter client should not be nil")
		t.Fail()
	}
	if client.Instagram == nil {
		t.Errorf("Macromeasures: Client: Error - %s", "instagram client should not be nil")
		t.Fail()
	}
	if client.apikey != apiKey {
		t.Errorf("Macromeasures: Client: Error - %s", "instagram client should not be nil")
		t.Fail()
	}
}

func TestTwitterUsername(t *testing.T) {
	resp, err := client.Twitter.Username("jack")
	if err != nil {
		t.Errorf("Macromeasures: Twitter: Username: Error - %s", "an error should not have occured")
		t.Fail()
	}
	if !resp.Complete {
		t.Errorf("Macromeasures: Twitter: Username: Error - %s", "response is not complete/finished")
		t.Fail()
	}
	if resp.Labels == nil || len(resp.Labels) == 0 {
		t.Errorf("Macromeasures: Twitter: Username: Error - %s", "user response is missing labels")
		t.Fail()
	}
}

func TestTwitterUserIDSuccess(t *testing.T) {
	resp, err := client.Twitter.UserID("12")
	if err != nil {
		t.Errorf("Macromeasures: Twitter: UserID: Error - %s", "an error should not have occured")
		t.Fail()
	}
	if !resp.Complete {
		t.Errorf("Macromeasures: Twitter: UserID: Error - %s", "response is not complete/finished")
		t.Fail()
	}
	if resp.Labels == nil || len(resp.Labels) == 0 {
		t.Errorf("Macromeasures: Twitter: UserID: Error - %s", "user response is missing labels")
		t.Fail()
	}
}

func TestTwitterUserIDFail(t *testing.T) {
	resp, err := client.Twitter.UserID("jack")
	if err != nil {
		t.Errorf("Macromeasures: Twitter: UserID: Error - %s", "an error should not have occured")
		t.Fail()
	}
	if resp.Labels["jack"].Valid {
		t.Errorf("Macromeasures: Twitter: UserID: Error - %s", "should be an invalid value (username instead of social id)")
		t.Fail()
	}
}
