package top_question

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (head *ListNode) PrintAll() {
	next := head
	for next != nil {
		fmt.Print(next.Val, " ")
		next = next.Next
	}
	fmt.Println()
}

func NewListNode(arr []int) *ListNode {
	head := &ListNode{}
	next := head
	for _, el := range arr {
		next.Next = &ListNode{Val: el}
		next = next.Next
	}
	return head.Next
}
