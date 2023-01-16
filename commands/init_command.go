package commands

import "os"

type InitCommand struct{}

func (obj *InitCommand) Execute(params []string) {
	path, _ := os.Getwd()
	_ = os.WriteFile(path+"/pcom.config.json", []byte("{\n    \"scripts\": []\n}"), 0644)
}
