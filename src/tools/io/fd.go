package tools/io

import (
	"os"
)

func ReadLine() (line string, err error) {
	
}

func FileSize(path) (size int, err error) {
	finfo, err := os.Stat(path)
	if err != nil {
		return
	}

	size := int(finfo.Size())
	return
}