package bashscript

import (
	"os/exec"
)

type BashScriptPayload struct {
	Project   string   `json:"project"`
	ParamsStr string   `json:"params"`
	Params    []string `json:"-"`
}

func (b *BashScriptPayload) Execute() error {
	cmd := exec.Command("./scripts/scripts.sh", b.Params...)
	return cmd.Run()
}
