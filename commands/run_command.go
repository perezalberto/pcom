package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

type PcomCommand struct {
	Name    string `json:"name"`
	Command string `json:"command"`
}

type PcomConfig struct {
	Commands []PcomCommand `json:"commands"`
}

type RunCommand struct{}

func (obj *RunCommand) Execute(params []string) {

	content, err := ioutil.ReadFile("./pcom.config.json")
	if err != nil {
		return
	}

	var payload PcomConfig
	err = json.Unmarshal(content, &payload)
	if err != nil {
		return
	}

	for i := 0; len(payload.Commands) > i; i++ {
		if payload.Commands[i].Name == params[0] {
			scriptArray := strings.Split(payload.Commands[i].Command, " ")
			cmd := exec.Command(scriptArray[0], scriptArray[1:]...)
			cmd.SysProcAttr = &syscall.SysProcAttr{Setsid: true}
			cmd.Stdout = os.Stdout
			cmd.Stdin = os.Stdin
			cmd.Stderr = os.Stderr
			cmd.Env = os.Environ()
			cmd.Run()
			return
		}
	}
	fmt.Println("Command not found!!")
}
