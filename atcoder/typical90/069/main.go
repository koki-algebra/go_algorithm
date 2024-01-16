package main

import "fmt"

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
		fmt.Println(k * (k - 1) % MOD * bitPower(k-2, n-2) % MOD)
	}
}

func bitPower(a, b int) int {
	ret := 1
	for b != 0 {
		if b%2 == 1 {
			ret = ret * a % MOD
		}
		a = a * a % MOD
		b /= 2
	}

	return ret
}
