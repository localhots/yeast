package core

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

type (
	Supervisor struct {
		pythonBin     string
		pythonWrapper string
	}
)

// XXX: We're about to spawn hundreds of Python processes
func (s *Supervisor) StartAll(units []string) {
	for _, name := range units {
		s.Start(name)
		time.Sleep(500 * time.Millisecond) // Don't spawn processes too fast
	}
}

func (s *Supervisor) Start(name string) {
	fmt.Println("Starting unit: " + name)
	cmd := exec.Command(s.pythonBin, s.pythonWrapper, name)
	cmd.Stdout = os.Stdout // Sorry
	if err := cmd.Start(); err != nil {
		fmt.Println("Failed to start unit: ", name)
		fmt.Println(err.Error())
	}
}
