package top_question

/// Construct Binary Tree from Preorder and Inorder Traversal
// Given two integer arrays preorder and inorder where preorder is the preorder traversal of a binary tree and inorder is the inorder traversal of the same tree,
// construct and return the binary tree.
// Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
// Output: [3,9,20,null,null,15,7]
// Input: preorder = [-1], inorder = [-1]
// Output: [-1]

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func BuildTreeFromPreorderIndorder(preorder []int, inorder []int) *TreeNode {
	indexInorderMap := make(map[int]int)
	for i, el := range inorder {
		indexInorderMap[el] = i
	}
	index := 0
	return buildSubTreeFromPreorderIndorder(preorder, 0, len(preorder)-1, &index, indexInorderMap)
}

func buildSubTreeFromPreorderIndorder(preorder []int, inorderStart int, inorderEnd int, preorderIndex *int, indexInorderMap map[int]int) *TreeNode {
	if inorderStart > inorderEnd {
		return nil
	}
	curr := preorder[*preorderIndex]
	root := &TreeNode{Val: curr}
	*preorderIndex++

	if inorderStart == inorderEnd {
		return root
	}
	inorderIndex := indexInorderMap[curr]
	root.Left = buildSubTreeFromPreorderIndorder(preorder, inorderStart, inorderIndex-1, preorderIndex, indexInorderMap)
	root.Right = buildSubTreeFromPreorderIndorder(preorder, inorderIndex+1, inorderEnd, preorderIndex, indexInorderMap)

	return root
}
