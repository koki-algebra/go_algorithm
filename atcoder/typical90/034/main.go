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

	n, k := NextInt(sc), NextInt(sc)
	a := make([]int, n)
	for i := range a {
		a[i] = NextInt(sc)
	}

	var (
		ans   = -1
		right = 0
		m     = make(map[int]int)
		cnt   = 0
	)

	for left := 0; left < n; left++ {
		for right < n {
			if m[a[right]] == 0 {
				if cnt == k {
					break
				}
				cnt++
			}
			m[a[right]]++
			right++
		}

		Chmax(&ans, right-left)
		if m[a[left]] == 1 {
			cnt--
		}
		m[a[left]]--
	}

	fmt.Println(ans)
}

func NextInt(sc *bufio.Scanner) int {
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

func Chmax(a *int, b int) bool {
	if *a < b {
		*a = b
		return true
	}
	return false
}
