# ダイクストラ法 (Dijkstra's algorithm)
**ダイクストラ法 (Dijkstra's algorithm)** はグラフ理論におけるエッジの重みが非負数の場合の単一始点最短経路問題を解くための最良優先探索によるアルゴリズムである.

## アルゴリズム
```
ダイクストラ法

1. 始点のコストを 0 とする.
2. 未確定の地点から最も小さな値を持つノードを一つ選び, その値を確定させる.
3. 2 で確定したノードの隣接する未確定なノードに対してコストを計算し, 既存のコストよりも小さければ更新する.
4. 全てのノードのコストが確定していれば終了, そうでなければ 2 に戻る.
```

## 実装例
https://github.com/koki-algebra/go_algorithm/blob/main/pkg/libs/dijkstra.go

## 計算量
グラフ `G(V, E)` に対して, 優先度付きキュー (Priority Queue) を用いたダイクストラ法の計算量は `O((|V|+|E|)log|V|)` である.

まず, Priority Queue の Pop の計算量は `O(log|V|)` であり, 最悪の場合これをノードの数だけ実行するので, `O(|V|log|V|)` となる.

また, Priority Queue の Push の計算量は `O(log|V|)` であり, 最悪の場合これをエッジの数だけ実行するので, `O(|E|log|V|)` となる.

これらの合計が `O((|V|+|E|)log|V|)` となる.