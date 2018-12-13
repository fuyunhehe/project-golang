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

type categoriesResponse struct {
	Categories []category `json:"categories"`
}

func main() {
	aa, cc := &categoriesResponse{
		[]category{
			{
				Name: "aaa",
				Children: []category{
					{
						"bbb",
						nil,
					},
				},
			},
		},
	}, &categoriesResponse{}

	bb, _ := json.Marshal(aa)
	fmt.Println(string(bb))
	_ = json.Unmarshal(bb, cc)
	fmt.Printf("%++v\n", cc)

	d := struct {
		Uuids []string `json:"uuids"`
	}{}
	a, _ := json.Marshal(map[string]interface{}{"uuids": []string{"aaaaa"}})
	json.Unmarshal(a, &d)
	fmt.Println(d)
}
