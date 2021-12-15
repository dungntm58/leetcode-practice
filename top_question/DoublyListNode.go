package top_question

import "fmt"

type DoublyListNode struct {
	Val      int
	Previous *DoublyListNode
	Next     *DoublyListNode
}

func (head *DoublyListNode) PrintAll() {
	next := head
	for next != nil {
		fmt.Print(next.Val, " ")
		next = next.Next
	}
	fmt.Println()
}

func NewDoublyListNode(arr []int) *DoublyListNode {
	head := &DoublyListNode{}
	next := head
	for _, el := range arr {
		newNode := &DoublyListNode{Val: el}
		next.Next = newNode
		newNode.Previous = next
		next = next.Next
	}
	return head.Next
}

func NewDoublyListNodeConsecutiveInRange(lower int, upper int) *DoublyListNode {
	head := &DoublyListNode{}
	next := head
	for i := lower; i <= upper; i++ {
		newNode := &DoublyListNode{Val: i}
		next.Next = newNode
		newNode.Previous = next
		next = next.Next
	}
	return head.Next
}
