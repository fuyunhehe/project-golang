package main

import (
	"fmt"
)

type AA interface {
	Println(a ...interface{}) (n int, err error)
}

type Fmt struct {
	Name string
}

func (f *Fmt) Println(a ...interface{}) (n int, err error) {
	n, err = fmt.Println(f.Name, a)
	return
}

func main() {
	aaa := &Fmt{Name: "Fmtobj"}
	//var bb AA
	var bb AA = aaa
	//bb.Println("bb")
	if a, ok := bb.(*Fmt); ok {
		a.Println("a")
	}
}
