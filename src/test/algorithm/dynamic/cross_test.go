package dynamic

import (
	"testing"
)

var cases = []struct{
	s1,s2,s3 string
	want bool
}{
	{"abcde", "fghij", "abcdefghij", true},
	{"abcde", "fghij", "afbgchdiej", true},
	{"abcde", "fghij", "afbbchdiej", false},
	{"abbccddee", "aabccdee", "aababbccccdededee", true},
	{"abbccddee", "aabccdee", "aababbccccdedeeed", false},
}

func TestIsCrossChan(t *testing.T) {
	for _, c := range cases{
		got := IsCrossChan(c.s1, c.s2, c.s3)
		if got != c.want {
			t.Error(c, got)
		}
	}
}

func TestIsCrossArr(t *testing.T) {
	for _, c := range cases{
		got := IsCrossArr(c.s1, c.s2, c.s3)
		if got != c.want {
			t.Error(c, got)
		}
	}
}