package main

import (
	"fmt"
	"math"
)

const MOD = 1000000007

func main() {
	var l, r uint
	fmt.Scan(&l, &r)

	power10 := make([]uint, 20)
	power10[0] = 1
	for i := 0; i < 19; i++ {
		power10[i+1] = power10[i] * 10
	}

	ans := 0
	for i := 1; i <= 19; i++ {
		vl := max(l, power10[i-1])
		vr := min(r, power10[i]-1)
		if vl > vr {
			continue
		}
		val := (f(int(vr), MOD) - f(int(vl)-1, MOD) + MOD) % MOD

		ans += 1 * i * int(val)
		ans %= MOD
	}

	fmt.Println(ans)
}

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

func countBinaryDigits(x int) int {
	if x == 0 {
		return 0
	}

	return int(math.Log2(float64(x))) + 1
}

func div(a, b, mod int) int {
	return a * Pow(b, mod-2, mod) % mod
}

func f(n, mod int) int {
	v1 := n % mod
	v2 := (n + 1) % mod
	v := v1 * v2 % mod
	return div(v, 2, mod)
}

func max(a, b uint) uint {
	if a < b {
		return b
	}
	return a
}

func min(a, b uint) uint {
	if a < b {
		return a
	}
	return b
}
