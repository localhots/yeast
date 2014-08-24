package input

import (
	"flag"

	"code.google.com/p/go.net/context"
)

func FromFlag(ctx context.Context) context.Context {
	var num int
	flag.IntVar(&num, "num", 0, "Pass this number")
	flag.Parse()

	ctx = context.WithValue(ctx, "num", num)

	return ctx
}
