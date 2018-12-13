package point

type A struct {
	Name string
}

func PassVal(a A)  {
	a.Name = "passVal"
}

func PassPtr(a *A)  {
	a.Name = "passPtr"
}