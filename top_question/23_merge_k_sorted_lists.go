package top_question

/// Merge k Sorted Lists
// You are given an array of k linked-lists lists, each linked-list is sorted in ascending order.
// Merge all the linked-lists into one sorted linked-list and return it.
// Input: lists = [[1,4,5],[1,3,4],[2,6]]
// Output: [1,1,2,3,4,4,5,6]
// Explanation: The linked-lists are:
// [
//   1->4->5,
//   1->3->4,
//   2->6
// ]
// merging them into one sorted list:
// 1->1->2->3->4->4->5->6
// Input: lists = []
// Output: []
// Input: lists = [[]]
// Output: []

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func MergeKLists(lists []*ListNode) *ListNode {
	n := len(lists)
	switch n {
	case 0:
		return nil
	case 1:
		return lists[0]
	default:
		leftResult := MergeKLists(lists[0 : n/2])
		rightResult := MergeKLists(lists[n/2 : n])
		return merge2Lists(leftResult, rightResult)
	}
}

func merge2Lists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	head := &ListNode{}
	next := head
	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			next.Next = list2
			list2 = list2.Next
		} else {
			next.Next = list1
			list1 = list1.Next
		}
		next = next.Next
	}
	if list1 == nil {
		next.Next = list2
	} else if list2 == nil {
		next.Next = list1
	}
	return head.Next
}
