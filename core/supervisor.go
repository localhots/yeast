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
	}
)

// XXX: We're about to spawn hundreds of Python processes
func (s *Supervisor) Start(units ...string) {
	for _, name := range units {
		log.Printf("Starting unit: %s", name)

		cmd := exec.Command(s.Bin, s.Wrapper, name)
		cmd.Stdout = os.Stderr // Sorry
		if err := cmd.Start(); err != nil {
			log.Printf("Failed to start unit: %s (%s)", name, err.Error())
		}

		time.Sleep(200 * time.Millisecond) // Don't spawn processes too fast
	}
}
