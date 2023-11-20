package main

import "fmt"

func main() {
	var n, max, argmax int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var h int
		fmt.Scan(&h)
		if chmax(&max, h) {
			argmax = i + 1
		}
	}

	fmt.Println(argmax)
}

func chmax(a *int, b int) bool {
	if *a < b {
		*a = b
		return true
	}

	return false
}
