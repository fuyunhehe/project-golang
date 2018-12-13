package rangeT

import (
	"math/rand"
	"time"
	"fmt"
	"sync"
)

func Call1() {
	a := make([]int, 10, 10)
	for i := 0; i < 10; i++ {
		a[i] = i
	}

	wg := sync.WaitGroup{}
	wg.Add(2)
	tmp := func(n int) {
		for i := range a {
			r := rand.Intn(3)
			time.Sleep(time.Millisecond * 100 * time.Duration(r))
			fmt.Println(n, i)
		}
		wg.Done()
	}
	go tmp(1)
	go tmp(2)
	wg.Wait()
}

func ForSlice(s []int) {
	len := len(s)
	for i := 0; i < len; i++ {
		s[i] = s[i] + i
	}
}

func RangeForSlice(s []int) {
	for i, v := range s {
		s[i] = v + i
	}
}
