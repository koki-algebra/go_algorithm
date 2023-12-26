# いもす法
累積和のアルゴリズムを多次元, 多次数に拡張したもの.

https://imoz.jp/algorithms/imos_method.html

## 1 次元 0 次いもす法

### 問題例
あなたは喫茶店を経営している. あなたの喫茶店を訪れたそれぞれのお客さん `i (0 ≤ i ≤ C)` について, 入店時刻 `S[i]` と出店時刻 `E[i]` が与えられる `(0 ≤ S[i] ≤ E[i] ≤ T)`. 同時刻にお店にいた客の数の最大値 `M` を求めよ. ただし, 同時刻に出店と入店がある場合は出店が先に行われるものとする.

### ナイーブな解法
それぞれのお客さんについて, 滞在した各時刻のカウントを1足す方法が挙げられる. 計算量は `O(CT)` である.

```golang
table := make([]int, T)

// O(CT)
for i := 0; i < C; i++ {
	for j := S[i]; j < E[i]; j++ {
		table[j]++
	}
}

M := 0
for i := 0; i < T; i++ {
	Chmax(&M, table[i])
}
```

### いもす法による解法
出入り口でのカウントのみをするという発想で, 入店時に +1, 出店時に -1 を記録しておき, 答えを求める直前に全体をシミュレートする. 記録には `O(C)`, シミュレートには `O(T)` かかるので全体としては `O(C + T)` となる.

```golang
table := make([]int, T)

// O(C)
for i := 0; i < C; i++ {
	table[S[i]]++ // 入店処理
	table[E[i]]++ // 出店処理
}

// O(T)
for i := 1; i < T; i++ {
	table[i] += table[i-1]
}

M := 0
for i := 0; i < T; i++ {
	Chmax(&M, table[i])
}
```

## 二元一次いもす法

### 問題例
あなたはさまざまな種類のモンスターを捕まえるゲームをしている. 今あなたがいるのは `W x H` のタイルからなる草むらである. この草むらでは `N` 種類のモンスターが現れる. モンスター `i` は `A[i] ≤ x < B[i], C[i] ≤ y < D[i]` の範囲にしか現れない. このゲームを効率的に進めるため, 1つのタイル上で現れるモンスターの種類の最大値が知りたい.

### ナイーブな解法
それぞれのモンスターについて, 現れる矩形の範囲に含まれる全てのタイルについて 1 を足す方法が挙げられる. 計算量は `O(NWH)` である.

```golang
tiles := make([][]int, H)
for i := 0; i < H; i++ {
	tiles[i] = make([]int, W)
}

for i := 0; i < N; i++ {
	for y := C[i]; y < D[i]; y++ {
		for x := A[i]; x < B[i]; x++ {
			tiles[y][x]++
		}
	}
}

max := 0
for y := 0; y < H; y++ {
	for x := 0; x < W; x++ {
		Chmax(&max, tiles[y][x])
	}
}
```

### いもす法による解法
いもす法では, 矩形の左上 `(A[i], C[i])` と 右下 `(B[i], D[i])` に +1, 右上 `(A[i], D[i])` と左下 `(B[i], C[i])` に -1 を加算し, 答えを求める直前に累積和を計算する. 記録には `O(N)`, 累積和の計算には `O(WH)` かかるので, 全体で `O(N + WH)` かかる.

```golang
tiles := make([][]int, H)
for i := 0; i < H; i++ {
	tiles[i] = make([]int, W)
}

// 重みの加算
for i := 0; i < B; i++ {
	tiles[C[i]][A[i]]++
	tiles[C[i]][B[i]]--
	tiles[D[i]][A[i]]--
	tiles[D[i]][B[i]]++
}

// 横方向の累積和
for y := 0; y < H; y++ {
	for x := 1; w < W; x++ {
		tiles[y][x] += tiles[y][x-1]
	}
}

// 縦方向の累積和
for y := 1; y < H; y++ {
	for x := 0; x < W; x++ {
		tiles[y][x] += tiles[y-1][x]
	}
}

max := 0
for y := 0; y < H; y++ {
	for x := 0; x < W; x++ {
		Chmax(&max, tiles[y][x])
	}
}
```
