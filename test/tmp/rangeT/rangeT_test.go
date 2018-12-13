package rangeT

import "testing"

const N = 1000

func TestCall1(t *testing.T) {
	Call1()
}

func initSlice() []int {
	s := make([]int, N)
	for i := 0; i < N; i++ {
		s[i] = i
	}
	return s
}

func BenchmarkForSlice(b *testing.B) {
	s := initSlice()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ForSlice(s)
	}
}

func BenchmarkRangeForSlice(b *testing.B) {
	s := initSlice()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RangeForSlice(s)
	}
}
