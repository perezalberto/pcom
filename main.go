package main

import (
	"commands"
	"core"
	"os"
)

func main() {
	commandBase := core.CommandList{}

	commandBase.AddCommand("run", new(commands.RunCommand))
	commandBase.AddCommand("init", new(commands.InitCommand))

	if len(os.Args[1:]) == 0 {
		return
	}
	commandBase.RunCommand(os.Args[1], os.Args[2:])
}
