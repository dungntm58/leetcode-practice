package top_question

/// Minimum Domino Rotations For Equal Row
// In a row of dominoes, tops[i] and bottoms[i] represent the top and bottom halves of the ith domino.
// (A domino is a tile with two numbers from 1 to 6 - one on each half of the tile.)
// We may rotate the ith domino, so that tops[i] and bottoms[i] swap values.
// Return the minimum number of rotations so that all the values in tops are the same,
// or all the values in bottoms are the same.
// If it cannot be done, return -1.
// Input: tops = [2,1,2,4,2,2], bottoms = [5,2,6,2,3,2]
// Output: 2
// Explanation:
// The first figure represents the dominoes as given by tops and bottoms: before we do any rotations.
// If we rotate the second and fourth dominoes,
// we can make every value in the top row equal to 2, as indicated by the second figure.
// Input: tops = [3,5,1,2,3], bottoms = [3,6,3,3,4]
// Output: -1
// Explanation:
// In this case, it is not possible to rotate the dominoes to make one row of values equal.

func MinDominoRotations(tops []int, bottoms []int) int {
	faceA, faceB, same := [7]int{}, [7]int{}, [7]int{}
	for i := 0; i < len(tops); i++ {
		faceA[tops[i]]++
		faceB[bottoms[i]]++
		if tops[i] == bottoms[i] {
			same[tops[i]]++
		}
	}

	for i := 1; i < 7; i++ {
		if faceA[i]+faceB[i]-same[i] == len(tops) {
			if faceA[i] < faceB[i] {
				return faceA[i] - same[i]
			} else {
				return faceB[i] - same[i]
			}
		}
	}
	return -1
}
