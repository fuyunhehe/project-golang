package main

import (
	"fmt"
	"reflect"
)

func main() {
	var (
		arr = make(map[string]interface{}, 4)
	)
	arr["a"] = []string{"aaa", "bbb"}
	arr["b"] = int64(10)
	arr["c"] = 10
	arr["d"] = map[string]string{"aaa": "bbb"}
	for k, v := range arr {
		//switch v.(type) {
		//case types.Slice:
		//	fmt.Println("slice", k)
		//case []string:
		//	fmt.Println("[]string", k)
		//case []interface{}:
		//	fmt.Println("[]interface", k)
		//}
		fmt.Println(k, reflect.TypeOf(v), reflect.TypeOf(v).Kind().String(), reflect.TypeOf(v).Name())
	}
}
