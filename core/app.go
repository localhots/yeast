package core

import (
	"github.com/localhots/yeast/chain"
	"github.com/localhots/yeast/unit"
)

type (
	App struct {
		config *Config
		chains *chain.Bank
	}
)

func NewApp() *App {
	a := &App{
		config: &Config{},
	}
	a.config.Init()

	ub := unit.NewBank(a.Conf().UnitsConfig)
	a.chains = chain.NewBank(a.Conf().ChainsConfig, ub)
	a.chains.Reload()

	return a
}

func (a *App) Conf() Config {
	return a.config.conf.Config().(Config)
}
