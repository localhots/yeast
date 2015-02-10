package core

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

type (
	Supervisor struct{}
)

// XXX: We're about to spawn hundreds of Python processes
func (s *Supervisor) StartAll() {
	for unit, _ := range Units {
		s.Start(unit)
		time.Sleep(500 * time.Millisecond) // Don't spawn processes too fast
	}
}

func (s *Supervisor) Start(unit string) {
	fmt.Println("Starting unit: " + unit)
	cmd := exec.Command(Conf().Python.BinPath, Conf().Python.WrapperPath, unit)
	cmd.Stdout = os.Stdout // Sorry
	if err := cmd.Start(); err != nil {
		fmt.Println("Failed to start unit: ", unit)
		fmt.Println(err.Error())
	}
}
