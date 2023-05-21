package main

import "fmt"

func main() {
	var (
		n int
		s string
	)
	fmt.Scanf("%d\n%s", &n, &s)

	ans := -1
	chmax(&ans, solve(s))
	chmax(&ans, solve(reverse(s)))
	fmt.Println(ans)
}

func solve(s string) int {
	ans := -1
	count := 0
	for i := range s {
		if s[i] == '-' {
			if count != 0 {
				chmax(&ans, count)
			}
			count = 0
		} else {
			count++
		}
	}
	return ans
}

func chmax(a *int, b int) {
	if *a < b {
		*a = b
	}
}

func reverse(s string) string {
	new := []rune(s)
	for i, j := 0, len(new)-1; i < j; i, j = i+1, j-1 {
		new[i], new[j] = new[j], new[i]
	}
	return string(new)
}
