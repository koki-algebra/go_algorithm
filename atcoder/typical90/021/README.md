# 021 - Come Back in One Piece（★5）

## テーマ: 強連結成分分解

## 解法

1. グラフの適当な頂点から深さ優先探索 (DFS) を行い, 帰りがけ順 (post-order) にノードを配列 `history` に記録する.
2. グラフの全ての有向辺の向きを逆にし, `history` の後ろから順に DFS を行い, 行き止まったところまでを1つの強連結成分とする.
3. 各強連結成分の頂点数を `k` とするとき, 頂点の組は `k(k-1)/2` 個あるため, 全ての強連結成分についての総和が答えとなる.

これにより `O(N + M)` で
