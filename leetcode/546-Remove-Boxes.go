/*
546. Remove Boxes
Given several boxes with different colors represented by different positive numbers.
You may experience several rounds to remove boxes until there is no box left. Each time you can choose some continuous
boxes with the same color (composed of k boxes, k >= 1), remove them and get k*k points.
Find the maximum points you can get.

Example 1:
Input:

[1, 3, 2, 2, 2, 3, 4, 3, 1]
Output:
23
Explanation:
[1, 3, 2, 2, 2, 3, 4, 3, 1]
----> [1, 3, 3, 4, 3, 1] (3*3=9 points)
----> [1, 3, 3, 3, 1] (1*1=1 points)
----> [1, 1] (3*3=9 points)
----> [] (2*2=4 points)
Note: The number of boxes n would not exceed 100.
*/
package main

import "fmt"

func main() {
	fmt.Println(removeBoxes([]int{1, 3, 2, 2, 2, 3, 4, 3, 1}))
}

type CBox struct {
	color, count int
}

func removeBoxes(boxes []int) int {
	if len(boxes) == 0 {
		return 0
	}
	cboxes := []CBox{}
	for i, c := range boxes {
		if i == 0 || c != boxes[i-1] {
			cboxes = append(cboxes, CBox{c, 1})
		} else {
			cboxes[len(cboxes)-1].count++
		}
	}
	// fmt.Printf("%+v\n", cboxes)
	return removeBoxes_helper(cboxes)
}

func removeBoxes_helper(cboxes []CBox) int {
	fmt.Println(cboxes)
	if len(cboxes) == 0 {
		return 0
	}
	lastBox := cboxes[len(cboxes)-1]
	cboxes = cboxes[:len(cboxes)-1]
	var ans int = lastBox.count*lastBox.count + removeBoxes_helper(cboxes)
	for i := range cboxes {
		if cboxes[i].color == lastBox.color {
			cboxes[i].count += lastBox.count
			ansN := removeBoxes_helper(cboxes[:i+1]) + removeBoxes_helper(cboxes[i+1:]) // 最后一个合并过来+剩下的
			if ansN > ans {
				ans = ansN
			}
			cboxes[i].count -= lastBox.count
		}
	}
	return ans
}
