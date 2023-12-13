package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	fmt.Scanf("%d %d %d", &a, &b, &c)

	gcd := Gcd(c, Gcd(a, b))
	fmt.Println(getCutNum(a, gcd) + getCutNum(b, gcd) + getCutNum(c, gcd))
}

func getCutNum(x, y int) int {
	return x/y - 1
}

func Gcd(x, y int) int {
	if y == 0 {
		return x
	}
	return Gcd(y, x%y)
}
