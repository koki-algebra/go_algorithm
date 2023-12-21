package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var (
		T, L, X, Y float64
		q          int
	)
	fmt.Scan(&T, &L, &X, &Y, &q)

	for i := 0; i < q; i++ {
		t := NextFloat(sc)

		// 時刻 t における E869120 君の y, z 座標
		y := -(L / 2) * math.Sin(2*math.Pi*t/T)
		z := (L / 2) * (1 - math.Cos(2*math.Pi*t/T))

		// 水平距離
		dist := math.Sqrt(X*X + (Y-y)*(Y-y))

		fmt.Fprintln(w, math.Atan2(z, dist)*(180/math.Pi))
	}
}

func NextFloat(sc *bufio.Scanner) float64 {
	if sc.Scan() {
		n, err := strconv.ParseFloat(sc.Text(), 64)
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
