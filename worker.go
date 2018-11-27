package main

import (
	"log"
	"os/exec"
)

type Worker struct {
	Command string
	Args    string
	Output  chan string
}

func (cmd *Worker) Run() {
	out, err := exec.Command(cmd.Command, cmd.Args).Output()
	if err != nil {
		log.Fatal(err)
	}

	cmd.Output <- string(out)
}
