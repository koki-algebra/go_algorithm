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
	m := NextInt()
	b := make([]int, m)
	for i := range b {
		b[i] = NextInt()
	}
	x := NextInt()

	available := make([]bool, x+1)
	for i := range available {
		available[i] = true
	}
	for i := range b {
		available[b[i]] = false
	}

	dp := make([]bool, x+1)
	dp[0] = true
	for i := 0; i <= x; i++ {
		if dp[i] {
			for j := 0; j < n; j++ {
				if i+a[j] <= x {
					if available[i+a[j]] {
						dp[i+a[j]] = true
					}
				}
			}
		}
	}

	if dp[x] {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func NextInt() int {
	sc.Scan()
	n, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return n
}
