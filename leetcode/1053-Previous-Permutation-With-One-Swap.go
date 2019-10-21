/*
1053. Previous Permutation With One Swap
Given an array A of positive integers (not necessarily distinct), return the lexicographically largest permutation that
is smaller than A, that can be made with one swap (A swap exchanges the positions of two numbers A[i] and A[j]).  If it
cannot be done, then return the same array.



Example 1:

Input: [3,2,1]
Output: [3,1,2]
Explanation: Swapping 2 and 1.
Example 2:

Input: [1,1,5]
Output: [1,1,5]
Explanation: This is already the smallest permutation.
Example 3:

Input: [1,9,4,6,7]
Output: [1,7,4,6,9]
Explanation: Swapping 9 and 7.
Example 4:

Input: [3,1,1,3]
Output: [1,3,1,3]
Explanation: Swapping 1 and 3.
*/
package main

import "fmt"

func main() {
	fmt.Println(prevPermOpt1([]int{1, 9, 4, 6, 7}))
}
func prevPermOpt1(A []int) []int {
	lenA := len(A)
	for i := lenA - 2; i > -1; i-- {
		tmpMax := A[i+1]
		tmpMaxIndex := i + 1
		for j := i + 2; j < lenA; j++ {
			if tmpMax < A[j] && A[j] < A[i] {
				tmpMax = A[j]
				tmpMaxIndex = j
			}
		}
		if tmpMax < A[i] {
			// exchange
			A[i], A[tmpMaxIndex] = A[tmpMaxIndex], A[i]
			break
		}
	}

	return A
}
