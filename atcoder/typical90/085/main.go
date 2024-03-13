package main

import (
	"fmt"
	"sort"
)

func main() {
	var k int
	fmt.Scan(&k)

	var divisors []int
	for i := 1; i*i <= k; i++ {
		if k%i != 0 {
			continue
		}
		divisors = append(divisors, i)
		if k/i != i {
			divisors = append(divisors, k/i)
		}
	}

	sort.Ints(divisors)

	ans := 0
	for i := range divisors {
		for j := i; j < len(divisors); j++ {
			if k/divisors[i] < divisors[j] {
				continue
			}
			if k%(divisors[i]*divisors[j]) != 0 {
				continue
			}
			if divisors[j] <= k/(divisors[i]*divisors[j]) {
				ans++
			}
		}
	}

	fmt.Println(ans)
}
