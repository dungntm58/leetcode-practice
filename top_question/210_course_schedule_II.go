package top_question

/// Course Schedule II
// There are a total of numCourses courses you have to take, labeled from 0 to numCourses - 1.
// You are given an array prerequisites where prerequisites[i] = [ai, bi] indicates
// that you must take course bi first if you want to take course ai.
// For example, the pair [0, 1], indicates that to take course 0 you have to first take course 1.
// Return the ordering of courses you should take to finish all courses.
// If there are many valid answers, return any of them. If it is impossible to finish all courses, return an empty array.
// Input: numCourses = 2, prerequisites = [[1,0]]
// Output: [0,1]
// Explanation: There are a total of 2 courses to take. To take course 1 you should have finished course 0. So the correct course order is [0,1].
// Input: numCourses = 4, prerequisites = [[1,0],[2,0],[3,1],[3,2]]
// Output: [0,2,1,3]
// Explanation: There are a total of 4 courses to take. To take course 3 you should have finished both courses 1 and 2. Both courses 1 and 2 should be taken after you finished course 0.
// So one correct course order is [0,1,2,3]. Another correct ordering is [0,2,1,3].
// Input: numCourses = 1, prerequisites = []
// Output: [0]

func FindOrder(numCourses int, prerequisites [][]int) []int {
	r := []int{}
	graph := make([]map[int]struct{}, numCourses)
	for i := 0; i < numCourses; i++ {
		graph[i] = map[int]struct{}{}
	}
	for _, pre := range prerequisites {
		graph[pre[0]][pre[1]] = struct{}{}
	}
	onpath, visited := make([]bool, numCourses), make([]bool, numCourses)
	for i := 0; i < numCourses; i++ {
		if !visited[i] && dfsSchedule(graph, i, onpath, visited, &r) {
			return []int{}
		}
	}
	return r
}

func dfsSchedule(graph []map[int]struct{}, node int, onpath, visited []bool, toposort *[]int) bool {
	if visited[node] {
		return false
	}
	visited[node] = true
	onpath[node] = true
	for neighbor := range graph[node] {
		if onpath[neighbor] || dfsSchedule(graph, neighbor, onpath, visited, toposort) {
			return true
		}
	}
	*toposort = append(*toposort, node)
	onpath[node] = false
	return false
}
