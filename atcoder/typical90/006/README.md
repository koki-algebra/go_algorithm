# 006 - Smallest Subsequence（★5）

## テーマ: 辞書順最小は前から貪欲法

## 解法
https://drken1215.hatenablog.com/entry/2021/10/10/195200

## 別解 (ヒープによる解法)
S を左から順に走査し, 文字とインデックスのペアをヒープ (min heap) に入れていく. N-K+1, N-K+2,... 個入れ終わったときに, ヒープから要素を取り出す. しかし, 前回取り出した位置 (インデックス) よりも前の位置の要素が取り出されることがあるので, 適切なものが出てくるまで取り出し続ける. これによって O(NlogN) で得ことができる.