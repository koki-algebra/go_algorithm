package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	count := 0
	for i := 0; i < n; i++ {
		var s string
		fmt.Scan(&s)
		if s == "For" {
			count++
		}
	}

	if n/2 < count {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
