package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64
	fmt.Scan(&a, &b)
	fmt.Printf("%.3f\n", roundToDecimalPlaces(b/a, 4))
}

// 小数点第 n 位で四捨五入
func roundToDecimalPlaces(v float64, n int) float64 {
	scale := math.Pow10(n - 1)
	return math.Round(v*scale) / scale
}
