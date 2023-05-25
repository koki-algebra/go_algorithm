package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	sc = bufio.NewScanner(os.Stdin)
)

func main() {
	sc.Split(bufio.ScanWords)
	h, _ := NextInt(), NextInt()
	ans := 0
	for i := 0; i < h; i++ {
		s := NextLine()
		ans += strings.Count(s, "#")
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

func NextLine() string {
	sc.Scan()
	return sc.Text()
}
