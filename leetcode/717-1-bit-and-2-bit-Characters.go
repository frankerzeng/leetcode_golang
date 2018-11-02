/*
717. 1-bit and 2-bit Characters
We have two special characters. The first character can be represented by one bit 0. The second character can be
represented by two bits (10 or 11).

Now given a string represented by several bits. Return whether the last character must be a one-bit character or not.
he given string will always end with a zero.

Example 1:
Input:
	bits = [1, 0, 0]
Output:
	True
Explanation:
	The only way to decode it is two-bit character and one-bit character. So the last character is one-bit character.

Example 2:
Input:
	bits = [1, 1, 1, 0]
Output:
	False
Explanation:
	The only way to decode it is two-bit character and two-bit character. So the last character is NOT one-bit character.

Note:
1 <= len(bits) <= 1000.
bits[i] is always 0 or 1.
*/
package main

import (
	"fmt"
)

func main() {
	bits := []int{0}
	fmt.Println(isOneBitCharacter(bits))

}
func isOneBitCharacter(bits []int) bool {
	if bits[len(bits)-1] == 0 {
		return canEnd(bits[0 : len(bits)-1])
	}
	return false
}

func canEnd(bits []int) bool {
	bitsLen := len(bits)
	if bitsLen == 0 {
		return true
	} else if bitsLen == 1 {
		if bits[0] == 0 {
			return true
		}
	} else if bitsLen == 2 && bits[0] == 1 {
		return true
	} else {
		if bits[bitsLen-1] == 0 && canEnd(bits[0:bitsLen-1]) {
			return true
		}
		if bits[bitsLen-2] == 1 && canEnd(bits[0:bitsLen-2]) {
			return true
		}
	}
	return false
}
