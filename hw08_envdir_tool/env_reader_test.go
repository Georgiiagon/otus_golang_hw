package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadDir(t *testing.T) {
	envMap, err := ReadDir("testdata/env")
	require.NoError(t, err)
	require.Equal(t, envMap["BAR"], EnvValue{Value: "bar", NeedRemove: false})
	require.Equal(t, envMap["EMPTY"], EnvValue{Value: "", NeedRemove: true})
	require.Equal(t, envMap["FOO"], EnvValue{Value: "   foo\nwith new line", NeedRemove: false})
	require.Equal(t, envMap["HELLO"], EnvValue{Value: "\"hello\"", NeedRemove: false})
	require.Equal(t, envMap["UNSET"], EnvValue{Value: "", NeedRemove: true})
}
