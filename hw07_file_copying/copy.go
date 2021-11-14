package main

import (
	"errors"
	"io"
	"io/fs"
	"os"

	"github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
	fileMode                 = os.ModeDir
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	file, err := os.OpenFile(fromPath, os.O_RDONLY, fileMode)
	if err != nil {
		return err
	}
	defer file.Close()

	fileStat, err := fileStat(file, offset)
	if err != nil {
		return err
	}

	count := fileStat.Size()
	if limit == 0 {
		limit = count
	}
	step := int64(100)
	bar := pb.StartNew(int(count/step) + 1)
	defer bar.Finish()

	newFile, err := os.Create(toPath)
	if err != nil {
		return err
	}
	defer newFile.Close()

	if offset != 0 {
		file.Seek(offset, 0)
	}

	for {
		if limit < step {
			step = limit
		}
		_, err := io.CopyN(newFile, file, step)

		if errors.Is(err, io.EOF) || limit == 0 {
			break
		}
		bar.Increment()
		limit -= step
	}

	return nil
}

func fileStat(file *os.File, offset int64) (fs.FileInfo, error) {
	fileStat, err := file.Stat()
	if err != nil {
		return nil, err
	}

	if fileStat.Size() == 0 {
		return nil, ErrUnsupportedFile
	}

	if fileStat.Size() < offset {
		return nil, ErrOffsetExceedsFileSize
	}

	return fileStat, nil
}
