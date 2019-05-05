package mreader

import (
	"io/ioutil"
	"io"
	"bytes"
	"strings"
	"fmt"
)

func Mreader() (err error) {
	var (
		body1, body2 io.Reader
		str1, str2         []byte
	)
	str := strings.Repeat("a", 256) + strings.Repeat("b", 256) + strings.Repeat("c", 256) + strings.Repeat("d", 256)
	body1 = strings.NewReader(str)
	b, _ := ioutil.ReadAll(io.LimitReader(body1, 512))
	body2 = io.MultiReader(bytes.NewReader(b), body1)
	str1, _ = ioutil.ReadAll(body1)
	str2, _ = ioutil.ReadAll(body2)
	fmt.Println(len(str), len(b), len(str1), len(str2))
	return nil
}
