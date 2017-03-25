package tools\io

import (
	"testing"
)

func TestFileSize(t *testing.T) {
	cases := []struct{
		in string
	}{
		{"fd.go"},
	}

	for _, c := range cases {
		got, err := FileSize(c.in)
		if err != nil {
			t.Errorf("in:%s\terr:%s\n", c.in, err.Error())
		}
		t.Logf("in:%s\tgot:%d\n", c.in, got)
	}
}