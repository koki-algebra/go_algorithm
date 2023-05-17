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
	n, k := NextInt(), NextInt()
	a, b := make([]int, n), make([]int, n)
	for i := range a {
		a[i] = NextInt()
	}
	for i := range b {
		b[i] = NextInt()
	}

	diff := calcDiff(a, b)
	if diff > k {
		fmt.Println("No")
		return
	}

	if (k-diff)%2 == 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func calcDiff(a, b []int) (ans int) {
	if len(a) != len(b) {
		return -1
	}
	for i := range a {
		ans += abs(a[i] - b[i])
	}
	return
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
