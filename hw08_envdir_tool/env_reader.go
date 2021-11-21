package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	env := make(Environment)
	dirInfo, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, fileStat := range dirInfo {
		file, err := os.OpenFile(dir+"/"+fileStat.Name(), os.O_RDONLY, fileStat.Mode())
		if err != nil {
			log.Fatal(err)
		}
		reader := bufio.NewReader(file)
		line, _, _ := reader.ReadLine()
		file.Close()

		line = prepareValue(line)

		env[fileStat.Name()] = EnvValue{
			Value:      string(line),
			NeedRemove: len(line) == 0,
		}
	}

	return env, nil
}

func prepareValue(line []byte) []byte {
	for i, byte := range line {
		if byte == 0 {
			line[i] = '\n'
		}
	}

	if len(line) > 0 && line[len(line)-1] == ' ' {
		line = line[:len(line)-1]
	}

	return line
}
