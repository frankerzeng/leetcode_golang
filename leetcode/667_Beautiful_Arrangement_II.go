/*
667. Beautiful Arrangement II
Given two integers n and k, you need to construct a list which contains n different positive integers ranging from 1 to n
and obeys the following requirement:
Suppose this list is [a1, a2, a3, ... , an], then the list [|a1 - a2|, |a2 - a3|, |a3 - a4|, ... , |an-1 - an|] has exactly
k distinct integers.

If there are multiple answers, print any of them.

Example 1:
	Input: n = 3, k = 1
	Output: [1, 2, 3]
	Explanation:
		The [1, 2, 3] has three different positive integers ranging from 1 to 3, and the [1, 1] has exactly 1 distinct integer: 1.

Example 2:
	Input: n = 3, k = 2
	Output: [1, 3, 2]
	Explanation:
		The [1, 3, 2] has three different positive integers ranging from 1 to 3, and the [2, 1] has exactly 2 distinct integers: 1 and 2.
Note:
	The n and k are in the range 1 <= k < n <= 104.
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	// 1 2 3 4 5 6 7 8
	// 8 1 2 3 4 5 6 7
	// 8 1 7 2 3 4 5 6
	t1 := time.Now()
	fmt.Println(constructArray(9999, 9998))
	elapsed := time.Since(t1)
	fmt.Println("running time : ", elapsed)
}
func constructArray(n int, k int) []int {
	var rstArr []int
	for i := 0; i < n; i++ {
		rstArr = append(rstArr, i+1)
	}

	// insert last numeric of arr into index of arr
	insert := func(arr []int, index int) []int {
		lastNum := arr[len(arr)-1]
		for i := len(arr) - 1; i > index; i-- {
			arr[i] = arr[i-1]
		}
		arr[index] = lastNum
		return arr
	}

	//is even numeric
	isEvenNum := (k - 1) % 2

	if isEvenNum == 0 {
		for i := 0; i < (k-1)/2; i++ {
			/*
				move the last numeric to middle of arr ,we get 1+2 distinct integer [last-middle,last-middle-1,1]
				and if move the last numeric to the second of arr,we so get 1+2 distinct integer [last-first,last-third,1]
				but in second way ,we possible get distinct integer never be the same we had used
				for example:
				arr = 1,2,3,4,5,6,7
				change to:
				1 : 	1,7,2,3,4,5,6 	[6,5,1]
				2 : 	1,7,2,6,3,4,5 	[6,5,4,3,1]
				3 : 	1,7,2,6,3,5,4 	[6,5,4,3,2,1]
			*/
			rstArr = insert(rstArr, i*2+1)
		}
	} else {
		// move the last numeric to the font of arr,so we get 2 distinct integer [last-first,1]
		rstArr = insert(rstArr, 0)
		for i := 0; i < (k-2)/2; i++ {
			rstArr = insert(rstArr, i*2+2)
		}
	}

	return rstArr
}
