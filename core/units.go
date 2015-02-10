package core

import (
	"encoding/json"

	"github.com/localhots/yeast/impl"
	"github.com/localhots/yeast/unit"
)

var (
	Units = map[string]Caller{}
)

func LoadUnits(b []byte) {
	var conf map[string]map[string]interface{}
	json.Unmarshal(b, &conf)

	for name, meta := range conf {
		// Check for unit implementation and create a unit if there is none
		if imp := impl.New(meta["impl"].(string)); imp != nil {
			Units[name] = imp
		} else {
			Units[name] = &unit.Unit{
				Name: name,
			}
		}
	}
}
