package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	testFileName = "testdata/input.txt"
	newFileName  = "tmp/new_input.txt"
)

func TestCopy(t *testing.T) {
	tempFile, _ := os.CreateTemp("tmp", "new_input_*.txt")
	Copy(testFileName, tempFile.Name(), 2, 0)
	defer os.Remove(tempFile.Name())
	testFile, _ := os.OpenFile(testFileName, os.O_RDONLY, fileMode)
	newFile, _ := os.OpenFile(tempFile.Name(), os.O_RDONLY, fileMode)
	testFileStat, _ := testFile.Stat()
	newFileStat, _ := newFile.Stat()
	require.Equal(t, testFileStat.Size()-2, newFileStat.Size())
}

func TestNegativeLimitCopy(t *testing.T) {
	tempFile, _ := os.CreateTemp("tmp", "new_input_*.txt")
	Copy(testFileName, tempFile.Name(), 2, -1)
	defer os.Remove(tempFile.Name())
	newFile, _ := os.OpenFile(tempFile.Name(), os.O_RDONLY, fileMode)
	newFileStat, _ := newFile.Stat()
	require.FileExists(t, tempFile.Name())
	require.Equal(t, int64(0), newFileStat.Size())
}

func TestNegativeOffsetCopy(t *testing.T) {
	tempFile, _ := os.CreateTemp("tmp", "new_input_*.txt")
	Copy(testFileName, tempFile.Name(), -1, 0)
	defer os.Remove(tempFile.Name())
	testFile, _ := os.OpenFile(testFileName, os.O_RDONLY, fileMode)
	newFile, _ := os.OpenFile(tempFile.Name(), os.O_RDONLY, fileMode)
	testFileStat, _ := testFile.Stat()
	newFileStat, _ := newFile.Stat()
	require.FileExists(t, tempFile.Name())
	require.Equal(t, testFileStat.Size(), newFileStat.Size())
}

func TestTooBigOffset(t *testing.T) {
	testFile, _ := os.OpenFile(testFileName, os.O_RDONLY, fileMode)
	testFileStat, _ := testFile.Stat()
	err := Copy(testFileName, newFileName, testFileStat.Size()+1, 0)
	require.NoFileExists(t, newFileName)
	require.EqualError(t, err, "offset exceeds file size")
}

func TestCopyWithoutFile(t *testing.T) {
	notExistsFile := "testdata/no_such_file.txt"
	err := Copy(notExistsFile, newFileName, 0, 0)
	require.NoFileExists(t, newFileName)
	require.EqualError(t, err, "open "+notExistsFile+": no such file or directory")
}

func TestCantCreateFile(t *testing.T) {
	tmpFileName := "tmp/tmp/new_input.txt"
	err := Copy(testFileName, tmpFileName, 0, 0)
	require.NoFileExists(t, tmpFileName)
	require.EqualError(t, err, "open "+tmpFileName+": no such file or directory")
}

func TestDifferentCases(t *testing.T) {
	tests := []struct {
		offset       int
		limit        int
		testFileName string
	}{
		{
			offset:       0,
			limit:        0,
			testFileName: "out_offset0_limit0.txt",
		},
		{
			offset:       0,
			limit:        10,
			testFileName: "out_offset0_limit10.txt",
		},
		{
			offset:       0,
			limit:        1000,
			testFileName: "out_offset0_limit1000.txt",
		},
		{
			offset:       0,
			limit:        10000,
			testFileName: "out_offset0_limit10000.txt",
		},
		{
			offset:       100,
			limit:        1000,
			testFileName: "out_offset100_limit1000.txt",
		},
		{
			offset:       6000,
			limit:        1000,
			testFileName: "out_offset6000_limit1000.txt",
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.testFileName, func(t *testing.T) {
			tempFile, _ := os.CreateTemp("tmp", "*_"+tc.testFileName)

			Copy(testFileName, tempFile.Name(), int64(tc.offset), int64(tc.limit))
			defer os.Remove(tempFile.Name())
			testFile, _ := os.OpenFile("testdata/"+tc.testFileName, os.O_RDONLY, fileMode)
			newFile, _ := os.OpenFile(tempFile.Name(), os.O_RDONLY, fileMode)
			testFileStat, _ := testFile.Stat()
			newFileStat, _ := newFile.Stat()
			require.Equal(t, testFileStat.Size(), newFileStat.Size())
		})
	}
}
