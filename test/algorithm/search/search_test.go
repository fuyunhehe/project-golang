package search

import (
	"math/rand"
	"time"
	"testing"
	"test/algorithm/sort"
	"fmt"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func TestBinarySearch(t *testing.T) {
	cases := sort.MakeCases(10, 15)
	for _, c := range cases {
		sort.Quick(c.In)
		tmp := rand.Intn(len(c.In))
		n := c.In[tmp] + rand.Intn(3) - 1
		i := BinarySearch(c.In, n)
		fmt.Println(c.In, n, i)
	}
}
