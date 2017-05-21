package io

import (
	"bufio"
	"os"
	//"io"
	//"log"
)

type File struct {
	Path string
	file *os.File
	br   *bufio.Reader
}

func NewFile(path string) *File {
	return &File{Path: path}
}

func (f *File) Size() (size int64, err error) {
	finfo, err := os.Stat(f.Path)
	if err != nil {
		return
	}
	size = finfo.Size()
	return
}

func (f *File) Open() (err error) {
	f.file, err = os.Open(f.Path)
	if err != nil {
		return
	}
	f.br = bufio.NewReader(f.file)
	return
}

func (f *File) Close() error {
	if f.file != nil {
		return f.file.Close()
	}
	return nil
}

func (f *File) ReadLine() (line string, err error) {
	if f.br == nil {
		err = f.Open()
		if err != nil {
			return
		}
	}

	for {
		pbytes, isPrefix, err := f.br.ReadLine()
		if err != nil {
			return line, err
		}

		line += string(pbytes)
		if isPrefix {
			continue
		}
		line += "\n"
		return line, err
	}
	return
}
