package top_question

/// Maximum Difference Between Node and Ancestor
// Given the root of a binary tree, find the maximum value v for which there exist different nodes a and b
// where v = |a.val - b.val| and a is an ancestor of b.
// A node a is an ancestor of b if either: any child of a is equal to b or any child of a is an ancestor of b.
// Input: root = [8,3,10,1,6,null,14,null,null,4,7,13]
// Output: 7
// Explanation: We have various ancestor-node differences, some of which are given below :
// |8 - 3| = 5
// |3 - 7| = 4
// |8 - 1| = 7
// |10 - 13| = 3
// Among all possible differences, the maximum value of 7 is obtained by |8 - 1| = 7.
// Input: root = [1,null,2,null,0,3]
// Output: 3

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func MaxAncestorDiff(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return ancestorDiff(root, root.Val, root.Val)
}

func ancestorDiff(root *TreeNode, currMax, currMin int) int {
	if root == nil {
		return currMax - currMin
	}
	if currMax < root.Val {
		currMax = root.Val
	}
	if currMin > root.Val {
		currMin = root.Val
	}
	r := ancestorDiff(root.Left, currMax, currMin)
	right := ancestorDiff(root.Right, currMax, currMin)
	if r < right {
		return right
	}
	return r
}
