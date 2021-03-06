package top_question

/// Populating Next Right Pointers in Each Node
// You are given a perfect binary tree where all leaves are on the same level, and every parent has two children.
// The binary tree has the following definition:
// struct Node {
//   int val;
//   Node *left;
//   Node *right;
//   Node *next;
// }
// Populate each next pointer to point to its next right node.
// If there is no next right node, the next pointer should be set to NULL.
// Initially, all next pointers are set to NULL.
// Input: root = [1,2,3,4,5,6,7]
// Output: [1,#,2,3,#,4,5,6,7,#]
// Explanation: Given the above perfect binary tree (Figure A),
// your function should populate each next pointer to point to its next right node, just like in Figure B.
// The serialized output is in level order as connected by the next pointers, with '#' signifying the end of each level.
// Input: root = []
// Output: []

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func Connect(root *Node) *Node {
	if root == nil || root.Left == nil {
		return root
	}
	if root.Next != nil {
		root.Right.Next = root.Next.Left
	}
	root.Left.Next = root.Right
	Connect(root.Left)
	Connect(root.Right)
	return root
}

// Use an auxiliary space
// func Connect(root *Node) *Node {
// 	if root == nil {
// 		return root
// 	}
// 	queue := []*Node{root}
// 	sameLevelCount := 1
// 	for len(queue) > 0 {
// 		for i := 0; i < sameLevelCount; i++ {
// 			if queue[i].Left == nil {
// 				break
// 			}
// 			queue = append(queue, queue[i].Left, queue[i].Right)
// 		}
// 		queue = queue[sameLevelCount:]
// 		sameLevelCount *= 2
// 		for i := 0; i < len(queue)-1; i++ {
// 			queue[i].Next = queue[i+1]
// 		}
// 	}
// 	return root
// }
