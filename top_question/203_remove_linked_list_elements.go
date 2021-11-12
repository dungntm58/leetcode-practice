package top_question

/// Remove Linked List Elements
// Given the head of a linked list and an integer val,
// remove all the nodes of the linked list that has Node.val == val, and return the new head.
// Input: head = [1,2,6,3,4,5,6], val = 6
// Output: [1,2,3,4,5]
// Input: head = [], val = 1
// Output: []
// Input: head = [7,7,7,7], val = 7
// Output: []
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func RemoveElements(head *ListNode, val int) *ListNode {
	travesed := head
	for travesed != nil && travesed.Val == val {
		travesed = travesed.Next
	}
	if travesed == nil {
		return nil
	}
	head = travesed
	for travesed.Next != nil {
		if travesed.Next.Val == val {
			travesed.Next = travesed.Next.Next
		} else {
			travesed = travesed.Next
		}
	}
	return head
}
