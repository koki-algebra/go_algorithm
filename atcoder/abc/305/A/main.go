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

	ans := 5
	min := 1 << 60
	for i := 0; i <= 100; i += 5 {
		if chmin(&min, abs(n-i)) {
			ans = i
		}
	}

	fmt.Println(ans)
}

func chmin(a *int, b int) bool {
	if *a > b {
		*a = b
		return true
	}
	return false
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
