package chain

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"

	"github.com/localhots/yeast/unit"
)

var (
	chains = map[string]*Chain{}
)

func LoadChains(path string) {
	f, err := os.Open(path)
	if err != nil {
		panic("Failed to open chains config: " + path)
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		panic("Failed to read chains config: " + path)
	}

	var schema map[string]interface{}
	if err := json.Unmarshal(b, &schema); err != nil {
		panic("Failed to parse chains config: " + path)
	}

	for name, c := range schema {
		chains[name] = Parse(interface{}(c))
	}
}

func Parse(conf interface{}) *Chain {
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
				subchain := Parse(interface{}(link))
				if len(subchain.Links) > 0 {
					c.Links = append(c.Links, unit.Caller(subchain))
				}
			case reflect.String:
				name := link.(string)
				if caller := unit.New(name); caller != nil {
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
