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

	n, k := NextInt(sc), NextInt(sc)
	a, b := make([]int, n), make([]int, n)
	points := make([]int, 2*n)

	for i := 0; i < n; i++ {
		a[i], b[i] = NextInt(sc), NextInt(sc)
		points[i] = b[i]
	}
	for i := 0; i < n; i++ {
		points[n+i] = a[i] - b[i]
	}

	sort.Slice(points, func(i, j int) bool { return points[i] > points[j] })

	ans := 0
	for i := 0; i < k; i++ {
		ans += points[i]
	}

	fmt.Println(ans)
}

func Chmax(a *int, b int) bool {
	if *a < b {
		*a = b
		return true
	}
	return false
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
