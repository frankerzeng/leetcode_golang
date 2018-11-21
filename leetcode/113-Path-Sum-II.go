/*
113. Path Sum II
Given a binary tree and a sum, find all root-to-leaf paths where each path's sum equals the given sum.

Note: A leaf is a node with no children.

Example:
	Given the below binary tree and sum = 22,

		  5
		 / \
		4   8
	   /   / \
	  11  13  4
	 /  \    / \
	7    2  5   1

Return:
[
   [5,4,11,2],
   [5,8,4,5]
]
*/
package main

import (
	"../Common"
	"fmt"
)

type TreeNode = Common.TreeNode

func main() {
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 11,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 13,
			},
			Right: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 5,
				},
				Right: &TreeNode{
					Val: 1,
				},
			},
		},
	}
	sum := 22

	// root = &TreeNode{}
	// sum = 1
	fmt.Println(pathSum(root, sum))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, sum int) [][]int {
	var rst [][]int
	if root != nil {
		if root.Right == nil && root.Left == nil {
			if sum == root.Val {
				rst = append(rst, []int{root.Val})
			}
		} else {
			var tmp [][]int
			tmp = pathSum(root.Right, sum-root.Val)
			if len(tmp) > 0 {
				for _, v := range tmp {
					rst = append(rst, append([]int{root.Val}, v...))
				}
			}
			tmp = pathSum(root.Left, sum-root.Val)
			if len(tmp) > 0 {
				for _, v := range tmp {
					rst = append(rst, append([]int{root.Val}, v...))
				}
			}
		}
	}
	return rst
}
