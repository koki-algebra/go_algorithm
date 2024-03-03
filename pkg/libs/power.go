package libs

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
func countBinaryDigits(n int) int {
	if n == 0 {
		return 1
	}

	cnt := 0
	for n > 0 {
		n >>= 1
		cnt++
	}

	return cnt
}
