package main

import (
	"fmt"
	"math"
)

const (
	MAX     = 9
	epsilon = 1e-9
)

type Coordinate struct {
	r, c float64
}

func NewCoordinate(r, c float64) Coordinate {
	return Coordinate{r: r, c: c}
}

func main() {
	cs := make(map[Coordinate]map[Coordinate]float64)
	for i := 0; i < MAX; i++ {
		var s string
		fmt.Scan(&s)
		for j := range s {
			if s[j] == '#' {
				cs[NewCoordinate(float64(i+1), float64(j+1))] = make(map[Coordinate]float64)
			}
		}
	}

	for c1 := range cs {
		for c2 := range cs {
			if c1 != c2 {
				cs[c1][c2] = d(c1, c2)
			}
		}
	}

	count := 0
	for c, m := range cs {
		for c1, v1 := range m {
			for c2, v2 := range m {
				if !isMidpoint(c, c1, c2) && c1 != c2 && v1 == v2 { // c が c1, c2 の中点の場合は除く
					target := NewCoordinate(c1.r+c2.r-c.r, c1.c+c2.c-c.c)
					if _, ok := m[target]; ok && areAlmostEqual(d(c, target), v1*math.Sqrt2) { // 4つ目の点が存在して, その4点が正方形をなす
						count++
					}
				}
			}
		}
	}
	// 同じ点で2回, 同じ長方形を4回数えているので, 8回分の重複を削除する
	fmt.Println(count / 8)
}

func d(c1, c2 Coordinate) float64 {
	return math.Sqrt((c1.r-c2.r)*(c1.r-c2.r) + (c1.c-c2.c)*(c1.c-c2.c))
}

func areAlmostEqual(a, b float64) bool {
	return math.Abs(a-b) < epsilon
}

func isMidpoint(c, c1, c2 Coordinate) bool {
	mr := (c1.r + c2.r) / 2
	mc := (c1.c + c2.c) / 2
	return c.r == mr && c.c == mc
}
