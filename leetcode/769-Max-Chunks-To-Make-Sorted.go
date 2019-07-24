/*
Given an array arr that is a permutation of [0, 1, ..., arr.length - 1], we split the array into some number of "chunks"
(partitions), and individually sort each chunk.  After concatenating them, the result equals the sorted array.

What is the most number of chunks we could have made?

Example 1:

Input: arr = [4,3,2,1,0]
Output: 1
Explanation:
Splitting into two or more chunks will not return the required result.
For example, splitting into [4, 3], [2, 1, 0] will result in [3, 4, 0, 1, 2], which isn't sorted.
Example 2:

Input: arr = [1,0,2,3,4]
Output: 4
Explanation:
We can split into two chunks, such as [1, 0], [2, 3, 4].
However, splitting into [1, 0], [2], [3], [4] is the highest number of chunks possible.
Note:

arr will have length in range [1, 10].
arr[i] will be a permutation of [0, 1, ..., arr.length - 1].

*/
package main

func main() {
	maxChunksToSorted([]int{4, 3, 2, 1, 0})
	maxChunksToSorted([]int{1, 0, 2, 3, 4})
}

func maxChunksToSorted(arr []int) int {
	rst := 1
	lenArr := len(arr)
	var spliteArr []int
	for i := 0; i < lenArr-1; i++ {
		spliteArr = append(spliteArr, 0) // 1 respect arr[i] and arr[i+1] is union together
	}

	for i := 0; i < lenArr; i++ { // the item less then arr[i] can't splitting
		tmp := i
		for ii := i + 1; ii < lenArr; ii++ {
			if arr[ii] < arr[i] {
				tmp = ii
			}
		}
		if tmp != i {
			for ii := i; ii < tmp; ii++ {
				spliteArr[ii] = 1
			}
		}
	}

	for _, v := range spliteArr {
		if v == 0 {
			rst++
		}
	}

	return rst
}
