package main

import (
	"fmt"

	"github.com/cli/go-gh"
)

type Response struct {
	Subject struct {
		Title string
		URL   string
	}
	Reason string
}

func main() {
	client, err := gh.RESTClient(nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	response := []Response{}
	err = client.Get("notifications", &response)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, notification := range response {
		if notification.Reason == "mention" {
			fmt.Printf("%s %s\n", notification.Subject.Title, notification.Subject.URL)
		}
	}
}

// For more examples of using go-gh, see:
// https://github.com/cli/go-gh/blob/trunk/example_gh_test.go
