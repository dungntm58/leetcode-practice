package top_question

/// Reorder List
// You are given the head of a singly linked-list. The list can be represented as:
// L(0) → L(1) → … → L(n - 1) → Ln
// Reorder the list to be on the following form:
// L(0) → L(n) → L(1) → L(n - 1) → L(2) → L(n - 2) → …
// You may not modify the values in the list's nodes.
// Only nodes themselves may be changed.
// Input: head = [1,2,3,4]
// Output: [1,4,2,3]
// Input: head = [1,2,3,4,5]
// Output: [1,5,2,4,3]

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func ReorderList(head *ListNode) {
	if head.Next == nil {
		return
	}
	mid := head
	fast := head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		mid = mid.Next
	}

	oddReverse := mid.Next
	mid.Next = nil
	oddReverse = reverseLinkedList(oddReverse)

	next := head
	for next != nil && oddReverse != nil {
		oldNext := next.Next
		oldOddReverse := oddReverse.Next

		next.Next = oddReverse
		oddReverse.Next = oldNext

		oddReverse = oldOddReverse
		next = oldNext
	}
}

func reverseLinkedList(head *ListNode) *ListNode {
	var pre, next *ListNode
	curr := head

	for curr != nil {
		next = curr.Next
		curr.Next = pre
		pre, curr = curr, next
	}
	return pre
}
