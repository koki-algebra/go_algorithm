package main

import (
	"fmt"
	"math"
)

var (
	T, L float64
	Q    int
)

type Coordinate struct {
	X, Y, Z float64
}

func main() {
	var x, y, t float64
	fmt.Scanf("%f\n%f %f %f\n%d", &T, &L, &x, &y, &Q)
	statue := Coordinate{X: x, Y: y, Z: 0}

	for i := 0; i < Q; i++ {
		fmt.Scan(&t)
		y, z := pos(t)
		current := Coordinate{X: 0, Y: y, Z: z}

		fmt.Println(angle(current, statue))
	}
}

func pos(t float64) (y, z float64) {
	theta := 2 * math.Pi * t / T
	y = -L / 2 * math.Sin(theta)
	z = L / 2 * (1 - math.Cos(theta))
	return
}

func angle(current, statue Coordinate) float64 {
	dx := math.Abs(current.X - statue.X)
	dy := math.Abs(current.Y - statue.Y)

	rad := math.Atan2(current.Z, math.Sqrt(dx*dx+dy*dy))

	return rad * (180 / math.Pi)
}
