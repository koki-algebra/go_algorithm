package main

import "fmt"

func main() {
	var n, a, b int
	fmt.Scan(&n, &a, &b)
	for i := 0; i < n; i++ {
		var c int
		fmt.Scan(&c)
		if c == a+b {
			fmt.Println(i + 1)
		}
	}
}
