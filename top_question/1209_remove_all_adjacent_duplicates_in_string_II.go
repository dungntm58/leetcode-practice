package top_question

/// Remove All Adjacent Duplicates in String II
// You are given a string s and an integer k, a k duplicate removal consists of choosing k adjacent
// and equal letters from s and removing them,
// causing the left and the right side of the deleted substring to concatenate together.
// We repeatedly make k duplicate removals on s until we no longer can.
// Return the final string after all such duplicate removals have been made. It is guaranteed that the answer is unique.
// Input: s = "abcd", k = 2
// Output: "abcd"
// Explanation: There's nothing to delete.
// Input: s = "deeedbbcccbdaa", k = 3
// Output: "aa"
// Explanation:
// First delete "eee" and "ccc", get "ddbbbdaa"
// Then delete "bbb", get "dddaa"
// Finally delete "ddd", get "aa"
// Input: s = "pbbcggttciiippooaais", k = 2
// Output: "ps"

func RemoveAdjacentDuplicates(s string, k int) string {
	pairs := [][]int{{int(s[0]), 1}}
	for i := 1; i < len(s); i++ {
		if len(pairs) > 0 && int(s[i]) == pairs[len(pairs)-1][0] {
			if pairs[len(pairs)-1][1] < k-1 {
				pairs[len(pairs)-1][1]++
			} else {
				pairs = pairs[:len(pairs)-1]
			}
		} else {
			pairs = append(pairs, []int{int(s[i]), 1})
		}
	}
	res := make([]byte, 0)
	for _, pair := range pairs {
		for i := 0; i < pair[1]; i++ {
			res = append(res, byte(pair[0]))
		}
	}
	return string(res)
}
