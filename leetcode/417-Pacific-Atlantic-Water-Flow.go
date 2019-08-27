/*
417. Pacific Atlantic Water Flow
Given an m x n matrix of non-negative integers representing the height of each unit cell in a continent, the "Pacific ocean"
touches the left and top edges of the matrix and the "Atlantic ocean" touches the right and bottom edges.

Water can only flow in four directions (up, down, left, or right) from a cell to another one with height equal or lower.

Find the list of grid coordinates where water can flow to both the Pacific and Atlantic ocean.

Note:

The order of returned grid coordinates does not matter.
Both m and n are less than 150.


Example:

Given the following 5x5 matrix:

  Pacific ~   ~   ~   ~   ~
       ~  1   2   2   3  (5) *
       ~  3   2   3  (4) (4) *
       ~  2   4  (5)  3   1  *
       ~ (6) (7)  1   4   5  *
       ~ (5)  1   1   2   4  *
          *   *   *   *   * Atlantic

Return:

[[0, 4], [1, 3], [1, 4], [2, 2], [3, 0], [3, 1], [4, 0]] (positions with parentheses in above matrix).
*/
package main

import "fmt"

func main() {
	matrix := [][]int{
		{1, 2, 2, 3, 5},
		{3, 2, 3, 4, 4},
		{2, 4, 5, 3, 1},
		{6, 7, 1, 4, 5},
		{5, 1, 1, 2, 4},
	}
	matrix = [][]int{
		{1, 1},
		{1, 1},
		{1, 1},
	}
	matrix = [][]int{
		{3, 3, 3, 3, 3, 3},
		{3, 0, 3, 3, 0, 3},
		{3, 3, 3, 3, 3, 3},
	}
	matrix = [][]int{
		{1, 2, 3, 4},
		{12, 13, 14, 5},
		{11, 16, 15, 6},
		{10, 9, 8, 7}}

	rst := pacificAtlantic(matrix)
	fmt.Println(rst)
}

