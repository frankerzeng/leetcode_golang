/*
69. Sqrt(x)
Implement int sqrt(int x).

Compute and return the square root of x, where x is guaranteed to be a non-negative integer.

Since the return type is an integer, the decimal digits are truncated and only the integer part of the result is returned.

Example 1:

Input: 4
Output: 2
Example 2:

Input: 8
Output: 2
Explanation: The square root of 8 is 2.82842..., and since
             the decimal part is truncated, 2 is returned.
*/
package main

import "fmt"

func main() {
	fmt.Println(mySqrt(8))
}

func mySqrt(x int) int {
	var rst int
	i := 0
	for true {
		if i*i == x {
			return i
		} else if i*i > x && (i-1)*(i-1) < x {
			return i - 1
		} else {
			i++
		}
	}
	return rst
}
