# 048 - I will not drop out（★3）

## テーマ: 上界と下界を見積もる

## 解法

得点の上界を考えてみる. 1 分使うことで以下のうちのいずれかの手段を選べる.

1. まだ部分点を取っていない問題を解き, `b[i]` 点を得る
1. 部分点を取った問題で満点を取り, `a[i]-b[i]` 点を得る

2n 個の手段で得られる点数

```
b[0],..., b[n-1], a[0] - b[0],..., a[n-1] - b[n-1]
```

を降順にソートして, 1 ~ k 番目の要素を足した点数を超えることはできない.

そして, 貪欲に k 個選ぶと答えとなる.
