package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name   string
	Career string
	Age    int64
}

func (p *Person) MyName() string {
	return p.Name
}

func NewPerson(name string, career string, age int64) *Person {
	return &Person{name, career, age}
}

func main() {
	var num uint64 = 23
	t := reflect.TypeOf(num)
	fmt.Printf("name:%s,kind:%s,methodnum:%d,\n", t.Name(), t.Kind(), t.NumMethod())

	youyou := Person{Name: "liufei", Age: 22, Career: "Coder"}
	t2 := reflect.TypeOf(youyou)
	fmt.Printf("name:%s,kind:%s,nummethod:%d\n", t2.Name(), t2.Kind(), t2.NumMethod())

	fmt.Println(youyou.MyName())

	p2 := NewPerson("liufei2", "Coder", 24)
	fmt.Println(p2.MyName())
}
