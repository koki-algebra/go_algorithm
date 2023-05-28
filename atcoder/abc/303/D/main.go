package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	INF = 1 << 60
)

var (
	sc = bufio.NewScanner(os.Stdin)
)

func main() {
	sc.Split(bufio.ScanWords)
	buf := make([]byte, 4096)
	sc.Buffer(buf, 3e6)
	x, y, z := NextInt(), NextInt(), NextInt()
	s := NextLine()

	n := len(s)

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, 2)
		for j := range dp[i] {
			dp[i][j] = INF
		}
	}
	dp[0][0] = 0

	for i := 0; i < n; i++ {
		var cur int
		if s[i] == 'A' {
			cur = 1
		}

		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				v := dp[i][j]
				if j != k {
					v += z
				}
				if cur == k {
					v += x
				} else {
					v += y
				}
				chmin(&dp[i+1][k], v)
			}
		}
	}

	fmt.Println(min(dp[n][0], dp[n][1]))
}

func chmin(a *int, b int) {
	if *a > b {
		*a = b
	}
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func NextInt() int {
	sc.Scan()
	n, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return n
}

func NextLine() string {
	sc.Scan()
	return sc.Text()
}
