/*
32. Longest Valid Parentheses
Given a string containing just the characters '(' and ')', find the length of the longest valid (well-formed) parentheses substring.

Example 1:

Input: "(()"
Output: 2
Explanation: The longest valid parentheses substring is "()"
Example 2:

Input: ")()())"
Output: 4
Explanation: The longest valid parentheses substring is "()()"
*/
package main

import "fmt"

func main() {
	fmt.Println(longestValidParentheses("()()"))
}
func longestValidParentheses(s string) int {
	var rst int
	fmt.Println("enter")
	var f string
	var k string
	fmt.Scanln(&f, &k)
	fmt.Scanln(&f, &k)
	fmt.Println(f, k)

	return rst
}
