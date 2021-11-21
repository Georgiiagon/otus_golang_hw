package main

import (
	"errors"
	"log"
	"os"
)

var errArguments = errors.New("expected at least 2 arguments: path to env dir and command to execute")

func main() {
	args := os.Args[1:]
	if len(args) < 2 {
		log.Fatal(errArguments)
	}
	env, err := ReadDir(args[0])
	if err != nil {
		log.Fatal(err)
	}
	RunCmd(args[1:], env)
}
