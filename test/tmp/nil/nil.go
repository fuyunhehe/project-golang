package main

import (
	"fmt"
	"reflect"
)

type myerror interface {
	MyError() string
}

type Error struct{}

func checkError(err myerror) {
	//if err != nil {
	//	panic(err)
	//}
}

func (e *Error) MyError() string {
	return string("aaaaError")
}

func (e *Error) String() string {
	return string("aaaaString")
}

type mInt int

func (e mInt) MyError() string {
	return string("aaaa mInt\n")
}

//func SS(i interface{}) {
//	switch f:= i.(type) {
//		case mInt:
//			fmt.Println(uint64(f))
//		case 
//	}
//}

func BB(b []string) {
	a := make([]string, 0, 0)
	fmt.Println(a, len(a), a == nil)
	if b == nil {
		fmt.Println("nilnil")
	}
	fmt.Println(b, len(b))
}

func main() {
	var e *Error
	var a mInt
	fmt.Printf("%T\t%T\n", e, a)
	fmt.Println(e, a)
	fmt.Println(reflect.ValueOf(e))
	BB(nil)
	//SS(e)
	//n, err := fmt.Println(e, a)
	//fmt.Println("-----------", n, err)
	//checkError(e)
}
