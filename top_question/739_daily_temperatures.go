package top_question

/// Daily Temperatures
// Given an array of integers temperatures represents the daily temperatures,
// return an array answer such that answer[i] is the number of days you have to wait after the ith day to get a warmer temperature.
// If there is no future day for which this is possible, keep answer[i] == 0 instead.
// Input: temperatures = [73,74,75,71,69,72,76,73]
// Output: [1,1,4,2,1,1,0,0]
// Input: temperatures = [30,40,50,60]
// Output: [1,1,1,0]
// Input: temperatures = [30,60,90]
// Output: [1,1,0]

func DailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	result := make([]int, n)
	maxTemp := 0

	for i := n - 1; i >= 0; i-- {
		curr := temperatures[i]
		if curr >= maxTemp {
			maxTemp = curr
			continue
		}
		days := 1
		for temperatures[i+days] <= curr {
			days += result[i+days]
		}
		result[i] = days
	}

	return result
}
