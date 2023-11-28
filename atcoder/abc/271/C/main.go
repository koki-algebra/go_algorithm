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

	var (
		n   = NextInt(sc)
		a   = make([]int, n)
		vol = make([]bool, n+2) // i巻を持っているかどうか
	)

	sold := 0
	for i := 0; i < n; i++ {
		a[i] = NextInt(sc)
		if a[i] >= len(vol) {
			sold++
		} else if vol[a[i]] {
			sold++
		} else {
			vol[a[i]] = true
		}
	}

	var (
		l = 1     // 現在持っていない巻のうち最小のもの
		r = n + 1 // 現在持っている巻のうち最大のもの
	)
	for {
		for vol[l] {
			l++
		}
		for r != 0 && !vol[r] {
			r--
		}
		if sold >= 2 {
			sold -= 2
			vol[l] = true
		} else {
			if l >= r {
				break
			}
			vol[r] = false
			sold++
		}
	}

	fmt.Println(l - 1)
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
