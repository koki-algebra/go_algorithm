package main

import "fmt"

func main() {
	var s, t string
	fmt.Scanf("%s\n%s", &s, &t)

	for i := range t {
		if i != len(t)-1 {
			if s[i] != t[i] {
				fmt.Println(i + 1)
				return
			}
		}
	}
	fmt.Println(len(t))
}
