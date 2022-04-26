package top_question

/// Min Cost to Connect All Points
// You are given an array points representing integer coordinates of some points on a 2D-plane, where points[i] = [xi, yi].
// The cost of connecting two points [xi, yi] and [xj, yj] is the manhattan distance between them: |xi - xj| + |yi - yj|,
// where |val| denotes the absolute value of val.
// Return the minimum cost to make all points connected.
// All points are connected if there is exactly one simple path between any two points.
// Input: points = [[0,0],[2,2],[3,10],[5,2],[7,0]]
// Output: 20
// Explanation:
// We can connect the points as shown above to get the minimum cost of 20.
// Notice that there is a unique path between every pair of points.
// Input: points = [[3,12],[-2,5],[-4,1]]
// Output: 18

func MinCostConnectPoints(points [][]int) int {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	n := len(points)
	mstCost, usedEdgesCount := 0, 0
	inMst := make([]bool, n)
	minDist := make([]int, n)
	minDist[0] = 0
	for i := 0; i < n; i++ {
		minDist[i] = 1e9
	}
	for usedEdgesCount < n {
		currMinEdge, currNode := int(1e9), -1
		for i := 0; i < n; i++ {
			if !inMst[i] && minDist[i] < currMinEdge {
				currMinEdge = minDist[i]
				currNode = i
			}
		}
		mstCost += currMinEdge
		usedEdgesCount++
		inMst[currNode] = true

		for i := 0; i < n; i++ {
			weight := abs(points[currNode][0]-points[i][0]) + abs(points[currNode][1]-points[i][1])
			if !inMst[i] && weight < minDist[i] {
				minDist[i] = weight
			}
		}
	}

	return mstCost
}
