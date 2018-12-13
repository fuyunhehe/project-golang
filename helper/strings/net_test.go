package strings

import (
	"testing"
)

func TestIp2Int(t *testing.T) {
	cases := []struct {
		in   string
		want uint32
	}{
		{"192.168.1.1", 3232235777},
		{"255.255.255.255", 4294967295},
		{"0.0.0.0", 0},
		{"00.0.0", 0},
	}
	for _, c := range cases {
		if out, err := Ip2Uint32(c.in); out != c.want {
			t.Error(c, out, err)
		}
	}
}
