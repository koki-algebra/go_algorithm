package main

import (
	"fmt"
	"strconv"
)

func main() {
	n := 8

	for bit := 0; bit < 1<<n; bit++ {
		fmt.Printf("%08s\n", strconv.FormatInt(int64(bit), 2))

		var selected []int
		for i := 0; i < n; i++ {
			if (bit >> i & 1) == 1 {
				// 左から i 番目にビットが立っている (i = 0,...,n-1)
				selected = append(selected, i)
			}
		}
		fmt.Println(selected)
	}
}
