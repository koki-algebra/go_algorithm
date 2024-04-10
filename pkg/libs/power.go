package libs

// Pow returns base^exponent (mod mod)
func Pow(base, exp, mod int) int {
	ret := 1
	x := base
	rem := exp

	for rem > 0 {
		if rem&1 == 1 {
			ret = ret * x % mod
		}
		x = x * x % mod
		rem >>= 1
	}

	return ret
}
