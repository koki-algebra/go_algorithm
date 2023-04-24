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

const (
	INF = 1 << 60
)

func NextInt() int {
	sc.Scan()
	n, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return n
}

func main() {
	sc.Split(bufio.ScanWords)
	n := NextInt()

	for bit := 0; bit < 1<<n; bit++ {
		var s string
		for i := n - 1; i >= 0; i-- {
			if (bit >> i & 1) == 1 {
				s += ")"
			} else {
				s += "("
			}
		}
		if isValid(s) {
			fmt.Println(s)
		}
	}
}

func isValid(s string) bool {
	score := 0
	for i := range s {
		switch string(s[i]) {
		case "(":
			score++
		case ")":
			score--
		}
		if score < 0 {
			return false
		}
	}

	return score == 0
}
