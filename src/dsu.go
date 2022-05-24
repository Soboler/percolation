package main

type DSU struct {
	parent []int
	rank   []int
}

func (dsu *DSU) MakeSet(v int) {
	dsu.parent = append(dsu.parent, v)
	dsu.rank = append(dsu.rank, 0)
}

func (dsu *DSU) FindSet(v int) int {
	if v == dsu.parent[v] {
		return v
	}
	dsu.parent[v] = dsu.FindSet(dsu.parent[v])
	return dsu.parent[v]
}

func (dsu *DSU) UnionSets(a, b int) {
	a = dsu.FindSet(a)
	b = dsu.FindSet(b)
	if a != b {
		if dsu.rank[a] < dsu.rank[b] {
			a, b = b, a
		}
		dsu.parent[b] = a
		if dsu.rank[a] == dsu.rank[b] {
			dsu.rank[a] += 1
		}
	}
}
