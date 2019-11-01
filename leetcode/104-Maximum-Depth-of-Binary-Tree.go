/*
104. Maximum Depth of Binary Tree
Given a binary tree, find its maximum depth.

The maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.

Note: A leaf is a node with no children.

Example:

Given binary tree [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
*/
package main

import "fmt"

func main() {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 10,
			},
		},
	}
	fmt.Println(maxDepth(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxLeft := maxDepth(root.Left)
	maxRight := maxDepth(root.Right)
	if maxLeft > maxRight {
		return 1 + maxLeft
	}
	return 1 + maxRight
}
