/*
452-Minimum-Number-of-Arrows-to-Burst-Balloons
There are a number of spherical balloons spread in two-dimensional space. For each balloon, provided input is the start
and end coordinates of the horizontal diameter. Since it's horizontal, y-coordinates don't matter and hence the x-coordinates
of start and end of the diameter suffice. Start is always smaller than end. There will be at most 104 balloons.

An arrow can be shot up exactly vertically from different points along the x-axis. A balloon with xstart and xend bursts
by an arrow shot at x if xstart ≤ x ≤ xend. There is no limit to the number of arrows that can be shot. An arrow once
shot keeps travelling up infinitely. The problem is to find the minimum number of arrows that must be shot to burst all
balloons.

Example:

Input:
[[10,16], [2,8], [1,6], [7,12]]

Output:
2

Explanation:
One way is to shoot one arrow for example at x = 6 (bursting the balloons [2,8] and [1,6]) and another arrow at x = 11
(bursting the other two balloons).
*/
package main

import (
	"fmt"
	"sort"
)

func main() {
	points := [][]int{
		{0, 9}, {1, 8}, {7, 8}, {1, 6}, {9, 16}, {7, 13}, {7, 10}, {6, 11}, {6, 9}, {9, 13},
	}
	fmt.Println(findMinArrowShots(points))

}

type intInterface [][]int

func (intInter intInterface) Len() int {
	return len(intInter)
}
func (intInter intInterface) Swap(i, j int) {
	intInter[i], intInter[j] = intInter[j], intInter[i]
}
func (intInter intInterface) Less(i, j int) bool {
	return intInter[i][1] < intInter[j][1]
}
func findMinArrowShots(points intInterface) int {
	var rst = 1
	if len(points) == 0 {
		return 0
	}
	sort.Sort(points)
	start := points[0][1]
	for i := 1; i < len(points); i++ {
		if points[i][0] > start {
			rst++
			start = points[i][1]
		}
	}

	return rst
}
