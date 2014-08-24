package tools

import (
	"code.google.com/p/go.net/context"
)

func SyncronizeParallelChain(ctx context.Context, fn func(context.Context)) {
	var (
		pcontexts = ctx.Value("parallel_contexts").(chan context.Context)
		pfinished = ctx.Value("parallel_finished").(chan struct{})
		working   = true
	)

	for len(pcontexts) > 0 || working {
		select {
		case pctx := <-pcontexts:
			fn(pctx)
		case <-pfinished:
			working = false
		}
	}
}
