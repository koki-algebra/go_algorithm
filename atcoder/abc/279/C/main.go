package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var h, w int
	fmt.Scan(&h, &w)

	sc := bufio.NewScanner(os.Stdin)
	var buf []byte
	sc.Buffer(buf, 500000)

	s := make([][]string, w)
	t := make([][]string, w)
	for i := 0; i < w; i++ {
		s[i] = make([]string, h)
		t[i] = make([]string, h)
	}

	for i := 0; i < h; i++ {
		if !sc.Scan() {
			panic("buffer overflow")
		}
		slice := strings.Split(sc.Text(), "")
		for j := 0; j < w; j++ {
			s[j][i] = slice[j]
		}
	}
	for i := 0; i < h; i++ {
		if !sc.Scan() {
			panic("buffer overflow")
		}
		slice := strings.Split(sc.Text(), "")
		for j := 0; j < w; j++ {
			t[j][i] = slice[j]
		}
	}

	m := make(map[string]int)
	for i := range s {
		str := strings.Join(s[i], "")
		m[str]++
	}

	for i := range t {
		str := strings.Join(t[i], "")
		m[str]--
		if m[str] < 0 {
			fmt.Println("No")
			return
		}
	}

	fmt.Println("Yes")
}
