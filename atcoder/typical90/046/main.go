package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const MOD = 46

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	a := make([]int, MOD)
	b := make([]int, MOD)
	c := make([]int, MOD)

	n := NextInt(sc)
	for i := 0; i < n; i++ {
		v := NextInt(sc)
		a[v%MOD]++
	}
	for i := 0; i < n; i++ {
		v := NextInt(sc)
		b[v%MOD]++
	}
	for i := 0; i < n; i++ {
		v := NextInt(sc)
		c[v%MOD]++
	}

	ans := 0
	for x := 0; x < MOD; x++ {
		for y := 0; y < MOD; y++ {
			for z := 0; z < MOD; z++ {
				if (x+y+z)%MOD == 0 {
					ans += a[x] * b[y] * c[z]
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
