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
	fmt.Println(reverseVowels("hello"))
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
