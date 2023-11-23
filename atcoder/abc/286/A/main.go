package main

import "fmt"

func main() {
	var n, p, q, r, s int
	fmt.Scan(&n, &p, &q, &r, &s)
	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}
	p--
	q--
	r--
	s--

	diff := r - p

	for k := p; k <= q; k++ {
		a[k], a[k+diff] = a[k+diff], a[k]
	}

	for i := range a {
		if i != len(a)-1 {
			fmt.Printf("%d ", a[i])
		} else {
			fmt.Println(a[i])
		}
	}
}
