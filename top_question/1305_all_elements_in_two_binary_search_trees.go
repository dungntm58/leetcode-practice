package top_question

import "sort"

/// All Elements in Two Binary Search Trees
// Given two binary search trees root1 and root2,
// return a list containing all the integers from both trees sorted in ascending order.
// Input: root1 = [2,1,4], root2 = [1,0,3]
// Output: [0,1,1,2,3,4]
// Input: root1 = [1,null,8], root2 = [8,1]
// Output: [1,1,8,8]

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func GetAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	res := []int{}
	recursiveGetAllElements(root1, &res)
	recursiveGetAllElements(root2, &res)
	sort.Ints(res)
	return res
}

// func GetAllElements(root1 *TreeNode, root2 *TreeNode) []int {
// 	res1 := []int{}
// 	res2 := []int{}
// 	recursiveGetAllElements(root1, &res1)
// 	recursiveGetAllElements(root2, &res2)
// 	res := []int{}
// 	i, j := 0, 0
// 	for i < len(res1) && j < len(res2) {
// 		if res1[i] < res2[j] {
// 			res = append(res, res1[i])
// 			i++
// 		} else {
// 			res = append(res, res2[j])
// 			j++
// 		}
// 	}
// 	if i < len(res1) {
// 		res = append(res, res1[i:]...)
// 	}
// 	if j < len(res2) {
// 		res = append(res, res2[j:]...)
// 	}
// 	return res
// }

func recursiveGetAllElements(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	recursiveGetAllElements(root.Left, res)
	*res = append(*res, root.Val)
	recursiveGetAllElements(root.Right, res)
}
