# 008 - AtCounter（★4）

## テーマ: 動的計画法 (Dynamic Programming; DP)

## 解法

以下のような DP を考える.

`dp[i][j]`: 長さ N の文字列 S の最初の i 文字から文字列かを抜き出してつなげる方法のうち, それが長さ M の文字列 T の最初の j 文字まで一致するような方法の個数 (i = 0,1,...,N, j = 0,1,...,M). ただし `dp[0][0]` = 0 とする.

`dp[i][j]` は以下のように更新する. (MOD は省略)

```golang
// S[i] を選ばないとき
dp[i+1][j] += dp[i][j]

// S[i] を選ぶとき
if S[i] == T[j] {
	dp[i+1][j+1] += dp[i][j]
}
```

<https://drken1215.hatenablog.com/entry/2022/04/01/121700>
