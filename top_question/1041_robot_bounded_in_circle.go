package top_question

/// Robot Bounded In Circle
// On an infinite plane, a robot initially stands at (0, 0) and faces north. The robot can receive one of three instructions:
// - "G": go straight 1 unit;
// - "L": turn 90 degrees to the left;
// - "R": turn 90 degrees to the right.
// The robot performs the instructions given in order, and repeats them forever.
// Return true if and only if there exists a circle in the plane such that the robot never leaves the circle.
// Input: instructions = "GGLLGG"
// Output: true
// Explanation: The robot moves from (0,0) to (0,2), turns 180 degrees, and then returns to (0,0).
// When repeating these instructions, the robot remains in the circle of radius 2 centered at the origin.
// Input: instructions = "GG"
// Output: false
// Explanation: The robot moves north indefinitely.
// Input: instructions = "GL"
// Output: true
// Explanation: The robot moves from (0, 0) -> (0, 1) -> (-1, 1) -> (-1, 0) -> (0, 0) -> ...

func IsRobotBounded(instructions string) bool {
	point := make([]int, 2)
	direction := 0
	for _, s := range instructions {
		switch s {
		case 'G':
			switch direction {
			case 0:
				point[1]++
			case 1:
				point[0]--
			case 2:
				point[1]--
			case 3:
				point[0]++
			}
		case 'L':
			direction = (direction + 1) % 4
		case 'R':
			direction = (direction + 3) % 4
		}
	}
	return point[0] == 0 && point[1] == 0 || direction != 0
}
