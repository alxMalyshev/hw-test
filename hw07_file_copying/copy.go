package main

import (
	"errors"
	"io"
	"log"
	"os"

	"github.com/cheggaaa/pb"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func fileCloser(files ...*os.File) {
	for _, file := range files {
		err := file.Close()
		if err != nil {
			log.Panic("error with close file ", err)
		}
	}
}

func Copy(fromPath, toPath string, offset, limit int64) error {
	fileInfo, err := os.Stat(fromPath)
	if err != nil {
		log.Panic("error to get file info:", err)
		return err
	}

	if !fileInfo.Mode().IsRegular() {
		return ErrUnsupportedFile
	}

	if offset > fileInfo.Size() {
		return ErrOffsetExceedsFileSize
	}

	if limit+offset > fileInfo.Size() {
		limit = fileInfo.Size() - offset
	} else if limit == 0 || limit > fileInfo.Size() {
		limit = fileInfo.Size()
	}

	srcFile, err := os.Open(fromPath)
	if err != nil {
		log.Panic("faild to open file:", err)
		return err
	}

	dstFile, err := os.Create(toPath)
	if err != nil {
		log.Panic("faild to create file:", err)
	}

	_, err = srcFile.Seek(offset, io.SeekStart)
	if err != nil {
		log.Panic(err)
	}

	bar := pb.New64(limit).SetUnits(pb.U_BYTES)
	bar.Start()

	buf := make([]byte, limit)
	reader := bar.NewProxyReader(srcFile)

	_, err = io.ReadFull(reader, buf)
	if err != nil {
		log.Panic("faild to read file ", err)
	}

	_, err = dstFile.Write(buf)
	if err != nil {
		log.Panic("faild to write file:", err)
	}

	bar.Finish()

	fileCloser(dstFile, srcFile)

	return nil
}
