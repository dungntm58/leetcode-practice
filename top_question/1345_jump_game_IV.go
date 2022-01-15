package top_question

/// Jump Game IV
// Given an array of integers arr, you are initially positioned at the first index of the array.
// In one step you can jump from index i to index:
// - i + 1 where: i + 1 < arr.length.
// - i - 1 where: i - 1 >= 0.
// - j where: arr[i] == arr[j] and i != j.
// Return the minimum number of steps to reach the last index of the array.
// Notice that you can not jump outside of the array at any time.
// Input: arr = [100,-23,-23,404,100,23,23,23,3,404]
// Output: 3
// Explanation: You need three jumps from index 0 --> 4 --> 3 --> 9. Note that index 9 is the last index of the array.
// Input: arr = [7]
// Output: 0
// Explanation: Start index is the last index. You do not need to jump.
// Input: arr = [7,6,9,6,9,6,9,7]
// Output: 1
// Explanation: You can jump directly from index 0 to index 7 which is last index of the array.

func MinJumps_IV(arr []int) int {
	n := len(arr)
	if n == 1 {
		return 0
	}
	graph := map[int][]int{}
	for i, v := range arr {
		if list, ok := graph[v]; ok {
			graph[v] = append(list, i)
		} else {
			graph[v] = []int{i}
		}
	}
	curs := map[int]struct{}{}
	curs[0] = struct{}{}
	visited := map[int]struct{}{}
	visited[0] = struct{}{}
	visited[n-1] = struct{}{}

	others := map[int]struct{}{}
	others[n-1] = struct{}{}

	res := 0
	for len(curs) > 0 {
		if len(curs) > len(others) {
			curs, others = others, curs
		}
		nexts := map[int]struct{}{}
		for v := range curs {
			for _, n := range graph[arr[v]] {
				if _, ok := others[n]; ok {
					return res + 1
				}

				if _, ok := visited[n]; !ok {
					nexts[n] = struct{}{}
					visited[n] = struct{}{}
				}
			}
			graph[arr[v]] = []int{}
			if _, ok := others[v+1]; ok {
				return res + 1
			}
			if _, ok := others[v-1]; ok {
				return res + 1
			}
			if v+1 < n {
				if _, ok := visited[v+1]; !ok {
					visited[v+1] = struct{}{}
					nexts[v+1] = struct{}{}
				}
			}
			if v-1 >= 0 {
				if _, ok := visited[v-1]; !ok {
					visited[v-1] = struct{}{}
					nexts[v-1] = struct{}{}
				}
			}
		}

		curs = nexts
		res++
	}
	return -1
}
