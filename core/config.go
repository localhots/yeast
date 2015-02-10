package core

import (
	"encoding/json"
)

type (
	Config struct {
		ChainsConfig string `json:"chains_config_path" attrs:"required" title:"Chains config path"`
		UnitsConfig  string `json:"units_config_path" attrs:"required" title:"Units config path"`
		Python       Python `json:"python" title:"Python"`
	}
	Python struct {
		BinPath     string `json:"bin_path" attrs:"required" title:"Python 3 binary path"`
		WrapperPath string `json:"wrapper_path" attrs:"required" title:"Path to wrapper.py"`
	}
)

func ConfigDecoder(b []byte) interface{} {
	var newConf Config
	if err := json.Unmarshal(b, &newConf); err != nil {
		panic(err)
	}
	return newConf
}
