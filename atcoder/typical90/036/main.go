package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	INF = 1 << 60
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var (
		n, q = NextInt(sc), NextInt(sc)
		p    = make([]Point, n)
		xmin = INF
		xmax = -INF
		ymin = INF
		ymax = -INF
	)

	for i := 0; i < n; i++ {
		_x, _y := NextInt(sc), NextInt(sc)
		x := _x - _y
		y := _x + _y
		p[i] = NewPoint(x, y)

		Chmin(&xmin, x)
		Chmax(&xmax, x)
		Chmin(&ymin, y)
		Chmax(&ymax, y)
	}

	for i := 0; i < q; i++ {
		j := NextInt(sc) - 1
		x, y := p[j].x, p[j].y

		ans := -1
		Chmax(&ans, abs(x-xmin))
		Chmax(&ans, abs(x-xmax))
		Chmax(&ans, abs(y-ymin))
		Chmax(&ans, abs(y-ymax))

		fmt.Fprintln(w, ans)
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

func Chmax(a *int, b int) bool {
	if *a < b {
		*a = b
		return true
	}
	return false
}

func Chmin(a *int, b int) bool {
	if *a > b {
		*a = b
		return true
	}
	return false
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

type Point struct {
	x, y int
}

func NewPoint(x, y int) Point {
	return Point{x, y}
}
