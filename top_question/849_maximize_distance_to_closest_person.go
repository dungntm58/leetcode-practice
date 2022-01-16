package top_question

/// Maximize Distance to Closest Person
// You are given an array representing a row of seats where seats[i] = 1 represents a person sitting in the ith seat,
// and seats[i] = 0 represents that the ith seat is empty (0-indexed).
// There is at least one empty seat, and at least one person sitting.
// Alex wants to sit in the seat such that the distance between him and the closest person to him is maximized.
// Return that maximum distance to the closest person.
// Input: seats = [1,0,0,0,1,0,1]
// Output: 2
// Explanation:
// If Alex sits in the second open seat (i.e. seats[2]), then the closest person has distance 2.
// If Alex sits in any other open seat, the closest person has distance 1.
// Thus, the maximum distance to the closest person is 2.
// Input: seats = [1,0,0,0]
// Output: 3
// Explanation:
// If Alex sits in the last seat (i.e. seats[3]), the closest person is 3 seats away.
// This is the maximum distance possible, so the answer is 3.
// Input: seats = [0,1]
// Output: 1

func MaxDistToClosest(seats []int) int {
	startOccupied, endOccupied, res := -1, len(seats)-1, 0
	for i, s := range seats {
		if s == 1 {
			startOccupied = i
			break
		}
	}
	if res < startOccupied {
		res = startOccupied
	}
	for i := len(seats) - 1; i >= startOccupied; i-- {
		if seats[i] == 1 {
			endOccupied = i
			break
		}
	}
	if res < len(seats)-1-endOccupied {
		res = len(seats) - 1 - endOccupied
	}
	if startOccupied == len(seats)-1 {
		return res
	}
	for i := startOccupied + 1; i <= endOccupied; i++ {
		if seats[i] == 1 {
			if res < (i-startOccupied)/2 {
				res = (i - startOccupied) / 2
			}
			startOccupied = i
		}
	}
	return res
}
