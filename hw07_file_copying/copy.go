package main

import (
	"errors"
	"os"
	"log"
	"io"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func fileCloser (files ...*os.File) {
	for _,file := range files {
		err := file.Close()
		if err != nil {
			log.Fatal("error with close file ", err)
		}
	}
}

func Copy(fromPath, toPath string, offset, limit int64) error {
	fileInfo, err := os.Stat(fromPath)
	if err != nil {
		log.Fatal("error to get file info:", err)
		return err
	}

	if offset > fileInfo.Size() {
		return ErrOffsetExceedsFileSize
	}

	if limit > fileInfo.Size() {
		limit = fileInfo.Size()
	}

	srcFile, err := os.Open(fromPath)
	if err != nil {
		log.Fatal("faild to open file:", err)
		return err
	}
	
	dstFile, err := os.Create(toPath)
	if err != nil {
		log.Fatal("faild to create file:",err)
	}


	_, err = srcFile.Seek(offset, io.SeekStart)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(limit)
	buf := make([]byte, limit)
	_, err = io.ReadFull(srcFile,buf)
	if err != nil {
		log.Fatal("faild to read file:", err)
	}

	_, err = dstFile.Write(buf)
	if err != nil {
		log.Fatal("faild to write file:", err)
	}

	fileCloser(dstFile,srcFile)


	return nil
}
