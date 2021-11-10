package contest

import (
	"sort"
)

/// Count Number of Pairs With Absolute Difference K

func CountKDifference(nums []int, k int) int {
	var count = 0

	for index, curr := range nums {
		for _, next := range nums[index+1:] {
			if next-curr == k || curr-next == k {
				count++
			}
		}
	}
	return count
}

/// Find Original Array From Doubled Array
// An integer array original is transformed into a doubled array changed by appending twice the value of every element in original,
// and then randomly shuffling the resulting array.
// Given an array changed, return original if changed is a doubled array.
// If changed is not a doubled array, return an empty array. The elements in original may be returned in any order.
// Input: changed = [1,3,4,2,6,8]
// Output: [1,3,4]
// Input: changed = [6,3,0,1]
// Output: []
// Input: changed = [1]
// Output: []

func FindOriginalArray(changed []int) []int {
	length := len(changed)
	if length%2 == 1 {
		return []int{}
	}
	sort.Ints(changed)
	occurrencePairs := make(map[int]int)
	for _, curr := range changed {
		occurrencePairs[curr]++
	}
	uniqueArrLength := len(occurrencePairs)
	uniqueArr := make([]int, uniqueArrLength)
	uniqueArrI := 0

	for i, curr := range changed {
		if i == 0 || changed[i-1] < curr {
			uniqueArr[uniqueArrI] = curr
			uniqueArrI++
		}
	}
	result := make([]int, length/2)
	currI := 0
	for _, el := range uniqueArr {
		occur := occurrencePairs[el]
		if occurrencePairs[el*2] < occur {
			return []int{}
		}
		occurrencePairs[el*2] -= occur
		occurrencePairs[el] = 0
		if el == 0 {
			if occur%2 == 1 {
				return []int{}
			}
			for i := 0; i < occur/2; i++ {
				result[currI+i] = el
			}
			currI += occur / 2
		} else {
			for i := 0; i < occur; i++ {
				result[currI+i] = el
			}
			currI += occur
		}
	}
	for _, occur := range occurrencePairs {
		if occur > 0 {
			return []int{}
		}
	}

	return result
}

/// Maximum Earnings From Taxi
// There are n points on a road you are driving your taxi on.
// The n points on the road are labeled from 1 to n in the direction you are going,
// and you want to drive from point 1 to point n to make money by picking up passengers.
// You cannot change the direction of the taxi.
// The passengers are represented by a 0-indexed 2D integer array rides,
// where rides[i] = [starti, endi, tipi] denotes the ith passenger requesting a ride from point starti to point endi who is willing to give a tipi dollar tip.
// For each passenger i you pick up, you earn endi - starti + tipi dollars. You may only drive at most one passenger at a time.
// Given n and rides, return the maximum number of dollars you can earn by picking up the passengers optimally.
// Note: You may drop off a passenger and pick up a different passenger at the same point.
// Input: n = 5, rides = [[2,5,4],[1,5,1]]
// Output: 7
// Explanation: We can pick up passenger 0 to earn 5 - 2 + 4 = 7 dollars.
// Input: n = 20, rides = [[1,6,1],[3,10,2],[10,12,3],[11,12,2],[12,15,2],[13,18,1]]
// Output: 20
// Explanation: We will pick up the following passengers:
// - Drive passenger 1 from point 3 to point 10 for a profit of 10 - 3 + 2 = 9 dollars.
// - Drive passenger 2 from point 10 to point 12 for a profit of 12 - 10 + 3 = 5 dollars.
// - Drive passenger 5 from point 13 to point 18 for a profit of 18 - 13 + 1 = 6 dollars.
// We earn 9 + 5 + 6 = 20 dollars in total.

func totalAmount(ride []int) int64 {
	return int64(ride[1]) - int64(ride[0]) + int64(ride[2])
}

func MaxTaxiEarnings(n int, rides [][]int) int64 {
	sameDestinationRides := map[int][][]int{}
	for _, ride := range rides {
		if ride[1] <= n {
			sameDestinationRides[ride[1]] = append(sameDestinationRides[ride[1]], ride)
		}
	}

	optimizedEarningsDestinations := make([]int64, n+1)
	var result int64 = 0

	for i := 1; i <= n; i++ {
		if len(sameDestinationRides[i]) == 0 {
			optimizedEarningsDestinations[i] = result
		} else {
			optimizedEarningsDestinations[i] = optimizedEarningsDestinations[i-1]
			for _, d := range sameDestinationRides[i] {
				totalBill := totalAmount(d) + optimizedEarningsDestinations[d[0]]
				if totalBill > optimizedEarningsDestinations[i] {
					optimizedEarningsDestinations[i] = totalBill
				}
				if totalBill > result {
					result = totalBill
				}
			}
		}
	}

	return result
}

/// Minimum Number of Operations to Make Array Continuous

// You are given an integer array nums. In one operation, you can replace any element in nums with any integer.
// nums is considered continuous if both of the following conditions are fulfilled:
// All elements in nums are unique.
// The difference between the maximum element and the minimum element in nums equals nums.length - 1.
// For example, nums = [4, 2, 5, 3] is continuous, but nums = [1, 2, 3, 5, 6] is not continuous.
// Return the minimum number of operations to make nums continuous.
// Input: nums = [4,2,5,3]
// Output: 0
// Explanation: nums is already continuous.
// Input: nums = [1,2,3,5,6]
// Output: 1
// Explanation: One possible solution is to change the last element to 4.
// The resulting array is [1,2,3,5,4], which is continuous.
// Input: nums = [1,10,100,1000]
// Output: 3
// Explanation: One possible solution is to:
// - Change the second element to 2.
// - Change the third element to 3.
// - Change the fourth element to 4.
// The resulting array is [1,2,3,4], which is continuous.

func MinOperations(nums []int) int {
	result := 0
	return result
}
