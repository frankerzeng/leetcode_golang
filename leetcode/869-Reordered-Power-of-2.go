/**
869. Reordered Power of 2
Starting with a positive integer N, we reorder the digits in any order (including the original order) such that the leading digit is not zero.

Return true if and only if we can do this in a way such that the resulting number is a power of 2.
题目理解：给一个数如 46 ，随意排列组合使得其值等于2的x次方，64=2^6所以返回true
Example 1:
Input: 1
Output: true

Example 2:
Input: 10
Output: false

Example 3:
Input: 16
Output: true

Example 4:
Input: 24
Output: false

Example 5:
Input: 46
Output: true


Note:
1 <= N <= 10^9
*/
package main

import (
	"fmt"
)

func main() {
	fmt.Println(reorderedPowerOf2(1) == true)
	fmt.Println(reorderedPowerOf2(10) == false)
	fmt.Println(reorderedPowerOf2(16) == true)
	fmt.Println(reorderedPowerOf2(24) == false)
	fmt.Println(reorderedPowerOf2(46) == true)
}

func reorderedPowerOf2(N int) bool {
	numTimesMapN := make(map[int]int, 10) // N 里面0-9的数字出现的次数
	for true {
		if N >= 10 {
			numTimesMapN[N%10]++
			N = N / 10
		} else {
			numTimesMapN[N]++
			break
		}
	}

	for i := 0; i < 31; i++ { // 10^9不大于2^30
		numTimesMap := make(map[int]int, 10)
		powr := 1 << uint(i)
		for true {
			if powr >= 10 {
				numTimesMap[powr%10]++
				powr = powr / 10
			} else {
				numTimesMap[powr]++
				break
			}
		}
		flag := true
		for k := 0; k < 10; k++ { // 比较数字出现的次数是否一样
			if numTimesMap[k] != numTimesMapN[k] {
				flag = false
				break
			}
		}
		if flag {
			return true
		}

	}
	return false

}
