# 069 - Colorful Blocks 2（★3）

## テーマ: $a^b \pmod m$ は繰り返し二乗法

## 解法

<https://github.com/E869120/kyopro_educational_90/blob/main/editorial/069.jpg>

## 繰り返し二乗法

$x^n \pmod m$ を愚直に計算すると $O(n)$ かかるが, 指数部 $n$ を **2 の冪乗の和に分解**することで $O(\log n)$ で計算できる.

例:

$$
3^{22} = 3^{16} \times 3^{4} \times 3^{2} = 3^{2^4} \times 3^{2^2} \times 3^{2^1}
$$

つまり, 指数部 $n$ を二進数で表したときにビットが立っているものを反映すれば良い.

```golang
func pow(x, n, m int) int {
	ret := 1
	// x を二進数で表したときの桁数
	digits := countBinaryDigits(n)

	// O(log(n))
	for i := 0; i < digits; i++ {
		// i 桁目にビットが立っているならば x^{2^i} を乗算する
		if n >> i & 1 == 1 {
			ret = ret * x % m
		}

		// x を更新
		x = x * x % m
	}

	return ret
}
```
