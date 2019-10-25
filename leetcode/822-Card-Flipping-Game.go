/*
822. Card Flipping Game
On a table are N cards, with a positive integer printed on the front and back of each card (possibly different).

We flip any number of cards, and after we choose one card.

If the number X on the back of the chosen card is not on the front of any card, then this number X is good.

What is the smallest number that is good?  If no number is good, output 0.

Here, fronts[i] and backs[i] represent the number on the front and back of card i.

A flip swaps the front and back numbers, so the value on the front is now on the back and vice versa.

Example:

Input: fronts = [1,2,4,4,7], backs = [1,3,4,1,3]
Output: 2
Explanation: If we flip the second card, the fronts are [1,3,4,4,7] and the backs are [1,2,4,1,3].
We choose the second card, which has number 2 on the back, and it isn't on the front of any card, so 2 is good.
*/
package main

import (
	"fmt"
)

func main() {
	fmt.Println(flipgame([]int{1, 1}, []int{2, 1}))
}

func flipgame(fronts []int, backs []int) int {
	var rst int

	if len(fronts) == 0 {
		return rst
	}

	minMap := make(map[int]int)    // 候选的最小值
	doubleMap := make(map[int]int) // 两面都是一样数字的排除掉
	for k := range fronts {
		if fronts[k] != backs[k] {
			if _, ok := doubleMap[fronts[k]]; !ok {
				minMap[fronts[k]] = fronts[k]
			}
			if _, ok := doubleMap[backs[k]]; !ok {
				minMap[backs[k]] = backs[k]
			}
		} else {
			doubleMap[fronts[k]] = fronts[k]
			if _, ok := minMap[fronts[k]]; ok {
				delete(minMap, fronts[k])
			}
		}
	}

	for _, v := range minMap {
		if rst == 0 || rst > v {
			rst = v
		}
	}

	return rst
}
