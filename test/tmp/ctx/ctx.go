package main

import (
	"reflect"
	"fmt"
)

type mySubType struct {
	n int
}

type myType struct {
	n int
	s string
	t *mySubType
}

func (m *myType) Comparable() bool {
	return true
}

func main()  {
	rt := reflect.TypeOf(myType{2, "hello", nil})
	//rt := reflect.TypeOf(myType{2, "hello", map[string]string{"k":"v"}})
	fmt.Println(rt.Comparable())
}
