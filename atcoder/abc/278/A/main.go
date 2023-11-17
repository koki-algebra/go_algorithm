package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	a := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	for i := 0; i < n; i++ {
		if i < n-k {
			fmt.Printf("%d ", a[i+k])
		} else if i == n-1 {
			fmt.Printf("%d\n", 0)
		} else {
			fmt.Printf("%d ", 0)
		}
	}
}
