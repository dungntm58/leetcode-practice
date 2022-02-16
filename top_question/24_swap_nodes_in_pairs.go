package top_question

/// Swap Nodes in Pairs
// Given a linked list, swap every two adjacent nodes and return its head.
// You must solve the problem without modifying the values in the list's nodes (i.e., only nodes themselves may be changed.)
// Input: head = [1,2,3,4]
// Output: [2,1,4,3]
// Input: head = []
// Output: []
// Input: head = [1]
// Output: [1]

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func SwapPairs(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	tempHead := &ListNode{Next: head}
	ptr := [3]*ListNode{tempHead, head, head.Next}
	for ptr[0] != nil && ptr[1] != nil {
		if ptr[2] != nil {
			ptr[0].Next = ptr[2]
		}
		if ptr[2] != nil {
			ptr[1].Next, ptr[2].Next = ptr[2].Next, ptr[1]
		}
		ptr[0], ptr[1] = ptr[1], ptr[1].Next
		if ptr[1] != nil {
			ptr[2] = ptr[1].Next
		}
	}
	return tempHead.Next
}
