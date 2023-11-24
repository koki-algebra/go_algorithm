package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	n := getInt(sc)
	a := make([]int, n)
	m := make(map[int]int)
	for i := range a {
		a[i] = getInt(sc)
		m[a[i]]++
	}

	var b []int
	for key := range m {
		b = append(b, key)
	}
	sort.Slice(b, func(i, j int) bool { return b[i] > b[j] })

	for k := 0; k < n; k++ {
		if k < len(b) {
			fmt.Fprintln(w, m[b[k]])
		} else {
			fmt.Fprintln(w, 0)
		}
	}
}

func getInt(sc *bufio.Scanner) int {
	if sc.Scan() {
		n, err := strconv.Atoi(sc.Text())
		if err != nil {
			panic(err)
		}
		return n
	}

	if err := sc.Err(); err != nil {
		panic(err)
	}

	return -1
}
