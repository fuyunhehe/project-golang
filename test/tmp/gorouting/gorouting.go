package gorouting

import (
	"sync"
	"fmt"
)

func WaitGroup() {
	defer func() {
		if msg := recover(); msg != nil {
			fmt.Println("recover", msg)
		}
	}()

	n := 2
	wg := sync.WaitGroup{}
	wg.Add(n)
	go func() {
		wg.Done()
	}()
	go func() {
		defer func() {
			if msg := recover(); msg != nil {
				fmt.Println("go recover", msg)
			}
		}()
		panic("aaaa")
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("wait end")
}
