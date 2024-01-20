# 038 - Large LCM（★3）

## テーマ: オーバーフローに注意

## 解法

整数 a と b の**最小公倍数 (Least Common Multiple; LCM)** `l := LCM(a, b)` は a と b の**最大公約数 (Gratest Common Devisor; GCD)** g := `GCD(a, b)`を用いて

```
l = a * b / g
```

と表される. GCD(a, b) はユークリッドの互除法によって `O(log(a + b))` で求めることができる.

ただし, オーバーフロー対策のために

```
l = (a / g) * b
```

のように計算する. また, `INF = 1e18` を超える場合の判定でも, `(a / g) * b > INF` を判定する代わりに `a / g > INF / b` を判定することでオーバーフローを回避できる.
