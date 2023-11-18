package main

import "fmt"

func main() {
	var n, x int
	fmt.Scan(&n, &x)

	for i := 0; i < n; i++ {
		var p int
		fmt.Scan(&p)
		if p == x {
			fmt.Println(i + 1)
			return
		}
	}
}
