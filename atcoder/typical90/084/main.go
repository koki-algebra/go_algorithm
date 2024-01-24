package main

import "fmt"

func main() {
	var (
		n int
		s string
	)
	fmt.Scan(&n, &s)

	d := []int{1}

	prev := s[0]
	k := 0
	for i := 1; i < n; i++ {
		if s[i] != prev {
			d = append(d, 1)
			k++
			prev = s[i]
		} else {
			d[k]++
		}
	}

	ret := 0
	for i := range d {
		ret += d[i] * (d[i] + 1) / 2
	}

	fmt.Println(n*(n+1)/2 - ret)
}
