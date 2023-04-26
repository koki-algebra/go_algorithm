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

	sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })
	for j := 0; j < q; j++ {
		b := NextInt(sc)
		// binary search
		i := sort.Search(n, func(i int) bool { return a[i] >= b })

		if i == n { // does not exist
			fmt.Println(dissatisfaction(a[i-1], b))
		} else if i == 0 {
			fmt.Println(dissatisfaction(a[i], b))
		} else {
			fmt.Println(min(dissatisfaction(a[i], b), dissatisfaction(a[i-1], b)))
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// 不満度
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
