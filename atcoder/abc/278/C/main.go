package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n, q int
	fmt.Scan(&n, &q)

	sc := bufio.NewScanner(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	m := make(map[int]map[int]bool)

	for i := 0; i < q; i++ {
		if sc.Scan() {
			s := strings.Split(sc.Text(), " ")
			t, _ := strconv.Atoi(s[0])
			a, _ := strconv.Atoi(s[1])
			b, _ := strconv.Atoi(s[2])

			if _, ok := m[a]; !ok {
				m[a] = make(map[int]bool)
			}

			switch t {
			case 1:
				m[a][b] = true
			case 2:
				m[a][b] = false
			case 3:
				if m[a][b] && m[b][a] {
					fmt.Fprintln(w, "Yes")
				} else {
					fmt.Fprintln(w, "No")
				}
			}
		}
		if err := sc.Err(); err != nil {
			panic(err)
		}
	}

	if err := w.Flush(); err != nil {
		panic(err)
	}
}
