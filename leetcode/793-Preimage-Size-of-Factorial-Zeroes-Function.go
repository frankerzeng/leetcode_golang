/*
793. Preimage Size of Factorial Zeroes Function

Let f(x) be the number of zeroes at the end of x!. (Recall that x! = 1 * 2 * 3 * ... * x, and by convention, 0! = 1.)
For example, f(3) = 0 because 3! = 6 has no zeroes at the end, while f(11) = 2 because 11! = 39916800 has 2 zeroes at the end.
Given K, find how many non-negative integers x have the property that f(x) = K.

Example 1:
	Input: K = 0
	Output: 5
	Explanation: 0!, 1!, 2!, 3!, and 4! end with K = 0 zeroes.

Example 2:
	Input: K = 5
	Output: 0
	Explanation: There is no x such that x! ends in K = 5 zeroes.

Note:
	K will be an integer in the range [0, 10^9].

*/
package main

import "fmt"

func main() {
	fmt.Println(preimageSizeFZF(5))
}
func preimageSizeFZF(K int) int {
	current_k := 0
	current_num := 0
	rst_num := 1

	return_num := 0 // return result

	for true {
		fmt.Println(current_num)
		if current_num > 0 {
			rst_num = rst_num * current_num
		}
		fmt.Println("]]]]]]]]]]]]]]]]]", rst_num)

		if rst_num%10 == 0 {
			// rst_num = rst_num / 10
			current_k++
		}
		fmt.Println("========current_k", current_k)
		if current_k == 55 {
			fmt.Println("--------------")
			return_num++
		} else if current_k > 55 {
			return return_num
		}
		current_num++
	}
	return 10
}
