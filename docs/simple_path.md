# Simple Path

## テーマ

[ABC270 C問題](https://atcoder.jp/contests/abc270/tasks/abc270_c)

木 T 上の任意の2つのノード X から Y への単純パス (simple path)上のノードを順に列挙する問題.

## 解法

[解答例](https://github.com/koki-algebra/go_algorithm/blob/main/atcoder/abc/270/C/main.go)

DFS と二重終端キュー (**d**ouble-**e**nded **que**ue; deque) 用いる. deque は Go の `list.List` で表現できる.

ノード Y が見つかるまで deque に探索済みのノードを `PushBack` していく. このとき, 帰りがけに deque から `PopBack` する. Y が見つかったタイミングで探索を終了し, deque から順に `PopFront` したものが X から Y への単純パスとなる.
