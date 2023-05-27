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
	s := NextLine()
	t := NextLine()
	for i := 0; i < n; i++ {
		if !isValid(string(s[i]), string(t[i])) {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}

func isValid(x, y string) bool {
	if x == y {
		return true
	}
	if x == "1" && y == "l" {
		return true
	}
	if x == "l" && y == "1" {
		return true
	}
	if x == "0" && y == "o" {
		return true
	}
	if x == "o" && y == "0" {
		return true
	}

	return false
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
