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
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	n := NextInt(sc)
	p := make([]int, n)
	for i := range p {
		p[i] = NextInt(sc)
	}

	cnt := make([]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < 3; j++ {
			cnt[(p[i]-i+j-1+n)%n]++
		}
	}

	ans := 0
	for i := 0; i < n; i++ {
		Chmax(&ans, cnt[i])
	}
	fmt.Fprintln(w, ans)
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

func Chmax(a *int, b int) bool {
	if *a < b {
		*a = b
		return true
	}
	return false
}
