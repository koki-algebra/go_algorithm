package main

import "fmt"

var (
	memo map[int]int
)

func main() {
	var n int
	fmt.Scan(&n)
	memo = make(map[int]int)

	fmt.Println(f(n))
}

func f(x int) int {
	if x == 0 {
		return 1
	}
	if v, ok := memo[x]; ok {
		return v
	}

	memo[x] = f(x/2) + f(x/3)
	return memo[x]
}
