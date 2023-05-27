package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	sc = bufio.NewScanner(os.Stdin)
)

func main() {
	sc.Split(bufio.ScanWords)
	n, p, q := NextInt(), NextInt(), NextInt()
	a := make([]int, n)
	for i := range a {
		a[i] = NextInt()
	}

	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			for k := 0; k < j; k++ {
				for l := 0; l < k; l++ {
					for m := 0; m < l; m++ {
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

func NextInt() int {
	sc.Scan()
	n, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return n
}
