/*
703-Kth-Largest-Element-in-a-Stream

Design a class to find the kth largest element in a stream. Note that it is the kth largest element in the sorted order, not the kth distinct element.

Your KthLargest class will have a constructor which accepts an integer k and an integer array nums, which contains initial elements from the stream. For each call to the method KthLargest.add, return the element representing the kth largest element in the stream.

Example:

int k = 3;
int[] arr = [4,5,8,2];
KthLargest kthLargest = new KthLargest(3, arr);
kthLargest.add(3);   // returns 4
kthLargest.add(5);   // returns 5
kthLargest.add(10);  // returns 5
kthLargest.add(9);   // returns 8
kthLargest.add(4);   // returns 8
Note:
You may assume that nums' length ≥ k-1 and k ≥ 1.
*/
package main

import (
	"fmt"
	"sort"
)

func main() {
	obj := Constructor(3, []int{4, 5, 8, 2})
	fmt.Println(obj.Add(3))
	fmt.Println(obj.Add(5))
	fmt.Println(obj.Add(10))
	fmt.Println(obj.Add(9))
	fmt.Println(obj.Add(4))
}

type KthLargest struct {
	arrSortKth []int
	kth        int
}

func Constructor(k int, nums []int) KthLargest {
	sort.Ints(nums)
	return KthLargest{nums, k}
}

func (this *KthLargest) Add(val int) int {
	if len(this.arrSortKth) == 0 {
		this.arrSortKth = append(this.arrSortKth, val)
	} else {
		// 二分法查找
		binarySearch(&this.arrSortKth, val)
	}
	if len(this.arrSortKth) > this.kth {
		this.arrSortKth = this.arrSortKth[len(this.arrSortKth)-this.kth:]
	}
	return this.arrSortKth[len(this.arrSortKth)-this.kth]
}
func binarySearch(s *[]int, k int) {
	if len(*s) == 1 {
		if (*s)[0] > k {
			*s = append([]int{k}, *s...)
		} else {
			*s = append(*s, k)
		}
		return
	}
	if (*s)[0] >= k {
		*s = append([]int{k}, *s...)
		return
	}
	if (*s)[len(*s)-1] <= k {
		*s = append(*s, k)
		return
	}

	lo, hi := 0, len(*s)-1
	for lo+1 != hi {
		m := (lo + hi) >> 1
		if (*s)[m] < k {
			lo = m
		} else if (*s)[m] > k {
			hi = m
		} else {
			tmpS := make([]int, len(*s))
			copy(tmpS, *s)
			tmp := append(tmpS[0:m], k)
			*s = append(tmp, (*s)[m:]...)
			return
		}
	}
	tmpS := make([]int, len(*s))
	copy(tmpS, *s)
	tmp := append(tmpS[0:hi], k)
	*s = append(tmp, (*s)[hi:]...)
	return
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
