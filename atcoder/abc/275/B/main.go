package main

import (
	"fmt"
)

const MOD = 998244353

func main() {
	var a, b, c, d, e, f int
	fmt.Scan(&a, &b, &c, &d, &e, &f)

	v1 := ((((a % MOD) * (b % MOD)) % MOD) * (c % MOD)) % MOD
	v2 := ((((d % MOD) * (e % MOD)) % MOD) * (f % MOD)) % MOD
	fmt.Println((v1 + MOD - v2) % MOD)
}
