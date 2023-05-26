package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	MOD = 46
)

var (
	sc = bufio.NewScanner(os.Stdin)
)

func main() {
	sc.Split(bufio.ScanWords)
	n := NextInt()
	a := make(map[int]int)
	b := make(map[int]int)
	c := make(map[int]int)

	for i := 0; i < n; i++ {
		a[NextInt()%MOD]++
	}
	for i := 0; i < n; i++ {
		b[NextInt()%MOD]++
	}
	for i := 0; i < n; i++ {
		c[NextInt()%MOD]++
	}

	ans := 0
	for keyA, valA := range a {
		for keyB, valB := range b {
			for keyC, valC := range c {
				if (keyA+keyB+keyC)%MOD == 0 {
					ans += valA * valB * valC
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
