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

type PcomScript struct {
	Name   string `json:"name"`
	Script string `json:"script"`
}

type PcomConfig struct {
	Scripts []PcomScript `json:"scripts"`
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

	for i := 0; len(payload.Scripts) > i; i++ {
		if payload.Scripts[i].Name == params[0] {
			scriptArray := strings.Split(payload.Scripts[i].Script, " ")
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
