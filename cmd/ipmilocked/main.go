package main

import (
	"os"
	"os/exec"

	"github.com/FoxDenHome/superfan/drivers/control"
)

func main() {
	cmd := exec.Command(os.Args[1], os.Args[2:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err := control.RunLockedVoid(cmd.Run)
	exitCode := 1
	if cmd.ProcessState != nil {
		exitCode = cmd.ProcessState.ExitCode()
	}
	if err != nil && exitCode == 0 {
		exitCode = 1
	}
	os.Exit(exitCode)
}
