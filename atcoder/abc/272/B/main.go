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
	x := make([][]int, m)
	for i := range x {
		k := NextInt(sc)
		x[i] = make([]int, k)
		for j := range x[i] {
			x[i][j] = NextInt(sc)
		}
	}

	seen := make([]map[int]struct{}, n)
	for i := 0; i < n; i++ {
		seen[i] = make(map[int]struct{})
	}

	for i := range x {
		for j := range x[i] {
			for k := range x[i] {
				if j != k {
					seen[x[i][j]-1][x[i][k]] = struct{}{}
				}
			}
		}
	}

	for i := range seen {
		if len(seen[i]) != n-1 {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
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
