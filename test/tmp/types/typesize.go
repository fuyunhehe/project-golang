package types

import (
	"fmt"
	"unsafe"
)

type TT struct {
	A int
	B bool
	//C string
	//D []string
}

func SizeOfTypes() {
	fmt.Println(unsafe.Sizeof(int(1)), unsafe.Sizeof(string("EDDYfffffffCJY")), unsafe.Sizeof(int8(1)), unsafe.Sizeof(TT{}))
	//fmt.Printf("bool align: %d\n", unsafe.Alignof(bool(true)))
	//fmt.Printf("int32 align: %d\n", unsafe.Alignof(int32(0)))
	//fmt.Printf("int8 align: %d\n", unsafe.Alignof(int8(0)))
	//fmt.Printf("int64 align: %d\n", unsafe.Alignof(int64(0)))
	//fmt.Printf("byte align: %d\n", unsafe.Alignof(byte(0)))
	fmt.Printf("string align: %d\n", unsafe.Alignof("EDDYfffffffCJY"))
	fmt.Printf("string align: %d\n", unsafe.Alignof(string("EDDYfffffffCJY")))
	fmt.Printf("map align: %d\n", unsafe.Alignof(map[string]string{}))
	fmt.Printf("slice align: %d\n", unsafe.Alignof([]string{"aaa","bbb"}))
	fmt.Printf("TT align: %d\n", unsafe.Alignof(TT{}))
}
