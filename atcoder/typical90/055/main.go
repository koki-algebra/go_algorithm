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

	n, p, q := NextInt(sc), NextInt(sc), NextInt(sc)
	a := make([]int, n)
	for i := range a {
		a[i] = NextInt(sc)
	}

	ans := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				for l := k + 1; l < n; l++ {
					for m := l + 1; m < n; m++ {
						if a[i]*a[j]%p*a[k]%p*a[l]%p*a[m]%p == q {
							ans++
						}
					}
				}
			}
		}
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
