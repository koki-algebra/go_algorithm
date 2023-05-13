package main

import "fmt"

func main() {
	var (
		n int
		s string
	)
	fmt.Scan(&n)
	fmt.Scan(&s)

	var t, a int
	for i := range s {
		switch string(s[i]) {
		case "T":
			t++
		case "A":
			a++
		}
		if i == len(s)-1 {
			if t > a {
				fmt.Println("T")
			} else if t < a {
				fmt.Println("A")
			} else {
				if string(s[i]) == "T" {
					fmt.Println("A")
				} else {
					fmt.Println("T")
				}
			}
		}
	}
}
