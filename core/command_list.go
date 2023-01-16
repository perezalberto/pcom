package core

import (
	"errors"
	"fmt"
)

type CommandList struct {
	commands []ICommand
	names    []string
}

func (obj *CommandList) AddCommand(name string, command ICommand) {
	obj.names = append(obj.names, name)
	obj.commands = append(obj.commands, command)
}

func (obj *CommandList) GetCommand(name string) (ICommand, error) {
	for i := 0; obj.Size() > i; i++ {
		if obj.names[i] == name {
			return obj.commands[i], nil
		}
	}
	return nil, errors.New("Command not found!!")
}

func (obj *CommandList) Size() int {
	return len(obj.commands)
}

func (obj *CommandList) RunCommand(name string, params []string) {
	command, err := obj.GetCommand(name)

	if err != nil {
		fmt.Println("Command not found!!")
		return
	}

	command.Execute(params)
}
