package mreader

import "testing"

func TestMreader(t *testing.T) {
	if err := Mreader(); err != nil {
		t.Error(err)
	}
}
