package main

import (
	"bufio"
	"bytes"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

var errInvalidChar = errors.New("string contains invalid character")

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

		if strings.Contains(string(line), "=") || strings.Contains(fileStat.Name(), "=") {
			return nil, errInvalidChar
		}

		env[fileStat.Name()] = EnvValue{
			Value:      string(line),
			NeedRemove: len(line) == 0,
		}
	}

	return env, nil
}

func prepareValue(line []byte) []byte {
	line = bytes.ReplaceAll(line, []byte{0}, []byte{'\n'})

	line = []byte(strings.TrimRight(string(line), " "))

	return line
}
