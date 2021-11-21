package top_question

/// Construct Binary Tree from Inorder and Postorder Traversal
// Given two integer arrays inorder and postorder where inorder is the inorder traversal of a binary tree and postorder is the postorder traversal of the same tree,
// construct and return the binary tree.
// Input: inorder = [9,3,15,20,7], postorder = [9,15,7,20,3]
// Output: [3,9,20,null,null,15,7]
// Input: inorder = [-1], postorder = [-1]
// Output: [-1]

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func BuildTreeFromPostorderInorder(inorder []int, postorder []int) *TreeNode {
	indexInorderMap := map[int]int{}
	for i, el := range inorder {
		indexInorderMap[el] = i
	}
	index := len(inorder) - 1
	return buildSubTreeFromPostorderInorder(postorder, 0, index, &index, indexInorderMap)
}

func buildSubTreeFromPostorderInorder(postorder []int, inorderStart int, inorderEnd int, postOrderIndex *int, indexInorderMap map[int]int) *TreeNode {
	if inorderStart > inorderEnd {
		return nil
	}
	curr := postorder[*postOrderIndex]
	root := &TreeNode{Val: curr}
	*postOrderIndex--

	if inorderStart == inorderEnd {
		return root
	}
	inorderIndex := indexInorderMap[curr]
	root.Right = buildSubTreeFromPostorderInorder(postorder, inorderIndex+1, inorderEnd, postOrderIndex, indexInorderMap)
	root.Left = buildSubTreeFromPostorderInorder(postorder, inorderStart, inorderIndex-1, postOrderIndex, indexInorderMap)

	return root
}
