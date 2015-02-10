package main

import (
	"flag"
	"fmt"

	"github.com/localhots/yeast/unit"
)

func main() {
	var (
		name string
		msg  string
	)

	flag.StringVar(&name, "unit", "", "Unit name")
	flag.StringVar(&msg, "msg", "", "Message")
	flag.Parse()

	u := unit.New(name)

	fmt.Println("Sending message:", msg)
	resp, _ := u.Call([]byte(msg))
	fmt.Println("Reply recieved:", string(resp))
}
