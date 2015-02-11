package main

import (
	"flag"

	"github.com/kr/pretty"
	"github.com/localhots/confection"
	"github.com/localhots/yeast/core"
)

func init() {
	confection.SetupFlags()
	flag.Parse()
}

func main() {
	app := core.NewApp()
	pretty.Println(app.Conf())

	println("Waiting")
	select {}
}
