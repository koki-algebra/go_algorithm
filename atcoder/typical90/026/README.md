# 026 - Independent Set on a Tree（★4）

## テーマ: 二部グラフ

## 解法

木は必ず二部グラフである. 木を2色で彩色したとき, 少なくとも片方の色は N/2 個以上の頂点を占める. したがって, 超点数の多い方の色を N/2 個選んで出力すれば良い.

グラフの彩色は深さ優先探索 (Depth First Search; DFS) を用いて実装できる.
