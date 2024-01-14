package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	h, w := NextInt(sc), NextInt(sc)
	p := make([][]int, h)
	for i := range p {
		p[i] = make([]int, w)
		for j := range p[i] {
			p[i][j] = NextInt(sc)
		}
	}

	ans := -1

	for bit := 0; bit < 1<<h; bit++ {
		var r []int
		for j := 0; j < w; j++ {
			var (
				num   = -1
				valid = true
			)
			for i := 0; i < h; i++ {
				if (bit >> i & 1) == 1 {
					if num == -1 {
						num = p[i][j]
					} else if num != p[i][j] {
						valid = false
					}
				}
			}
			if valid && num != -1 {
				r = append(r, num)
			}
		}

		cntH := bits.OnesCount(uint(bit))
		cntW := maximumSame(r)
		Chmax(&ans, cntH*cntW)
	}

	fmt.Println(ans)
}

func maximumSame(r []int) int {
	var (
		m   = make(map[int]int)
		ret = -1
	)

	for i := range r {
		m[r[i]]++
		Chmax(&ret, m[r[i]])
	}

	return ret
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

func Chmax(a *int, b int) bool {
	if *a < b {
		*a = b
		return true
	}
	return false
}
