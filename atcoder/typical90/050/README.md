# 050 - Stair Jump（★3）

## テーマ: 漸化式を立てて DP をしよう

## 解法

`dp[i]: i 段目にたどり着くまでの移動方法の数`
として DP を行う. 以下のような漸化式を立てると `O(n)` で解くことができる.

```
// MOD 計算は省略

dp[0] = 1
dp[i] = dp[i-1]           (0 < i < l)
dp[i] = dp[i-1] + dp[i-l] (i ≥ l)
```
