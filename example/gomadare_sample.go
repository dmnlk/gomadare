package main

import (
	"fmt"
	"github.com/dmnlk/gomadare"
)

func main() {
	client := gomadare.NewClient("Do84oBb6LKEms34UpTEhHxcdQ",
		"r1WE3xKbDkbai1Bc6FPwX9353K4f4zO9dcs9eqHUMZuJMifs1P",
		"85270957-ghFGzlk8tuxY4boACsAA8xdJlv3EpEFwGdNkTvkHn",
		"u5QnST5CxBP53dy3vUDYHjRwrvBi5BYbP834QwiktIFWd")
	client.GetUserStream(nil)
	fmt.Println("aa")
}
