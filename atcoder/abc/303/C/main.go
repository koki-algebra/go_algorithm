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

func main() {
	sc.Split(bufio.ScanWords)
	buf := make([]byte, 4096)
	sc.Buffer(buf, 2e6)

	n, m, h, k := NextInt(), NextInt(), NextInt(), NextInt()
	s := NextLine()
	items := make(map[Position]struct{})
	for i := 0; i < m; i++ {
		x, y := NextInt(), NextInt()
		items[Position{x: x, y: y}] = struct{}{}
	}

	current := Position{x: 0, y: 0}
	for i := 0; i < n; i++ {
		h--
		if h < 0 {
			fmt.Println("No")
			return
		}

		switch s[i] {
		case 'R':
			current.x++
		case 'L':
			current.x--
		case 'U':
			current.y++
		case 'D':
			current.y--
		}

		if _, ok := items[current]; ok {
			if h < k {
				h = k
				delete(items, current)
			}
		}
	}

	fmt.Println("Yes")
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
