package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	fmt.Println(string(s[(len(s)+1)/2-1]))
}
