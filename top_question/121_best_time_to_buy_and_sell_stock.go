package top_question

/// Best Time to Buy and Sell Stock II
// You are given an array prices where prices[i] is the price of a given stock on the ith day.
// You want to maximize your profit by choosing a single day to buy one stock
// and choosing a different day in the future to sell that stock.
// Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.
// Input: prices = [7,1,5,3,6,4]
// Output: 5
// Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
// Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.
// Input: prices = [7,6,4,3,1]
// Output: 0
// Explanation: In this case, no transactions are done and the max profit = 0.

func MaxProfit(prices []int) int {
	n := len(prices)
	maxArr := make([]int, n)
	maxValue := prices[n-1]
	maxArr[n-1] = maxValue
	for i := n - 2; i > 0; i-- {
		if prices[i] > maxValue {
			maxValue = prices[i]
			maxArr[i] = maxValue
		} else {
			maxArr[i] = maxArr[i+1]
		}
	}
	res := 0
	for i := 0; i < n-1; i++ {
		r := maxArr[i+1] - prices[i]
		if r > res {
			res = r
		}
	}
	return res
}
