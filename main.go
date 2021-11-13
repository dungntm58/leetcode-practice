package main

import (
	"fmt"
	"main/contest"
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
