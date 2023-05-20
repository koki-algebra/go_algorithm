package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var (
	sc = bufio.NewScanner(os.Stdin)
)

func main() {
	sc.Split(bufio.ScanWords)
	n, _ := NextInt(), NextInt()
	s := make([]string, n)
	for i := range s {
		s[i] = NextLine()
	}

	sort.Strings(s)
	if isValid(s) {
		fmt.Println("Yes")
		return
	}
	for i := 1; NextPermutation(sort.StringSlice(s)); i++ {
		if isValid(s) {
			fmt.Println("Yes")
			return
		}
	}
	fmt.Println("No")
}

func isValid(s []string) bool {
	n := len(s)
	for i := 0; i < n-1; i++ {
		if !isOneLetterDifferent(s[i], s[i+1]) {
			return false
		}
	}
	return true
}

func isOneLetterDifferent(a, b string) bool {
	count := 0
	m := len(a)
	for i := 0; i < m; i++ {
		if a[i] != b[i] {
			count++
		}
	}
	return count == 1
}

func NextInt() int {
	sc.Scan()
	n, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return n
}

func NextLine() string {
	sc.Scan()
	return sc.Text()
}

func NextPermutation(x sort.Interface) bool {
	n := x.Len() - 1
	if n < 1 {
		return false
	}
	j := n - 1
	for ; !x.Less(j, j+1); j-- {
		if j == 0 {
			return false
		}
	}
	l := n
	for !x.Less(j, l) {
		l--
	}
	x.Swap(j, l)
	for k, l := j+1, n; k < l; k, l = k+1, l-1 {
		x.Swap(k, l)
	}
	return true
}
