package top_question

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (node *TreeNode) PreorderTraverse() {
	fmt.Print(node.Val, " ")
	if node.Left != nil {
		node.Left.PreorderTraverse()
	}
	if node.Right != nil {
		node.Right.PreorderTraverse()
	}
}

func (node *TreeNode) PostorderTraverse() {
	if node.Left != nil {
		node.Left.PostorderTraverse()
	}
	if node.Right != nil {
		node.Right.PostorderTraverse()
	}
	fmt.Print(node.Val, " ")
}

func (node *TreeNode) InorderTraverse() {
	if node.Left != nil {
		node.Left.InorderTraverse()
	}
	fmt.Print(node.Val, " ")
	if node.Right != nil {
		node.Right.InorderTraverse()
	}
}
