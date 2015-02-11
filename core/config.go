package core

import (
	"encoding/json"
	"log"
	"os"

	"github.com/localhots/confection"
)

type (
	Config struct {
		conf         *confection.Manager
		ChainsConfig string `json:"chains_config_path" attrs:"required" title:"Chains config path"`
		UnitsConfig  string `json:"units_config_path" attrs:"required" title:"Units config path"`
		Python       Python `json:"python" title:"Python"`
	}
	Python struct {
		BinPath     string `json:"bin_path" attrs:"required" title:"Python 3 binary path"`
		WrapperPath string `json:"wrapper_path" attrs:"required" title:"Path to wrapper.py"`
	}
)

func (c *Config) Init() {
	log.SetOutput(os.Stderr)
	log.SetFlags(log.Ltime)
	log.SetPrefix("YEAST @ ")

	c.conf = confection.New(*c, c.decoder)
	go c.conf.StartServer()
	c.conf.RequireConfig()
}

func (c *Config) decoder(b []byte) interface{} {
	var newConf Config
	if err := json.Unmarshal(b, &newConf); err != nil {
		panic(err)
	}
	return newConf
}
