package main

import (
	"log"
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	command := exec.Command(cmd[0], cmd[1:]...) //nolint:gosec
	command.Stdout = os.Stdout
	prepareEnv(env)
	if err := command.Run(); err != nil {
		log.Fatal(err)
	}

	return command.ProcessState.ExitCode()
}

func prepareEnv(env Environment) {
	for name, value := range env {
		os.Setenv(name, value.Value)
	}
}
