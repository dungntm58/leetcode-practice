package top_question

/// Remove Duplicates from Sorted List II
// Given the head of a sorted linked list, delete all nodes that have duplicate numbers,
// leaving only distinct numbers from the original list.
// Return the linked list sorted as well.
// Input: head = [1,2,3,3,4,4,5]
// Output: [1,2,5]
// Input: head = [1,1,1,2,3]
// Output: [2,3]

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func DeleteDuplicates(head *ListNode) *ListNode {
	sentinel := &ListNode{0, head}
	prev := head
	for head != nil {
		if head.Next != nil && head.Val == head.Next.Val {
			for head.Next != nil && head.Val == head.Next.Val {
				head = head.Next
			}
			prev.Next = head.Next
		} else {
			prev = head
		}
		head = head.Next
	}
	return sentinel.Next
}
