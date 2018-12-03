/*
938. Range Sum of BST
Given the root node of a binary search tree, return the sum of values of all nodes with value between L and R (inclusive).

The binary search tree is guaranteed to have unique values.



Example 1:

Input: root = [10,5,15,3,7,null,18], L = 7, R = 15
Output: 32
Example 2:

Input: root = [10,5,15,3,7,13,18,1,null,6], L = 6, R = 10
Output: 23


Note:

The number of nodes in the tree is at most 10000.
The final answer is guaranteed to be less than 2^31.
*/
package main

import (
	"../Common"
	"fmt"
	"time"
)

type TreeNode = Common.TreeNode

func main() {
	root := &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 1,
				},
			},
			Right: &TreeNode{
				Val: 7,
				Left: &TreeNode{
					Val: 6,
				},
			},
		},
		Right: &TreeNode{
			Val: 15,
			Left: &TreeNode{
				Val: 13,
			},
			Right: &TreeNode{
				Val: 18,
			},
		},
	}
	L, R := 7, 15
	t := time.Now()
	for i := 0; i < 100001; i++ {
		rangeSumBST(root, L, R)
	}
	fmt.Println(time.Since(t))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rangeSumBST(root *TreeNode, L int, R int) int {
	var rstInt int

	tmpList := []*TreeNode{root}

	for true {
		if len(tmpList) == 0 {
			break
		}
		var tmpList_ []*TreeNode
		for _, v := range tmpList {
			if v.Val <= R && v.Val >= L {
				rstInt += v.Val
				if v.Left != nil {
					tmpList_ = append(tmpList_, v.Left)
				}
				if v.Right != nil {
					tmpList_ = append(tmpList_, v.Right)
				}
			} else if v.Val > R {
				if v.Left != nil {
					tmpList_ = append(tmpList_, v.Left)
				}
			} else {
				if v.Right != nil {
					tmpList_ = append(tmpList_, v.Right)
				}
			}
		}
		tmpList = tmpList_
	}

	return rstInt
}
