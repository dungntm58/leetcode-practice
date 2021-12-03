package top_question

/// Split Linked List in Parts
// Given the head of a singly linked list and an integer k, split the linked list into k consecutive linked list parts.
// The length of each part should be as equal as possible: no two parts should have a size differing by more than one.
// This may lead to some parts being null.
// The parts should be in the order of occurrence in the input list,
// and parts occurring earlier should always have a size greater than or equal to parts occurring later.
// Return an array of the k parts.
// Input: head = [1,2,3], k = 5
// Output: [[1],[2],[3],[],[]]
// Explanation:
// The first element output[0] has output[0].val = 1, output[0].next = null.
// The last element output[4] is null, but its string representation as a ListNode is [].
// Input: head = [1,2,3,4,5,6,7,8,9,10], k = 3
// Output: [[1,2,3,4],[5,6,7],[8,9,10]]
// Explanation:
// The input has been split into consecutive parts with size difference at most 1,
// and earlier parts are a larger size than the later parts.

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func SplitListToParts(head *ListNode, k int) []*ListNode {
	if k == 1 {
		return []*ListNode{head}
	}
	if head == nil {
		return make([]*ListNode, k)
	}
	curr := head
	n := 0
	for curr != nil {
		n++
		curr = curr.Next
	}

	longerCount, eachLength := n%k, n/k
	curr = head
	r := make([]*ListNode, k)

	arrI, upper := 0, head.Val+eachLength-1
	if longerCount > 0 {
		upper++
	}
	for curr != nil {
		if r[arrI] == nil {
			r[arrI] = curr
		}

		next := curr.Next
		if curr.Val >= upper {
			upper += eachLength
			if arrI < longerCount-1 {
				upper++
			}
			arrI++
			curr.Next = nil
		}
		curr = next
	}
	return r
}
