package sort

import (
	"math/rand"
	"sort"
)

func exchage(ori []int, i, j int) {
	ori[i], ori[j] = ori[j], ori[i]
}

type SortCase struct {
	In []int
}

func (c *SortCase) match(out []int) bool {
	if len(c.In) != len(out) {
		return false
	}
	want := make([]int, len(c.In))
	copy(want, c.In)
	sort.Ints(want)
	return compare(want, out)
}

func MakeCases(n int, max int) (scs []SortCase) {
	scs = make([]SortCase, n, n)
	for i := 0; i < n; i++ {
		var l int
		if max > 0 {
			l = rand.Intn(max)
		} else {
			l = rand.Int()
		}
		l += 1
		var sc = SortCase{
			In: make([]int, l, l),
		}
		for j := 0; j < l; j++ {
			sc.In[j] = rand.Intn(l<<3) + 1
		}
		scs[i] = sc
	}
	return
}

func compare(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
