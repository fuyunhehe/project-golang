package chanT

import (
	"fmt"
	"time"
)

func ChanCall1() {
	c := make(chan int, 1)
	defer close(c)
	go func(c chan int) {
		time.Sleep(time.Second * 5)
		<-c
		fmt.Println("go func")
	}(c)
	fmt.Println("start")
	c <- 1
	fmt.Println("main")
	time.Sleep(time.Second * 10)
}

func ChanCall2() {
	c := make(chan int)
	go func(c chan int) {
		time.Sleep(time.Second * 5)
		n := <-c
		fmt.Println("in", n)
		n = <-c
		fmt.Println("in after", n)
	}(c)
	c <- 1
	fmt.Println("out before close")
	close(c)
	fmt.Println("out after close")
}
