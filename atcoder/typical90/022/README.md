# 022 - Cubic Cake（★2）

## テーマ: 最大公約数 (Gratest Common Dvisor; GCD)

## 解法
まず A, B, C の最大公約数 G を求める. ある辺の長さを N とするとき, カット数は `N/G - 1` で求まるため, `A/G + B/G + C/G - 3` が答えとなる.

したがって, `O(log(A + B + C))` で求めることができる.
