# 001 - Yokan Party（★4）

## テーマ: 最小値の最大化には二分探索

## 解法

[解答例](https://github.com/koki-algebra/go_algorithm/tree/main/atcoder/tipical90/001)

「切断されてできる $K+1$ 個のようかんの長さをすべて $x$ 以上とすることは可能か」という判定問題を $O(N)$ で解くことができるので, $x$ の最大値を $O(\log{L})$ で二分探索することで $O(N\log{L})$ で解くことができる.

<https://drken1215.hatenablog.com/entry/2021/06/12/020300>
