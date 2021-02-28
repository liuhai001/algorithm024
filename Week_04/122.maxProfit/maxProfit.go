package main

import "fmt"

//股票买卖的最佳时机，贪心算法
func maxProfit(prices []int) int {
	sum := 0
	for i := 1; i <= len(prices)-1; i++ {
		sum += max(0, prices[i]-prices[i-1])
	}

	return sum

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}
