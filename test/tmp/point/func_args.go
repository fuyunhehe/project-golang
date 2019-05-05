package point

type B struct {
	C string
}

type A struct {
	Name string
	B    B
	B1   *B
}

func PassVal(a A) {
	a.Name = "passVal"
}

func PassPtr(a *A) {
	a.Name = "passPtr"
}

func PassVal2(a A) (*B, *B) {
	return &a.B, a.B1
}

func PassPtr2(a *A) (*B, *B) {
	return &a.B, a.B1
}
