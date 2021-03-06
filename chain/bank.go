package chain

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"

	"github.com/localhots/yeast/unit"
)

type (
	Bank struct {
		config string
		chains map[string]*Chain
		units  *unit.Bank
	}
)

func NewBank(chainsConfig, unitsConfig string) *Bank {
	return &Bank{
		config: chainsConfig,
		chains: map[string]*Chain{},
		units:  unit.NewBank(unitsConfig),
	}
}

func (b *Bank) Chain(name string) (c *Chain, ok bool) {
	c, ok = b.chains[name]
	return
}

func (b *Bank) Reload() {
	b.units.Reload()

	f, err := os.Open(b.config)
	if err != nil {
		panic("Failed to open chains config: " + b.config)
	}
	bs, err := ioutil.ReadAll(f)
	if err != nil {
		panic("Failed to read chains config: " + b.config)
	}

	var schema map[string]interface{}
	if err := json.Unmarshal(bs, &schema); err != nil {
		panic("Failed to parse chains config: " + b.config)
	}

	b.chains = map[string]*Chain{}
	for name, c := range schema {
		b.chains[name] = b.parse(interface{}(c))
	}
}

func (b *Bank) parse(conf interface{}) *Chain {
	c := &Chain{
		Links: []unit.Caller{},
	}

	for f, links := range conf.(map[string]interface{}) {
		if flow := FlowOf(f); flow != UnknownFlow {
			c.Flow = flow
		} else {
			panic("Unknown chain flow: " + f)
		}

		for _, link := range links.([]interface{}) {
			val := reflect.ValueOf(link)

			switch val.Kind() {
			case reflect.Map:
				subchain := b.parse(interface{}(link))
				if len(subchain.Links) > 0 {
					c.Links = append(c.Links, unit.Caller(subchain))
				}
			case reflect.String:
				name := link.(string)
				if caller, ok := b.units.Unit(name); ok {
					c.Links = append(c.Links, caller)
				} else {
					fmt.Println("Unknown unit:", name)
				}
			default:
				panic("Unexpected chain element: " + val.Kind().String())
			}
		}
	}

	return c
}
