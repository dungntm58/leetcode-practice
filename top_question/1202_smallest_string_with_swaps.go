package top_question

import "sort"

/// Smallest String With Swaps
// You are given a string s, and an array of pairs of indices in the string pairs
// where pairs[i] = [a, b] indicates 2 indices(0-indexed) of the string.
// You can swap the characters at any pair of indices in the given pairs any number of times.
// Return the lexicographically smallest string that s can be changed to after using the swaps.
// Input: s = "dcab", pairs = [[0,3],[1,2]]
// Output: "bacd"
// Explaination:
// Swap s[0] and s[3], s = "bcad"
// Swap s[1] and s[2], s = "bacd"
// Input: s = "dcab", pairs = [[0,3],[1,2],[0,2]]
// Output: "abcd"
// Explaination:
// Swap s[0] and s[3], s = "bcad"
// Swap s[0] and s[2], s = "acbd"
// Swap s[1] and s[2], s = "abcd"
// Input: s = "cba", pairs = [[0,1],[1,2]]
// Output: "abc"
// Explaination:
// Swap s[0] and s[1], s = "bca"
// Swap s[1] and s[2], s = "bac"
// Swap s[0] and s[1], s = "abc"

func SmallestStringWithSwaps(s string, pairs [][]int) string {
	n := len(s)
	visited := make([]bool, n)
	adj := make([][]int, n)
	for i := 0; i < n; i++ {
		adj[i] = make([]int, 0)
	}
	for _, pair := range pairs {
		adj[pair[0]] = append(adj[pair[0]], pair[1])
		adj[pair[1]] = append(adj[pair[1]], pair[0])
	}
	res := make([]byte, n)
	for i := 0; i < n; i++ {
		if !visited[i] {
			chars := make([]byte, 0)
			indices := make([]int, 0)

			dfsSmallestStringWithSwaps(s, i, adj, visited, &chars, &indices)

			sort.Slice(chars, func(i, j int) bool {
				return chars[i] < chars[j]
			})
			sort.Ints(indices)
			for j := 0; j < len(chars); j++ {
				res[indices[j]] = chars[j]
			}
		}
	}
	return string(res)
}

func dfsSmallestStringWithSwaps(s string, vertex int, adj [][]int, visited []bool, chars *[]byte, indices *[]int) {
	*chars = append(*chars, s[vertex])
	*indices = append(*indices, vertex)
	visited[vertex] = true
	for _, v := range adj[vertex] {
		if !visited[v] {
			dfsSmallestStringWithSwaps(s, v, adj, visited, chars, indices)
		}
	}
}
