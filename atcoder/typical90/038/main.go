package main

import "fmt"

const (
	large = 1000000000000000000
)

func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)

	r := b / Gcd(a, b)
	if r > large/a {
		fmt.Println("Large")
	} else {
		fmt.Println(r * a)
	}
}

func Gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return Gcd(b, a%b)
}
