# 010 - Score Sum Queries（★2）

## テーマ: 累積和

## 解法
各クラス毎に学籍番号 i の生徒までの点数の累積和 sum[i] (i = 0, 1,..., N) を事前に計算しておく. ただし sum[0] = 0 である. これによって各クエリ j に対して, sum[r] - sum[l-1] によって答えを O(1) で求めることができる.