package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)

	m := make(map[string]struct{})

	for i := 0; i < n; i++ {
		var s string
		fmt.Scan(&s)

		// 重複チェック
		if _, ok := m[s]; ok {
			fmt.Println("No")
			return
		}
		m[s] = struct{}{}

		slice := strings.Split(s, "")

		// 一文字目
		switch slice[0] {
		case "H", "D", "C", "S":
		default:
			fmt.Println("No")
			return
		}

		// 二文字目
		switch slice[1] {
		case "A", "2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K":
		default:
			fmt.Println("No")
			return
		}
	}

	fmt.Println("Yes")
}
