package main

import (
	"os"

	"github.com/dmnlk/gomadare"
	"github.com/k0kubun/pp"
)

func main() {
	ck := os.Getenv("ck")
	cs := os.Getenv("cs")
	at := os.Getenv("at")
	as := os.Getenv("as")
	client := gomadare.NewClient(ck, cs, at, as)
	client.GetUserStream(nil, func (s Status, e Event) {
		if s != nil {
			pp.Println(s)
		}
		if e != nil {
			pp.Print(e)
		}
	})
}

