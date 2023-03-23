package leetcode

func find(x int, parent []int) int {
	if x != parent[x] {
		parent[x] = find(parent[x], parent)
	}
	return parent[x]
}

func disjoint(a int, b int, parent []int) {
	a = find(a, parent)
	b = find(b, parent)
	parent[b] = a
}

func same(a int, b int, parent []int) bool {
	return find(a, parent) == find(b, parent)
}

func makeConnected(n int, connections [][]int) int {
	cableNum := len(connections)
	parent := make([]int, n)
	group := n
	if cableNum < n-1 {
		return -1
	}
	//initialize
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	//union find
	for i := 0; i < cableNum; i++ {
		if !same(connections[i][0], connections[i][1], parent) {
			disjoint(connections[i][0], connections[i][1], parent)
			group--
		}
	}
	return group - 1
}
