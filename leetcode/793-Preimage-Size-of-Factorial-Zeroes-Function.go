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

import (
	"fmt"
	"time"
)

func main() {
	timeTmp := time.Now()
	fmt.Println(preimageSizeFZF(80502705))
	fmt.Println(time.Since(timeTmp))
}

func f(x int) int {
	k := 0
	for {
		x = x / 5
		if x == 0 {
			break
		}
		k = k + x
	}
	return k
}

func preimageSizeFZF(K int) int {
	var mid int
	left := K
	right := K * 5
	for left <= right {
		mid = left + (right-left)/2
		tmp := f(mid)
		if tmp < K {
			left = mid + 1
		} else if tmp > K {
			right = mid - 1
		} else {
			return 5
		}
	}
	return 0
}
