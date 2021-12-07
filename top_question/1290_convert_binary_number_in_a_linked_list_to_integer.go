package top_question

/// Convert Binary Number in a Linked List to Integer
// Given head which is a reference node to a singly-linked list.
// The value of each node in the linked list is either 0 or 1.
// The linked list holds the binary representation of a number.
// Return the decimal value of the number in the linked list.
// Input: head = [1,0,1]
// Output: 5
// Explanation: (101) in base 2 = (5) in base 10
// Input: head = [0]
// Output: 0
// Input: head = [1]
// Output: 1
// Input: head = [1,0,0,1,0,0,1,1,1,0,0,0,0,0,0]
// Output: 18880
// Input: head = [0,0]
// Output: 0

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func GetDecimalValue(head *ListNode) int {
	curr := head
	r := 0
	for curr != nil {
		r = r << 1
		r += curr.Val
		curr = curr.Next
	}
	return r
}
