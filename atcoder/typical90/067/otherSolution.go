package main

import (
	"fmt"
	"strconv"
	"strings"
)

func OtherSolution() {
	var (
		n string
		k int
	)
	fmt.Scan(&n, &k)

	for i := 0; i < k; i++ {
		digit, err := strconv.ParseInt(n, 8, 0)
		if err != nil {
			panic(err)
		}
		n = strconv.FormatInt(digit, 9)

		n = strings.ReplaceAll(n, "8", "5")
	}

	fmt.Println(n)
}
