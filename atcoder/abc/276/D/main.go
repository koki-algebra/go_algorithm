package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}

	g := 0
	for i := 0; i < n; i++ {
		g = gcd(g, a[i])
	}

	count := 0
	for i := 0; i < n; i++ {
		a[i] /= g
		for a[i]%2 == 0 {
			a[i] /= 2
			count++
		}
		for a[i]%3 == 0 {
			a[i] /= 3
			count++
		}
		if a[i] != 1 {
			fmt.Println(-1)
			return
		}
	}

	fmt.Println(count)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
