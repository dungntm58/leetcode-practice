package top_question

/// Balanced Binary Tree
// Given a binary tree, determine if it is height-balanced.
// For this problem, a height-balanced binary tree is defined as:
// a binary tree in which the left and right subtrees of every node differ in height by no more than 1.
// Input: root = [3,9,20,null,null,15,7]
// Output: true
// Input: root = [1,2,2,3,3,null,null,4,4]
// Output: false
// Input: root = []
// Output: true

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func IsBalanced(root *TreeNode) bool {
	var result bool = true
	checkHeight(root, &result)
	return result
}

func checkHeight(root *TreeNode, isBalanced *bool) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	leftHeight, rightHeight := checkHeight(root.Left, isBalanced), checkHeight(root.Right, isBalanced)
	variant := leftHeight - rightHeight
	if variant > 1 || variant < -1 {
		*isBalanced = *isBalanced && false
	}
	if variant > 0 {
		return leftHeight + 1
	}
	return rightHeight + 1
}
