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

const (
	INF = 1 << 60
)

type Person struct {
	name string
	age  int
}

func main() {
	sc.Split(bufio.ScanWords)
	n := NextInt()
	people := make([]Person, n)
	for i := range people {
		name := NextLine()
		age := NextInt()
		people[i] = Person{name, age}
	}

	var argmin int
	min := INF
	for i := range people {
		if chmin(&min, people[i].age) {
			argmin = i
		}
	}

	for i := 0; i < n; i++ {
		fmt.Println(people[(argmin+i)%n].name)
	}
}

func chmin(a *int, b int) bool {
	if *a > b {
		*a = b
		return true
	}
	return false
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
