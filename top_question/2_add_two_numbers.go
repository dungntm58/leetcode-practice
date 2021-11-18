package top_question

/// Add Two Numbers
// You are given two non-empty linked lists representing two non-negative integers.
// The digits are stored in reverse order, and each of their nodes contains a single digit.
// Add the two numbers and return the sum as a linked list.
// You may assume the two numbers do not contain any leading zero, except the number 0 itself.
// Input: l1 = [2,4,3], l2 = [5,6,4]
// Output: [7,0,8]
// Explanation: 342 + 465 = 807.
// Input: l1 = [0], l2 = [0]
// Output: [0]
// Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
// Output: [8,9,9,9,0,0,0,1]

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	digitSum := l1.Val + l2.Val
	remember := digitSum / 10

	result := &ListNode{Val: digitSum % 10}
	sum := result
	l1Curr, l2Curr := l1.Next, l2.Next

	for l1Curr != nil || l2Curr != nil {
		var l1Digit = 0
		if l1Curr != nil {
			l1Digit = l1Curr.Val
			l1Curr = l1Curr.Next
		}
		var l2Digit = 0
		if l2Curr != nil {
			l2Digit = l2Curr.Val
			l2Curr = l2Curr.Next
		}
		digitSum = l1Digit + l2Digit + remember
		remember = digitSum / 10
		sum.Next = &ListNode{Val: digitSum % 10}
		sum = sum.Next
	}
	if remember > 0 {
		sum.Next = &ListNode{Val: 1}
	}

	return result
}
