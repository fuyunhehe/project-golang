package main

import "fmt"

type AA struct {
	n int
	name string
}

type BB struct {
	name string
	A1 AA
	AList []AA
}

func main()  {
	bb1 := BB{
		name:"bb1",
		A1:    AA{
			n:    1,
			name: "A1",
		},
		AList: []AA{
			{
				n:    2,
				name: "AList1",
			},
		},
	}
	bb2 := bb1
	fmt.Println(bb1, bb2)
	bb1.name = "bb11"
	bb1.A1.name = "bb11A1"
	bb1.AList[0].name = "bb11Alist1"
	fmt.Println(bb1, bb2)
	bb2.AList = []AA{{n:3,name:"bb2Alist1"}}
	fmt.Println(bb1, bb2)
}