/*
思路：
某个点是否可以通过自己上下左右的路径走过不高于自己的点，从而到达上边或左边同时到达下边或右边，如流水流动的路径
广度优先法
*/
func pacificAtlantic(matrix [][]int) [][]int {
	var rst [][]int
	jMax := len(matrix) // 行
	if jMax == 0 {
		return rst
	}
	iMax := len(matrix[0]) // 列

	// j,i 点通过 向上或向左 是否能走到上边界和左边界
	// j,i 点通过 向下或向右 是否能走到下边界和右边界
	// canFlow[j][i]=[pacific,atlantic] ，1 表示可以，0 表示没有配置
	var canFlow [][][]int

	for i := 0; i < jMax; i++ {
		var tmp [][]int
		for j := 0; j < iMax; j++ {
			tmp = append(tmp, []int{0, 0})
		}
		canFlow = append(canFlow, tmp)
	}

	for i := 0; i < jMax; i++ {
		canFlow[i][0][0] = 1
		canFlow[i][iMax-1][1] = 1
	}
	for i := 0; i < iMax; i++ {
		canFlow[0][i][0] = 1
		canFlow[jMax-1][i][1] = 1
	}

	var checkCell [][]int // 由这些点开始向外蔓延
	for i := 0; i < jMax; i++ {
		checkCell = append(checkCell, []int{i, 0})
	}
	for i := 1; i < iMax; i++ {
		checkCell = append(checkCell, []int{0, i})
	}

	checkFlag := true // 是否可以继续查找
	for checkFlag {   // 先查找能流到pacific的点
		checkFlag = false
		for _, v := range checkCell {
			if canFlow[v[0]][v[1]][0] == 1 {
				if v[0]+1 < jMax && matrix[v[0]+1][v[1]] >= matrix[v[0]][v[1]] && canFlow[v[0]+1][v[1]][0] == 0 { // 向下能否流动
					appendFlag := true
					for _, vc := range checkCell {
						if vc[0] == v[0]+1 && vc[1] == v[1] {
							appendFlag = false
							break
						}
					}
					if appendFlag {
						checkCell = append(checkCell, []int{v[0] + 1, v[1]})
					}
					canFlow[v[0]+1][v[1]][0] = 1
					checkFlag = true
				}
				if v[0] > 0 && matrix[v[0]-1][v[1]] >= matrix[v[0]][v[1]] && canFlow[v[0]-1][v[1]][0] == 0 { // 向上能否流动
					appendFlag := true
					for _, vc := range checkCell {
						if vc[0] == v[0]-1 && vc[1] == v[1] {
							appendFlag = false
							break
						}
					}
					if appendFlag {
						checkCell = append(checkCell, []int{v[0] - 1, v[1]})
					}
					canFlow[v[0]-1][v[1]][0] = 1
					checkFlag = true
				}
				if v[1]+1 < iMax && matrix[v[0]][v[1]+1] >= matrix[v[0]][v[1]] && canFlow[v[0]][v[1]+1][0] == 0 { // 向右能否流动
					appendFlag := true
					for _, vc := range checkCell {
						if vc[0] == v[0] && vc[1] == v[1]+1 {
							appendFlag = false
							break
						}
					}
					if appendFlag {
						checkCell = append(checkCell, []int{v[0], v[1] + 1})
					}
					canFlow[v[0]][v[1]+1][0] = 1
					checkFlag = true
				}
				if v[1] > 0 && matrix[v[0]][v[1]-1] >= matrix[v[0]][v[1]] && canFlow[v[0]][v[1]-1][0] == 0 { // 向左能否流动
					appendFlag := true
					for _, vc := range checkCell {
						if vc[0] == v[0] && vc[1] == v[1]-1 {
							appendFlag = false
							break
						}
					}
					if appendFlag {
						checkCell = append(checkCell, []int{v[0], v[1] - 1})
					}
					canFlow[v[0]][v[1]-1][0] = 1
					checkFlag = true
				}
			}
		}

	}

	checkCell = [][]int{}
	for i := 0; i < jMax-1; i++ {
		checkCell = append(checkCell, []int{i, iMax - 1})
	}
	for i := 0; i < iMax; i++ {
		checkCell = append(checkCell, []int{jMax - 1, i})
	}

	checkFlag = true
	for checkFlag {
		checkFlag = false

		for _, v := range checkCell {
			if canFlow[v[0]][v[1]][1] == 1 {
				if v[0]+1 < jMax && matrix[v[0]+1][v[1]] >= matrix[v[0]][v[1]] && canFlow[v[0]+1][v[1]][1] == 0 {
					appendFlag := true
					for _, vc := range checkCell {
						if vc[0] == v[0]+1 && vc[1] == v[1] {
							appendFlag = false
							break
						}
					}
					if appendFlag {
						checkCell = append(checkCell, []int{v[0] + 1, v[1]})
					}
					canFlow[v[0]+1][v[1]][1] = 1
					checkFlag = true
				}
				if v[0] > 0 && matrix[v[0]-1][v[1]] >= matrix[v[0]][v[1]] && canFlow[v[0]-1][v[1]][1] == 0 {
					appendFlag := true
					for _, vc := range checkCell {
						if vc[0] == v[0]-1 && vc[1] == v[1] {
							appendFlag = false
							break
						}
					}
					if appendFlag {
						checkCell = append(checkCell, []int{v[0] - 1, v[1]})
					}
					canFlow[v[0]-1][v[1]][1] = 1
					checkFlag = true
				}
				if v[1]+1 < iMax && matrix[v[0]][v[1]+1] >= matrix[v[0]][v[1]] && canFlow[v[0]][v[1]+1][1] == 0 {
					appendFlag := true
					for _, vc := range checkCell {
						if vc[0] == v[0] && vc[1] == v[1]+1 {
							appendFlag = false
							break
						}
					}
					if appendFlag {
						checkCell = append(checkCell, []int{v[0], v[1] + 1})
					}
					canFlow[v[0]][v[1]+1][1] = 1
					checkFlag = true
				}
				if v[1] > 0 && matrix[v[0]][v[1]-1] >= matrix[v[0]][v[1]] && canFlow[v[0]][v[1]-1][1] == 0 {
					appendFlag := true
					for _, vc := range checkCell {
						if vc[0] == v[0] && vc[1] == v[1]-1 {
							appendFlag = false
							break
						}
					}
					if appendFlag {
						checkCell = append(checkCell, []int{v[0], v[1] - 1})
					}
					canFlow[v[0]][v[1]-1][1] = 1
					checkFlag = true
				}
			}
		}

		for _, v := range checkCell {
			if (!(v[0]+1 < jMax) || v[0]+1 < jMax && canFlow[v[0]+1][v[1]][1] == 1) &&
				(!(v[0] > 0) || v[0] > 0 && canFlow[v[0]-1][v[1]][1] == 1) &&
				(!(v[1]+1 < iMax) || v[1]+1 < iMax && canFlow[v[0]][v[1]+1][1] == 1) &&
				(!(v[1] > 0) || v[1] > 0 && canFlow[v[0]][v[1]-1][1] == 1) {
				for k, vv := range checkCell {
					if vv[0] == v[0] && vv[1] == v[1] {
						checkCell = append(checkCell[0:k], checkCell[k+1:]...)
						break
					}
				}
			}
		}
	}

	// for _, v := range canFlow {
	// 	fmt.Println(v)
	// }
	// fmt.Println("--------------", checkFlag)

	for i, v := range canFlow {
		for j, vv := range v {
			if vv[0] == 1 && vv[1] == 1 {
				rst = append(rst, []int{i, j})
			}
		}
	}

	return rst
}
