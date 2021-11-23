package top_question

type DisjointSet struct {
	parents []int
	rank    []int
}

func NewDisjointSet(n int) *DisjointSet {
	parents := make([]int, n)
	for i := range parents {
		parents[i] = i
	}
	rank := make([]int, n)
	return &DisjointSet{
		parents: parents,
		rank:    rank,
	}
}

func (set *DisjointSet) Find(x int) int {
	if set.parents[x] != x {
		set.parents[x] = set.Find(set.parents[x])
	}
	return set.parents[x]
}

func (set *DisjointSet) Union(x, y int) bool {
	xSet, ySet := set.Find(x), set.Find(y)
	if xSet == ySet {
		return false
	}
	if set.rank[xSet] < set.rank[ySet] {
		set.parents[xSet] = ySet
	} else if set.rank[xSet] > set.rank[ySet] {
		set.parents[ySet] = xSet
	} else {
		set.parents[ySet] = xSet
		set.rank[xSet]++
	}
	return true
}

func (set DisjointSet) GetRank() []int {
	return set.rank
}
