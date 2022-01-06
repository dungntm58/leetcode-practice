package top_question

/// Car Pooling
// There is a car with capacity empty seats.
// The vehicle only drives east (i.e., it cannot turn around and drive west).
// You are given the integer capacity and an array trips where trip[i] = [numPassengersi, fromi, toi]
// indicates that the ith trip has numPassengersi passengers
// and the locations to pick them up and drop them off are fromi and toi respectively.
// The locations are given as the number of kilometers due east from the car's initial location.
// Return true if it is possible to pick up and drop off all passengers for all the given trips,
// or false otherwise.
// Input: trips = [[2,1,5],[3,3,7]], capacity = 4
// Output: false
// Input: trips = [[2,1,5],[3,3,7]], capacity = 5
// Output: true

// Optimize a bit
func CarPooling(trips [][]int, capacity int) bool {
	destination := 0
	for _, trip := range trips {
		if destination < trip[2] {
			destination = trip[2]
		}
	}
	passengers := make([]int, destination+1)
	for _, trip := range trips {
		passengers[trip[1]] += trip[0]
		passengers[trip[2]] -= trip[0]
	}
	carPool := 0
	for _, p := range passengers {
		carPool += p
		if carPool > capacity {
			return false
		}
	}
	return true
}

// func CarPooling(trips [][]int, capacity int) bool {
// 	destination := 0
// 	for _, trip := range trips {
// 		if destination < trip[2] {
// 			destination = trip[2]
// 		}
// 	}
// 	passengers := make([]int, destination+1)
// 	for _, trip := range trips {
// 		for i := trip[1]; i <= trip[2]; i++ {
// 			passengers[i] += trip[0]
// 		}
// 		passengers[trip[2]] -= trip[0]
// 	}
// 	for _, p := range passengers {
// 		if p > capacity {
// 			return false
// 		}
// 	}
// 	return true
// }
