package top_question

// Given a binary tree root and a linked list with head as the first node.
// Return True if all the elements in the linked list starting from the head correspond to some downward path connected in the binary tree otherwise return False.
// In this context downward path means a path that starts at some node and goes downwards.

// Example 1:
//          1
//    4                4
//       2          2
// 1           6          8
//                     1     3

// Input: head = [4,2,8], root = [1,4,4,null,2,2,null,1,null,6,8,null,null,null,null,1,3]
// Output: true
// Explanation: Nodes in blue form a subpath in the binary Tree.

// Example 2:
//          1
//    4                4
//       2          2
// 1           6          8
//                     1     3
// Input: head = [1,4,2,6], root = [1,4,4,null,2,2,null,1,null,6,8,null,null,null,null,1,3]
// Output: true

// Example 3:
// Input: head = [1,4,2,6,8], root = [1,4,4,null,2,2,null,1,null,6,8,null,null,null,null,1,3]
// Output: false
// Explanation: There is no path in the binary tree that contains all the elements of the linked list from head.

// Constraints:
// - The number of nodes in the tree will be in the range [1, 2500].
// - The number of nodes in the list will be in the range [1, 100].
// - 1 <= Node.val <= 100 for each node in the linked list and binary tree.

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func IsSubPath(head *ListNode, root *TreeNode) bool {
	return checkPath(head, root)
}

func checkPath(head *ListNode, root *TreeNode) bool {
	if root == nil {
		return false
	}
	if dfsIsSubPath(head, root) {
		return true
	}
	return checkPath(head, root.Left) || checkPath(head, root.Right)
}

func dfsIsSubPath(head *ListNode, root *TreeNode) bool {
	if head == nil {
		return true
	}
	if root == nil {
		return false
	}
	if head.Val != root.Val {
		return false
	}
	return dfsIsSubPath(head.Next, root.Left) || dfsIsSubPath(head.Next, root.Right)
}
