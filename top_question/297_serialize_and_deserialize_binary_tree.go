package top_question

import (
	"fmt"
	"strconv"
	"strings"
)

/// Serialize and Deserialize Binary Tree
// Serialization is the process of converting a data structure or object into a sequence of bits
// so that it can be stored in a file or memory buffer,
// or transmitted across a network connection link to be reconstructed later in the same or another computer environment.
// Design an algorithm to serialize and deserialize a binary tree.
// There is no restriction on how your serialization/deserialization algorithm should work.
// You just need to ensure that a binary tree can be serialized to a string
// and this string can be deserialized to the original tree structure.
// https://support.leetcode.com/hc/en-us/articles/360011883654-What-does-1-null-2-3-mean-in-binary-tree-representation-
// Input: root = [1,2,3,null,null,4,5]
// Output: [1,2,3,null,null,4,5]
// Input: root = []
// Output: []
// Input: root = [1]
// Output: [1]
// Input: root = [1,2]
// Output: [1,2]

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
}

func ConstructorTreeCodec() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) Serialize(root *TreeNode) string {
	queue := []*TreeNode{root}
	r := ""
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		if curr != nil {
			r += fmt.Sprintf("%d,", curr.Val)
			queue = append(queue, curr.Left)
			queue = append(queue, curr.Right)
		} else {
			r += "N,"
		}
	}
	return r
}

// Deserializes your encoded data to tree.
func (this *Codec) Deserialize(data string) *TreeNode {
	values := []*int{}
	for len(data) > 0 {
		commaIndex := strings.Index(data, ",")
		i, err := strconv.Atoi(string(data[:commaIndex]))
		data = data[commaIndex+1:]
		if err != nil {
			values = append(values, nil)
		} else {
			values = append(values, &i)
		}
	}

	if values[0] == nil {
		return nil
	}
	i := 1
	root := &TreeNode{Val: *values[0]}
	queue := []*TreeNode{root}
	for i < len(values) {
		curr := queue[0]
		queue = queue[1:]
		if values[i] != nil {
			curr.Left = &TreeNode{Val: *values[i]}
			queue = append(queue, curr.Left)
		}
		i++
		if values[i] != nil {
			curr.Right = &TreeNode{Val: *values[i]}
			queue = append(queue, curr.Right)
		}
		i++
	}
	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
