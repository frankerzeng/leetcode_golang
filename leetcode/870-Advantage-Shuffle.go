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
	fmt.Println(advantageCount([]int{9, 1, 2, 4, 5}, []int{6, 2, 9, 1, 4}))
}

func advantageCount(A []int, B []int) []int {
	rst := make([]int, len(A))

	Asort := make([]int, len(A))
	copy(Asort, A)
	sort.Ints(Asort)

	higherIndexOfSlice := func(tag int) (v int) { // 找到最小的大于B[i]的元素
		k := 0
		flag := false
		for k, v = range Asort {
			if v > tag {
				flag = true
				break
			}
		}
		if !flag {
			k = 0 // 没找到就取最小的
		}

		v = Asort[k]
		Asort = append(Asort[:k], Asort[k+1:]...)
		return
	}

	for k, v := range B {
		rst[k] = higherIndexOfSlice(v)
	}

	return rst
}

// time limit
func advantageCount2(A []int, B []int) []int { // time limit
	rst := make([]int, len(A))

	AminIndex := -1
	AmaxIndex := -1

	BminIndex := -1
	BmaxIndex := -1

	getMinMax := func(slice []int, s byte) {
		fmt.Println(slice)
		min := -1
		max := -1
		minIndex := 0
		maxIndex := 0
		for k, v := range slice {
			if v == -1 {
				continue
			}
			if min == -1 {
				min = v
				minIndex = k
			}
			if max == -1 {
				max = v
				maxIndex = k
			}
			if v > max {
				max = v
				maxIndex = k
			}
			if v < min {
				min = v
				minIndex = k
			}
		}
		if s == 'A' {
			AminIndex = minIndex
			AmaxIndex = maxIndex
		} else {
			BminIndex = minIndex
			BmaxIndex = maxIndex
		}
	}
	for i := 0; i < len(A); i++ {
		if AminIndex == -1 || AmaxIndex == -1 {
			getMinMax(A, 'A')
		}
		if BminIndex == -1 || BmaxIndex == -1 {
			getMinMax(B, 'B')
		}
		if A[AminIndex] > B[BminIndex] {
			rst[BminIndex] = A[AminIndex]
			A[AminIndex] = -1
			B[BminIndex] = -1
			AminIndex = -1
			BminIndex = -1
		} else {
			rst[BmaxIndex] = A[AminIndex]
			A[AminIndex] = -1
			B[BmaxIndex] = -1
			AminIndex = -1
			BmaxIndex = -1
		}
	}

	return rst
}

// memory limit
func advantageCount1(A []int, B []int) []int {
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
