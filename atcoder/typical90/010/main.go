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

	n := NextInt(sc)
	p1, p2 := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		c, p := NextInt(sc), NextInt(sc)
		if c == 1 {
			p1[i] = p
		} else {
			p2[i] = p
		}
	}

	// 累積和のスライス
	s1, s2 := make([]int, n+1), make([]int, n+1)
	for i := 0; i < n; i++ {
		s1[i+1] = s1[i] + p1[i]
		s2[i+1] = s2[i] + p2[i]
	}

	q := NextInt(sc)
	for i := 0; i < q; i++ {
		l, r := NextInt(sc), NextInt(sc)
		fmt.Printf("%d %d\n", s1[r]-s1[l-1], s2[r]-s2[l-1])
	}
}

func NextInt(sc *bufio.Scanner) int {
	sc.Scan()
	n, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return n
}
