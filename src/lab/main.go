package main

import (
	"command"
	"git"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		subcommand := os.Args[1]
		if c, ok := command.LabCommands[subcommand]; ok {
			c(os.Args[1:]...)
		} else {
			git.RunGit(os.Args...)
		}
	} else {
		git.RunGit(os.Args...)
	}
}
