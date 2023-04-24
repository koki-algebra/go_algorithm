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
	INF = 1 << 60
)

func NextInt() int {
	sc.Scan()
	n, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return n
}

func main() {
	sc.Split(bufio.ScanWords)
	n, l, k := NextInt(), NextInt(), NextInt()
	a := make([]int, n)
	for i := range a {
		a[i] = NextInt()
	}

	// 切断されてできる K+1 個のようかんの長さを全て x 以上にすることは可能か
	check := func(x int) bool {
		num := 0
		pre := 0

		for i := range a {
			if a[i]-pre >= x {
				num++
				pre = a[i]
			}
		}

		if l-pre >= x {
			num++
		}

		return k+1 <= num
	}

	// x を二分探索
	left := -1
	right := l + 1
	for right-left > 1 {
		mid := (left + right) / 2
		if check(mid) {
			left = mid
		} else {
			right = mid
		}
	}

	fmt.Println(left)
}
