package top_question

/// Kth Smallest Element in a BST
// Given the root of a binary search tree, and an integer k,
// return the kth smallest value (1-indexed) of all the values of the nodes in the tree.
// Input: root = [3,1,4,null,2], k = 1
// Output: 1
// Input: root = [5,3,6,2,4,null,null,1], k = 3
// Output: 3

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func KthSmallest(root *TreeNode, k int) int {
	stack := []*TreeNode{}
	curr := root
	for {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}
		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k--
		if k == 0 {
			return curr.Val
		}
		curr = curr.Right
	}
}

// func KthSmallest(root *TreeNode, k int) int {
// 	count := nodeCount(root.Left)
// 	if count == k-1 {
// 		return root.Val
// 	} else if count > k-1 {
// 		return KthSmallest(root.Left, k)
// 	} else {
// 		return KthSmallest(root.Right, k-count-1)
// 	}
// }

// func nodeCount(root *TreeNode) int {
// 	if root == nil {
// 		return 0
// 	}
// 	return 1 + nodeCount(root.Left) + nodeCount(root.Right)
// }
