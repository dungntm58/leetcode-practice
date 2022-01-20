package top_question

/// Koko Eating Bananas
// Koko loves to eat bananas. There are n piles of bananas, the ith pile has piles[i] bananas.
// The guards have gone and will come back in h hours.
// Koko can decide her bananas-per-hour eating speed of k.
// Each hour, she chooses some pile of bananas and eats k bananas from that pile.
// If the pile has less than k bananas, she eats all of them instead and will not eat any more bananas during this hour.
// Koko likes to eat slowly but still wants to finish eating all the bananas before the guards return.
// Return the minimum integer k such that she can eat all the bananas within h hours.
// Input: piles = [3,6,7,11], h = 8
// Output: 4
// Input: piles = [30,11,23,4,20], h = 5
// Output: 30
// Input: piles = [30,11,23,4,20], h = 6
// Output: 23

func MinEatingSpeed(piles []int, h int) int {
	right := 0
	for _, p := range piles {
		if right < p {
			right = p
		}
	}
	left := 1
	for left < right {
		mid := (left + right) / 2

		hours := 0
		for _, p := range piles {
			if p%mid != 0 {
				hours += p/mid + 1
			} else {
				hours += p / mid
			}
			if hours > h {
				left = mid + 1
				break
			}
		}
		if hours > h {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}
