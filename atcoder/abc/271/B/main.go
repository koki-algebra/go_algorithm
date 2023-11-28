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
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	n, q := NextInt(sc), NextInt(sc)
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		l := NextInt(sc)
		a[i] = make([]int, l)
		for j := 0; j < l; j++ {
			a[i][j] = NextInt(sc)
		}
	}

	for k := 0; k < q; k++ {
		s, t := NextInt(sc), NextInt(sc)
		fmt.Fprintln(w, a[s-1][t-1])
	}
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
