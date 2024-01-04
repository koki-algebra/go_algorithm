package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
)

const INF = 1 << 60

type State struct {
	r, c, dir int
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer([]byte{}, 5e6)

	var h, w, rs, cs, rt, ct int
	fmt.Scan(&h, &w, &rs, &cs, &rt, &ct)
	rs--
	cs--
	rt--
	ct--
	s := make([]string, h)
	for i := range s {
		s[i] = NextLine(sc)
	}

	dist := make([][][]int, h)
	for r := 0; r < h; r++ {
		dist[r] = make([][]int, w)
		for c := 0; c < w; c++ {
			dist[r][c] = make([]int, 4)
			for dir := 0; dir < 4; dir++ {
				dist[r][c][dir] = INF
			}
		}
	}

	dr := []int{0, 1, 0, -1}
	dc := []int{1, 0, -1, 0}

	deque := list.New()
	for dir := 0; dir < 4; dir++ {
		dist[rs][cs][dir] = 0
		deque.PushBack(State{r: rs, c: cs, dir: dir})
	}

	for deque.Len() != 0 {
		v := deque.Remove(deque.Front()).(State)
		for dir := 0; dir < 4; dir++ {
			// Next row & column
			nr := v.r + dr[dir]
			nc := v.c + dc[dir]
			cost := dist[v.r][v.c][v.dir]
			if v.dir != dir {
				cost++
			}

			if nr >= 0 && nr < h && nc >= 0 && nc < w && s[nr][nc] == '.' && Chmin(&dist[nr][nc][dir], cost) {
				state := State{r: nr, c: nc, dir: dir}
				if v.dir != dir {
					deque.PushBack(state)
				} else {
					deque.PushFront(state)
				}
			}
		}
	}

	ans := INF
	for dir := 0; dir < 4; dir++ {
		Chmin(&ans, dist[rt][ct][dir])
	}
	fmt.Println(ans)
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

func Chmin(a *int, b int) bool {
	if *a > b {
		*a = b
		return true
	}
	return false
}
