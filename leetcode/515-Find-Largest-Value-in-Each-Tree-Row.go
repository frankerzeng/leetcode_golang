/*
515. Find Largest Value in Each Tree Row
You need to find the largest value in each row of a binary tree.

Example:
Input:
          1
         / \
        3   2
       / \   \
      5   3   9

Output:
	[1, 3, 9]
*/
package main

import "fmt"
import "../Common"

type TreeNode = Common.TreeNode

func main() {
	root := new(TreeNode)
	root = nil
	fmt.Println(largestValues(root))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func largestValues(root *TreeNode) []int {
	var rst []int
	if root == nil {
		return rst
	}
	listNode := []*TreeNode{root}

	for true {
		if len(listNode) == 0 {
			break
		}
		max := listNode[0].Val
		var nextListNode []*TreeNode
		for _, v := range listNode {
			if v.Val > max {
				max = v.Val
			}
			if v.Left != nil {
				nextListNode = append(nextListNode, v.Left)
			}
			if v.Right != nil {
				nextListNode = append(nextListNode, v.Right)
			}
		}
		listNode = nextListNode
		rst = append(rst, max)
	}

	return rst
}
