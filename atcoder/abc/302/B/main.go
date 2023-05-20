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

const (
	target = "snuke"
)

func main() {
	sc.Split(bufio.ScanWords)
	h, w := NextInt(), NextInt()
	s := make([]string, h)
	for i := 0; i < h; i++ {
		s[i] = NextLine()
	}

	dx := []int{-1, -1, -1, 0, 0, 1, 1, 1}
	dy := []int{-1, 0, 1, -1, 1, -1, 0, 1}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			for k := 0; k < 8; k++ {
				var str string
				for t := 0; t < 5; t++ {
					x := i + t*dx[k]
					y := j + t*dy[k]
					if x < 0 || x >= h || y < 0 || y >= w {
						break
					}
					str += string(s[x][y])
				}
				if str == target {
					for t := 0; t < 5; t++ {
						x := i + t*dx[k]
						y := j + t*dy[k]
						fmt.Printf("%d %d\n", x+1, y+1)
					}
					return
				}
			}
		}
	}
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
