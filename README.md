# Macromeasures Go Client

This is a Go library built to support the [Macromeasures](http://macromeasures.com) REST API.

## Initialize Client

```go
import "github.com/verticalmass/go-macromeasures"

client, err := macromeasures.NewClient("your-api-key")

resp, err := client.Twitter.Username("jack")

users, err := resp.Users()

for _, user := range users {
  fmt.Println(user)
}
```

## Twitter API

```go
  client.Twitter.Username(username string) (*macromeasures.UserResponse, error)
  client.Twitter.UserID(userID string) (*macromeasures.UserResponse, error)
```

## Instagram API

```go
  client.Instagram.Username(username string) (*macromeasures.UserResponse, error)
  client.Instagram.UserID(userID string) (*macromeasures.UserResponse, error)
```
