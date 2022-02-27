package top_question

/// Given the root of a binary tree, return the maximum width of the given tree.
// The maximum width of a tree is the maximum width among all levels.
// The width of one level is defined as the length between the end-nodes (the leftmost and rightmost non-null nodes),
// where the null nodes between the end-nodes are also counted into the length calculation.
// It is guaranteed that the answer will in the range of 32-bit signed integer.
// Input: root = [1,3,2,5,3,null,9]
// Output: 4
// Explanation: The maximum width existing in the third level with the length 4 (5,3,null,9).
// Input: root = [1,3,null,5,3]
// Output: 2
// Explanation: The maximum width existing in the third level with the length 2 (5,3).
// Input: root = [1,3,2,5]
// Output: 2
// Explanation: The maximum width existing in the second level with the length 2 (3,2).

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func WidthOfBinaryTree(root *TreeNode) int {
	res := 1
	type pair struct {
		node  *TreeNode
		index int
	}
	q := []pair{{root, 0}}
	for len(q) > 0 {
		cnt := len(q)
		start, end := q[0].index, q[cnt-1].index
		if res < end-start+1 {
			res = end - start + 1
		}

		for i := 0; i < cnt; i++ {
			p := q[0]
			idx := p.index - start
			q = q[1:]

			if p.node.Left != nil {
				q = append(q, pair{p.node.Left, 2*idx + 1})
			}

			if p.node.Right != nil {
				q = append(q, pair{p.node.Right, 2*idx + 2})
			}
		}
	}
	return res
}
