package main

import (
	"flag"
	"fmt"
)

func main() {
	var arg1, arg2 int
	flag.IntVar(&arg1, "arg1", 0, "first arg need to sum")
	flag.IntVar(&arg2, "arg2", 0, "second arg need to sum")
	flag.Parse()

	if arg1 == 0 && arg2 == 0 {
		flag.Usage()
	}

	fmt.Println(arg1, "+", arg2, "=", arg1+arg2)
}
