package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	INF = 1 << 60
)

var (
	h, w int
	c    [][]string
	dx   = []int{0, 1, 0, -1}
	dy   = []int{1, 0, -1, 0}
	seen [][]bool
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer([]byte{}, 5e6)

	fmt.Scan(&h, &w)
	c = make([][]string, h)
	seen = make([][]bool, h)
	for i := 0; i < h; i++ {
		seen[i] = make([]bool, w)
		line := NextLine(sc)
		c[i] = strings.Split(line, "")
	}

	ans := -1
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			Chmax(&ans, dfs(i, j, i, j))
		}
	}

	if ans <= 2 {
		ans = -1
	}

	fmt.Println(ans)
}

func dfs(sx, sy, px, py int) int {
	if sx == px && sy == py && seen[px][py] {
		return 0
	}

	seen[px][py] = true
	ret := -INF

	for i := 0; i < 4; i++ {
		nx := px + dx[i]
		ny := py + dy[i]
		if nx < 0 || ny < 0 || nx >= h || ny >= w || c[nx][ny] == "#" {
			continue
		}
		if (nx != sx || ny != sy) && seen[nx][ny] {
			continue
		}

		v := dfs(sx, sy, nx, ny)
		Chmax(&ret, v+1)
	}
	seen[px][py] = false

	return ret
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

func Chmax(a *int, b int) bool {
	if *a < b {
		*a = b
		return true
	}
	return false
}
