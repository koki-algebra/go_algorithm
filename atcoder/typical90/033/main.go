package main

import "fmt"

func main() {
	var h, w int
	fmt.Scan(&h, &w)

	if h == 1 || w == 1 {
		fmt.Println(h * w)
	} else {
		fmt.Println(((h + 1) / 2) * ((w + 1) / 2))
	}
}
