package top_question

import "math"

/// Shortest Path Visiting All Nodes
// You have an undirected, connected graph of n nodes labeled from 0 to n - 1.
// You are given an array graph where graph[i] is a list of all the nodes connected with node i by an edge.
// Return the length of the shortest path that visits every node.
// You may start and stop at any node, you may revisit nodes multiple times, and you may reuse edges.
// Input: graph = [[1,2,3],[0],[0],[0]]
// Output: 4
// Explanation: One possible path is [1,0,2,0,3]
// Input: graph = [[1],[0,2,4],[1,3,4],[2],[1,2]]
// Output: 4
// Explanation: One possible path is [0,1,4,2,3]

func ShortestPathLength(graph [][]int) int {
	n := len(graph)
	endingMask := (1 << n) - 1
	cache := make([][]int, n+1)
	for i := range cache {
		cache[i] = make([]int, endingMask+1)
	}
	res := math.MaxInt
	for node := 0; node < n; node++ {
		r := shortestPathDP(graph, cache, node, endingMask)
		if res > r {
			res = r
		}
	}
	return res
}

func shortestPathDP(graph [][]int, cache [][]int, node, mask int) int {
	if cache[node][mask] != 0 {
		return cache[node][mask]
	}
	if (mask & (mask - 1)) == 0 {
		return 0
	}

	cache[node][mask] = math.MaxInt - 1
	for _, neighbor := range graph[node] {
		if mask&(1<<neighbor) != 0 {
			betterOption, notVisited := shortestPathDP(graph, cache, neighbor, mask), shortestPathDP(graph, cache, neighbor, mask^(1<<node))
			if betterOption > notVisited {
				betterOption = notVisited
			}
			if cache[node][mask] > 1+betterOption {
				cache[node][mask] = 1 + betterOption
			}
		}
	}

	return cache[node][mask]
}
