package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	g := Gcd(Gcd(a, b), c)
	fmt.Println(a/g + b/g + c/g - 3)
}

func Gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return Gcd(b, a%b)
}
