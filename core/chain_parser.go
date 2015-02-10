package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
)

func ParseChains() (map[string]*Chain, error) {
	f, err := os.Open(Conf().ChainsConfig)
	if err != nil {
		panic("Failed to open chains config: " + Conf().ChainsConfig)
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		panic("Failed to parse chains config: " + Conf().ChainsConfig)
	}

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
				subchain := buildChain(interface{}(link))
				if len(subchain.Links) > 0 {
					c.Links = append(c.Links, Caller(subchain))
				}
			case reflect.String:
				name := link.(string)
				caller, ok := Units[name]
				if !ok {
					fmt.Println("Unknown unit:", name)
				} else {
					c.Links = append(c.Links, caller)
				}
			default:
				panic("Unexpected chain element: " + val.Kind().String())
			}
		}
	}

	return c
}
