package point

type aaa struct {
	n int
}

func (aa *aaa) Header() *aaa {
	return aa
}

//func main()  {
//	bb := &aaa{123}
//	cc := bb.Header()
//	fmt.Printf("%+v,%+v,%+v", bb, cc, bb==cc)
//}
