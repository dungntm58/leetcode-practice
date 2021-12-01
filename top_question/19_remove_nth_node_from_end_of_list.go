package top_question

/// Remove Nth Node From End of List
// Given the head of a linked list, remove the nth node from the end of the list and return its head.
// Input: head = [1,2,3,4,5], n = 2
// Output: [1,2,3,5]
// Input: head = [1], n = 1
// Output: []
// Input: head = [1,2], n = 1
// Output: [1]

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	fast := head
	for i := 0; i < n; i++ {
		if fast == nil {
			break
		}
		fast = fast.Next
	}
	if fast == nil {
		return head.Next
	}
	slow := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return head
}
