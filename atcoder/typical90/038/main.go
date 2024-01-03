package main

import "fmt"

const INF = 1e18

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	g := Gcd(a, b)

	if a/g > INF/b {
		fmt.Println("Large")
	} else {
		fmt.Println((a / g) * b)
	}
}

func Gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return Gcd(b, a%b)
}
