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
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		a[i] = NextInt()
	}

	// 尺取り方
	var (
		ans   = 0
		cr    = 1
		count = 0
		m     = make(map[int]int)
	)
	for i := 1; i <= n; i++ {
		for cr <= n {
			if m[a[cr]] == 0 {
				if count == k {
					break
				}
				count++
			}
			m[a[cr]]++
			cr++
		}
		chmax(&ans, cr-i)
		if m[a[i]] == 1 {
			count--
		}
		m[a[i]]--
	}
	fmt.Println(ans)
}

func chmax(a *int, b int) {
	if *a < b {
		*a = b
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
