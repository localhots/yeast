package main

import (
	"code.google.com/p/go.net/context"
	"github.com/localhots/yeast/core"
)

func main() {
	chain, ok := core.NewChain("calculate")
	if !ok {
		println("Bad chain: calculate")
		return
	}

	ctx := context.Background()
	core.ProcessChain(ctx, chain)
}
