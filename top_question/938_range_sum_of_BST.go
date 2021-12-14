package top_question

/// Range Sum of BST
// Given the root node of a binary search tree and two integers low and high,
// return the sum of values of all nodes with a value in the inclusive range [low, high].
// Input: root = [10,5,15,3,7,null,18], low = 7, high = 15
// Output: 32
// Explanation: Nodes 7, 10, and 15 are in the range [7, 15]. 7 + 10 + 15 = 32.
// Input: root = [10,5,15,3,7,13,18,1,null,6], low = 6, high = 10
// Output: 23
// Explanation: Nodes 6, 7, and 10 are in the range [6, 10]. 6 + 7 + 10 = 23.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func RangeSumBST(root *TreeNode, low int, high int) int {
	r := 0
	rangeSumBSTDfs(root, low, high, &r)
	return r
}

func rangeSumBSTDfs(root *TreeNode, low int, high int, ans *int) {
	if root == nil {
		return
	}
	if root.Val < low {
		rangeSumBSTDfs(root.Right, low, high, ans)
		return
	}
	if root.Val > high {
		rangeSumBSTDfs(root.Left, low, high, ans)
		return
	}
	*ans += root.Val
	rangeSumBSTDfs(root.Left, low, high, ans)
	rangeSumBSTDfs(root.Right, low, high, ans)
}
