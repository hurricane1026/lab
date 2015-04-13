package command

import (
	"fmt"
)

type Command func(args ...string)

var LabCommands map[string]Command

func printCommand(args ...string) {
	fmt.Println(args)
}

func init() {
	LabCommands = make(map[string]Command)
	LabCommands["sync"] = printCommand
	LabCommands["fork"] = printCommand
	LabCommands["pr"] = printCommand
	LabCommands["setup"] = setupCommand
	LabCommands["end"] = printCommand
	LabCommands["start"] = printCommand
	LabCommands["fix"] = printCommand
	LabCommands["create"] = printCommand

}
