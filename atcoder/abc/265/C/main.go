package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Point struct {
	r, c int
}

func NewPoint(r, c int) Point {
	return Point{r: r, c: c}
}

func (p *Point) Move(g [][]string) bool {
	h := len(g)
	w := len(g[0])

	switch g[p.r][p.c] {
	case "U":
		if p.r-1 < 0 {
			return false
		}
		p.r--
	case "D":
		if p.r+1 >= h {
			return false
		}
		p.r++
	case "L":
		if p.c-1 < 0 {
			return false
		}
		p.c--
	case "R":
		if p.c+1 >= w {
			return false
		}
		p.c++
	}

	return true
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer([]byte{}, 5e6)

	var h, w int
	fmt.Scan(&h, &w)
	g := make([][]string, h)
	for i := 0; i < h; i++ {
		line := NextLine(sc)
		g[i] = make([]string, w)
		g[i] = strings.Split(line, "")
	}

	seen := make(map[Point]struct{})
	current := NewPoint(0, 0)

	for {
		if _, ok := seen[current]; ok {
			fmt.Println(-1)
			return
		} else {
			seen[current] = struct{}{}
		}

		if !current.Move(g) {
			fmt.Println(current.r+1, current.c+1)
			return
		}
	}
}

func NextLine(sc *bufio.Scanner) string {
	if sc.Scan() {
		return sc.Text()
	}
	if err := sc.Err(); err != nil {
		panic(err)
	}

	return ""
}
