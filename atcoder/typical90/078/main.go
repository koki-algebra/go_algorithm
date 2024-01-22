package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	n, m := NextInt(sc), NextInt(sc)

	counts := make([]int, n)
	for i := 0; i < m; i++ {
		a, b := NextInt(sc)-1, NextInt(sc)-1

		if a > b {
			counts[a]++
		} else {
			counts[b]++
		}
	}

	ans := 0
	for i := range counts {
		if counts[i] == 1 {
			ans++
		}
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
