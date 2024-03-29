# ハッシュテーブル (hash table)

## ハッシュテーブルの考え方

$M$ を正の整数, $x$ を $0$ 以上 $M$ 未満の整数として, 以下のクエリを高速に処理する.

- Query1: $x$ をデータ構造に挿入する
- Query2: $x$ をデータ構造から削除する
- Query3: $x$ がデータ構造に含まれるかどうかを判定する

このとき, $x$ を添え字に持つ配列 $T[x]$ に対して

```
T[x] <- データ構造中に x が存在するかどうかを表す値 (true or false)
```

と定義する. この配列 $T$ を用いると　Query1~3 を $O(1)$ で処理できる. この配列 $T$ を **バケット (bucket)** と呼ぶことがある.

しかし, このままではクエリ対象となる要素 $x$ が $0$ 以上 $M$ 未満の整数値に限られてしまう. このアイデアに汎用性を持たせるために用いられるものが **ハッシュテーブル (hash table)** である. ハッシュテーブルでは, 整数とは限らない一般的なデータ集合 $S$ の各要素 $x$ に対し, $0 \leq h(x) < M$ を満たす整数 $h(x)$ に対応させることを考える. この $h(x)$ を **ハッシュ関数 (hash function)** と呼ぶ. また $x$ のことをハッシュテーブルの **キー (key)** と呼び, ハッシュ関数の値 $h(x)$ のことを **ハッシュ値 (hash value)** と呼ぶ. また, どのキー  $x \in S$ に対してもハッシュ値 $h(x)$ が異なるようなハッシュ関数を **完全ハッシュ関数 (perfect hash function)** と呼ぶ.

## ハッシュの衝突対策

現実的な用途において, 完全ハッシュ関数を設計することは困難である. 異なる要素に対してハッシュ値が等しくなることをハッシュの **衝突 (collision)** という. ハッシュの衝突を解決する方法は, 各ハッシュ値ごとに連結リストを構築する方法が代表的である.

まず, 配列 $T$ を次のように修正する. $S$ の各要素 $x$ に対し, ハッシュ値 $h(x)$ が等しいもの同士で連結リストを構築し, $T[h(x)]$ にはその連結リストの先頭を指すポインタを格納する. 要素 $x \in S$ を新たにハッシュテーブルに挿入するときには, ハッシュ値 $h(x)$ に対応する連結リストに $x$ を挿入し, $T[h(x)]$ をその先頭を指すポインタに書き換える. また, ハッシュテーブルから要素 $x$ を検索するときには, $T[h(x)]$ が指す連結リストをたどり, その連結リストの各ノードの中身を $x$ と照合する.

## ハッシュテーブルの計算量

連結リストを用いたハッシュテーブルの計算量について考える. 最悪ケースは, $N$ 個のキーが全て同一のハッシュ値を持つ場合であり、キーの探索に $O(N)$ の計算量を要してしまう. しかし, ハッシュ関数が十分によい性能を持つとき, 理想的には「任意のキーに対してハッシュ値が特定の値をとる確率が $1/M$ であり, 任意の2つのキーに対して, それらの類似性とは関係なくハッシュ値が衝突する確率が $1/M$ である」という **単純一様ハッシュ (simple uniform hashing)** の仮定を満たすとき, ハッシュテーブルの各要素にアクセスする計算量は平均的に $O(1 + N/M)$ となる. ここで, $a = N/M$ を **負荷率 (load factor)** と呼び, ハッシュテーブルの振る舞いを示す重要な指標である. 経験的には $a = 1/2$ 程度とすれば, 十分 $O(1)$ の計算量が達成できることが知られている.
