package core

import (
	"encoding/json"
	"io/ioutil"
	"reflect"
	"sync"

	"code.google.com/p/go.net/context"
)

func NewChain(name string) (chain interface{}, ok bool) {
	chains := readChains()
	chain, ok = chains[name]

	return
}

func ProcessChain(ctx context.Context, chain interface{}) context.Context {
	ctx = processSequentialChain(ctx, []interface{}{chain})

	return ctx
}

func processSequentialChain(ctx context.Context, chain []interface{}) context.Context {
	for _, link := range chain {
		val := reflect.ValueOf(link)

		switch val.Kind() {
		case reflect.Map:
			lmap := link.(map[string]interface{})
			if subchain, ok := lmap["s"]; ok {
				ctx = processSequentialChain(ctx, subchain.([]interface{}))
			} else if subchain, ok := lmap["p"]; ok {
				ctx = processParallelChain(ctx, subchain.([]interface{}))
			} else if subchain, ok := lmap["d"]; ok {
				go processSequentialChain(ctx, subchain.([]interface{}))
			}
		case reflect.String:
			unitName := val.String()
			unit, ok := Units[unitName]
			if !ok {
				panic("Unknown unit: " + unitName)
			}

			// Execution of a unit
			ctx = unit(ctx)
		default:
			panic("Unexpected chain element: " + val.Kind().String())
		}
	}

	return ctx
}

func processParallelChain(ctx context.Context, chain []interface{}) context.Context {
	var (
		wg       sync.WaitGroup
		contexts = make(chan context.Context, len(chain))
		finished = make(chan struct{})
	)

	for _, link := range chain {
		val := reflect.ValueOf(link)

		switch val.Kind() {
		case reflect.Map:
			lmap := link.(map[string]interface{})
			if subchain, ok := lmap["s"]; ok {
				wg.Add(1)
				go func(ctx context.Context) {
					contexts <- processSequentialChain(ctx, subchain.([]interface{}))
					wg.Done()
				}(ctx)
			} else if subchain, ok := lmap["p"]; ok {
				wg.Add(1)
				go func(ctx context.Context) {
					contexts <- processParallelChain(ctx, subchain.([]interface{}))
					wg.Done()
				}(ctx)
			} else if subchain, ok := lmap["d"]; ok {
				go processSequentialChain(ctx, subchain.([]interface{}))
			}
		case reflect.String:
			unitName := val.String()
			unit, ok := Units[unitName]
			if !ok {
				panic("Unknown unit: " + unitName)
			}

			// Execution of a unit
			wg.Add(1)
			go func(ctx context.Context) {
				contexts <- unit(ctx)
				wg.Done()
			}(ctx)
		default:
			panic("Unexpected chain element: " + val.Kind().String())
		}
	}

	go func() {
		wg.Wait()
		finished <- struct{}{}
	}()

	ctx = context.WithValue(ctx, "parallel_contexts", contexts)
	ctx = context.WithValue(ctx, "parallel_finished", finished)

	return ctx
}

func readChains() map[string]interface{} {
	b, err := ioutil.ReadFile("config/chains.json")
	if err != nil {
		panic(err)
	}

	var chains map[string]interface{}
	if err := json.Unmarshal(b, &chains); err != nil {
		panic(err)
	}

	return chains
}
