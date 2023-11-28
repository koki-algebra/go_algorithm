package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)

	first := toHex(n % 16)
	second := toHex(n / 16)

	fmt.Println(second + first)
}

func toHex(n int) string {
	switch n {
	case 15:
		return "F"
	case 14:
		return "E"
	case 13:
		return "D"
	case 12:
		return "C"
	case 11:
		return "B"
	case 10:
		return "A"
	default:
		return strconv.Itoa(n)
	}
}
