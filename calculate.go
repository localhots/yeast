package main

import (
	"flag"

	"code.google.com/p/go.net/context"
	"github.com/localhots/yeast/core"
)

func main() {
	var num int
	flag.IntVar(&num, "num", 0, "Pass this number")
	flag.Parse()

	chain, ok := core.NewChain("calculate")
	if !ok {
		println("Bad chain: calculate")
		return
	}

	ctx := context.WithValue(context.Background(), "num", num)
	core.ProcessChain(ctx, chain)
}
