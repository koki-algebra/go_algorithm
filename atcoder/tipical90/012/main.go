package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/koki-algebra/go_algorithm/pkg/libs"
)

var (
	sc   = bufio.NewScanner(os.Stdin)
	h, w int
)

func main() {
	sc.Split(bufio.ScanWords)
	h, w = NextInt(), NextInt()
	m := make([][]bool, h)
	for i := range m {
		m[i] = make([]bool, w)
	}

	q := NextInt()
	uf := libs.NewUnionFind(h * w)
	for i := 0; i < q; i++ {
		t := NextInt()
		if t == 1 {
			r, c := NextInt()-1, NextInt()-1
			m[r][c] = true
			// Union
			if r-1 >= 0 && m[r-1][c] {
				uf.Union(Index(r, c), Index(r-1, c))
			}
			if r+1 <= h-1 && m[r+1][c] {
				uf.Union(Index(r, c), Index(r+1, c))
			}
			if c-1 >= 0 && m[r][c-1] {
				uf.Union(Index(r, c), Index(r, c-1))
			}
			if c+1 <= w-1 && m[r][c+1] {
				uf.Union(Index(r, c), Index(r, c+1))
			}
		} else {
			ra, ca := NextInt()-1, NextInt()-1
			rb, cb := NextInt()-1, NextInt()-1
			if m[ra][ca] && m[rb][cb] && uf.Same(Index(ra, ca), Index(rb, cb)) {
				fmt.Println("Yes")
			} else {
				fmt.Println("No")
			}
		}
	}
}

func Index(i, j int) int {
	return (i + j) + (w-1)*i
}

func NextInt() int {
	sc.Scan()
	n, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return n
}
