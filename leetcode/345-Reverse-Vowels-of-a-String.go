/*
345. Reverse Vowels of a String
Write a function that takes a string as input and reverse only the vowels of a string.

Example 1:

Input: "hello"
Output: "holle"
Example 2:

Input: "leetcode"
Output: "leotcede"
Note:
The vowels does not include the letter "y".
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(reverseVowels1("hello"))
}

func reverseVowels(s string) string {
	// 元音 a e i o u
	var reverseIndex []int
	for i := 0; i < len(s); i++ {
		if strings.Contains("AEIOUaeiou", string(s[i])) {
			reverseIndex = append(reverseIndex, i)
		}
	}

	lenReverseIndex := len(reverseIndex)
	for i := 0; i < lenReverseIndex/2; i++ {
		s = s[0:reverseIndex[i]] + string(s[reverseIndex[lenReverseIndex-i-1]]) + s[reverseIndex[i]+1:reverseIndex[lenReverseIndex-i-1]] +
			string(s[reverseIndex[i]]) + s[reverseIndex[lenReverseIndex-i-1]+1:]
	}

	return s
}
func reverseVowels1(s string) string {
	left := -1
	right := len(s)
	exchange := -1
	for true {
		exchange = -1
		for i := left + 1; i < right; i++ {
			left = i
			if strings.Contains("AEIOUaeiou", string(s[i])) {
				exchange++
				break
			}
		}
		for i := right - 1; i > left; i-- {
			right = i
			if strings.Contains("AEIOUaeiou", string(s[i])) {
				exchange++
				break
			}
		}

		if exchange == 1 {
			s = s[0:left] + string(s[right]) + s[left+1:right] + string(s[left]) + s[right+1:]
		}

		if left == right || (left+1) == right {
			break
		}
	}

	return s
}
