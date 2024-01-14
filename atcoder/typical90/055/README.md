# 055 - Select 5（★2）

## テーマ: 定数倍を見積もる

## 解法
全探索の計算量は `Θ(n^5)` であり, 制約が n ≤ 100 のため, TLE しそうだが...
実はこの方法でも AC することができる.

なぜなら, 全探索の数は
```
nC5 = n(n-1)(n-2)(n-3)(n-4) / 5! ≒ n^5 / 120
```
と, 定数倍非常に軽いためである.