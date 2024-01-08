package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const MOD = 1000000007

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	n := NextInt(sc)
	a := make([][]int, n)
	for i := range a {
		a[i] = make([]int, 6)
		for j := range a[i] {
			a[i][j] = NextInt(sc)
		}
	}

	sums := make([]int, n)
	for i := range a {
		for j := range a[i] {
			sums[i] += a[i][j]
		}
	}

	ans := 1
	for i := range sums {
		ans *= sums[i]
		ans %= MOD
	}
	fmt.Println(ans)
}

func NextInt(sc *bufio.Scanner) int {
	if sc.Scan() {
		n, err := strconv.Atoi(sc.Text())
		if err != nil {
			panic(err)
		}
		return n
	}

	if err := sc.Err(); err != nil {
		panic(err)
	}

	return -1
}
