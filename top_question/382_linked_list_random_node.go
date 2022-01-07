package top_question

import "math/rand"

/// Linked List Random Node
// Given a singly linked list, return a random node's value from the linked list.
// Each node must have the same probability of being chosen.
// Implement the Solution class:
// Solution(ListNode head) Initializes the object with the integer array nums.
// int getRandom() Chooses a node randomly from the list and returns its value.
// All the nodes of the list should be equally likely to be choosen.
// Input
// ["Solution", "getRandom", "getRandom", "getRandom", "getRandom", "getRandom"]
// [[[1, 2, 3]], [], [], [], [], []]
// Output
// [null, 1, 3, 2, 2, 3]

// Explanation
// Solution solution = new Solution([1, 2, 3]);
// solution.getRandom(); // return 1
// solution.getRandom(); // return 3
// solution.getRandom(); // return 2
// solution.getRandom(); // return 2
// solution.getRandom(); // return 3
// getRandom() should return either 1, 2, or 3 randomly.
// Each element should have equal probability of returning.

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type Solution382 struct {
	head *ListNode
}

func ConstructorRandomLinkedList(head *ListNode) Solution382 {
	return Solution382{head: head}
}

func (this *Solution382) GetRandom() int {
	scope, chosenValue := 1, 0
	curr := this.head
	for curr != nil {
		if rand.Float32() < 1/float32(scope) {
			chosenValue = curr.Val
		}
		scope++
		curr = curr.Next
	}
	return chosenValue
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */
