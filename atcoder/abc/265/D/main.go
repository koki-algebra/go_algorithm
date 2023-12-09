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
	n, p, q, r := NextInt(sc), NextInt(sc), NextInt(sc), NextInt(sc)
	sums := make(map[int]bool)
	sums[0] = true
	sum := 0
	for i := 0; i < n; i++ {
		a := NextInt(sc)
		sum += a
		sums[sum] = true
	}

	for sx := range sums {
		if sums[sx+p] && sums[sx+p+q] && sums[sx+p+q+r] {
			fmt.Println("Yes")
			return
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
