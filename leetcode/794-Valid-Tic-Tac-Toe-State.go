/*
794. Valid Tic-Tac-Toe State
A Tic-Tac-Toe board is given as a string array board. Return True if and only if it is possible to reach this board position during the course of a valid tic-tac-toe game.

The board is a 3 x 3 array, and consists of characters " ", "X", and "O".  The " " character represents an empty square.

Here are the rules of Tic-Tac-Toe:

Players take turns placing characters into empty squares (" ").
The first player always places "X" characters, while the second player always places "O" characters.
"X" and "O" characters are always placed into empty squares, never filled ones.
The game ends when there are 3 of the same (non-empty) character filling any row, column, or diagonal.
The game also ends if all squares are non-empty.
No more moves can be played if the game is over.
Example 1:
Input: board = ["O  ", "   ", "   "]
Output: false
Explanation: The first player always plays "X".

Example 2:
Input: board = ["XOX", " X ", "   "]
Output: false
Explanation: Players take turns making moves.

Example 3:
Input: board = ["XXX", "   ", "OOO"]
Output: false

Example 4:
Input: board = ["XOX", "O O", "XOX"]
Output: true
Note:

board is a length-3 array of strings, where each string board[i] has length 3.
Each board[i][j] is a character in the set {" ", "X", "O"}.
*/
package main

import "fmt"

func main() {
	board := []string{"O  ", "   ", "   "}
	fmt.Println(validTicTacToe(board))
	board = []string{"XOX", " X ", "   "}
	fmt.Println(validTicTacToe(board))
	board = []string{"XXX", "   ", "OOO"}
	fmt.Println(validTicTacToe(board))
	board = []string{"XOX", "O O", "XOX"}
	fmt.Println(validTicTacToe(board))
}

func validTicTacToe(board []string) bool {

	oCount, xCount := 0, 0
	for _, v := range board {
		for i := 0; i < 3; i++ {
			if v[i] == 'X' {
				xCount++
			} else if v[i] == 'O' {
				oCount++
			}
		}
	}
	if !(xCount == oCount || xCount == oCount+1) {
		return false // x先手所以，x等于o或x等于o加一
	}

	rstO := threeFunc(board, 'O')
	rstX := threeFunc(board, 'X')

	// 不能存在x，o同时满足结束条件，连续三个子
	if rstO && rstX {
		return false
	}

	if rstO && xCount != oCount { // O赢的时候必须是最后一个下子
		return false
	}
	if rstX && xCount != oCount+1 { // X赢的时候必须是最后一个下子
		return false
	}

	return true
}

func threeFunc(board []string, t byte) bool {
	for i := 0; i < 3; i++ {
		if board[i][0] == t && board[i][1] == t && board[i][2] == t {
			return true
		}
		if board[0][i] == t && board[1][i] == t && board[2][i] == t {
			return true
		}
	}

	if board[0][0] == t && board[1][1] == t && board[2][2] == t {
		return true
	}

	if board[2][0] == t && board[1][1] == t && board[0][2] == t {
		return true
	}
	return false
}
