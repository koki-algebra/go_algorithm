package main

import "fmt"

func main() {
	var x, y, n int
	fmt.Scan(&x, &y, &n)

	if y/3 > x {
		fmt.Println(x * n)
	} else {
		fmt.Println((n/3)*y + (n%3)*x)
	}
}
