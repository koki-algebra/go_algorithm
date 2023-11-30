package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer([]byte{}, 5e6)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	a, b, c, d := -1, -1, -1, -1

	// create row slice
	rows := make([]string, 10)
	for i := 0; i < 10; i++ {
		rows[i] = NextLine(sc)
	}

	// row
	if strings.Contains(rows[0], "#") {
		a = 1
	}
	if strings.Contains(rows[9], "#") {
		b = 10
	}
	for i := 1; i < 10; i++ {
		if a == -1 && !strings.Contains(rows[i-1], "#") && strings.Contains(rows[i], "#") {
			a = i + 1
		} else if b == -1 && strings.Contains(rows[i-1], "#") && !strings.Contains(rows[i], "#") {
			b = i
		}
	}

	// create column slice
	cols := make([]string, 10)
	for i := range rows {
		ss := strings.Split(rows[i], "")
		for j := range cols {
			cols[j] += ss[j]
		}
	}

	// column
	if strings.Contains(cols[0], "#") {
		c = 1
	}
	if strings.Contains(cols[9], "#") {
		d = 10
	}
	for j := 1; j < 10; j++ {
		if c == -1 && !strings.Contains(cols[j-1], "#") && strings.Contains(cols[j], "#") {
			c = j + 1
		} else if d == -1 && strings.Contains(cols[j-1], "#") && !strings.Contains(cols[j], "#") {
			d = j
		}
	}

	fmt.Fprintf(w, "%d %d\n%d %d\n", a, b, c, d)
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
