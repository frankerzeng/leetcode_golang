/*
461. Hamming Distance
The Hamming distance between two integers is the number of positions at which the corresponding bits are different.

Given two integers x and y, calculate the Hamming distance.

Note:
0 ≤ x, y < 231.

Example:

Input: x = 1, y = 4

Output: 2

Explanation:
1   (0 0 0 1)
4   (0 1 0 0)
       ↑   ↑

The above arrows point to positions where the corresponding bits are different.
*/
package main

import "fmt"

func main() {
	fmt.Println(hammingDistance(3, 1))
}

// 异或
func hammingDistance(x int, y int) int {
	rst := 0
	XOR := x ^ y

	for q := XOR; q > 0; q = q / 2 {
		m := q % 2
		if m == 1 {
			rst++
		}
	}

	return rst
}
