package sort

import (
	"math/rand"
)

func Bubble(ori []int) {
	for i := len(ori); i > 0; i-- {
		for j := 0; j < i-1; j++ {
			if ori[j] > ori[j+1] {
				exchage(ori, j, j+1)
			}
		}
	}
	return
}

func Insert(ori []int) {
	//分配空间
	var sorted = make([]int, len(ori), len(ori))

	//第一层循环，挨个插入
	for i, v := range ori {
		//第二层循环，寻找插入位置，插入之前所有元素挨个后移
		for ; i > 0; i-- {
			if v < sorted[i-1] {
				sorted[i] = sorted[i-1]
			} else {
				break
			}
		}
		sorted[i] = v
	}
	copy(ori, sorted)
	return
}

func Select(ori []int) {
	for i := 0; i < len(ori); i++ {
		m := i
		for j := i; j < len(ori); j++ {
			if ori[j] < ori[m] {
				m = j
			}
		}
		exchage(ori, i, m)
	}
	return
}

func Quick(ori []int) {
	if len(ori) <= 1 {
		return
	}

	var l, h, m = 0, len(ori) - 1, rand.Intn(len(ori))
	for l <= h {
		if l < m {
			if ori[l] > ori[m] {
				exchage(ori, l, m)
				m = l
			}
			l++
		} else {
			if ori[h] < ori[m] {
				exchage(ori, h, m)
				m = h
			}
			h--
		}
	}

	if m > 0 {
		Quick(ori[0:m])
	}
	if m < len(ori)-1 {
		Quick(ori[m+1:])
	}
	return
}

func Merge(ori []int) {
	//终止条件
	if len(ori) <= 1 {
		return
	}
	//递归
	var m = len(ori) / 2
	Merge(ori[0:m])
	Merge(ori[m:])
	//归并
	var i, j, sorted = 0, m, make([]int, 0, len(ori))
	for i < m && j < len(ori) {
		if ori[i] < ori[j] {
			sorted = append(sorted, ori[i])
			i++
		} else {
			sorted = append(sorted, ori[j])
			j++
		}
	}
	if i < m {
		sorted = append(sorted, ori[i:m]...)
	} else {
		sorted = append(sorted, ori[j:]...)
	}
	copy(ori, sorted)
	return
}
