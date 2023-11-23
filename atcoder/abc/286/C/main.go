package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		n, a, b int
		str     string
	)
	fmt.Scan(&n, &a, &b, &str)

	s := strings.Split(str, "")
	s = append(s, s...)

	ans := 1 << 60
	for i := 0; i < n; i++ {
		tmp := a * i
		for j := 0; j < n/2; j++ {
			l := i + j
			r := i + n - 1 - j
			if s[l] != s[r] {
				tmp += b
			}
		}
		chmin(&ans, tmp)
	}

	fmt.Println(ans)
}

func chmin(a *int, b int) bool {
	if *a > b {
		*a = b
		return true
	}

	return false
}
