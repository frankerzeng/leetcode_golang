/*
849. Maximize Distance to Closest Person

In a row of seats, 1 represents a person sitting in that seat, and 0 represents that the seat is empty.

There is at least one empty seat, and at least one person sitting.

Alex wants to sit in the seat such that the distance between him and the closest person to him is maximized.

Return that maximum distance to closest person.

Example 1:

Input: [1,0,0,0,1,0,1]
Output: 2
Explanation:
If Alex sits in the second open seat (seats[2]), then the closest person has distance 2.
If Alex sits in any other open seat, the closest person has distance 1.
Thus, the maximum distance to the closest person is 2.
Example 2:

Input: [1,0,0,0]
Output: 3
Explanation:
If Alex sits in the last seat, the closest person is 3 seats away.
This is the maximum distance possible, so the answer is 3.
Note:

1 <= seats.length <= 20000
seats contains only 0s or 1s, at least one 0, and at least one 1.
*/
package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxDistToClosest([]int{1, 0, 0, 0}))
}

func maxDistToClosest(seats []int) int {
	var tmp int
	var needDouble = false
	var max int
	for k, v := range seats {
		if k == 0 && v == 0 {
			needDouble = true
		}
		if v == 0 {
			tmp++
		} else {
			tmp = 0
			needDouble = false
		}

		if tmp > 0 {
			if needDouble || (k == len(seats)-1 && v == 0) { // 头尾的0单独计算长度
				if tmp > max {
					max = tmp
				}
			} else {
				if ((tmp-1)/2)+1 > max {
					max = (tmp + 1) / 2
				}
			}
		}
	}

	return max
}
