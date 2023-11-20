package main

import (
	"fmt"
	"strings"
)

func main() {
	var h, w int
	fmt.Scan(&h, &w)
	a, b := make([][]string, h), make([][]string, h)
	for i := 0; i < h; i++ {
		a[i] = make([]string, w)
		var s string
		fmt.Scan(&s)
		a[i] = strings.Split(s, "")
	}
	for i := 0; i < h; i++ {
		b[i] = make([]string, w)
		var s string
		fmt.Scan(&s)
		b[i] = strings.Split(s, "")
	}

	for i := 0; i < h; i++ {
		verticalShift(a)
		for j := 0; j < w; j++ {
			horizontalShift(b)
			if areEqual(a, b) {
				fmt.Println("Yes")
				return
			}
		}
	}
	fmt.Println("No")
}

func verticalShift(a [][]string) {
	h := len(a)
	w := len(a[0])

	for j := 0; j < w; j++ {
		first := a[0][j]
		for i := 0; i < h-1; i++ {
			a[i][j] = a[i+1][j]
		}
		a[h-1][j] = first
	}
}

func horizontalShift(a [][]string) {
	h := len(a)
	w := len(a[0])

	for i := 0; i < h; i++ {
		first := a[i][0]
		for j := 0; j < w-1; j++ {
			a[i][j] = a[i][j+1]
		}
		a[i][w-1] = first
	}
}

func areEqual(a, b [][]string) bool {
	for i := range a {
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}
