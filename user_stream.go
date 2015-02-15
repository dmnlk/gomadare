package gomadare

import (
	"bufio"
	"encoding/json"

	"log"

	"github.com/k0kubun/pp"
)

// userstream url
const (
	STREAM_URL = "https://userstream.twitter.com/1.1/user.json"
)

// get User Stream and output std.out
func (client *Client) GetUserStream(params map[string]string) {
	//userStreamAPI叩く
	response, err := client.consumer.Get(STREAM_URL, params, client.accessToken)
	if err != nil {
		return
	}
	defer response.Body.Close()
	scanner := bufio.NewScanner(response.Body)

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
		msg := result.(map[string]interface{})
		//pp.Print(msg)
		if _, ok := msg["event"]; ok {
			// unmarshal event
			var e Event
			if err := json.Unmarshal(b, &e); err != nil {
				continue
			}
			pp.Print(e.Event)
			pp.Print(e.TargetObject.Text)
		}
		if _, ok := msg["user"]; ok {
			// unmarshal Status
			var s Status
			if err := json.Unmarshal(b, &s); err != nil {
				continue
			}
			pp.Print(s)
		}
	}
}
