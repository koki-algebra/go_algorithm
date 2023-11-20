package main

import (
	"fmt"
	"strings"
)

func main() {
	var h, w int
	fmt.Scan(&h, &w)
	c := make([][]string, h)
	for i := range c {
		c[i] = make([]string, w)
		var s string
		fmt.Scan(&s)
		c[i] = strings.Split(s, "")
	}

	n := min(h, w)
	for i := 1; i <= n; i++ {
		fmt.Printf("%d ", solve(c, i))
	}
	fmt.Println()
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func solve(c [][]string, n int) (count int) {
	h := len(c)
	w := len(c[0])

	for i := 0; i < h; i++ {
		rowMax, rowMin := i+n, i-n
		if rowMax > h-1 || rowMin < 0 {
			continue
		}
		for j := 0; j < w; j++ {
			colMax, colMin := j+n, j-n
			if colMax > w-1 || colMin < 0 || c[i][j] != "#" {
				continue
			}
			if c[rowMax][colMax] == "#" && c[rowMax][colMin] == "#" && c[rowMin][colMax] == "#" && c[rowMin][colMin] == "#" {
				isValid := false

				// サイズ n かどうか
				if rowMax+1 < h {
					if colMax+1 < w {
						if c[rowMax+1][colMax+1] != "#" { // 右下の次が存在して # ではない
							isValid = true
						}
					}
					if colMin-1 >= 0 {
						if c[rowMax+1][colMin-1] != "#" { // 左下の次が存在して # ではない
							isValid = true
						}
					}
				} else { // 右上と左上の次が存在しない
					isValid = true
				}

				if rowMin-1 >= 0 {
					if colMax+1 < w {
						if c[rowMin-1][colMax+1] != "#" { // 右上の次が存在して # ではない
							isValid = true
						}
					}
					if colMin-1 >= 0 {
						if c[rowMin-1][colMin-1] != "#" { // 左上の次が存在して # ではない
							isValid = true
						}
					}
				} else { //右下と左下の次が存在しない
					isValid = true
				}

				if !isValid {
					continue
				}

				// バツ印をなしているか
				for k := 1; k < n; k++ {
					if c[i+k][j+k] != "#" || c[i+k][j-k] != "#" || c[i-k][j+k] != "#" || c[i-k][j-k] != "#" {
						isValid = false
						break
					}
				}
				if isValid {
					count++
				}
			}
		}
	}

	return
}
