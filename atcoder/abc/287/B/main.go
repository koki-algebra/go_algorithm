package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer([]byte{}, 5e6)

	var n, m int
	fmt.Scan(&n, &m)

	s := make([]string, n)
	for i := 0; i < n; i++ {
		s[i] = NextLine(sc)
	}

	t := make(map[string]struct{})
	for i := 0; i < m; i++ {
		line := NextLine(sc)
		t[line] = struct{}{}
	}

	count := 0
	for i := range s {
		if _, ok := t[s[i][3:]]; ok {
			count++
		}
	}
	fmt.Println(count)
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
