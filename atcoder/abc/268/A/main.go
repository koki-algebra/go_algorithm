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

	m := make(map[int]struct{})
	m[NextInt(sc)] = struct{}{}
	m[NextInt(sc)] = struct{}{}
	m[NextInt(sc)] = struct{}{}
	m[NextInt(sc)] = struct{}{}
	m[NextInt(sc)] = struct{}{}

	fmt.Fprintln(w, len(m))
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
