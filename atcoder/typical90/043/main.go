package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
)

var (
	sc = bufio.NewScanner(os.Stdin)
	dx = []int{1, 0, -1, 0}
	dy = []int{0, 1, 0, -1}
)

const (
	INF = 1 << 60
)

type State struct {
	x, y, dir int
}

func main() {
	sc.Split(bufio.ScanWords)
	h, w := NextInt(), NextInt()
	sx, sy := NextInt()-1, NextInt()-1
	gx, gy := NextInt()-1, NextInt()-1
	s := make([]string, h)
	for i := range s {
		s[i] = NextLine()
	}

	dist := make([][][]int, h)
	for i := 0; i < h; i++ {
		dist[i] = make([][]int, w)
		for j := 0; j < w; j++ {
			dist[i][j] = make([]int, 4)
			dist[i][j][0] = INF
			dist[i][j][1] = INF
			dist[i][j][2] = INF
			dist[i][j][3] = INF
		}
	}

	l := list.New()
	for i := 0; i < 4; i++ {
		dist[sx][sy][i] = 0
		l.PushBack(State{x: sx, y: sy, dir: i})
	}

	for l.Len() != 0 {
		u := l.Remove(l.Front()).(State)
		for i := 0; i < 4; i++ {
			tx := u.x + dx[i]
			ty := u.y + dy[i]
			cost := dist[u.x][u.y][u.dir]
			if u.dir != i {
				cost++
			}

			if tx >= 0 && tx < h && ty >= 0 && ty < w && s[tx][ty] == '.' && dist[tx][ty][i] > cost {
				dist[tx][ty][i] = cost
				if u.dir != i {
					l.PushBack(State{x: tx, y: ty, dir: i})
				} else {
					l.PushFront(State{x: tx, y: ty, dir: i})
				}
			}
		}
	}

	ans := INF
	for i := 0; i < 4; i++ {
		chmin(&ans, dist[gx][gy][i])
	}
	fmt.Println(ans)
}

func chmin(a *int, b int) {
	if *a > b {
		*a = b
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
