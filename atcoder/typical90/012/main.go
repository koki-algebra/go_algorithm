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
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	h, w, q := NextInt(sc), NextInt(sc), NextInt(sc)

	uf := NewUnionFind(h * w)

	isRed := make([]bool, h*w)

	for i := 0; i < q; i++ {
		if t := NextInt(sc); t == 1 {
			var (
				r, c  = NextInt(sc), NextInt(sc)
				curr  = Index(w, r, c)
				up    = Index(w, r-1, c)
				down  = Index(w, r+1, c)
				left  = Index(w, r, c-1)
				right = Index(w, r, c+1)
			)

			isRed[curr] = true

			if r > 1 && isRed[up] {
				uf.Union(curr, up)
			}
			if r < h && isRed[down] {
				uf.Union(curr, down)
			}
			if c > 1 && isRed[left] {
				uf.Union(curr, left)
			}
			if c < w && isRed[right] {
				uf.Union(curr, right)
			}
		} else {
			ra, ca, rb, cb := NextInt(sc), NextInt(sc), NextInt(sc), NextInt(sc)
			a, b := Index(w, ra, ca), Index(w, rb, cb)

			if isRed[a] && isRed[b] && uf.Same(a, b) {
				fmt.Fprintln(writer, "Yes")
			} else {
				fmt.Fprintln(writer, "No")
			}
		}
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

func Index(w, r, c int) int {
	r--
	c--
	return r + c + (w-1)*r
}

type UnionFind struct {
	parents []int
	ranks   []int
}

func NewUnionFind(n int) *UnionFind {
	uf := new(UnionFind)
	uf.parents = make([]int, n)
	uf.ranks = make([]int, n)
	for i := 0; i < n; i++ {
		uf.parents[i] = -1
		uf.ranks[i] = 1
	}

	return uf
}

func (uf *UnionFind) Union(x, y int) {
	x = uf.Find(x)
	y = uf.Find(y)
	if x == y {
		return
	}
	if uf.ranks[x] > uf.ranks[y] {
		x, y = y, x
	}
	if uf.ranks[x] == uf.ranks[y] {
		uf.ranks[y]++
	}
	uf.parents[x] = y
}

func (uf *UnionFind) Find(x int) int {
	if uf.parents[x] < 0 {
		return x
	}
	uf.parents[x] = uf.Find(uf.parents[x])
	return uf.parents[x]
}

func (uf *UnionFind) Same(x, y int) bool {
	return uf.Find(x) == uf.Find(y)
}
