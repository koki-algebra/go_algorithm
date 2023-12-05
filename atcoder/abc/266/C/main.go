package main

import (
	"fmt"
)

type Point struct {
	x, y int
}

type Vector struct {
	x, y int
}

func NewVector(p1, p2 Point) Vector {
	return Vector{
		x: p2.x - p1.x,
		y: p2.y - p1.y,
	}
}

func CrossProduct(v1, v2 Vector) float64 {
	return float64(v1.x*v2.y - v2.x*v1.y)
}

func main() {
	var a, b, c, d Point
	fmt.Scan(&a.x, &a.y, &b.x, &b.y, &c.x, &c.y, &d.x, &d.y)

	ab := NewVector(a, b)
	bc := NewVector(b, c)
	cd := NewVector(c, d)
	da := NewVector(d, a)

	if CrossProduct(ab, bc) < 0 || CrossProduct(bc, cd) < 0 || CrossProduct(cd, da) < 0 || CrossProduct(da, ab) < 0 {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}
}
