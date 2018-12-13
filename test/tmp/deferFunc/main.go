package main

import "fmt"

func deferTest() (aa string) {
	aa = "bbbb"
	defer func() {
		fmt.Println(aa)
		aa = "eee"
	}()
	aa = "ccc"
	return "dddd"
}

func main()  {
	ff := deferTest()
	fmt.Println(ff)
}
