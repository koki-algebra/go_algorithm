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
	n := NextInt()
	a := make([][]int, n)
	for i := range a {
		a[i] = make([]int, n)
		for j := range a[i] {
			a[i][j] = NextInt()
		}
	}
	m := NextInt()
	taboosMap := make(map[int][]int)
	for i := 0; i < m; i++ {
		x, y := NextInt()-1, NextInt()-1
		taboosMap[x] = append(taboosMap[x], y)
		taboosMap[y] = append(taboosMap[y], x)
	}

	order := make([]int, n)
	for i := range order {
		order[i] = i
	}

	min := 1 << 60
	if isValid(taboosMap, order) {
		sum := getSum(a, order)
		if sum < min {
			min = sum
		}
	}
	for i := 1; NextPermutation(sort.IntSlice(order)); i++ {
		if !isValid(taboosMap, order) {
			continue
		}
		sum := getSum(a, order)
		if sum < min {
			min = sum
		}
	}

	if min == 1<<60 {
		fmt.Println(-1)
	} else {
		fmt.Println(min)
	}
}

func getSum(a [][]int, order []int) (sum int) {
	for i := range order {
		sum += a[order[i]][i]
	}
	return sum
}

func isValid(taboosMap map[int][]int, order []int) bool {
	n := len(order)
	for i := range order {
		if i != n-1 {
			next := order[i+1]
			taboos, ok := taboosMap[order[i]]
			if ok {
				for _, taboo := range taboos {
					if next == taboo {
						return false
					}
				}
			}
		}
	}

	return true
}

func NextInt() int {
	sc.Scan()
	n, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return n
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
