package main

import "fmt"

func main() {
	var n, k int
	fmt.Scanf("%d %d", &n, &k)

	numPrimeFactor := make([]int, n+1)
	for i := 2; i <= n; i++ {
		if numPrimeFactor[i] != 0 {
			continue
		}
		for j := i; j <= n; j += i {
			numPrimeFactor[j]++
		}
	}

	ans := 0
	for i := range numPrimeFactor {
		if numPrimeFactor[i] >= k {
			ans++
		}
	}

	fmt.Println(ans)
}
