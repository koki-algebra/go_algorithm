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
	a := make([]int, n)
	for i := range a {
		a[i] = NextInt(sc)
	}

	top := 0

	for i := 0; i < q; i++ {
		t := NextInt(sc)
		x, y := NextInt(sc)-1, NextInt(sc)-1
		x = (x + top) % n
		y = (y + top) % n

		switch t {
		case 1:
			a[x], a[y] = a[y], a[x]
		case 2:
			top = (top - 1 + n) % n
			fmt.Println("top:", top)
		case 3:
			fmt.Fprintln(w, a[x])
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
