package point

import (
	"testing"
)

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
		}, Want: "init"},
	}
	for _, c := range cases {
		PassVal(c.In)
		if c.In.Name != c.Want {
			t.Error(c)
		}
	}
}

func TestPassVal2(t *testing.T) {
	cases := []struct {
		In A
	}{
		{
			In: A{
				Name: "init",
				B:    B{"B1"},
				B1:   &B{"B2"},
			},
		},
	}
	for _, c := range cases {
		v1, v2 := PassVal2(cases[0].In)
		if &c.In.B == v1 || c.In.B1 != v2 {
			t.Errorf("%p\t%p\t%p\t%p\n", &cases[0].In.B, cases[0].In.B1, v1, v2)
		}
	}
}

func TestPassPtr2(t *testing.T) {
	cases := []struct {
		In A
	}{
		{
			In: A{
				Name: "init",
				B:    B{"B1"},
				B1:   &B{"B2"},
			},
		},
	}
	for _, c := range cases {
		v1, v2 := PassPtr2(&c.In)
		if &c.In.B != v1 || c.In.B1 != v2 {
			t.Errorf("%p\t%p\t%p\t%p\n", &c.In.B, c.In.B1, v1, v2)
		}
	}
}

func TestPassPtr3(t *testing.T) {
	var (
		i interface{}
	)
	a := &A{
		Name: "AA",
	}
	i = a
	if b, ok := i.(A); !ok {
		t.Error("interface *A not A")
	} else {
		t.Log(b)
	}
}
