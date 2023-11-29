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

	a, b := NextInt(sc), NextInt(sc)
	resultA := isSolved(a)
	resultB := isSolved(b)

	ans := 0
	if resultA[0] || resultB[0] {
		ans += 1
	}
	if resultA[1] || resultB[1] {
		ans += 2
	}
	if resultA[2] || resultB[2] {
		ans += 4
	}

	fmt.Fprintln(w, ans)
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

func isSolved(n int) []bool {
	ans := make([]bool, 3)
	switch n {
	case 1:
		ans[0] = true
	case 2:
		ans[1] = true
	case 3:
		ans[0] = true
		ans[1] = true
	case 4:
		ans[2] = true
	case 5:
		ans[0] = true
		ans[2] = true
	case 6:
		ans[1] = true
		ans[2] = true
	case 7:
		ans[0] = true
		ans[1] = true
		ans[2] = true
	}

	return ans
}
