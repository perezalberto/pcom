package commands

import "os"

type InitCommand struct{}

func (obj *InitCommand) Execute(params []string) {
	path, _ := os.Getwd()
	_ = os.WriteFile(path+"/pcom.config.json", []byte("{\n    \"commands\": [\n        {\n            \"name\": \"\",\n            \"command\": \"\"\n        }\n    ]\n}"), 0644)
}
