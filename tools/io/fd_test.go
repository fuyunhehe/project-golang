package io

import (
	"testing"
	"io"
)

func TestFileSize(t *testing.T) {
	cases := []struct {
		in string
	}{
		{"fd.go"},
	}

	for _, c := range cases {
		f := NewFile(c.in)
		got, err := f.Size()
		if err != nil {
			t.Errorf("in:%s\terr:%s\n", c.in, err.Error())
		}
		t.Logf("in:%s\tgot:%d\n", c.in, got)
	}
}

func TestReadLine(t *testing.T) {
	cases := []struct {
		in string
	}{
		{"fd.go"},
		{"aaa.txt"},
	}

	var line string = ""
	var err error = nil
	for _, c := range cases {
		f := NewFile(c.in)
		defer f.Close()
		i := 0
		var size int = 0
		for {
			i++
			line, err = f.ReadLine()
			if err != nil {
				break
			} else {
				t.Logf("%d:\t%s\n", i, line)
			}
			size += len(line)
		}
		if err != io.EOF {
			t.Errorf("%d=>err:%s,line:%d\n", i, err.Error(), len(line))
		}
		t.Logf("in:%s,out size:%d\n", c.in, size)
	}
}
