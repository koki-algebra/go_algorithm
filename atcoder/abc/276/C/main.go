package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)
	a := make(sort.IntSlice, n)
	for i := range a {
		fmt.Scan(&a[i])
	}

	PrevPermutation(a)
	for i := range a {
		if i == len(a)-1 {
			fmt.Println(a[i])
		} else {
			fmt.Printf("%d ", a[i])
		}
	}
}

func PrevPermutation(x sort.Interface) bool {
	n := x.Len() - 1
	if n < 1 {
		return false
	}

	j := n - 1
	for ; x.Less(j, j+1); j-- {
		if j == 0 {
			return false
		}
	}

	l := n
	for x.Less(j, l) {
		l--
	}
	x.Swap(j, l)
	for k, l := j+1, n; k < l; k, l = k+1, l-1 {
		x.Swap(k, l)
	}
	return true
}
