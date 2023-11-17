package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	var h, m int
	fmt.Scan(&h, &m)

	t := time.Date(0, 0, 0, h, m, 0, 0, time.UTC)
	for {
		if isAnswer(t) {
			fmt.Printf("%d %d\n", t.Hour(), t.Minute())
			return
		}
		t = t.Add(time.Minute)
	}
}

func isAnswer(t time.Time) bool {
	h := fmt.Sprintf("%02d", t.Hour())
	m := fmt.Sprintf("%02d", t.Minute())

	s1 := strings.Split(h, "")
	s2 := strings.Split(m, "")

	a, _ := strconv.Atoi(s1[0])
	b, _ := strconv.Atoi(s1[1])
	c, _ := strconv.Atoi(s2[0])

	if a == 2 {
		if c > 3 {
			return false
		}
	}

	if b < 6 {
		return true
	}

	return false
}
