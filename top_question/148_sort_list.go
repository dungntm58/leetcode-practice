package top_question

/// Sort List
// Given the head of a linked list, return the list after sorting it in ascending order.
// Input: head = [4,2,1,3]
// Output: [1,2,3,4]
// Input: head = [-1,5,3,4,0]
// Output: [-1,0,3,4,5]
// Input: head = []
// Output: []

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func SortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := (*ListNode)(nil), head
	for fast != nil && fast.Next != nil {
		if slow == nil {
			slow = head
		} else {
			slow = slow.Next
		}
		fast = fast.Next.Next
	}
	mid := slow.Next
	slow.Next = nil
	return merge2Lists(SortList(head), SortList(mid))
}
