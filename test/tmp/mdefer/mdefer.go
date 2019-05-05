package mdefer

import "fmt"

func Defer1() (aa string) {
	aa = "bbbb"
	defer func() {
		fmt.Println(aa)
		aa = "eee"
	}()
	aa = "ccc"
	return "dddd"
}

func PanicDefer(m, n int) (code int, err error) {
	defer func() {
		code = n
		if e := recover(); e != nil {
			if ee, ok := e.(error); ok {
				err = ee
			}
		}
	}()
	code = m
	panic("panic")
	//code = n
	return
}
