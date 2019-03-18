/*
594. Longest Harmonious Subsequence

We define a harmonious array is an array where the difference between its maximum value and its minimum value is exactly 1.

Now, given an integer array, you need to find the length of its longest harmonious subsequence among all its possible subsequences.

Example 1:
Input: [1,3,2,2,5,2,3,7]
Output: 5
Explanation: The longest harmonious subsequence is [3,2,2,2,3].
Note: The length of the input array will not exceed 20,000.
*/
package main

import "fmt"

func main() {
	fmt.Println(findLHS([]int{1, 3, 2, 2, 5, 2, 3, 7}))
}

func findLHS(nums []int) int {
	var rst int

	mMap := make(map[int]int)

	for _, v := range nums {
		mMap[v]++
	}

	for k, _ := range mMap {
		if mMap[k+1] != 0 {
			tmp := mMap[k] + mMap[k+1]
			if rst < tmp {
				rst = tmp
			}
		}
	}

	return rst
}
