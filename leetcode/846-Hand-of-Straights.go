/*
846. Hand of Straights
Alice has a hand of cards, given as an array of integers.

Now she wants to rearrange the cards into groups so that each group is size W, and consists of W consecutive cards.

Return true if and only if she can.



Example 1:

Input: hand = [1,2,3,6,2,3,4,7,8], W = 3
Output: true
Explanation: Alice's hand can be rearranged as [1,2,3],[2,3,4],[6,7,8].
Example 2:

Input: hand = [1,2,3,4,5], W = 4
Output: false
Explanation: Alice's hand can't be rearranged into groups of 4.


Note:

1 <= hand.length <= 10000
0 <= hand[i] <= 10^9
1 <= W <= hand.length
*/
package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(isNStraightHand([]int{1, 2}, 2))

}
func isNStraightHand(hand []int, W int) bool {
	if len(hand)%W != 0 {
		return false
	}
	sort.Ints(hand)

	// [['加入个数','当前值']]
	alice := make([][]int, len(hand)/W)
	for k := range alice {
		alice[k] = []int{0, 0}
	}

	for i := 0; i < len(hand); i++ {
		flag := false // 是否使用了hand[i]
		for k := range alice {
			if alice[k][0] == 0 {
				alice[k][0]++
				alice[k][1] = hand[i]
				flag = true
				break
			}
			if alice[k][0] == W {
				continue
			}
			if alice[k][1]+1 == hand[i] {
				alice[k][0]++
				alice[k][1] = hand[i]
				flag = true
				break
			}
		}

		if !flag {
			return false
		}
	}

	return true
}
