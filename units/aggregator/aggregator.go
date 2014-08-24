package aggregator

import (
	"fmt"

	"code.google.com/p/go.net/context"
	"github.com/localhots/yeast/tools"
	"github.com/localhots/yeast/units/power"
)

func Call(ctx context.Context) context.Context {
	results := []string{}

	tools.SyncronizeParallelChain(ctx, func(ctx context.Context) {
		r := ctx.Value("power_result").(power.PowerResult)
		results = append(results, fmt.Sprintf("%d ^ %d = %d", r.Num, r.Power, r.Result))
	})

	ctx = context.WithValue(ctx, "power_results", results)

	return ctx
}
