package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	t := 1
	for i := 0; i < b; i++ {
		t *= c
	}
	if a < t {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
