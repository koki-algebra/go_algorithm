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

	n := NextInt(sc)
	sum1 := make([]int, n+1)
	sum2 := make([]int, n+1)
	for i := 0; i < n; i++ {
		c, p := NextInt(sc), NextInt(sc)
		if c == 1 {
			sum1[i+1] = sum1[i] + p
			sum2[i+1] = sum2[i]
		} else {
			sum1[i+1] = sum1[i]
			sum2[i+1] = sum2[i] + p
		}
	}

	q := NextInt(sc)
	for i := 0; i < q; i++ {
		l, r := NextInt(sc), NextInt(sc)
		ans1 := sum1[r] - sum1[l-1]
		ans2 := sum2[r] - sum2[l-1]
		fmt.Fprintln(w, ans1, ans2)
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
