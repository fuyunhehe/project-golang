package sort

import (
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func testCases(t testing.TB, fn func([]int), n, max int, see bool) {
	cases := MakeCases(n, max)
	for i, c := range cases {
		out := make([]int, len(c.In))
		copy(out, c.In)
		fn(out)
		if !c.match(out) {
			t.Errorf("%d => In:%+v, out:%+v\n", i, c.In, out)
		} else if see {
			t.Logf("%d => In:%+v, out:%+v\n", i, c.In, out)
		}
	}
}

func TestBubble(t *testing.T) {
	testCases(t, Bubble, 100, 10, false)
}

func TestInsert(t *testing.T) {
	testCases(t, Insert, 100, 10, false)
}

func TestSelect(t *testing.T) {
	testCases(t, Select, 100, 10, false)
}

func TestQuick(t *testing.T) {
	testCases(t, Quick, 100, 10, false)
}

func TestMerge(t *testing.T) {
	testCases(t, Merge, 100, 10, false)
}
