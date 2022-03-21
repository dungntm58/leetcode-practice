package top_question

/// Partition Labels
// You are given a string s. We want to partition the string into as many parts as possible
// so that each letter appears in at most one part.
// Note that the partition is done so that after concatenating all the parts in order, the resultant string should be s.
// Return a list of integers representing the size of these parts.
// Input: s = "ababcbacadefegdehijhklij"
// Output: [9,7,8]
// Explanation:
// The partition is "ababcbaca", "defegde", "hijhklij".
// This is a partition so that each letter appears in at most one part.
// A partition like "ababcbacadefegde", "hijhklij" is incorrect, because it splits s into less parts.
// Input: s = "eccbbbbdec"
// Output: [10]

func PartitionLabels(s string) []int {
	lastChar := [26]int{}
	for i := 0; i < len(s); i++ {
		lastChar[s[i]-'a'] = i
	}
	res, anchor, j := []int{}, 0, 0
	for i := 0; i < len(s); i++ {
		if j < lastChar[s[i]-'a'] {
			j = lastChar[s[i]-'a']
		}
		if i == j {
			res = append(res, i-anchor+1)
			anchor = i + 1
		}
	}
	return res
}
