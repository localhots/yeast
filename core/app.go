package core

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/localhots/yeast/chain"
)

type (
	App struct {
		config *Config
		chains *chain.Bank
		sv     *Supervisor
	}
)

func NewApp() *App {
	conf := &Config{}
	conf.Init()

	a := &App{
		config: conf,
		chains: chain.NewBank(conf.C().ChainsConfig, conf.C().UnitsConfig),
		sv:     NewSupervisor(conf.C().Python.BinPath, conf.C().Python.WrapperPath),
	}
	a.chains.Reload()

	return a
}

func (a *App) Conf() Config {
	// This is terrible
	return a.config.conf.Config().(Config)
}

func (a *App) Call(chainName string, data []byte) (resp []byte, err error) {
	if c, ok := a.chains.Chain(chainName); ok {
		return c.Call(data)
	} else {
		return nil, fmt.Errorf("Unknown chain: %s", chainName)
	}
}

func (a *App) HandleInterrupt() {
	shutdown := make(chan os.Signal)
	signal.Notify(shutdown, os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		<-shutdown
		log.Print("Interrupt recieved. Shutting down...")
		a.sv.Shutdown()
		os.Exit(0)
	}()
}

func (a *App) BootChain(name string) {
	if c, ok := a.chains.Chain(name); ok {
		a.sv.Start(c.Units()...)
		return
	}
	panic(fmt.Errorf("Unknown chain: %s", name))
}

func (a *App) ChainUnits(name string) []string {
	if c, ok := a.chains.Chain(name); ok {
		return c.Units()
	}
	panic(fmt.Errorf("Unknown chain: %s", name))
}
