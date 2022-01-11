package top_question

/// Sum of Root To Leaf Binary Numbers
// You are given the root of a binary tree where each node has a value 0 or 1.
// Each root-to-leaf path represents a binary number starting with the most significant bit.
// For example, if the path is 0 -> 1 -> 1 -> 0 -> 1, then this could represent 01101 in binary, which is 13.
// For all leaves in the tree, consider the numbers represented by the path from the root to that leaf.
// Return the sum of these numbers.
// The test cases are generated so that the answer fits in a 32-bits integer.
// Input: root = [1,0,1,0,1,0,1]
// Output: 22
// Explanation: (100) + (101) + (110) + (111) = 4 + 5 + 6 + 7 = 22
// Input: root = [0]
// Output: 0

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func SumRootToLeaf(root *TreeNode) int {
	res := 0
	sumDfsRootToLeaf(root, 0, &res)
	return res
}

func sumDfsRootToLeaf(root *TreeNode, curr int, res *int) {
	if root == nil {
		return
	}
	curr = curr*2 + root.Val
	if root.Left == nil && root.Right == nil {
		*res += curr
	}
	sumDfsRootToLeaf(root.Left, curr, res)
	sumDfsRootToLeaf(root.Right, curr, res)
}
