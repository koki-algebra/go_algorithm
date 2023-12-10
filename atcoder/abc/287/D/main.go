package main

import "fmt"

func main() {
	var s, t string
	fmt.Scan(&s, &t)

	// S と T の先頭 i 文字が一致するかどうか
	pre := make([]bool, len(s)+1)
	// S と T の末尾 i　文字が一致するかどうか
	suf := make([]bool, len(s)+1)

	pre[0] = true
	suf[0] = true

	for i := 0; i < len(t); i++ {
		if !isMatch(s[i], t[i]) {
			break
		}
		pre[i+1] = true
	}

	s = ReverseString(s)
	t = ReverseString(t)

	for i := 0; i < len(t); i++ {
		if !isMatch(s[i], t[i]) {
			break
		}
		suf[i+1] = true
	}

	for x := 0; x <= len(t); x++ {
		if pre[x] && suf[len(t)-x] {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}

func isMatch(c1, c2 byte) bool {
	return c1 == c2 || c1 == '?' || c2 == '?'
}

func ReverseSlice[T any](slice []T) []T {
	new := make([]T, len(slice))
	copy(new, slice)
	for i, j := 0, len(new)-1; i < j; i, j = i+1, j-1 {
		new[i], new[j] = new[j], new[i]
	}
	return new
}

func ReverseString(s string) string {
	return string(ReverseSlice([]rune(s)))
}
