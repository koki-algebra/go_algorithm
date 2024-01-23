package libs

import "math"

// Pow returns base^exponent (mod mod)
func Pow(base, exponent, mod int) int {
	ret := 1
	digits := countBinaryDigits(exponent)

	for i := 0; i < digits; i++ {
		if exponent>>i&1 == 1 {
			ret = ret * base % mod
		}
		base = base * base % mod
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
