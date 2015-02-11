package main

import (
	"flag"

	"github.com/kr/pretty"
	"github.com/localhots/confection"
	"github.com/localhots/yeast/chain"
	"github.com/localhots/yeast/core"
	"github.com/localhots/yeast/unit"
)

func init() {
	confection.SetupFlags()
	flag.Parse()
}

func main() {
	core.InitConfig()

	ub := unit.NewBank(core.Conf().UnitsConfig)
	ub.Reload()

	cb := chain.NewBank(core.Conf().ChainsConfig, ub)
	cb.Reload()

	pretty.Println(core.Conf())

	println("Waiting")
	select {}
}
