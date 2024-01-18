package main

import (
	"fmt"
	"math"
)

const MOD = 1000000007

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	if k == 1 {
		if n == 1 {
			fmt.Println(1)
		} else {
			fmt.Println(0)
		}
	} else if n == 1 {
		fmt.Println(k % MOD)
	} else if n == 2 {
		fmt.Println(k * (k - 1) % MOD)
	} else {
		fmt.Println(k * (k - 1) % MOD * Pow(k-2, n-2, MOD) % MOD)
	}
}

// Pow returns a^b (mod m)
func Pow(a, b, m int) int {
	ret := 1
	digits := countBinaryDigits(b)

	for i := 0; i < digits; i++ {
		if b>>i&1 == 1 {
			ret = ret * a % m
		}
		a = a * a % m
	}

	return ret
}

// countBinaryDigits returns the number of digits when x is expressed in binary
func countBinaryDigits(x int) int {
	if x == 0 {
		return 0
	}

	return int(math.Log2(float64(x))) + 1
}
