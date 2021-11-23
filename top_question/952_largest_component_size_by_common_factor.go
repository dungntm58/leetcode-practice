package top_question

import (
	"math"
)

/// Largest Component Size by Common Factor
// You are given an integer array of unique positive integers nums.
// Consider the following graph:
// - There are nums.length nodes, labeled nums[0] to nums[nums.length - 1],
// - There is an undirected edge between nums[i] and nums[j] if nums[i] and nums[j] share a common factor greater than 1.
// Return the size of the largest connected component in the graph.
// Input: nums = [4,6,15,35]
// Output: 4
// Input: nums = [20,50,9,63]
// Output: 2
// Input: nums = [2,3,6,7,4,12,21,39]
// Output: 8

func LargestComponentSize(nums []int) int {
	n := len(nums)
	maxNum := nums[0]
	for i := 1; i < n; i++ {
		if maxNum < nums[i] {
			maxNum = nums[i]
		}
	}
	set := NewDisjointSet(maxNum + 1)
	r := 0
	primeMap := map[int]int{}
	for _, num := range nums {
		for f := 2; f <= int(math.Sqrt(float64(num))); f++ {
			if num%f == 0 {
				set.Union(num, f)
				set.Union(num, num/f)
			}
		}
	}
	for _, num := range nums {
		v := set.Find(num)
		primeMap[v]++
		if r < primeMap[v] {
			r = primeMap[v]
		}
	}
	return r
}

// func sieve(n int) []int {
// 	primes := make([]int, n+1)
// 	for i := range primes {
// 		primes[i] = i
// 	}
// 	for i := 2; i*i < n; i++ {
// 		if primes[i] == i {
// 			for j := i * i; j < n; j += i {
// 				if primes[j] == j {
// 					primes[j] = i
// 				}
// 			}
// 		}
// 	}
// 	return primes
// }

// func getPrimes(x int, primeFactors []int) []int {
// 	primes := []int{}
// 	for x != 1 {
// 		p := primeFactors[x]
// 		primes = append(primes, p)
// 		for x%p == 0 {
// 			x /= p
// 		}
// 	}
// 	return primes
// }
