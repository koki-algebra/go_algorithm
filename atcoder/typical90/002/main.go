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
		cnt := 0
		for i := 0; i < n; i++ {
			if (bit >> i & 1) == 1 {
				s = "(" + s
				cnt--
			} else {
				s = ")" + s
				cnt++
			}

			if cnt < 0 {
				break
			}
		}

		if cnt == 0 {
			fmt.Fprintln(w, s)
		}
	}
}
