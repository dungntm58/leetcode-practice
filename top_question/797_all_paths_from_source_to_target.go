package top_question

/// All Paths From Source to Target
// Given a directed acyclic graph (DAG) of n nodes labeled from 0 to n - 1,
// find all possible paths from node 0 to node n - 1 and return them in any order.
// The graph is given as follows: graph[i] is a list of all nodes you can visit from node i
// (i.e., there is a directed edge from node i to node graph[i][j]).
// Input: graph = [[1,2],[3],[3],[]]
// Output: [[0,1,3],[0,2,3]]
// Explanation: There are two paths: 0 -> 1 -> 3 and 0 -> 2 -> 3.
// Input: graph = [[4,3,1],[3,2,4],[3],[4],[]]
// Output: [[0,4],[0,3,4],[0,1,3,4],[0,1,2,3,4],[0,1,4]]
// Input: graph = [[1],[]]
// Output: [[0,1]]
// Input: graph = [[1,2,3],[2],[3],[]]
// Output: [[0,1,2,3],[0,2,3],[0,3]]
// Input: graph = [[1,3],[2],[3],[]]
// Output: [[0,1,2,3],[0,3]]

func AllPathsSourceTarget(graph [][]int) [][]int {
	r := [][]int{}
	stack := []int{}
	stack = append(stack, 0)
	return buildPath(graph, r, 0, stack)
}

func buildPath(graph [][]int, ans [][]int, currentIndex int, stack []int) [][]int {
	if currentIndex >= len(graph)-1 {
		r := make([]int, len(stack))
		copy(r, stack)
		ans = append(ans, r)
		return ans
	}
	for _, i := range graph[currentIndex] {
		stack = append(stack, i)
		ans = buildPath(graph, ans, i, stack)
		stack = stack[:len(stack)-1]
	}
	return ans
}
