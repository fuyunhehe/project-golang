package closure

import "fmt"

func CreateFunc() []func() {
	var a = make([]func(), 0, 10)
	for i := 0; i < 10; i++ {
		a = append(a, func(i int) func() {
			return func() {
				fmt.Println(i)
			}
		}(i))
	}
	return a
}

func Performance() int {
	f := p([]int{5000, 4000}, []string{"5000", "4000"})
	return f()
}

func Performance2() int {
	return p2(5000, 5000)
}

func p(n []int, m []string) func() int {
	return func() int {
		j := 0
		for i := 0; i < len(n)+len(m); i++ {
			j = -(j + i)
		}
		return j
	}
}

func p2(n, m int) int {
	j := 0
	for i := 0; i < n+m; i++ {
		j = -(j + i)
	}
	return j
}
