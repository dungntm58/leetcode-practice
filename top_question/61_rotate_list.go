package top_question

// Rotate List
// Given the head of a linked list, rotate the list to the right by k places.
// Input: head = [1,2,3,4,5], k = 2
// Output: [4,5,1,2,3]
// Input: head = [0,1,2], k = 4
// Output: [2,0,1]

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func RotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}
	// find the length of the list
	length := 0
	cur := head
	for cur != nil {
		length++
		cur = cur.Next
	}
	k = k % length
	if k == 0 {
		return head
	}
	// find the new head
	cur = head
	for i := 0; i < length-k-1; i++ {
		cur = cur.Next
	}
	newHead := cur.Next
	cur.Next = nil
	// connect the new head to the old tail
	cur = newHead
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = head
	return newHead
}
