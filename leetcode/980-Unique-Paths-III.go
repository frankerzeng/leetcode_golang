/*
980. Unique Paths III
On a 2-dimensional grid, there are 4 types of squares:

1 represents the starting square.  There is exactly one starting square.
2 represents the ending square.  There is exactly one ending square.
0 represents empty squares we can walk over.
-1 represents obstacles that we cannot walk over.
Return the number of 4-directional walks from the starting square to the ending square, that walk over every non-obstacle square exactly once.



Example 1:

Input: [[1,0,0,0],[0,0,0,0],[0,0,2,-1]]
Output: 2
Explanation: We have the following two paths:
1. (0,0),(0,1),(0,2),(0,3),(1,3),(1,2),(1,1),(1,0),(2,0),(2,1),(2,2)
2. (0,0),(1,0),(2,0),(2,1),(1,1),(0,1),(0,2),(0,3),(1,3),(1,2),(2,2)
Example 2:

Input: [[1,0,0,0],[0,0,0,0],[0,0,0,2]]
Output: 4
Explanation: We have the following four paths:
1. (0,0),(0,1),(0,2),(0,3),(1,3),(1,2),(1,1),(1,0),(2,0),(2,1),(2,2),(2,3)
2. (0,0),(0,1),(1,1),(1,0),(2,0),(2,1),(2,2),(1,2),(0,2),(0,3),(1,3),(2,3)
3. (0,0),(1,0),(2,0),(2,1),(2,2),(1,2),(1,1),(0,1),(0,2),(0,3),(1,3),(2,3)
4. (0,0),(1,0),(2,0),(2,1),(1,1),(0,1),(0,2),(0,3),(1,3),(1,2),(2,2),(2,3)
Example 3:

Input: [[0,1],[2,0]]
Output: 0
Explanation:
There is no path that walks over every empty square exactly once.
Note that the starting and ending square can be anywhere in the grid.


Note:

1 <= grid.length * grid[0].length <= 20
*/
package main

import "fmt"

func main() {
	fmt.Println(uniquePathsIII([][]int{{1, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 2, -1}}))
}

func uniquePathsIII(grid [][]int) int {
	step := 1 // from 1 to 2 at least 1 step
	x, y := 0, 0
	for i := 0; i < len(grid); i++ {
		for ii := 0; ii < len(grid[0]); ii++ {
			if grid[i][ii] == 0 {
				step++
			}
			if grid[i][ii] == 1 {
				x, y = i, ii // start cell
			}
		}
	}
	Grid = grid
	return dfsFunc(x, y, step)
}

var Grid [][]int

func dfsFunc(i, j, step int) int {
	if i < 0 || j < 0 || i >= len(Grid) || j >= len(Grid[0]) || Grid[i][j] == -1 {
		return 0
	}
	if Grid[i][j] == 2 {
		if step == 0 {
			return 1
		}
		return 0
	}

	rst := 0
	Grid[i][j] = -1                                               //  don't go back
	for _, v := range [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} { // From four directions
		rst += dfsFunc(i+v[0], j+v[1], step-1)
	}
	Grid[i][j] = 0
	return rst
}
