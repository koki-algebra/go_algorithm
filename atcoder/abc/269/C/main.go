package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const (
	INF = 1 << 60
)

func main() {
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n int
	fmt.Scan(&n)

	var a []int
	for i := 0; i < 60; i++ {
		if (n >> i & 1) == 1 {
			a = append(a, i)
		}
	}
	k := len(a)

	var ans []int
	for i := 0; i < 1<<k; i++ {
		cur := 0
		for j := 0; j < k; j++ {
			if (i >> j & 1) == 1 {
				cur |= 1 << a[j]
			}
		}
		ans = append(ans, cur)
	}

	sort.Ints(ans)
	for i := range ans {
		fmt.Fprintln(w, ans[i])
	}
}
