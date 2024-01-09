package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const MAX_SIZE = 100000

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	cl, cr := MAX_SIZE, MAX_SIZE
	a := make([]int, 2*MAX_SIZE)

	q := NextInt(sc)
	for i := 0; i < q; i++ {
		t, x := NextInt(sc), NextInt(sc)
		switch t {
		case 1:
			cl--
			a[cl] = x
		case 2:
			a[cr] = x
			cr++
		case 3:
			fmt.Fprintln(w, a[cl+x-1])
		}
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
