package main

import (
	//_ "runtime/pprof"
	"github.com/fuyunhehe/project-golang/test/tmp/minit"
)

//func myHandler(w http.ResponseWriter, r *http.Request) {
//	w.Write([]byte("Hello World!"))
//}
//
//func httpRun() {
//	http.HandleFunc("/", myHandler)
//	http.ListenAndServe(":8080", nil)
//}
//
//func timeParse(t string) {
//	if tt, err := time.Parse(time.RFC3339, t); err != nil {
//		fmt.Println(err.Error())
//	} else {
//		fmt.Println(tt)
//	}
//}
//
//type AA struct {
//	Name string
//}
//
//type BB struct {
//	AList []AA
//}
//
//func sliceTest() {
//	bb := &BB{
//		AList: []AA{
//			{Name: "a1"},
//			{Name: "a2"},
//		},
//	}
//	fmt.Println(bb.AList)
//	for i, a := rangeT bb.AList {
//		a.Name = "b"
//		bb.AList[i] = a
//	}
//	fmt.Println(bb.AList)
//}
//
//func channelTest() {
//	c := make(chanT string, 3)
//	var i = 0
//	for {
//		select {
//		case s := <-c:
//			fmt.Println(s)
//			return
//		default:
//			if i < 3 {
//				fmt.Println(i)
//				i++
//			} else {
//				c <- "Hello Channel"
//			}
//		}
//	}
//}
//
//func slice2ptr() {
//	var o []byte
//	var ptr unsafe.Pointer
//	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&o))
//	sliceHeader.Len = 10
//	sliceHeader.Cap = 20
//	sliceHeader.Data = uintptr(ptr)
//
//}
//
//func round(n, a uintptr) uintptr {
//	return (n + a - 1) &^ (a - 1)
//}
//
//func f(result int) int {
//	defer func() {
//		result += 1
//		fmt.Println("defer", result)
//	}()
//	fmt.Println("f", result)
//	return result
//}

func main() {
	//var a, b = 2, 3
	//sliceTest()
	//fmt.Println(a == b)
	//pwd := sha256.Sum256([]byte("thiefa"))
	//a := pwd[:]
	//fmt.Println(hex.EncodeToString(a))
	//fmt.Println(round((1<<14)+1<<10, 1<<13), 1<<14, 1<<13)
	//fmt.Println("main", f(5))
	//a := 4 << (^uintptr(0) >> 63)
	//fmt.Println(a)

	minit.AA()
}
