// 结构体
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	var node121 *TreeNode = new(TreeNode)
	var node122 *TreeNode = new(TreeNode)
	var node12 *TreeNode = new(TreeNode)
	var node11 *TreeNode = new(TreeNode)
	var node1 *TreeNode = new(TreeNode)
	node121.Val = 15
	node122.Val = 7
	node12.Val = 20
	node12.Left = node121
	node12.Right = node122
	node11.Val = 9
	node1.Val = 3
	node1.Left = node11
	node1.Right = node12

	zigzagLevelOrder(node1)
}
func zigzagLevelOrder(root *TreeNode) [][]int {
	var returnD [][]int
	var currentNods, currentNodsTmp []*TreeNode
	currentNods = append(currentNods, root)
	var returnD_child []int

	if root != nil {
		for true {
			if len(currentNods) == 0 {
				return returnD
			}
			currentNodsTmp = currentNods
			currentNods = currentNods[:0]

			if len(returnD_child) > 0 {
				returnD_child = returnD_child[:0]
			}

			if len(currentNodsTmp)>1 {
				fmt.Println(currentNodsTmp[1])
			}
			for _, v := range currentNodsTmp {
				fmt.Println(v)

				returnD_child = append(returnD_child, v.Val)

				if v.Right != nil {
					currentNods = append(currentNods, v.Right)
				}

				if v.Left != nil {
					currentNods = append(currentNods, v.Left)
				}

			}

			if len(returnD_child) > 0 {
				returnD = append(returnD, returnD_child)
			}
			fmt.Println(returnD)
			fmt.Println("----------------------------returnD_child")
		}
	}
	return returnD
}
