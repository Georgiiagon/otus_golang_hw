package main

import (
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	command := exec.Command(cmd[0], cmd[1:]...) //nolint:gosec
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	prepareEnv(env)
	command.Run()

	return command.ProcessState.ExitCode()
}

func prepareEnv(env Environment) {
	for name, value := range env {
		os.Setenv(name, value.Value)
	}
}
