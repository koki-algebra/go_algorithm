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

const (
	MAX = 1000
)

func main() {
	sc.Split(bufio.ScanWords)
	n := NextInt()
	var area [MAX + 1][MAX + 1]int
	for i := 0; i < n; i++ {
		lx, ly, rx, ry := NextInt(), NextInt(), NextInt(), NextInt()
		area[lx][ly]++
		area[lx][ry]--
		area[rx][ry]++
		area[rx][ly]--
	}

	// x方向の累積和
	for x := 0; x <= MAX; x++ {
		for y := 1; y <= MAX; y++ {
			area[x][y] += area[x][y-1]
		}
	}
	// y方向の累積和
	for x := 1; x <= MAX; x++ {
		for y := 0; y <= MAX; y++ {
			area[x][y] += area[x-1][y]
		}
	}

	ans := make([]int, n+1)
	for x := 0; x <= MAX; x++ {
		for y := 0; y <= MAX; y++ {
			ans[area[x][y]]++
		}
	}
	for i := 1; i <= n; i++ {
		fmt.Println(ans[i])
	}
}

func NextInt() int {
	sc.Scan()
	n, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return n
}
