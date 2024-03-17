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

	s := NextLine(sc)
	t := NextLine(sc)

	// S と T の先頭 i 文字が一致するかどうか
	pre := make([]bool, len(s)+1)
	pre[0] = true
	// S と T の末尾 i　文字が一致するかどうか
	suf := make([]bool, len(s)+1)
	suf[0] = true

	for i := 1; i <= len(t); i++ {
		pre[i] = isMatch(s[i-1], t[i-1]) && pre[i-1]
		suf[i] = isMatch(s[len(s)-i], t[len(t)-i]) && suf[i-1]
	}

	for x := 0; x <= len(t); x++ {
		if pre[x] && suf[len(t)-x] {
			fmt.Fprintln(w, "Yes")
		} else {
			fmt.Fprintln(w, "No")
		}
	}
}

func isMatch(c1, c2 byte) bool {
	return c1 == c2 || c1 == '?' || c2 == '?'
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
