package gomadare

import (
	"bufio"
	"encoding/json"
	"fmt"

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
	fmt.Println("get")
	//defer response.Body.Close()
	//レスポンスのbodyをscannerで読み込む
	fmt.Println(response.StatusCode)
	scanner := bufio.NewScanner(response.Body)
	for {
		fmt.Println("into forloop")
		//都度scanして新しく受信してたら
		if ok := scanner.Scan(); !ok {
			fmt.Println("scan error")
			continue
		}
		var result interface{}
		//json化
		if err := json.Unmarshal(scanner.Bytes(), &result); err != nil {
			fmt.Println(err)
			fmt.Println(scanner.Bytes())
			continue
		}
		pp.Print(result)
		msg := result.(map[string]interface{})
		if val, ok := msg["event"]; ok {
			pp.Print(val)
		}
	}
}
