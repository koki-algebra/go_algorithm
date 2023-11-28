package main

import (
	"bufio"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer([]byte{}, 5e6)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
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
