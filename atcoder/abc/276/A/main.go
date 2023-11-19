package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	n := len(s)
	for i := n - 1; i >= 0; i-- {
		if s[i] == 'a' {
			fmt.Println(i + 1)
			return
		}
	}
	fmt.Println(-1)
}
