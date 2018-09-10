package main

import "fmt"

func createFunc() []func() {
	var a = make([]func(), 0, 10)
	for i := 0; i < 10; i++ {
		a = append(a, func(i int) func(){return func() {
			fmt.Println(i)
		}}(i))
	}
	return a
}

func main()  {
	fList := createFunc()
	//fmt.Println(fList)
	for _, f := range fList{
		f()
	}
}