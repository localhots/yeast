package power

import (
	"code.google.com/p/go.net/context"
)

type (
	PowerResult struct {
		Num    int
		Power  int
		Result int
	}
)

func Power2(ctx context.Context) context.Context {
	v := ctx.Value("num").(int)
	res := PowerResult{Num: v, Power: 2, Result: v * v}
	ctx = context.WithValue(ctx, "power_result", res)

	return ctx
}

func Power3(ctx context.Context) context.Context {
	v := ctx.Value("num").(int)
	res := PowerResult{Num: v, Power: 3, Result: v * v * v}
	ctx = context.WithValue(ctx, "power_result", res)

	return ctx
}

func Power5(ctx context.Context) context.Context {
	v := ctx.Value("num").(int)
	res := PowerResult{Num: v, Power: 5, Result: v * v * v * v * v}
	ctx = context.WithValue(ctx, "power_result", res)

	return ctx
}
