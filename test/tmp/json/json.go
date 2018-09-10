package main

import (
	"encoding/json"
	"fmt"
)

const (
	AA = `{"name":"bbb","children":[]}`
)

type category struct {
	Name     string     `json:"name"`
	Children []category `json:"children"`
}

type CategoriesResponse struct {
	Categories []category `json:"categories"`
}

func main()  {
	aa, cc := &CategoriesResponse{
		[]category{
			{
				Name:"aaa",
				Children:[]category{
					{
						"bbb",
						nil,
					},
				},
			},
		},
	}, &CategoriesResponse{}

	bb, _ := json.Marshal(aa)
	fmt.Println(string(bb))
	_ = json.Unmarshal(bb, cc)
	fmt.Println(cc)
}