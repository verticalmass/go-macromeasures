package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/verticalmass/go-macromeasures"
)

func fatalOnErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	client, err := macromeasures.NewClient("your-api-key", 60)
	fatalOnErr(err)
	resp, err := client.Twitter.Username("jack")
	fatalOnErr(err)
	users, err := resp.Users()
	fatalOnErr(err)
	body, err := json.Marshal(users)
	fatalOnErr(err)
	fmt.Println(string(body))
}
