# 014 - We Used to Sing a Song Together（★3）

## テーマ: ソートして貪欲法

## 解法
A, B をそれぞれ昇順にソートしたものを (A[1],...,A[N]), (B[1],...,B[N]) とすると, 答えは
<div style="text-align: center;">
	|A[1] - B[1]| + ... + |A[N] - B[N]|
</div>
となる. したがって, O(NlogN) で解くことができる.
