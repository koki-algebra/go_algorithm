package main

import "fmt"

func main() {
	var r, c int
	fmt.Scan(&r, &c)

	if max(abc(8-r), abc(8-c))%2 == 0 {
		fmt.Println("white")
	} else {
		fmt.Println("black")
	}
}

func abc(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
