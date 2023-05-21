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
	n, t := NextInt(), NextInt()
	c := make([]int, n)
	r := make([]int, n)
	for i := range c {
		c[i] = NextInt()
	}
	for i := range r {
		r[i] = NextInt()
	}

	flag := false
	max := -1
	maxP1 := -1
	ans := 0
	for i := 0; i < n; i++ {
		if !flag {
			if c[i] == t {
				flag = true
				max = r[i]
				ans = i + 1
			} else if c[0] == c[i] {
				if chmax(&maxP1, r[i]) {
					ans = i + 1
				}
			}
		} else if flag && c[i] == t {
			if chmax(&max, r[i]) {
				ans = i + 1
			}
		}
	}

	fmt.Println(ans)
}

func chmax(a *int, b int) bool {
	if *a < b {
		*a = b
		return true
	}
	return false
}

func NextInt() int {
	sc.Scan()
	n, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return n
}
