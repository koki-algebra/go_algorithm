package main

import "fmt"

func main() {
	var l, r int
	fmt.Scan(&l, &r)

	s := "atcoder"

	fmt.Println(s[l-1 : r])
}
