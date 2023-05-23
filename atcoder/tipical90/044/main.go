package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	sc = bufio.NewScanner(os.Stdin)
)

func main() {
	sc.Split(bufio.ScanWords)
	n, q := NextInt(), NextInt()
	a := make([]int, n)
	for i := range a {
		a[i] = NextInt()
	}

	shifts := 0
	for i := 0; i < q; i++ {
		t, x, y := NextInt(), NextInt()-1, NextInt()-1
		switch t {
		case 1:
			_x := (x + shifts) % n
			_y := (y + shifts) % n
			a[_x], a[_y] = a[_y], a[_x]
		case 2:
			shifts = (shifts + n - 1) % n
		case 3:
			fmt.Println(a[(x+shifts)%n])
		}
	}
}

func NextInt() int {
	sc.Scan()
	n, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return n
}
