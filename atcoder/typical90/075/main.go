package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	if IsPrime(n) {
		fmt.Println(0)
		return
	}

	// prime factorize
	factors := PrimeFactorize(n)

	// the number of prime factors of n
	cnt := 0
	for _, exp := range factors {
		cnt += exp
	}

	for x := 1; ; x++ {
		if float64(x) >= math.Log2(float64(cnt)) {
			fmt.Println(x)
			return
		}
	}
}

func IsPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func PrimeFactorize(n int) map[int]int {
	ret := make(map[int]int)
	for p := 2; p*p <= n; p++ {
		if n%p != 0 {
			continue
		}

		e := 0
		for n%p == 0 {
			e++
			n /= p
		}

		ret[p] = e
	}

	if n != 1 {
		ret[n] = 1
	}

	return ret
}
