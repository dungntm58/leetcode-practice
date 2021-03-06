package top_question

/// Can Place Flowers
// You have a long flowerbed in which some of the plots are planted, and some are not.
// However, flowers cannot be planted in adjacent plots.
// Given an integer array flowerbed containing 0's and 1's,
// where 0 means empty and 1 means not empty, and an integer n,
// return if n new flowers can be planted in the flowerbed without violating the no-adjacent-flowers rule.
// Input: flowerbed = [1,0,0,0,1], n = 1
// Output: true
// Input: flowerbed = [1,0,0,0,1], n = 2
// Output: false

func CanPlaceFlowers(flowerbed []int, n int) bool {
	count := 0
	for i, bed := range flowerbed {
		if bed == 0 && (i == 0 || flowerbed[i-1] == 0) && (i == len(flowerbed)-1 || flowerbed[i+1] == 0) {
			flowerbed[i] = 1
			count++
		}
		if count >= n {
			return true
		}
	}
	return false
}
