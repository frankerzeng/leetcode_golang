/*
881. Boats to Save People

The i-th person has weight people[i], and each boat can carry a maximum weight of limit.

Each boat carries at most 2 people at the same time, provided the sum of the weight of those people is at most limit.

Return the minimum number of boats to carry every given person.  (It is guaranteed each person can be carried by a boat.)

Example 1:
	Input: people = [1,2], limit = 3
	Output: 1
	Explanation: 1 boat (1, 2)

Example 2:
	Input: people = [3,2,2,1], limit = 3
	Output: 3
	Explanation: 3 boats (1, 2), (2) and (3)

Example 3:
	Input: people = [3,5,3,4], limit = 5
	Output: 4
	Explanation: 4 boats (3), (3), (4), (5)
Note:
	1 <= people.length <= 50000
	1 <= people[i] <= limit <= 30000
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	people := []int{2, 1, 4, 7, 3, 6}

	limit := 3
	ti := time.Now()
	fmt.Println(numRescueBoats(people, limit))
	fmt.Println(time.Since(ti))
}

func quickSort(values []int, left, right int) {
	temp := values[left]
	p := left
	i, j := left, right

	for i <= j {
		for j >= p && values[j] >= temp {
			j--
		}
		if j >= p {
			values[p] = values[j]
			p = j
		}

		if values[i] <= temp && i <= p {
			i++
		}
		if i <= p {
			values[p] = values[i]
			p = i
		}
	}
	values[p] = temp
	if p-left > 1 {
		quickSort(values, left, p-1)
	}
	if right-p > 1 {
		quickSort(values, p+1, right)
	}
}
func numRescueBoats(people []int, limit int) int {
	boats := 0

	// sort by weight

	// time limit
	//sortFun := func() {
	//	for i := 0; i < len(people); i++ {
	//		for j := i + 1; j < len(people); j++ {
	//			if people[i] > people[j] {
	//				people[i], people[j] = people[j], people[i]
	//			}
	//		}
	//	}
	//}
	//
	quickSort(people, 0, len(people)-1)

	maxIndex := len(people) - 1

	for i := 0; i <= maxIndex; i++ {
		if i == maxIndex {
			boats++
			break
		}
		for j := maxIndex; j > i; j-- {
			maxIndex--
			boats++
			if people[i]+people[j] > limit {
				if i == maxIndex {
					boats++
					return boats
				}
			} else {
				break
			}
		}
	}

	return boats
}
