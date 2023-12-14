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
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	h, w := NextInt(sc), NextInt(sc)
	a := make([][]int, h)
	rowSum := make([]int, h)
	colSum := make([]int, w)

	for i := range a {
		a[i] = make([]int, w)
		for j := range a[i] {
			a[i][j] = NextInt(sc)
			rowSum[i] += a[i][j]
		}
	}

	for j := 0; j < w; j++ {
		for i := 0; i < h; i++ {
			colSum[j] += a[i][j]
		}
	}

	for i := range a {
		for j := range a[i] {
			fmt.Fprintf(writer, "%d ", rowSum[i]+colSum[j]-a[i][j])
		}
		fmt.Fprintln(writer)
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
