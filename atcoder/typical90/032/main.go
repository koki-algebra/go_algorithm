package main

import (
	"bufio"
	"fmt"
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
	a := make([][]int, n)
	for i := range a {
		a[i] = make([]int, n)
		for j := range a[i] {
			a[i][j] = NextInt(sc)
		}
	}

	m := NextInt(sc)

	// 選手 x と y が険悪ならば serious[x][y] == true
	serious := make([][]bool, n)
	for i := range serious {
		serious[i] = make([]bool, n)
	}

	for i := 0; i < m; i++ {
		x, y := NextInt(sc)-1, NextInt(sc)-1
		serious[x][y] = true
		serious[y][x] = true
	}

	// 第 i 区を選手 order[i] が担当する
	order := make(sort.IntSlice, n)
	for i := range order {
		order[i] = i
	}

	ans := INF
	for {
		if isValid(order, serious) {
			sum := 0
			for i := 0; i < n; i++ {
				sum += a[order[i]][i]
			}
			Chmin(&ans, sum)
		}
		if !NextPermutation(order) {
			break
		}
	}

	if ans != INF {
		fmt.Println(ans)
	} else {
		fmt.Println(-1)
	}
}

func isValid(x sort.IntSlice, serious [][]bool) bool {
	n := x.Len()
	for i := 1; i < n; i++ {
		if serious[x[i-1]][x[i]] || serious[x[i]][x[i-1]] {
			return false
		}
	}

	return true
}

func NextPermutation(x sort.Interface) bool {
	n := x.Len() - 1
	if n < 1 {
		return false
	}
	j := n - 1
	for ; !x.Less(j, j+1); j-- {
		if j == 0 {
			return false
		}
	}
	l := n
	for !x.Less(j, l) {
		l--
	}
	x.Swap(j, l)
	for k, l := j+1, n; k < l; k, l = k+1, l-1 {
		x.Swap(k, l)
	}
	return true
}

func Chmin(a *int, b int) bool {
	if *a > b {
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
