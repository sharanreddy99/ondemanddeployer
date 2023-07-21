package bashscript

import (
	"ondemanddeployer/utils"
	"os/exec"
	"sync"
	"time"
)

type BashScriptPayload struct {
	Project   string   `json:"project"`
	Params    []string `json:"params"`
	TimeStamp string   `json:"timestamp"`
}

var bashScriptQueue []BashScriptPayload
var bashScriptLength int = 10
var lock sync.Mutex

func (b *BashScriptPayload) AddToQueue() {
	if len(bashScriptQueue) == bashScriptLength {
		bashScriptQueue = bashScriptQueue[1:]
	}

	bashScriptQueue = append(bashScriptQueue, *b)
}

func Execute() error {
	lock.Lock()
	defer lock.Unlock()

	var task BashScriptPayload = getNextTask()
	if len(task.Params) == 0 {
		return nil
	}

	utils.Log("Executing Task: ", task)

	cmd := exec.Command("./scripts/scripts.sh", task.Params...)
	err := cmd.Run()
	time.Sleep(1 * time.Minute)
	return err
}

func getNextTask() BashScriptPayload {
	if len(bashScriptQueue) == 0 {
		return BashScriptPayload{}
	}

	task := bashScriptQueue[0]
	bashScriptQueue = bashScriptQueue[1:]
	return task
}

func init() {
	bashScriptQueue = make([]BashScriptPayload, 10)

	ticker := time.NewTicker(2 * time.Minute)

	go func() {
		for range ticker.C {
			Execute()
		}
	}()
}
