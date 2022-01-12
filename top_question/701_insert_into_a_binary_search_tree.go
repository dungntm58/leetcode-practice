package top_question

/// Insert into a Binary Search Tree
// You are given the root node of a binary search tree (BST) and a value to insert into the tree.
// Return the root node of the BST after the insertion.
// It is guaranteed that the new value does not exist in the original BST.
// Notice that there may exist multiple valid ways for the insertion,
// as long as the tree remains a BST after insertion. You can return any of them.
// Input: root = [4,2,7,1,3], val = 5
// Output: [4,2,7,1,3,5]
// Explanation: Another accepted tree is: [5,2,7,1,3,null,null,null,4]
// Input: root = [40,20,60,10,30,50,70], val = 25
// Output: [40,20,60,10,30,50,70,null,null,25]
// Input: root = [4,2,7,1,3,null,null,null,null,null,null], val = 5
// Output: [4,2,7,1,3,5]

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func InsertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	pre := (*TreeNode)(nil)
	cur := root
	for cur != nil {
		if val < cur.Val {
			pre = cur
			cur = cur.Left
		} else {
			pre = cur
			cur = cur.Right
		}
	}
	if pre.Val < val {
		pre.Right = &TreeNode{Val: val}
	} else {
		pre.Left = &TreeNode{Val: val}
	}
	return root
}
