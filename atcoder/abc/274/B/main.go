package main

import (
	"fmt"
	"strings"
)

func main() {
	var h, w int
	fmt.Scan(&h, &w)

	c := make([][]string, h)
	for i := range c {
		c[i] = make([]string, w)
		var s string
		fmt.Scan(&s)
		c[i] = strings.Split(s, "")
	}

	for j := 0; j < w; j++ {
		count := 0
		for i := 0; i < h; i++ {
			if c[i][j] == "#" {
				count++
			}
		}
		if j != w-1 {
			fmt.Printf("%d ", count)
		} else {
			fmt.Println(count)
		}
	}
}
