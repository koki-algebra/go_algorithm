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
	n, m := NextInt(), NextInt()
	a := make([][]int, m)
	for i := range a {
		a[i] = make([]int, n)
		for j := range a[i] {
			a[i][j] = NextInt() - 1
		}
	}

	b := make([]map[int]struct{}, n)
	for i := 0; i < n; i++ {
		b[i] = make(map[int]struct{})
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if j >= 1 {
				b[a[i][j]][a[i][j-1]] = struct{}{}
			}
			if j <= n-2 {
				b[a[i][j]][a[i][j+1]] = struct{}{}
			}
		}
	}

	ans := 0
	for i := 0; i < n; i++ {
		ans += n - len(b[i]) - 1
	}
	ans /= 2
	fmt.Println(ans)
}

func NextInt() int {
	sc.Scan()
	n, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return n
}

func NextLine() string {
	sc.Scan()
	return sc.Text()
}
