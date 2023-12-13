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

	n, l, k := NextInt(sc), NextInt(sc), NextInt(sc)
	a := make([]int, n)
	for i := range a {
		a[i] = NextInt(sc)
	}

	// すべての長さを x 以上にすることが可能か？
	check := func(x int) bool {
		num := 0 // 何個に切れたか
		pre := 0 // 前回の切れ目
		for i := 0; i < n; i++ {
			if a[i]-pre >= x {
				num++
				pre = a[i]
			}
		}

		if l-pre >= x {
			num++
		}

		return num >= k+1
	}

	// 二分探索
	left, right := -1, l+1
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
