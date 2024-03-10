package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func publishMessage(channel string, content string) {
	url := "https://rest-hz.goeasy.io/v2/pubsub/publish"
	data := map[string]string{
		"appkey":  "BC-1e157115606e4516bb6db14d9ad90639",
		"channel": channel,
		"content": content,
	}
	jsonValue, _ := json.Marshal(data)
	response, err := http.Post(url, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s", err)
	} else {
		fmt.Println(response.Status)
	}
}

func main() {
	publishMessage("test_channel", "Hello, GoEasy!")
}
