package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunCmd(t *testing.T) {
	env, err := ReadDir("testdata/env")
	require.NoError(t, err)
	code := RunCmd([]string{"ls", "-l", "-a"}, env)

	require.Equal(t, code, 1)
}
