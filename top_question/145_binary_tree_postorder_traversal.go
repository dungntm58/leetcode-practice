package top_question

// Given the root of a binary tree, return the postorder traversal of its nodes' values.

// Example 1:

// Input: root = [1,null,2,3]
// Output: [3,2,1]
// Example 2:

// Input: root = []
// Output: []
// Example 3:

// Input: root = [1]
// Output: [1]

// Constraints:
// - The number of the nodes in the tree is in the range [0, 100].
// - -100 <= Node.val <= 100

// /**
//   - Definition for a binary tree node.
//   - type TreeNode struct {
//   - Val int
//   - Left *TreeNode
//   - Right *TreeNode
//   - }
//     */

func PostorderTraversal(root *TreeNode) []int {
	var traversal func(*TreeNode)
	res := []int{}
	traversal = func(node *TreeNode) {
		if node.Left != nil {
			traversal(node.Left)
		}
		if node.Right != nil {
			traversal(node.Right)
		}
		res = append(res, node.Val)
	}

	traversal(root)
	return res
}
