/*
718. Maximum Length of Repeated Subarray

Given two integer arrays A and B, return the maximum length of an subarray that appears in both arrays.

Example 1:

Input:
	A: [1,2,3,2,1]
	B: [3,2,1,4,7]
Output:
	3
Explanation:
	The repeated subarray with maximum length is [3, 2, 1].
Note:
	1: 1 <= len(A), len(B) <= 1000
	2: 0 <= A[i], B[i] < 100
*/
package main

import "fmt"

func main() {
	A := []int{1, 2, 4, 5, 6, 7}
	B := []int{1, 2, 4, 6, 7}
	fmt.Println(findLength(A, B))
}
func findLength(A []int, B []int) int {
	rInt := 0
	lenA := len(A)
	lenB := len(B)

	for i := 0; i < lenA; i++ {
		for j := 0; j < lenB; j++ {
			tmp := 0
			min := lenB - j
			if lenB-i < min {
				min = lenB - i
			}
			for tmp < min && A[i+tmp] == B[j+tmp] {
				tmp++
			}
			if tmp > rInt {
				rInt = tmp
			}
		}
	}

	return rInt
}
