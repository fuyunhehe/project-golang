package mdefer

import (
	"testing"
	"fmt"
)

func TestDefer1(t *testing.T) {
	ff := Defer1()
	fmt.Println(ff)
}

func TestPanicDefer(t *testing.T) {
	c, e := PanicDefer(5, 10)
	t.Log(c, e)
}
