package point

import "testing"

func TestPassPtr(t *testing.T) {
	cases := []struct {
		In   A
		Want string
	}{
		{In: A{
			Name: "init",
		}, Want: "passPtr"},
	}
	for _, c := range cases {
		PassPtr(&c.In)
		if c.In.Name != c.Want {
			t.Error(c)
		}
	}
}

func TestPassVal(t *testing.T) {
	cases := []struct {
		In   A
		Want string
	}{
		{In: A{
			Name: "init",
		}, Want: "passPtr"},
	}
	for _, c := range cases {
		PassVal(c.In)
		if c.In.Name != c.Want {
			t.Error(c)
		}
	}
}
