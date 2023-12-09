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

	n, m, t := NextInt(sc), NextInt(sc), NextInt(sc)
	a := make([]int, n-1)
	for i := range a {
		a[i] = NextInt(sc)
	}

	bonus := make(map[int]int)
	for i := 0; i < m; i++ {
		x, y := NextInt(sc)-1, NextInt(sc)
		bonus[x] = y
	}

	for i := range a {
		t -= a[i]
		if t <= 0 {
			fmt.Fprintln(w, "No")
			return
		}
		if v, ok := bonus[i+1]; ok {
			t += v
		}
	}

	fmt.Fprintln(w, "Yes")
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
