package core

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/localhots/yeast/unit"
)

type (
	Supervisor struct{}
)

// XXX: We're about to spawn hundreds of Python processes
func (s *Supervisor) StartAll() {
	for _, name := range unit.Units() {
		s.Start(name)
		time.Sleep(500 * time.Millisecond) // Don't spawn processes too fast
	}
}

func (s *Supervisor) Start(name string) {
	fmt.Println("Starting unit: " + name)
	conf := Conf().Python
	cmd := exec.Command(conf.BinPath, conf.WrapperPath, name)
	cmd.Stdout = os.Stdout // Sorry
	if err := cmd.Start(); err != nil {
		fmt.Println("Failed to start unit: ", name)
		fmt.Println(err.Error())
	}
}
