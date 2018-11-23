/*
572. Subtree of Another Tree
Given two non-empty binary trees s and t, check whether tree t has exactly the same structure and node values with a
subtree of s. A subtree of s is a tree consists of a node in s and all of this node's descendants. The tree s could also
be considered as a subtree of itself.

Example 1:
Given tree s:
     3
    / \
   4   5
  / \
 1   2
Given tree t:
   4
  / \
 1   2

Return true, because t has the same structure and node values with a subtree of s.

Example 2:
Given tree s:
     3
    / \
   4   5
  / \
 1   2
    /
   0
Given tree t:
   4
  / \
 1   2
Return false.

*/
package main

import (
	"../Common"
	"fmt"
	"strconv"
	"strings"
)

type TreeNode = Common.TreeNode

func main() {
	s := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   0,
					Left:  nil,
					Right: nil,
				},
				Right: nil,
			},
			Right: &TreeNode{
				Val:   2,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   5,
			Left:  nil,
			Right: nil,
		},
	}
	t := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
	}

	fmt.Println(isSubtree(s, t))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubtree(s *TreeNode, t *TreeNode) bool {
	return strings.Contains(leftTraversing(s, true), leftTraversing(t, true))
}

// left traversing binary tree
func leftTraversing(s *TreeNode, isLeft bool) string {
	if s == nil {
		if isLeft {
			return "null(L)"
		} else {
			return "null(R)"
		}
	}
	return "Node:" + strconv.Itoa(s.Val) + leftTraversing(s.Left, true) + leftTraversing(s.Right, false)
}
