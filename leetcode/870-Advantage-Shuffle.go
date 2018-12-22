/*
870. Advantage Shuffle
Given two arrays A and B of equal size, the advantage of A with respect to B is the number of indices i for which A[i] > B[i].

Return any permutation of A that maximizes its advantage with respect to B.



Example 1:

Input: A = [2,7,11,15], B = [1,10,4,11]
Output: [2,11,7,15]
Example 2:

Input: A = [12,24,8,32], B = [13,25,32,11]
Output: [24,32,8,12]


Note:

1 <= A.length = B.length <= 10000
0 <= A[i] <= 10^9
0 <= B[i] <= 10^9
*/
package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(advantageCount([]int{12, 24, 8, 32}, []int{13, 25, 32, 11}))
}
func advantageCount(A []int, B []int) []int {
	rst := make([]int, len(A))

	Asort := make([]int, len(A))
	copy(Asort, A)
	sort.Ints(Asort)

	Bsort := make([]int, len(B))
	copy(Bsort, B)
	sort.Ints(Bsort)

	firstIndexOfSlice := func(tag int, slice []int) (k int) {
		v := 0
		for k, v = range slice {
			if v == tag {
				fmt.Println(tag, k)
				return
			}
		}
		return
	}

	fmt.Println(Asort)
	fmt.Println(Bsort)

	lastIndex := len(Asort) - 1
	firstIndex := 0
	for _, v := range Asort {
		if v > Bsort[firstIndex] {
			rst[firstIndexOfSlice(Bsort[firstIndex], B)] = v
			firstIndex++
		} else {
			rst[firstIndexOfSlice(Bsort[lastIndex], B)] = v
			lastIndex--
		}
	}

	return rst
}
