package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// 测试数据
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

	fmt.Println(zigzagLevelOrder(node1))
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	var returnD [][]int
	var currentNods, currentNodsTmp []*TreeNode
	currentNods = append(currentNods, root)
	var returnChild []int
	right := true // 取值方向，从左到右
	if root != nil {
		for true {
			if len(currentNods) == 0 {
				return returnD
			}
			currentNodsTmp = make([]*TreeNode, len(currentNods)) // 初始化清空切片
			returnChild = make([]int, 0)                         // 初始化清空切片
			copy(currentNodsTmp, currentNods)
			currentNods = currentNods[:0] //currentNods 当前层的所有节点
			for _, v := range currentNodsTmp {
				returnChild = append(returnChild, v.Val)
				if v.Right != nil {
					currentNods = append(currentNods, v.Right)
				}
				if v.Left != nil {
					currentNods = append(currentNods, v.Left)
				}
			}
			if len(returnChild) > 0 {
				if right { // 反转切片
					for i, j := 0, len(returnChild)-1; i < j; i, j = i+1, j-1 {
						returnChild[i], returnChild[j] = returnChild[j], returnChild[i]
					}
				}
				returnD = append(returnD, returnChild)
			}
			right = !right
		}
	}
	return returnD
}
