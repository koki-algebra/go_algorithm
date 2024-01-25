package main

import (
	"fmt"
	"sort"
)

func main() {
	var k int
	fmt.Scan(&k)

	var devisors []int
	for i := 1; i*i <= k; i++ {
		if k%i != 0 {
			continue
		}
		devisors = append(devisors, i)
		if k/i != i {
			devisors = append(devisors, k/i)
		}
	}

	sort.Ints(devisors)

	ans := 0
	for i := range devisors {
		for j := i; j < len(devisors); j++ {
			if k/devisors[i] < devisors[j] {
				continue
			}
			if k%(devisors[i]*devisors[j]) != 0 {
				continue
			}
			if devisors[j] <= k/(devisors[i]*devisors[j]) {
				ans++
			}
		}
	}

	fmt.Println(ans)
}
