package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	sc = bufio.NewScanner(os.Stdin)
)

func main() {
	sc.Split(bufio.ScanWords)
	n := NextInt()
	a := make([]int, n)
	for i := range a {
		a[i] = NextInt()
	}

	for i := 0; i < n-1; i++ {
		diff := a[i] - a[i+1]
		fmt.Printf("%d ", a[i])
		if abs(diff) > 1 {
			if diff > 0 {
				for v := a[i] - 1; v > a[i+1]; v-- {
					fmt.Printf("%d ", v)
				}
			} else {
				for v := a[i] + 1; v < a[i+1]; v++ {
					fmt.Printf("%d ", v)
				}
			}
		}
	}
	fmt.Println(a[n-1])
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func NextInt() int {
	sc.Scan()
	n, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return n
}
