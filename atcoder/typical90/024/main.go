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

	n, k := NextInt(sc), NextInt(sc)
	a := make([]int, n)
	for i := range a {
		a[i] = NextInt(sc)
	}
	sum := 0
	for i := 0; i < n; i++ {
		b := NextInt(sc)
		sum += abc(a[i] - b)
	}

	if diff := k - sum; diff >= 0 && diff%2 == 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func abc(v int) int {
	if v < 0 {
		return -v
	}
	return v
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
