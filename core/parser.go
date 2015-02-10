package core

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func ParseChains(b []byte) (map[string]*Chain, error) {
	var schema map[string]interface{}
	if err := json.Unmarshal(b, &schema); err != nil {
		return nil, err
	}

	chains := map[string]*Chain{}
	for name, chain := range schema {
		chains[name] = buildChain(interface{}(chain))
	}

	return chains, nil
}

func buildChain(conf interface{}) *Chain {
	c := &Chain{
		Links: []Caller{},
	}

	for flow, links := range conf.(map[string]interface{}) {
		if f, ok := FlowMap[flow]; ok {
			c.Flow = f
		} else {
			panic("Unknown chain flow: " + flow)
		}

		for _, link := range links.([]interface{}) {
			val := reflect.ValueOf(link)

			switch val.Kind() {
			case reflect.Map:
				subchain := buildChain(interface{}(link))
				if len(subchain.Links) > 0 {
					c.Links = append(c.Links, Caller(subchain))
				}
			case reflect.String:
				unit, ok := Units[link.(string)]
				if !ok {
					fmt.Println("Unknown unit `" + link.(string) + "`")
				} else {
					c.Links = append(c.Links, unit)
				}
			default:
				panic("Unexpected chain element: " + val.Kind().String())
			}
		}
	}

	return c
}
