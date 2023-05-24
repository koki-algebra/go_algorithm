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
	n, t := NextInt(), NextInt()
	a := make([]int, n)
	for i := range a {
		a[i] = NextInt()
	}

	sum := 0
	for i := range a {
		sum += a[i]
	}

	r := t % sum
	for i := 0; ; i++ {
		if r-a[i] < 0 {
			fmt.Println(i+1, r)
			break
		}
		r -= a[i]
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
