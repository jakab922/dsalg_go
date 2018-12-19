package dsalg

type Utree struct {
	parent, rank []int
}

func NewUtree(n int) *Utree {
	parent := make([]int, n)
	rank := make([]int, n)

	for i := 0; i < n; i++ {
		parent[i] = i
		rank[i] = 0
	}

	ret := Utree{
		parent: parent,
		rank:   rank,
	}

	return &ret
}

func (u *Utree) Find(a int) int {
	if a != u.parent[a] {
		u.parent[a] = u.Find(u.parent[a])
	}
	return u.parent[a]
}

func (u *Utree) Union(a, b int) {
	pa := u.Find(a)
	pb := u.Find(b)

	if pa == pb {
		return
	}

	if u.rank[pa] > u.rank[pb] {
		u.parent[pb] = pa
	} else {
		u.parent[pa] = pb
		if u.rank[pa] == u.rank[pb] {
			u.rank[pb] += 1
		}
	}
}
