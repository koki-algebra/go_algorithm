package main

import "fmt"

func main() {
	var (
		N, K int
		S    string
		ans  string
	)
	fmt.Scanf("%d %d\n%s", &N, &K, &S)

	j := 0
	for i := 0; i < K; i++ {
	Loop:
		for c := 'a'; c <= 'z'; c++ {
			for k := j; len(S[k:]) >= K-i; k++ {
				if string(S[k]) == string(c) {
					j = k + 1
					ans += string(S[k])
					break Loop
				}
			}
		}
	}
	fmt.Println(ans)
}
