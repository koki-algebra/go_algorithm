package main

import "fmt"

func main() {
	var h, w int
	fmt.Scanf("%d %d", &h, &w)
	area := make([][]bool, h)
	for i := range area {
		area[i] = make([]bool, w)
	}

	var countH, countW int
	if h == 1 {
		fmt.Println(w)
		return
	}
	if w == 1 {
		fmt.Println(h)
		return
	}
	if h%2 == 0 {
		countH = h / 2
	} else {
		countH = h/2 + 1
	}
	if w%2 == 0 {
		countW = w / 2
	} else {
		countW = w/2 + 1
	}

	fmt.Println(countH * countW)
}
