package main

import "fmt"

func main() {
	var s, t string
	fmt.Scanf("%s\n%s", &s, &t)

	var (
		n  = len(s)
		ms = make(map[string]int)
		mt = make(map[string]int)
	)

	for i := 0; i < n; i++ {
		ms[string(s[i])]++
		mt[string(t[i])]++
	}

	if ok := solve(ms, mt); !ok {
		fmt.Println("No")
		return
	}
	if ok := solve(mt, ms); !ok {
		fmt.Println("No")
		return
	}

	for c, count := range ms {
		if v, ok := mt[c]; ok {
			if count != v {
				fmt.Println("No")
				return
			}
		} else {
			fmt.Println("No")
			return
		}
	}

	fmt.Println("Yes")
}

func solve(ma, mb map[string]int) bool {
	for c, count := range ma {
		if c != "@" {
			// @ で置き換える回数
			var diff int
			if v, ok := mb[c]; ok {
				diff = count - v
				if diff == 0 {
					continue
				}
			} else {
				diff = count
			}

			switch c {
			case "a", "t", "c", "o", "d", "e", "r":
			default:
				return false
			}

			if diff > 0 {
				if ok := replace(c, abs(diff), mb); !ok {
					return false
				}
			} else {
				if ok := replace(c, abs(diff), ma); !ok {
					return false
				}
			}
		}
	}

	return true
}

func replace(c string, diff int, m map[string]int) bool {
	v, ok := m["@"]
	if !ok {
		return false
	}
	if diff > v {
		return false
	}

	for i := 0; i < diff; i++ {
		m["@"]--
		m[c]++
	}

	return true
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}
