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
	n := NextInt(sc)

	var evens, odds []int
	for i := 0; i < n; i++ {
		a := NextInt(sc)
		if a%2 == 0 {
			evens = append(evens, a)
		} else {
			odds = append(odds, a)
		}
	}

	sort.Ints(evens)
	sort.Ints(odds)

	oddMax := -1
	if len(odds) > 1 {
		oddMax = odds[len(odds)-1] + odds[len(odds)-2]
	}
	evenMax := -1
	if len(evens) > 1 {
		evenMax = evens[len(evens)-1] + evens[len(evens)-2]
	}
	if oddMax < evenMax {
		fmt.Println(evenMax)
	} else {
		fmt.Println(oddMax)
	}
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
