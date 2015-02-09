package main

import (
	"fmt"

	"github.com/localhots/yeast/unit"
)

func main() {
	u := unit.New("uuid")
	msg := []byte("{}")

	fmt.Println("Sending message:", string(msg))
	resp, _ := u.Send(msg)
	fmt.Println("Reply recieved:", string(resp))
}
