package top_question

/// Design Linked List
// Design your implementation of the linked list. You can choose to use a singly or doubly linked list.
// A node in a singly linked list should have two attributes: val and next.
// val is the value of the current node, and next is a pointer/reference to the next node.
// If you want to use the doubly linked list,
// you will need one more attribute prev to indicate the previous node in the linked list.
// Assume all nodes in the linked list are 0-indexed.
// Implement the MyLinkedList class:
// - MyLinkedList() Initializes the MyLinkedList object.
// - int get(int index) Get the value of the indexth node in the linked list. If the index is invalid, return -1.
// - void addAtHead(int val) Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list.
// - void addAtTail(int val) Append a node of value val as the last element of the linked list.
// - void addAtIndex(int index, int val) Add a node of value val before the indexth node in the linked list. If index equals the length of the linked list, the node will be appended to the end of the linked list. If index is greater than the length, the node will not be inserted.
// - void deleteAtIndex(int index) Delete the indexth node in the linked list, if the index is valid.

type MyLinkedList struct {
	head  *ListNode
	tail  *ListNode
	count int
}

func ConstructorLinkedList() MyLinkedList {
	return MyLinkedList{}
}

func (this *MyLinkedList) Get(index int) int {
	if this.count == 0 || index > this.count-1 {
		return -1
	}
	if index == this.count-1 {
		return this.tail.Val
	}
	curr := this.head
	for i := 0; i < index; i++ {
		curr = curr.Next
	}
	return curr.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	oldHead := this.head
	this.head = &ListNode{Val: val}
	this.head.Next = oldHead
	if this.tail == nil {
		this.tail = this.head
	}
	this.count++
}

func (this *MyLinkedList) AddAtTail(val int) {
	if this.tail == nil {
		this.tail = &ListNode{Val: val}
	} else {
		this.tail.Next = &ListNode{Val: val}
		this.tail = this.tail.Next
	}
	if this.head == nil {
		this.head = this.tail
	}
	this.count++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.count {
		return
	}
	if index == 0 {
		this.AddAtHead(val)
		return
	}
	if index == this.count {
		this.AddAtTail(val)
		return
	}
	curr := this.head
	var prev = curr
	for i := 0; i < index; i++ {
		prev = curr
		curr = curr.Next
	}
	prev.Next = &ListNode{Val: val}
	prev.Next.Next = curr
	this.count++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index > this.count-1 {
		return
	}
	if index == 0 {
		this.head = this.head.Next
		if this.head == nil {
			this.tail = nil
		}
		this.count--
		return
	}
	curr := this.head
	var prev = curr
	for i := 0; i < index; i++ {
		prev = curr
		curr = curr.Next
	}
	if curr != nil {
		prev.Next = curr.Next
	}
	if index == this.count-1 {
		this.tail = prev
	}
	this.count--
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */

func (this *MyLinkedList) PrintAll() {
	this.head.PrintAll()
}
