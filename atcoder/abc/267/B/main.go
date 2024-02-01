package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	isStanding := make([]bool, 10)
	for i := range s {
		if s[i] == '1' {
			isStanding[i] = true
		}
	}

	if isStanding[0] {
		fmt.Println("No")
		return
	}

	// 5 が倒れている
	if !isStanding[4] &&
		(isStanding[1] || isStanding[3] || isStanding[7] || isStanding[6]) &&
		(isStanding[2] || isStanding[5] || isStanding[8] || isStanding[9]) {
		fmt.Println("Yes")
		return
	}

	// 2 と 8 が倒れている
	if !isStanding[1] && !isStanding[7] &&
		(isStanding[3] || isStanding[6]) &&
		(isStanding[4] || isStanding[2] || isStanding[5] || isStanding[8] || isStanding[9]) {
		fmt.Println("Yes")
		return
	}

	// 4 が倒れている
	if !isStanding[3] && isStanding[6] &&
		(isStanding[1] || isStanding[2] ||
			isStanding[4] || isStanding[5] ||
			isStanding[7] || isStanding[8] || isStanding[9]) {
		fmt.Println("Yes")
		return
	}

	// 3 と 9 が倒れている
	if !isStanding[2] && !isStanding[8] &&
		(isStanding[5] || isStanding[9]) &&
		(isStanding[1] || isStanding[3] ||
			isStanding[4] || isStanding[6] || isStanding[7]) {
		fmt.Println("Yes")
		return
	}

	// 6 が倒れている
	if !isStanding[5] && isStanding[9] &&
		(isStanding[1] || isStanding[2] ||
			isStanding[3] || isStanding[4] ||
			isStanding[6] || isStanding[7] || isStanding[8]) {
		fmt.Println("Yes")
		return
	}

	fmt.Println("No")
}
