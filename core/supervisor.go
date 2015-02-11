package core

import (
	"log"
	"os"
	"os/exec"
	"time"
)

type (
	Supervisor struct {
		Bin     string
		Wrapper string
		procs   map[string]*exec.Cmd
	}
)

func NewSupervisor(bin, wrapper string) *Supervisor {
	return &Supervisor{
		Bin:     bin,
		Wrapper: wrapper,
		procs:   map[string]*exec.Cmd{},
	}
}

// XXX: We're about to spawn hundreds of Python processes
func (s *Supervisor) Start(units ...string) {
	for _, name := range units {
		log.Printf("Starting unit: %s", name)

		cmd := exec.Command(s.Bin, s.Wrapper, name)
		cmd.Stdout = os.Stderr // Sorry
		if err := cmd.Start(); err != nil {
			log.Printf("Failed to start unit: %s (%s)", name, err.Error())
		}
		s.procs[name] = cmd

		time.Sleep(200 * time.Millisecond) // Don't spawn processes too fast
	}
}
