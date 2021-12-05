package top_question

/// House Robber III
// The thief has found himself a new place for his thievery again. There is only one entrance to this area, called root.
// Besides the root, each house has one and only one parent house.
// After a tour, the smart thief realized that all houses in this place form a binary tree.
// It will automatically contact the police if two directly-linked houses were broken into on the same night.
// Given the root of the binary tree, return the maximum amount of money the thief can rob without alerting the police.
// Input: root = [3,2,3,null,3,null,1]
// Output: 7
// Explanation: Maximum amount of money the thief can rob = 3 + 3 + 1 = 7.
// Input: root = [3,4,5,1,3,null,1]
// Output: 9
// Explanation: Maximum amount of money the thief can rob = 4 + 5 = 9.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rob_III(root *TreeNode) int {
	robRoot, robChildren := recursiveRob(root)
	if robRoot < robChildren {
		return robChildren
	}
	return robRoot
}

func recursiveRob(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}
	notRobLeft, robLeft := recursiveRob(root.Left)
	notRobRight, robRight := recursiveRob(root.Right)
	maxRobLeft := notRobLeft
	if maxRobLeft < robLeft {
		maxRobLeft = robLeft
	}
	maxRobRight := notRobRight
	if maxRobRight < robRight {
		maxRobRight = robRight
	}
	return maxRobLeft + maxRobRight, root.Val + notRobLeft + notRobRight
}
