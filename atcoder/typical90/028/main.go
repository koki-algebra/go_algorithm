package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	MAX = 1000
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	n := NextInt(sc)
	a := make([][]int, MAX+1)
	for i := range a {
		a[i] = make([]int, MAX+1)
	}

	for i := 0; i < n; i++ {
		lx, ly, rx, ry := NextInt(sc), NextInt(sc), NextInt(sc), NextInt(sc)
		a[lx][ly]++
		a[lx][ry]--
		a[rx][ry]++
		a[rx][ly]--
	}

	// x 軸方向の累積和
	for y := 0; y <= MAX; y++ {
		for x := 0; x < MAX; x++ {
			a[x+1][y] += a[x][y]
		}
	}

	// y 軸方向の累積和
	for x := 0; x <= MAX; x++ {
		for y := 0; y < MAX; y++ {
			a[x][y+1] += a[x][y]
		}
	}

	ans := make([]int, n+1)
	for x := 0; x <= MAX; x++ {
		for y := 0; y <= MAX; y++ {
			ans[a[x][y]]++
		}
	}

	for i := 1; i <= n; i++ {
		fmt.Fprintln(w, ans[i])
	}
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
