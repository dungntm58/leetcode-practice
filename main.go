package main

import (
	"fmt"
	"main/contest"
	"main/top_question"
)

func main() {
	// fmt.Println(contest.FindOriginalArray([]int{0, 1, 0, 2}))
	// fmt.Println(contest.MaxTaxiEarnings(3, [][]int{{1, 3, 1}, {1, 2, 4}}))
	// fmt.Println(contest.MaxTaxiEarnings(20, [][]int{{1, 6, 1}, {3, 10, 2}, {10, 12, 3}, {11, 12, 2}, {12, 15, 2}, {13, 18, 1}}))
	// fmt.Println(contest.MaxTaxiEarnings(10, [][]int{{9, 10, 2}, {4, 5, 6}, {6, 8, 1}, {1, 5, 5}, {4, 9, 5}, {1, 6, 5}, {4, 8, 3}, {4, 7, 10}, {1, 9, 8}, {2, 3, 5}}))
	// fmt.Println(contest.SumOfBeauties([]int{6, 8, 3, 7, 8, 9, 4, 8}))

	// detectSquare()

	// fmt.Println(contest.MaximumDifference([]int{7, 1, 5, 4}))
	// fmt.Println(contest.GridGame([][]int{{2, 5, 4}, {1, 5, 1}}))
	// fmt.Println(contest.GridGame([][]int{{1, 3, 1, 15}, {1, 3, 3, 1}}))
	// fmt.Println(contest.GridGame([][]int{{20, 3, 20, 17, 2, 12, 15, 17, 4, 15}, {20, 10, 13, 14, 15, 5, 2, 3, 14, 3}}))
	// fmt.Println(top_question.MinStartValue([]int{-3, 2, -3, 4, 2}))
	// fmt.Println(top_question.Multiply("2", "3"))
	// fmt.Println(top_question.DailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))

	// iterate_combination()
	// fmt.Println(top_question.LargestDivisibleSubset([]int{1, 2, 4, 8, 9}))
	// fmt.Println(top_question.MaxProfit_II([]int{7, 1, 5, 3, 6, 4}))

	// top_question.MergeKLists([]*top_question.ListNode{
	// 	top_question.NewListNode([]int{1, 4, 5}),
	// 	top_question.NewListNode([]int{1, 3, 4}),
	// 	top_question.NewListNode([]int{2, 6}),
	// }).PrintAll()

	// list := top_question.NewListNode([]int{1, 2, 3, 4, 5})
	// top_question.ReorderList(list)
	// list.PrintAll()

	// tree := top_question.BuildTree([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3})
	// tree.PreorderTraverse()
	// fmt.Println()
	// tree.InorderTraverse()
	// fmt.Println()
	// tree.PostorderTraverse()
	// fmt.Println(top_question.LargestComponentSize([]int{4, 6, 15, 35}))
	// fmt.Println(top_question.Merge([][]int{{2, 3}, {4, 5}, {6, 7}, {1, 10}}))
	// fmt.Println(top_question.AllPathsSourceTarget([][]int{{4, 3, 1}, {3, 2, 4}, {3}, {4}, {}}))
	// for _, n := range top_question.SplitListToParts(top_question.NewListNodeConsecutiveInRange(0, 3), 3) {
	// 	n.PrintAll()
	// }
	// streamChecker()
	// fmt.Println(top_question.NthMagicalNumber(5, 2, 4))
	// fmt.Println(top_question.CanPartition([]int{1, 5, 11, 5}))

	// tree := top_question.BuildTreeFromPreorderIndorder([]int{1, 2, 3, 4, 6, 7, 5}, []int{2, 1, 6, 4, 7, 3, 5})
	// codec := top_question.ConstructorTreeCodec()
	// str := codec.Serialize(tree)
	// fmt.Println(str)
	// clonedTree := codec.Deserialize(str)
	// clonedTree.PreorderTraverse()
	// fmt.Println()
	// clonedTree.InorderTraverse()

	// board := [][]byte{{'o', 'a', 'a', 'n'}, {'e', 't', 'a', 'e'}, {'i', 'h', 'k', 'r'}, {'i', 'f', 'l', 'v'}}
	// fmt.Println(top_question.FindWords(board, []string{"oath", "pea", "eat", "rain"}))
	// board := [][]byte{{'a', 'a'}}
	// fmt.Println(top_question.FindWords(board, []string{"aaa"}))

	// grid := [][]byte{
	// 	{'1', '1', '0', '0', '0'},
	// 	{'1', '1', '0', '0', '0'},
	// 	{'0', '0', '1', '0', '0'},
	// 	{'0', '0', '0', '1', '1'},
	// }
	// fmt.Println(top_question.NumIslands(grid))

	// linkedList()

	// fmt.Println(top_question.LongestConsecutive([]int{100, 4, 200, 1, 3, 2}))
	// fmt.Println(top_question.AtMostNGivenDigitSet([]string{"1", "3", "5", "7"}, 100))
	// fmt.Println(top_question.DecodeString("3[a]2[bc]"))

	// fmt.Println(top_question.CombinationSum4([]int{6, 4, 2, 8}, 6))
}

func detectSquare() {
	actions := []string{"add", "add", "add", "add", "count", "add", "count"}
	points := [][]int{
		{3, 10},
		{11, 2},
		{3, 2},
		{11, 10},
		{14, 8},
		{11, 2},
		{11, 10},
	}
	obj := contest.Constructor()
	for i := 0; i < len(actions); i++ {
		switch actions[i] {
		case "add":
			obj.Add(points[i])
		case "count":
			fmt.Println(obj.Count(points[i]))
		}
	}
}

func iterate_combination() {
	actions := []string{"next", "hasNext", "next", "hasNext", "next", "hasNext", "next", "hasNext", "next", "hasNext"}
	obj := top_question.Constructor("abcd", 2)
	for i := 0; i < len(actions); i++ {
		switch actions[i] {
		case "next":
			fmt.Println(obj.Next())
		case "hasNext":
			fmt.Println(obj.HasNext())
		}
	}
}

func streamChecker() {
	obj := top_question.ConstructorStreamChecker([]string{"cd", "f", "kl"})
	params := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l'}
	for _, p := range params {
		fmt.Println(obj.Query(p))
	}
}

func linkedList() {
	myLinkedList := top_question.ConstructorLinkedList()
	myLinkedList.AddAtHead(2)
	myLinkedList.PrintAll()
	myLinkedList.AddAtHead(1)
	myLinkedList.PrintAll()
	myLinkedList.AddAtTail(3)
	myLinkedList.PrintAll()
	myLinkedList.DeleteAtIndex(0)
	myLinkedList.PrintAll()
}
