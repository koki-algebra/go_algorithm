package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	n := getInt(sc)

	// key: index, value: depth
	m := make(map[int]int)
	// the depth of root is 0
	m[1] = 0

	// current index
	index := 1
	for i := 0; i < n; i++ {
		a := getInt(sc)
		if depth, ok := m[a]; ok {
			m[index+1] = depth + 1
			m[index+2] = depth + 1
			index += 2
		}
	}

	depths := make([]int, 2*n+1)
	for index, depth := range m {
		depths[index-1] = depth
	}

	for i := range depths {
		fmt.Fprintln(w, depths[i])
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
