package mstruct

import (
	"encoding/json"
	"fmt"
)

func Mashal() {
	type Aaa struct {
		Bbb struct {
			Ccc struct {
				Ddd string `json:"d"`
			} `json:"c"`
		} `json:"b"`
	}
	t := Aaa{
		Bbb: struct {
			Ccc struct {
				Ddd string `json:"d"`
			} `json:"c"`
		}{
			Ccc: struct {
				Ddd string `json:"d"`
			}{
				Ddd: "fff",
			},
		},
	}
	s, _ := json.Marshal(t)
	fmt.Println(string(s))
}
