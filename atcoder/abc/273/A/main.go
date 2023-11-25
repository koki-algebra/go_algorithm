package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	fmt.Println(f(n))
}

func f(x int) int {
	if x == 0 {
		return 1
	}
	return x * f(x-1)
}
