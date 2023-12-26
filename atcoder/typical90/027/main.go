package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer([]byte{}, 5e6)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var (
		n     int
		users = make(map[string]struct{})
	)

	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		s := NextLine(sc)
		if _, ok := users[s]; !ok {
			users[s] = struct{}{}
			fmt.Fprintln(w, i+1)
		}
	}
}

func NextLine(sc *bufio.Scanner) string {
	if sc.Scan() {
		return sc.Text()
	}
	if err := sc.Err(); err != nil {
		panic(err)
	}

	return ""
}
