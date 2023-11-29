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

	x, y, z := NextInt(sc), NextInt(sc), NextInt(sc)

	// 0, x, y
	if x > 0 && x < y {
		fmt.Fprintln(w, x)
		return
	}
	// y, 0, x
	if y < 0 && x > 0 {
		fmt.Fprintln(w, x)
		return
	}
	// y, x, 0
	if x < 0 && x > y {
		fmt.Fprintln(w, -x)
		return
	}
	// x, 0, y
	if x < 0 && y > 0 {
		fmt.Fprintln(w, -x)
		return
	}

	// 0, z, y, x
	if z > 0 && x > y && y > z {
		fmt.Fprintln(w, x)
		return
	}
	// x, y, z, 0
	if z < 0 && x < y && y < z {
		fmt.Fprintln(w, -x)
		return
	}

	// z, 0, y, x
	if y > 0 && x > y && z < 0 {
		fmt.Fprintln(w, x-2*z)
		return
	}
	// x, y, 0, z
	if y < 0 && x < y && z > 0 {
		fmt.Fprintln(w, 2*z-x)
		return
	}

	fmt.Fprintln(w, -1)
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
