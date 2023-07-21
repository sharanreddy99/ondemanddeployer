package bashscript

import (
	"os/exec"
)

type BashScriptPayload struct {
	Project string   `json:"project"`
	Params  []string `json:"params"`
}

func (b *BashScriptPayload) Execute() error {
	cmd := exec.Command("./scripts/scripts.sh", b.Params...)
	return cmd.Run()
}
