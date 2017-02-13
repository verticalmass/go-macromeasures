package macromeasures

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/beefsack/go-rate"
)

const (
	// MaxIdleConnections for shared http client
	MaxIdleConnections = 10
	// RequestTimeout for shared http client
	RequestTimeout = 60
	// DefaultRate is the default requests per minute for macromeasures
	DefaultRateLimit = 60
)

var (
	apiDomain   = "api.macromeasures.com"
	libraryName = "Macromeasures Client: Go"
	errUnknown  = fmt.Errorf("%s: Error unknown", libraryName)
	errLibrary  = fmt.Errorf("%s: Library-specific error", libraryName)
)

// Client is the main Macromeasures client
type Client struct {
	apikey     string
	httpClient *http.Client
	Twitter    *api
	Instagram  *api
}

// NewClient returns a client for Macromeasures
func NewClient(apikey string, reqMin int) (client *Client, err error) {
	client = &Client{apikey: apikey}
	err = client.configure(reqMin)
	if err != nil {
		return nil, err
	}
	return client, nil
}

// configure sets up the shared http client and api endpoints
func (c *Client) configure(reqMin int) (err error) {
	c.httpClient = &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: MaxIdleConnections,
		},
		Timeout: time.Duration(RequestTimeout) * time.Second,
	}
	c.Twitter, err = newAPI(c, "twitter", 10)
	if err != nil {
		return err
	}
	c.Instagram, err = newAPI(c, "instagram", 20)
	if err != nil {
		return err
	}
	return nil
}

// api handles requests to macromeasures https://api.macromeasures.com/:provider/users
type api struct {
	apikey     string
	httpClient *http.Client
	provider   string
	rl         *rate.RateLimiter
}

// newAPI returns back a macromeasures api endpoint (twitter vs instagram)
func newAPI(c *Client, provider string, reqSec int) (*api, error) {
	return &api{apikey: c.apikey, httpClient: c.httpClient, provider: provider, rl: rate.New(100, time.Duration(reqSec)*time.Second)}, nil
}

// Username returns back an individual user based off of their social username
func (a *api) Username(id string) (resp *UserResponse, err error) {
	u := a.buildURL("usernames", []string{id})
	return a.get(u)
}

// User returns back an individual user based off of their social id
func (a *api) UserID(id string) (resp *UserResponse, err error) {
	u := a.buildURL("ids", []string{id})
	return a.get(u)
}

// buildURL prepares the URL based off of the search lookup (usernames vs ids)
func (a *api) buildURL(key string, values []string) *url.URL {
	u := &url.URL{
		Scheme: "http",
		Host:   apiDomain,
		Path:   fmt.Sprintf("%s/%s.json", a.provider, "users"),
	}
	v := url.Values{}
	v.Set("key", a.apikey)
	v.Set(key, strings.Join(values, ","))
	u.RawQuery = v.Encode()
	return u
}

// get returns back a UserResponse when the response is complete
func (a *api) get(u *url.URL) (resp *UserResponse, err error) {
Loop:
	for {
		resp, err = a.user(u)
		if err != nil {
			return nil, err
		}
		if resp.Complete {
			break Loop
		}
		time.Sleep(100 * time.Millisecond)
	}
	return resp, nil
}

// user returns back a user response from a GET request
func (a *api) user(u *url.URL) (user *UserResponse, err error) {
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	resp, err := a.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	user = &UserResponse{}
	if err := json.NewDecoder(resp.Body).Decode(user); err != nil {
		return nil, err
	}
	return user, nil
}
