package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	left, right := 1, n

	var s int
	for left <= right {
		mid := (left + right) / 2
		fmt.Printf("? %d\n", mid)
		fmt.Scan(&s)

		if s == 0 {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	fmt.Printf("! %d\n", right)
}
