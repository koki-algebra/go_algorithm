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
	n := NextInt()
	s := make([]int, n+1)
	for i := 1; i <= n; i++ {
		s[i] = NextInt()
	}

	for i := 0; i < n; i++ {
		fmt.Printf("%d ", s[i+1]-s[i])
	}
	fmt.Println()
}

func NextInt() int {
	sc.Scan()
	n, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return n
}

func NextLine() string {
	sc.Scan()
	return sc.Text()
}
