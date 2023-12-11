package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64
	fmt.Scan(&a, &b)

	// f(x) が極小値をとる x
	argmin := math.Pow(a/(2.0*b), 2.0/3.0) - 1.0

	l := max(argmin-5.0, 0)
	r := min(argmin+5, a/b)
	ans := a
	for i := l; i <= r; i++ {
		ans = min(ans, f(float64(i), a, b))
	}

	fmt.Println(ans)
}

func f(n, a, b float64) float64 {
	return b*n + a/math.Sqrt(n+1)
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}
