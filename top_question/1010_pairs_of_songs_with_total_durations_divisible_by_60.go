package top_question

/// Pairs of Songs With Total Durations Divisible by 60
// You are given a list of songs where the ith song has a duration of time[i] seconds.
// Return the number of pairs of songs for which their total duration in seconds is divisible by 60.
// Formally, we want the number of indices i, j such that i < j with (time[i] + time[j]) % 60 == 0.
// Input: time = [30,20,150,100,40]
// Output: 3
// Explanation: Three pairs have a total duration divisible by 60:
// (time[0] = 30, time[2] = 150): total duration 180
// (time[1] = 20, time[3] = 100): total duration 120
// (time[1] = 20, time[4] = 40): total duration 60
// Input: time = [60,60,60]
// Output: 3
// Explanation: All three pairs have a total duration of 120, which is divisible by 60.

func NumPairsDivisibleBy60(time []int) int {
	numMap := map[int]int{}
	ans := 0
	for _, num := range time {
		r := num % 60
		if times, ok := numMap[r]; ok {
			ans += times
		}
		if r == 0 {
			numMap[0]++
		} else {
			numMap[60-num%60]++
		}
	}
	return ans
}
