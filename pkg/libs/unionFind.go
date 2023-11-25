package libs

type UnionFind struct {
	parents []int
	ranks   []int
}

func NewUnionFind(n int) *UnionFind {
	uf := new(UnionFind)
	uf.parents = make([]int, n)
	uf.ranks = make([]int, n)
	for i := 0; i < n; i++ {
		uf.parents[i] = -1
		uf.ranks[i] = 1
	}

	return uf
}

func (uf *UnionFind) Union(x, y int) {
	x = uf.Find(x)
	y = uf.Find(y)
	if x == y {
		return
	}
	if uf.ranks[x] > uf.ranks[y] {
		x, y = y, x
	}
	if uf.ranks[x] == uf.ranks[y] {
		uf.ranks[y]++
	}
	uf.parents[x] = y
}

func (uf *UnionFind) Find(x int) int {
	if uf.parents[x] < 0 {
		return x
	}
	uf.parents[x] = uf.Find(uf.parents[x])
	return uf.parents[x]
}

func (uf *UnionFind) Same(x, y int) bool {
	return uf.Find(x) == uf.Find(y)
}
