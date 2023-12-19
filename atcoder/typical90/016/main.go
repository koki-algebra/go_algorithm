package main

import "fmt"

const (
	P = 9999
)

func main() {
	var n, a, b, c int
	fmt.Scan(&n, &a, &b, &c)

	ans := P + 1
	for x := 0; x <= P; x++ {
		for y := 0; y <= P; y++ {
			tmp := a*x + b*y
			if (n-tmp)%c == 0 && tmp <= n {
				z := (n - tmp) / c
				Chmin(&ans, x+y+z)
			}
		}
	}
	fmt.Println(ans)
}

func Chmin(a *int, b int) bool {
	if *a > b {
		*a = b
		return true
	}
	return false
}
