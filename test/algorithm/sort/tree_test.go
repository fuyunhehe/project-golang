package sort

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func TestBinaryTree_Insert(t *testing.T) {
	cases := MakeCases(5, 10)
	for _, c := range cases {

		var tree = &BinaryTree{
			length: 0,
			head:   nil,
		}
		for _, k := range c.In {
			tree.Insert(k)
		}
		//fmt.Println(c.In, tree.PreOrder())

		//fmt.Println(c)
		//tree.Insert(10)
		//tree.Insert(20)
		//tree.Insert(5)
		//fmt.Printf("%+v, %+v, %+v, %+v\n", tree, tree.head, tree.head.children[0], tree.head.children[1])
		//break
	}
}

func TestBinaryTree_PreOrder(t *testing.T) {
	cases := []struct {
		In   []int
		want []int
	}{
		{
			In:   []int{2, 52, 57, 10, 1, 14, 42, 25},
			want: []int{2, 1, 52, 10, 14, 42, 25, 57},
		},
		{
			In:   []int{3, 1},
			want: []int{3, 1},
		},
		{
			In:   []int{1},
			want: []int{1},
		},
		{
			In:   []int{2, 5},
			want: []int{2, 5},
		},
	}
	for _, c := range cases {
		var tree = &BinaryTree{
			length: 0,
			head:   nil,
		}
		for _, k := range c.In {
			tree.Insert(k)
		}
		out := tree.PreOrder(tree.head)
		if !compare(c.want, out) {
			t.Errorf("In:%+v, out:%+v, want:%+v", c.In, out, c.want)
		}
	}
}

func TestBinaryTree_InOrder(t *testing.T) {
	cases := MakeCases(10, 10)
	for _, c := range cases {
		var tree = &BinaryTree{
			length: 0,
			head:   nil,
		}
		for _, k := range c.In {
			tree.Insert(k)
		}
		want := make([]int, len(c.In))
		copy(want, c.In)
		sort.Ints(want)
		out := tree.InOrder(tree.head)
		if !compare(want, out) {
			t.Errorf("In:%+v, out:%+v, want:%+v", c.In, out, want)
		}
	}
}

func TestBinaryTree_PostOrder(t *testing.T) {
	cases := []struct {
		In   []int
		want []int
	}{
		{
			In:   []int{11, 4, 22, 14, 13, 53, 43, 34, 37, 61},
			want: []int{4, 13, 14, 37, 34, 43, 61, 53, 22, 11},
		},
		{
			In:   []int{2, 52, 57, 10, 1, 14, 42, 25},
			want: []int{1, 25, 42, 14, 10, 57, 52, 2},
		},
		{
			In:   []int{3, 1},
			want: []int{1, 3},
		},
		{
			In:   []int{1},
			want: []int{1},
		},
		{
			In:   []int{2, 5},
			want: []int{5, 2},
		},
	}
	for _, c := range cases {
		var tree = &BinaryTree{
			length: 0,
			head:   nil,
		}
		for _, k := range c.In {
			tree.Insert(k)
		}
		want := c.want
		out := tree.PostOrder(tree.head)
		if !compare(want, out) {
			t.Errorf("In:%+v, out:%+v, want:%+v", c.In, out, want)
		}
	}
}

//t.Logf("In:%+v, out:%+v, want:%+v", c.In, out, want)
