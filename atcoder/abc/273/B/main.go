package main

import (
	"fmt"
	"math"
)

func main() {
	var x, k int
	fmt.Scan(&x, &k)

	for i := 0; i < k; i++ {
		x = roundToPowerOfTen(x, i)
	}
	fmt.Println(x)
}

func roundToPowerOfTen(x int, i int) int {
	divisor := math.Pow10(i + 1)
	rounded := float64(x) / divisor
	rounded = math.Round(rounded)
	result := rounded * divisor
	return int(result)
}
