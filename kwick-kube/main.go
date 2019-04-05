package main

import (
	"log"
	"os/exec"

	"github.com/ngtrimble/executil"
)

func main() {
	cmd := ""
	runCmd(cmd)
}

func runCmd(cmdString string) {
	cmd := exec.Command("sh", "-c", cmdString)
	err := executil.StartWaitCombinedStdout(cmd)
	if err != nil {
		log.Fatalf("%v", err)
	}
}
