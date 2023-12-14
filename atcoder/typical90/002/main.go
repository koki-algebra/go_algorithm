package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n int
	fmt.Scan(&n)

	for bit := (1 << n) - 1; bit > 0; bit-- {
		s := ""
		for i := 0; i < n; i++ {
			if (bit >> i & 1) == 1 {
				s = "(" + s
			} else {
				s = ")" + s
			}
		}

		if isValid(s) {
			fmt.Fprintln(w, s)
		}
	}
}

func isValid(s string) bool {
	siz := 0
	for i := range s {
		if s[i] == '(' {
			siz++
		} else if s[i] == ')' {
			siz--
		}
		if siz < 0 {
			return false
		}
	}

	return siz == 0
}
