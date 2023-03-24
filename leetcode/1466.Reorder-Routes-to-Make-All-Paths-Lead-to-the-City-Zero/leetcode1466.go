package leetcode

import "math"

func dfs(reverse *int, node int, visited []bool, adjList [][]int) {
	visited[node] = true
	for _, neighbor := range adjList[node] {
		abs_neighbor := int(math.Abs(float64(neighbor)))
		if visited[abs_neighbor] == false {
			if neighbor < 0 {
				*reverse++
			}
			dfs(reverse, abs_neighbor, visited, adjList)
		}
	}
}
func minReorder(n int, connections [][]int) int {
	adjList := make([][]int, n)
	visited := make([]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = false
	}
	reverse := 0
	for _, con := range connections {
		from := con[0]
		to := con[1]
		adjList[from] = append(adjList[from], -to)
		adjList[to] = append(adjList[to], from)
	}
	dfs(&reverse, 0, visited, adjList)
	return reverse
}
