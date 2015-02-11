package unit

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/localhots/yeast/impl"
)

type (
	Bank struct {
		config string
		units  map[string]*Unit
	}
)

func NewBank(config string) *Bank {
	return &Bank{
		config: config,
		units:  map[string]*Unit{},
	}
}

func (b *Bank) Unit(name string) Caller {
	if u, ok := b.units[name]; ok {
		// Check for unit implementation and create a unit if there is none
		if imp := impl.New(u.Impl); imp != nil {
			return imp
		} else {
			return u
		}
	} else {
		return nil
	}
}

func (b *Bank) Reload() {
	f, err := os.Open(b.config)
	if err != nil {
		panic("Failed to open units config: " + b.config)
	}
	bs, err := ioutil.ReadAll(f)
	if err != nil {
		panic("Failed to read units config: " + b.config)
	}

	var conf map[string]map[string]interface{}
	if err := json.Unmarshal(bs, &conf); err != nil {
		panic("Failed to parse units config: " + b.config)
	}

	b.units = map[string]*Unit{}
	for name, meta := range conf {
		b.units[name] = &Unit{
			Name:   name,
			Impl:   meta["impl"].(string),
			Config: meta["config"],
		}
	}
}

func (b *Bank) Units() []string {
	list := []string{}
	for name, _ := range b.units {
		list = append(list, name)
	}
	return list
}
