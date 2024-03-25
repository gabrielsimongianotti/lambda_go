package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Event struct {
	Resource    string `json:"resource"`
	Path        string `json:"path"`
	HTTPMethod  string `json:"httpMethod"`
	ContentType string `json:"Content-type"`
	Body        string `json:"body"`
}

func main() {

	data := map[string]interface{}{
		"name":  "John",
		"price": 30,
	}

	body, err := json.Marshal(data)

	if err != nil {
		fmt.Printf("could not marshal json: %s\n", err)
		return
	}

	event := Event{
		Resource:    "/",
		Path:        "/lambda",
		HTTPMethod:  "POST",
		ContentType: "application/json",
		Body:        string(body),
	}

	fileName := "events/exemplo.json"

	file, err := os.Create(fileName)

	if err != nil {
		fmt.Printf("could not created file: %s\n", err)
		return
	}

	context, err := json.Marshal(event)

	if err != nil {
		fmt.Printf("could not marshal json: %s\n", err)
		return
	}

	file.Write([]byte(string(context)))
}
