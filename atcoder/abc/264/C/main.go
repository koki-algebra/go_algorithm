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

	h1, w1 := NextInt(sc), NextInt(sc)
	a := make([][]int, h1)
	for i := range a {
		a[i] = make([]int, w1)
		for j := range a[i] {
			a[i][j] = NextInt(sc)
		}
	}

	h2, w2 := NextInt(sc), NextInt(sc)
	b := make([][]int, h2)
	for i := range b {
		b[i] = make([]int, w2)
		for j := range b[i] {
			b[i][j] = NextInt(sc)
		}
	}

	for biti := 0; biti < 1<<h1; biti++ {
		for bitj := 0; bitj < 1<<w1; bitj++ {
			var hs, ws []int
			for k := 0; k < h1; k++ {
				if (biti >> k & 1) == 1 {
					hs = append(hs, k)
				}
			}
			for k := 0; k < w1; k++ {
				if (bitj >> k & 1) == 1 {
					ws = append(ws, k)
				}
			}
			if len(hs) != h2 || len(ws) != w2 {
				continue
			}

			match := true
			for k := 0; k < h2; k++ {
				for l := 0; l < w2; l++ {
					if a[hs[k]][ws[l]] != b[k][l] {
						match = false
						break
					}
				}
			}
			if match {
				fmt.Println("Yes")
				return
			}
		}
	}

	fmt.Println("No")
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
