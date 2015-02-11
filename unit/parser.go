package unit

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

var (
	units = map[string]*Unit{}
)

func LoadUnits(path string) {
	f, err := os.Open(path)
	if err != nil {
		panic("Failed to open units config: " + path)
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		panic("Failed to parse units config: " + path)
	}

	var conf map[string]map[string]interface{}
	json.Unmarshal(b, &conf)

	for name, meta := range conf {
		units[name] = &Unit{
			Name:   name,
			Impl:   meta["impl"].(string),
			Config: meta["config"],
		}
	}
}

func Units() []string {
	list := []string{}
	for name, _ := range units {
		list = append(list, name)
	}
	return list
}
