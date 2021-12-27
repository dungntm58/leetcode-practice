package top_question

/// Duplicate Zeros
// Given a fixed-length integer array arr,
// duplicate each occurrence of zero, shifting the remaining elements to the right.
// Note that elements beyond the length of the original array are not written.
// Do the above modifications to the input array in place and do not return anything.
// Input: arr = [1,0,2,3,0,4,5,0]
// Output: [1,0,0,2,3,0,0,4]
// Explanation: After calling your function, the input array is modified to: [1,0,0,2,3,0,0,4]
// Input: arr = [1,2,3]
// Output: [1,2,3]
// Explanation: After calling your function, the input array is modified to: [1,2,3]

// In-place
func DuplicateZeros(arr []int) {
	n, possibleDup := len(arr), 0
	for i := 0; i+possibleDup < n; i++ {
		if arr[i] == 0 {
			if i+possibleDup == n-1 {
				arr[n-1] = 0
				n--
				break
			}
			possibleDup++
		}
	}

	for i := n - 1 - possibleDup; i >= 0; i-- {
		if arr[i] == 0 {
			arr[i+possibleDup] = 0
			possibleDup--
			arr[i+possibleDup] = 0
		} else {
			arr[i+possibleDup] = arr[i]
		}
	}
}

// Use an auxiliary array to store the expected result

// func DuplicateZeros(arr []int) {
// 	n := len(arr)
// 	extraArr := make([]int, n)
// 	variance := 0
// 	for i, num := range arr {
// 		if i+variance >= n {
// 			break
// 		}
// 		if num == 0 {
// 			variance++
// 		} else {
// 			extraArr[i+variance] = num
// 		}
// 	}

// 	for i := 1; i < n; i++ {
// 		arr[i] = extraArr[i]
// 	}
// }
