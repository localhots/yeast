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
	core.InitConfig()
	core.LoadUnits()
	core.ParseChains()

	pretty.Println(core.Conf())
	select {}
}
