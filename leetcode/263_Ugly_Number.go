/**
263. Ugly Number
Write a program to check whether a given number is an ugly number.

Ugly numbers are positive numbers whose prime factors only include 2, 3, 5.

Example 1:
	Input: 6
	Output: true
	Explanation: 6 = 2 × 3

Example 2:
	Input: 8
	Output: true
	Explanation: 8 = 2 × 2 × 2

Example 3:
	Input: 14
	Output: false
	Explanation: 14 is not ugly since it includes another prime factor 7.

Note:
	1 is typically treated as an ugly number.
	2 Input is within the 32-bit signed integer range: [−231,  231 − 1].
*/
package main

func main() {
	println(isUgly(14))
}

func isUgly(num int) bool {
	dividedFun := func(dividend int, divider int) int {
		if dividend%divider == 0 {
			return dividend / divider
		} else {
			return 0
		}
	}
	if num < 1 {
		return false
	}
	if num == 1 {
		return true
	}
	// can divided by x 2 and y 3 and z 5
	for _, v := range []int{2, 3, 5} {
		for true {
			rst := dividedFun(num, v)
			if rst == 1 {
				return true
			}
			if rst > 0 {
				num = rst
				continue
			}
			break
		}
	}

	return false
}
