# 075 - Magic For Balls（★3）

## テーマ: $O(\sqrt{n})$ での素因数分解

## 解法

<https://github.com/E869120/kyopro_educational_90/blob/main/editorial/075.jpg>

$n$ を root とし, $n$ を素因数分解したときに素因数が leaf となるような binary tree を考える. 木の高さ $h$ が最小となるとき, $h$ がそのまま答えになる. つまり, $n$ の素因数の個数を $k$ とするとき, $2^h \geq k$ すなわち $h \geq \log_2{k}$ を満たす最小の整数 $h$ を求めれば良い.

整数 $n$ の素因数分解については, $2 \leq x \leq \sqrt{n}$ を満たす全ての整数 $x$ について $n$ を試し割すれば良い.

実装例:

```golang
// key: prime factor, value: exponent of prime factor
primeFactors := make(map[int]int)

for x := 2; x*x <= n; x++ {
	if n % x != 0 {
		continue
	}

	exp := 0
	for n % x == 0 {
		exp++
		n /= x
	}
	primeFactors[x] = exp
}

if n != 1 {
	primeFactors[n] = 1
}
```

以上から全体で $O(\sqrt{n})$ で解くことができる.
