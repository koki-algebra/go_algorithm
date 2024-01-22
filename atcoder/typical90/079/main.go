package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	h, w := NextInt(sc), NextInt(sc)
	a, b := make([][]int, h), make([][]int, h)
	for i := range a {
		a[i] = make([]int, w)
		for j := range a[i] {
			a[i][j] = NextInt(sc)
		}
	}
	for i := range b {
		b[i] = make([]int, w)
		for j := range b[i] {
			b[i][j] = NextInt(sc)
		}
	}

	cnt := 0
	for i := 0; i < h-1; i++ {
		for j := 0; j < w-1; j++ {
			d := b[i][j] - a[i][j]
			a[i][j] += d
			a[i][j+1] += d
			a[i+1][j] += d
			a[i+1][j+1] += d

			if d < 0 {
				cnt -= d
			} else {
				cnt += d
			}
		}
	}

	for i := 0; i < h; i++ {
		if a[i][w-1] != b[i][w-1] {
			fmt.Println("No")
			return
		}
	}
	for j := 0; j < w; j++ {
		if a[h-1][j] != b[h-1][j] {
			fmt.Println("No")
			return
		}
	}

	fmt.Println("Yes")
	fmt.Println(cnt)
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
