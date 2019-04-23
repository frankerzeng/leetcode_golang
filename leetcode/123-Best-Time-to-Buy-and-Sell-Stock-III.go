/*
123. Best Time to Buy and Sell Stock III
Say you have an array for which the ith element is the price of a given stock on day i.

Design an algorithm to find the maximum profit. You may complete at most two transactions.

Note: You may not engage in multiple transactions at the same time (i.e., you must sell the stock before you buy again).

Example 1:

Input: [3,3,5,0,0,3,1,4]
Output: 6
Explanation: Buy on day 4 (price = 0) and sell on day 6 (price = 3), profit = 3-0 = 3.
             Then buy on day 7 (price = 1) and sell on day 8 (price = 4), profit = 4-1 = 3.
Example 2:

Input: [1,2,3,4,5]
Output: 4
Explanation: Buy on day 1 (price = 1) and sell on day 5 (price = 5), profit = 5-1 = 4.
             Note that you cannot buy on day 1, buy on day 2 and sell them later, as you are
             engaging multiple transactions at the same time. You must sell before buying again.
Example 3:

Input: [7,6,4,3,1]
Output: 0
Explanation: In this case, no transaction is done, i.e. max profit = 0.
*/
package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{3, 3, 5, 0, 0, 3, 1, 4}))

}
func maxProfit(prices []int) int {
	k := 2
	if len(prices) <= 1 || k <= 0 {
		return 0
	}

	if k >= len(prices)/2 {
		max := 0
		for i := 1; i < len(prices); i++ {
			if prices[i] > prices[i-1] {
				max += prices[i] - prices[i-1]
			}
		}
		return max
	}

	var arr [][]int // arr[i][j] 表示最多i次交易截止到j天的最大收益
	for i := 0; i <= k+1; i++ {
		var arrTmp []int
		for j := 0; j <= len(prices); j++ {
			arrTmp = append(arrTmp, 0)
		}
		arr = append(arr, arrTmp)
	}

	for i := 1; i <= k; i++ {
		for j := 1; j < len(prices); j++ {
			subMaxTmp := prices[j] - prices[0] + arr[i-1][0]
			for jj := 1; jj < j; jj++ {
				subMaxTmp = max(subMaxTmp, prices[j]-prices[jj]+arr[i-1][jj])
			}

			arr[i][j] = max(arr[i][j-1], subMaxTmp)
		}
	}

	return arr[k][len(prices)-1]
}

func max(v1, v2 int) int {
	if v1 > v2 {
		return v1
	}
	return v2
}
