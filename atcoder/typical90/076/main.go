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

	n := NextInt(sc)
	b := make([]int, 2*n)
	for i := 0; i < n; i++ {
		a := NextInt(sc)
		b[i] = a
		b[i+n] = a
	}

	// prefix sum
	sums := make([]int, 2*n+1)
	for i := 0; i < 2*n; i++ {
		sums[i+1] = sums[i] + b[i]
	}

	if sums[n]%10 != 0 {
		fmt.Println("No")
		return
	}

	for i := 0; i < n; i++ {
		target := sums[i] + sums[n]/10

		// binary search
		left, right := -1, 2*n+1
		for right-left > 1 {
			mid := (left + right) / 2
			if sums[mid] == target {
				fmt.Println("Yes")
				return
			} else if sums[mid] < target {
				left = mid
			} else {
				right = mid
			}
		}
	}

	fmt.Println("No")
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
