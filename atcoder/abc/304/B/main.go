package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	if n < 1000 {
		fmt.Println(n)
	} else if n >= 1000 && n < 10000 {
		fmt.Println(floor(n, 1))
	} else if n >= 10000 && n < 100000 {
		fmt.Println(floor(n, 2))
	} else if n >= 100000 && n < 1000000 {
		fmt.Println(floor(n, 3))
	} else if n >= 1000000 && n < 10000000 {
		fmt.Println(floor(n, 4))
	} else if n >= 10000000 && n < 100000000 {
		fmt.Println(floor(n, 5))
	} else if n >= 100000000 && n < 1000000000 {
		fmt.Println(floor(n, 6))
	}
}

func floor(n int, place int) int {
	p := pow10(place)
	n /= p
	n *= p
	return n
}

func pow10(n int) int {
	ans := 1
	for i := 0; i < n; i++ {
		ans *= 10
	}
	return ans
}
