package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	s := strings.Split(NextLine(sc), " ")
	h, _ := strconv.Atoi(s[0])
	w, _ := strconv.Atoi(s[1])
	a := make([][]int, h)
	rowSum := make([]int, h)
	colSum := make([]int, w)

	for i := 0; i < h; i++ {
		a[i] = make([]int, w)
		s := strings.Split(NextLine(sc), " ")
		for j := 0; j < w; j++ {
			a[i][j], _ = strconv.Atoi(s[j])
			rowSum[i] += a[i][j]
			colSum[j] += a[i][j]
		}
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			fmt.Fprintf(writer, "%d ", rowSum[i]+colSum[j]-a[i][j])
		}
		fmt.Fprint(writer, "\n")
	}
}

func NextInt(sc *bufio.Scanner) int {
	sc.Scan()
	n, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return n
}

func NextLine(sc *bufio.Scanner) string {
	sc.Scan()
	return sc.Text()
}
