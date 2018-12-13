package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("------------" + string(i))
		{
			fmt.Println("aaaaaaaaaaaa")
			break
			fmt.Println("bbbbbbbbbbbb")
		}
		fmt.Println("+++++++++++++" + string(i))
	}
}
