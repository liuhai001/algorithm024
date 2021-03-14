package main

import (
	"fmt"
)

//分析问题：这里不能用贪心算法，因为金币面额不是固定的，有可能不是全局最优解
//只能用动态规划思想
//三步走：
//重复子问题（分治）
//bottom-up, dp[amount] = min(dp[amount-coins[i]]) + 1
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if coins[j] <= i {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
				fmt.Println(i, dp[i])
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	coins := []int{2,5,10,1}
	amount := 2
	coinChange(coins, amount)
}
