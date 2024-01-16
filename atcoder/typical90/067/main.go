package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var (
		n string
		k int
	)
	fmt.Scan(&n, &k)

	for i := 0; i < k; i++ {
		d, err := base8ToInt(n)
		if err != nil {
			panic(err)
		}
		n = intToBase9(d)
		n = strings.ReplaceAll(n, "8", "5")
	}

	fmt.Println(n)
}

func base8ToInt(s string) (int64, error) {
	var (
		ret int64
		x   int64 = 1
	)
	for i := len(s) - 1; i >= 0; i-- {
		d, err := strconv.ParseInt(string(s[i]), 0, 64)
		if err != nil {
			return -1, err
		}
		ret += d * x
		x *= 8
	}

	return ret, nil
}

func intToBase9(n int64) (ret string) {
	if n == 0 {
		return "0"
	}

	for n > 0 {
		ret = fmt.Sprint(n%9) + ret
		n /= 9
	}

	return
}
