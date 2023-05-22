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

type Position struct {
	x, y int
}

func NewPosition(x, y int) Position {
	return Position{x: x, y: y}
}

func main() {
	sc.Split(bufio.ScanWords)
	n, q := NextInt(), NextInt()
	positions := make([]Position, n)

	var (
		xmin = 1 << 60
		xmax = -(1 << 60)
		ymin = 1 << 60
		ymax = -(1 << 60)
	)

	for i := range positions {
		_x, _y := NextInt(), NextInt()
		x := _x - _y
		y := _x + _y
		chmin(&xmin, x)
		chmax(&xmax, x)
		chmin(&ymin, y)
		chmax(&ymax, y)
		positions[i] = NewPosition(x, y)
	}

	for i := 0; i < q; i++ {
		query := NextInt() - 1
		pos := positions[query]
		ans := max(max(abs(pos.x-xmin), abs(pos.x-xmax)), max(abs(pos.y-ymin), abs(pos.y-ymax)))
		fmt.Println(ans)
	}
}

func abs(v int) int {
	if v < 0 {
		v = -v
	}
	return v
}

func chmin(a *int, b int) {
	if *a > b {
		*a = b
	}
}

func chmax(a *int, b int) {
	if *a < b {
		*a = b
	}
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func NextInt() int {
	sc.Scan()
	n, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return n
}
