/*
852. Peak Index in a Mountain Array

Let's call an array A a mountain if the following properties hold:

A.length >= 3
There exists some 0 < i < A.length - 1 such that A[0] < A[1] < ... A[i-1] < A[i] > A[i+1] > ... > A[A.length - 1]
Given an array that is definitely a mountain, return any i such that A[0] < A[1] < ... A[i-1] < A[i] > A[i+1] > ... > A[A.length - 1].

Example 1:
Input: [0,1,0]
Output: 1

Example 2:
Input: [0,2,1,0]
Output: 1

Note:
	1: 3 <= A.length <= 10000
	2: 0 <= A[i] <= 10^6
	3: A is a mountain, as defined above.
*/
package main

import "fmt"

func main() {
	A := []int{0, 1, 0}
	fmt.Print(peakIndexInMountainArray(A))
}
func peakIndexInMountainArray(A []int) int {
	ALen := len(A)
	IndexMin := 0
	IndexMax := ALen - 1

	tmpIndex := 0

	for {
		if (IndexMax+IndexMin)%2 == 0 {
			tmpIndex = (IndexMax + IndexMin) / 2
		} else {
			if IndexMax == 1 && IndexMin == 0 {
				tmpIndex = 0
			} else {
				tmpIndex = (IndexMax + IndexMin + 1) / 2
			}
		}

		// first or end numeric
		if tmpIndex == 0 || tmpIndex == ALen-1 {
			return tmpIndex
		}

		// current numeric is lesser than next numeric ,so answer is in the second part
		if A[tmpIndex] < A[tmpIndex+1] {
			IndexMin = tmpIndex
			continue
		}
		// current numeric is bigger than next numeric ,so answer is in the second part
		if A[tmpIndex-1] > A[tmpIndex] {
			IndexMax = tmpIndex
			continue
		}

		return tmpIndex
	}
}
