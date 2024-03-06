package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	// the number of prime factors of n
	cnt := 0
	// remainder
	rem := n

	for p := 2; p*p <= n; p++ {
		for rem%p == 0 {
			rem /= p
			cnt++
		}
	}

	if rem != 1 {
		cnt++
	}

	for x := 0; ; x++ {
		if float64(x) >= math.Log2(float64(cnt)) {
			fmt.Println(x)
			return
		}
	}
}
