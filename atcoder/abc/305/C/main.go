package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	sc = bufio.NewScanner(os.Stdin)
)

func main() {
	sc.Split(bufio.ScanWords)
	h, w := NextInt(), NextInt()
	s := make([][]string, h)
	for i := 0; i < h; i++ {
		line := NextLine()
		s[i] = append(s[i], strings.Split(line, "")...)
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if s[i][j] == "." {
				if i+1 < h && i-1 >= 0 {
					if s[i+1][j] == "#" && s[i-1][j] == "#" {
						fmt.Println(i+1, j+1)
						return
					}
				}
				if j+1 < w && j-1 >= 0 {
					if s[i][j+1] == "#" && s[i][j-1] == "#" {
						fmt.Println(i+1, j+1)
						return
					}
				}

				if i+1 < h && j-1 >= 0 {
					if s[i+1][j] == "#" && s[i][j-1] == "#" {
						fmt.Println(i+1, j+1)
						return
					}
				}
				if i+1 < h && j+1 < w {
					if s[i+1][j] == "#" && s[i][j+1] == "#" {
						fmt.Println(i+1, j+1)
						return
					}
				}
				if j-1 >= 0 && i-1 >= 0 {
					if s[i][j-1] == "#" && s[i-1][j] == "#" {
						fmt.Println(i+1, j+1)
						return
					}
				}
				if j+1 < w && i-1 >= 0 {
					if s[i][j+1] == "#" && s[i-1][j] == "#" {
						fmt.Println(i+1, j+1)
						return
					}
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
