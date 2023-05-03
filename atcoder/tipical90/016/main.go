package main

import "fmt"

const (
	INF = 1 << 60
	L   = 9999
)

func main() {
	var n, a, b, c int
	fmt.Scanf("%d\n%d %d %d", &n, &a, &b, &c)

	ans := INF
	for x := 0; x <= L; x++ {
		for y := 0; x+y <= L; y++ {
			if (n-a*x-b*y)%c != 0 {
				continue
			}
			z := (n - a*x - b*y) / c
			if z < 0 {
				continue
			}
			if x+y+z < ans {
				ans = x + y + z
			}
		}
	}
	fmt.Println(ans)
}
