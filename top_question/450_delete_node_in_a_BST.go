package top_question

/// Delete Node in a BST
// Given a root node reference of a BST and a key, delete the node with the given key in the BST.
// Return the root node reference (possibly updated) of the BST.
// Basically, the deletion can be divided into two stages:
// - Search for a node to remove.
// - If the node is found, delete the node.
// Input: root = [5,3,6,2,4,null,7], key = 3
// Output: [5,4,6,2,null,null,7]
// Explanation: Given key to delete is 3. So we find the node with value 3 and delete it.
// One valid answer is [5,4,6,2,null,null,7], shown in the above BST.
// Please notice that another valid answer is [5,2,6,null,4,null,7] and it's also accepted.
// Input: root = [5,3,6,2,4,null,7], key = 0
// Output: [5,3,6,2,4,null,7]
// Explanation: The tree does not contain a node with value = 0.
// Input: root = [], key = 0
// Output: []

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func DeleteNodeStardard(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val > key {
		root.Left = DeleteNodeStardard(root.Left, key)
		return root
	} else if root.Val < key {
		root.Right = DeleteNodeStardard(root.Right, key)
		return root
	} else if root.Left == nil && root.Right == nil {
		return nil
	} else if root.Left == nil {
		return root.Right
	} else if root.Right == nil {
		return root.Left
	}
	minValueNode := minValueNode(root.Right)
	root.Val = minValueNode.Val
	root.Right = DeleteNodeStardard(root.Right, minValueNode.Val)
	return root
}
