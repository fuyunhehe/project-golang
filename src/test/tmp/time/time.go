package main

import (
	stdTime "time"
	"fmt"
)

func main()  {
	ts := stdTime.Now().Format("2006-01-02 15:04:05")
	fmt.Println(ts)
}
