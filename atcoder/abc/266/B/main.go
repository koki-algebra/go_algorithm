package main

import (
	"fmt"
)

const MOD = 998244353

func main() {
	var n int
	fmt.Scan(&n)

	n %= MOD
	if n < 0 {
		n += MOD
	}

	fmt.Println(n)
}
