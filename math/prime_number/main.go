package main

import (
	"flag"
	"fmt"
)

var goroutingCnt int = 0
var goroutingCntMax int = 0

func ProcessResult(seq chan int, wait chan struct{}, max int) {
	go func() {
		goroutingCnt++
		if goroutingCntMax < goroutingCnt {
			goroutingCntMax = goroutingCnt
		}
		//fmt.Println("goroutingCnt", goroutingCnt)
		prime, ok := <-seq
		if !ok {
			close(wait)
			//fmt.Println("goroutingCnt", goroutingCnt)
			goroutingCnt--
			return
		}
		fmt.Println(prime)
		if prime > max/2 {
			for num := range seq {
				fmt.Println(num)
			}
			close(wait)
		} else {
			n := 100
			if max/4 < n {
				n = max / 4
			}
			out := make(chan int, n)
			ProcessResult(out, wait, max)
			for num := range seq {
				if num%prime != 0 {
					out <- num
				}
			}
			close(out)
		}

		goroutingCnt--
		//fmt.Println("goroutingCnt", goroutingCnt)
	}()
}

func main() {
	var num int
	flag.IntVar(&num, "num", 0, "小于num的所有素数")
	flag.Parse()
	if num < 2 {
		flag.Usage()
	}

	seq, wait := make(chan int, 100), make(chan struct{})
	ProcessResult(seq, wait, num)
	for i := 2; i < num; i++ {
		seq <- i
	}
	close(seq)
	<-wait
	fmt.Println("max", goroutingCntMax)
}
