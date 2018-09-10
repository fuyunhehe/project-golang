package dynamic

import "fmt"

func IsCrossChan(s1, s2, s3 string) bool {
	fmt.Println(s1, s2, s3)
	if len(s3) != len(s1)+len(s2) {
		return false
	}
	type node struct {
		i1, i2, i3 int
	}
	var ch = make(chan *node, (len(s1)+len(s2))*1000)
	defer close(ch)
	ch <- &node{-1, -1, -1}
	var cur *node
	l := 0
	for l >= 0 {
		cur = <-ch
		if cur.i1+1 < len(s1) && s1[cur.i1+1] == s3[cur.i3+1] {
			ch <- &node{cur.i1 + 1, cur.i2, cur.i3 + 1}
			l++
		}
		if cur.i2+1 < len(s2) && s2[cur.i2+1] == s3[cur.i3+1] {
			ch <- &node{cur.i1, cur.i2 + 1, cur.i3 + 1}
			l++
		}
		l--
	}
	if cur.i3 < len(s3)-1 {
		return false
	}
	return true
}

func IsCrossArr(s1, s2, s3 string) bool {
	var (
		m    [][]bool
		i, j = 0, 0
	)
	caculate := func(m [][]bool, i, j int) bool {
		if i == 0 && j == 0 {
			return true
		}
		if i >= 1 && m[i-1][j] && s1[i-1] == s3[i+j-1] {
			return true
		} else if j >= 1 && m[i][j-1] && s2[j-1] == s3[i+j-1] {
			return true
		}
		return false
	}

	for i = 0; i <= len(s1); i++ {
		m = append(m, make([]bool, len(s2)+1))
		for j = 0; j <= len(s2); j++ {
			m[i][j] = caculate(m, i, j)
		}
	}
	return m[len(s1)][len(s2)]
}
