package top_question

/// Maximum Depth of Binary Tree
// Given the root of a binary tree, return its maximum depth.
// A binary tree's maximum depth is the number of nodes along the longest path
// from the root node down to the farthest leaf node.
// Input: root = [3,9,20,null,null,15,7]
// Output: 3
// Input: root = [1,null,2]
// Output: 2
// Input: root = []
// Output: 0
// Input: root = [0]
// Output: 1

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftHeight, rightHeight := MaxDepth(root.Left), MaxDepth(root.Right)
	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}
