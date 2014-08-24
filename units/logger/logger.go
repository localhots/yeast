package logger

import (
	"fmt"

	"code.google.com/p/go.net/context"
)

func Call(ctx context.Context) context.Context {
	results := ctx.Value("power_results").([]string)
	id := ctx.Value("id").(string)

	fmt.Println("Calculation result", id)
	fmt.Println("Power results are:")
	for _, r := range results {
		fmt.Println(r)
	}

	return ctx
}
