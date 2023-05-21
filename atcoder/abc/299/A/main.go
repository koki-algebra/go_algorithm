package main

import (
	"fmt"
)

func main() {
	var (
		n int
		s string
	)
	fmt.Scanf("%d\n%s", &n, &s)
	count := 0
	flag := false
	for i := range s {
		switch string(s[i]) {
		case "|":
			count++
		case "*":
			if count == 1 {
				flag = true
			}
		}
	}

	if count == 2 && flag {
		fmt.Println("in")
	} else {
		fmt.Println("out")
	}
}
