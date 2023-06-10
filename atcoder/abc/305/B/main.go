package main

import (
	"fmt"
)

func main() {
	var p, q string
	fmt.Scanf("%s %s", &p, &q)

	arr := []int{0, 3, 1, 4, 1, 5, 9}
	sum := 0
	sums := make([]int, 7)
	for i := range arr {
		sum += arr[i]
		sums[i] = sum
	}

	fmt.Println(abs(sums[index(p)] - sums[index(q)]))
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func index(s string) int {
	switch s {
	case "A":
		return 0
	case "B":
		return 1
	case "C":
		return 2
	case "D":
		return 3
	case "E":
		return 4
	case "F":
		return 5
	case "G":
		return 6
	}

	return -1
}
