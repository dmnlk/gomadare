package main

import (
	"os"

	"github.com/dmnlk/gomadare"
)

func main() {
	ck := os.Getenv("ck")
	cs := os.Getenv("cs")
	at := os.Getenv("at")
	as := os.Getenv("as")
	client := gomadare.NewClient(ck, cs, at, as)
	client.GetUserStream(nil)
}
