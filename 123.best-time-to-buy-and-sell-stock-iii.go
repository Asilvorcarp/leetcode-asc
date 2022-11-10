/*
 * @lc app=leetcode.cn id=123 lang=golang
 *
 * [123] Best Time to Buy and Sell Stock III
 */

package leetcode

// @lc code=start

import (
	"fmt"
	"math"
)

// with record
func maxProfit(prices []int) int {
	bou1 := -math.MaxInt
	sel1 := 0
	bou2 := -math.MaxInt
	sel2 := 0
	path := [4][4]int{} // all init as 0
	fmt.Printf("%v\n", path)
	for i, p := range prices {
		if -p > bou1 {
			path[0][0] = i
			bou1 = -p
		}
		if bou1+p > sel1 {
			path[1] = path[0]
			path[1][1] = i
			sel1 = bou1 + p
		}
		if sel1-p > bou2 {
			path[2] = path[1]
			path[2][2] = i
			bou2 = sel1 - p
		}
		if bou2+p > sel2 {
			path[3] = path[2]
			path[3][3] = i
			sel2 = bou2 + p
		}
	}
	fmt.Printf("%v\n", path)
	fmt.Printf("final path: %v\n", path[3])
	return sel2
}

// easy to understand O(n)/O(1) solution
func maxProfit_On(prices []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	bou1 := -prices[0]
	sel1 := 0
	bou2 := -prices[0]
	sel2 := 0
	for i := 1; i < len(prices); i++ {
		bou1 = max(bou1, -prices[i])
		sel1 = max(sel1, bou1+prices[i])
		bou2 = max(bou2, sel1-prices[i])
		sel2 = max(sel2, bou2+prices[i])
	}
	return sel2
}

// all np stored, O(n)/O(n)
func maxProfit_np(prices []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	type npElem struct {
		// none int // always 0
		bou1 int // bought 1
		sel1 int // sold 1
		bou2 int // bought 2
		sel2 int // sold 2
	}
	np := make([]npElem, len(prices))
	np[0] = npElem{
		bou1: -prices[0],
		sel1: 0,
		bou2: -prices[0],
		sel2: 0,
	}
	for i := 1; i < len(np); i++ {
		np[i] = npElem{
			bou1: max(np[i-1].bou1, -prices[i]),
			sel1: max(np[i-1].sel1, np[i-1].bou1+prices[i]),
			bou2: max(np[i-1].bou2, np[i-1].sel1-prices[i]),
			sel2: max(np[i-1].sel2, np[i-1].bou2+prices[i]),
		}
	}
	return np[len(np)-1].sel2
}

// @lc code=end
