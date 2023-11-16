package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	count := 0
	for i := range s {
		c := string(s[i])
		if c == "v" {
			count++
		} else if c == "w" {
			count += 2
		}
	}

	fmt.Println(count)
}
