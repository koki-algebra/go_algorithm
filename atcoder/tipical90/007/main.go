package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

const (
	INF = 1 << 60
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	n := NextInt(sc)
	a := make([]int, n)
	for i := range a {
		a[i] = NextInt(sc)
	}
	q := NextInt(sc)
	b := make([]int, q)
	for i := range b {
		b[i] = NextInt(sc)
	}

	sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })
	for j := range b {
		i := sort.Search(n, func(i int) bool { return a[i] >= b[j] })
		if i > len(a)-1 {
			fmt.Println(dissatisfaction(a[i-1], b[j]))
		} else if i == 0 {
			fmt.Println(dissatisfaction(a[i], b[j]))
		} else {
			if dissatisfaction(a[i], b[j]) < dissatisfaction(a[i-1], b[j]) {
				fmt.Println(dissatisfaction(a[i], b[j]))
			} else {
				fmt.Println(dissatisfaction(a[i-1], b[j]))
			}
		}
	}
}

func dissatisfaction(target, current int) int {
	val := math.Abs(float64(target - current))
	return int(val)
}

func NextInt(sc *bufio.Scanner) int {
	sc.Scan()
	n, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return n
}

func NextLine(sc *bufio.Scanner) string {
	sc.Scan()
	return sc.Text()
}
