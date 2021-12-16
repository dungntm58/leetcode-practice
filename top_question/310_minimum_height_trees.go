package top_question

/// Minimum Height Trees
// A tree is an undirected graph in which any two vertices are connected by exactly one path.
// In other words, any connected graph without simple cycles is a tree.
// Given a tree of n nodes labelled from 0 to n - 1,
// and an array of n - 1 edges where edges[i] = [ai, bi] indicates that there is an undirected edge between the two nodes ai and bi in the tree,
// you can choose any node of the tree as the root. When you select a node x as the root, the result tree has height h.
// Among all possible rooted trees, those with minimum height (i.e. min(h))  are called minimum height trees (MHTs).
// Return a list of all MHTs' root labels. You can return the answer in any order.
// The height of a rooted tree is the number of edges on the longest downward path between the root and a leaf.
// Input: n = 4, edges = [[1,0],[1,2],[1,3]]
// Output: [1]
// Explanation: As shown, the height of the tree is 1 when the root is the node with label 1 which is the only MHT.
// Input: n = 6, edges = [[3,0],[3,1],[3,2],[3,4],[5,4]]
// Output: [3,4]
// Input: n = 1, edges = []
// Output: [0]
// Input: n = 2, edges = [[0,1]]
// Output: [0,1]

func FindMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}
	degree, adj := make([]int, n), make([][]int, n)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
		degree[e[0]]++
		degree[e[1]]++
	}
	queue, r := []int{}, []int{}
	for i, d := range degree {
		if d == 1 {
			degree[i]--
			queue = append(queue, i)
		}
	}
	for len(queue) > 0 {
		queueLen := len(queue)
		r = []int{}
		for queueLen > 0 {
			v := queue[0]
			queue = queue[1:]
			r = append(r, v)
			for _, node := range adj[v] {
				degree[node]--
				if degree[node] == 1 {
					queue = append(queue, node)
				}
			}
			queueLen--
		}
	}
	return r
}

// Brutal force
// func FindMinHeightTrees(n int, edges [][]int) []int {
// 	minHeight := n
// 	r := []int{}
// 	for i := 0; i < n; i++ {
// 		h := 0
// 		findHeight(i, edges, &h)
// 		if h < minHeight {
// 			minHeight = h
// 			r = []int{i}
// 		} else if h == minHeight {
// 			r = append(r, i)
// 		}
// 	}
// 	return r
// }

// func findHeight(node int, edges [][]int, h *int) {
// 	if len(edges) == 0 {
// 		return
// 	}
// 	adjacentNodes := []int{}
// 	filteredEdges := [][]int{}
// 	for _, e := range edges {
// 		if e[0] == node {
// 			adjacentNodes = append(adjacentNodes, e[1])
// 		} else if e[1] == node {
// 			adjacentNodes = append(adjacentNodes, e[0])
// 		} else {
// 			filteredEdges = append(filteredEdges, e)
// 		}
// 	}
// 	if len(adjacentNodes) > 0 {
// 		*h++
// 	}
// 	curr := *h
// 	for _, n := range adjacentNodes {
// 		height := curr
// 		findHeight(n, filteredEdges, &height)
// 		if height > *h {
// 			*h = height
// 		}
// 	}
// }
