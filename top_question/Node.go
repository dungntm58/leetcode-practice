package top_question

type Node struct {
	Val       int
	Neighbors []*Node
	// 116
	Left  *Node
	Right *Node
	Next  *Node
}
