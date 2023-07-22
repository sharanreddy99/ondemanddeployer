package bashscript

import (
	"ondemanddeployer/utils"
	"os"
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
var ActiveProject string = ""

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
	ActiveProject = ""
	cmd := exec.Command("./scripts/scripts.sh", task.Params...)


	// Make test file
	testFile, err := os.Create("test.txt")
	if err != nil {
		utils.Log("Error creating file: ",err.Error())
	}

	defer testFile.Close()

	// Redirect the output here (this is the key part)
	cmd.Stdout = testFile

	err = cmd.Start(); if err != nil {
		utils.Log("Error running command: ",err.Error())
	}
	
	cmd.Wait()
	ActiveProject = task.Project
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
	bashScriptQueue = make([]BashScriptPayload, 0)

	ticker := time.NewTicker(30 * time.Second)

	go func() {
		for range ticker.C {
			Execute()
		}
	}()
}
