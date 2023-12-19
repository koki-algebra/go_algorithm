package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	n := NextInt(sc)
	a := make([]int, n)
	for i := range a {
		a[i] = NextInt(sc)
	}
	b := make([]int, n)
	for i := range b {
		b[i] = NextInt(sc)
	}

	sort.Ints(a)
	sort.Ints(b)

	ans := 0
	for i := range a {
		ans += distance(a[i], b[i])
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

func distance(u, v int) int {
	if u > v {
		return u - v
	}
	return v - u
}
