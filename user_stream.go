package gomadare

import (
	"bufio"
	"encoding/json"

	"log"
)

// userstream url
const (
	STREAM_URL = "https://userstream.twitter.com/1.1/user.json"
)

type Handler func(s Status, e Event) bool

// get User Stream and output std.out
func (client *Client) GetUserStream(params map[string]string, handle Handler) {
	//userStreamAPI叩く
	response, err := client.consumer.Get(STREAM_URL, params, client.accessToken)
	if err != nil {
		return
	}
	defer response.Body.Close()
	scanner := bufio.NewScanner(response.Body)
	// ignore friend list id
	scanner.Scan()
	for {
		if ok := scanner.Scan(); !ok {
			log.Fatal("scan error")
			continue
		}
		var result interface{}
		b := scanner.Bytes()
		//json化
		if err := json.Unmarshal(b, &result); err != nil {
			log.Println(err)
			continue
		}
		var e Event
		var s Status
		msg := result.(map[string]interface{})
		//pp.Print(msg)
		if _, ok := msg["event"]; ok {
			// unmarshal event
			if err := json.Unmarshal(b, &e); err != nil {
				continue
			}
		}
		if _, ok := msg["user"]; ok {
			// unmarshal Status

			if err := json.Unmarshal(b, &s); err != nil {
				continue
			}
		}
		handle(s, e)
	}
}
