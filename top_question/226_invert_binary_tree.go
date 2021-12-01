package top_question

/// Invert Binary Tree
// Given the root of a binary tree, invert the tree, and return its root.
// Input: root = [4,2,7,1,3,6,9]
// Output: [4,7,2,9,6,3,1]
// Input: root = [2,1,3]
// Output: [2,3,1]
// Input: root = []
// Output: []

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func InvertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	if root.Left == nil && root.Right == nil {
		return root
	}
	root.Left, root.Right = InvertTree(root.Right), InvertTree(root.Left)
	return root
}
