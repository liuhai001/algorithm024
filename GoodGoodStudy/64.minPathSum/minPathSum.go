package main

import "fmt"

//动态规划：三步走
//找重复性（分治）
//定义状态数组 f(i,j)
//dp方程 f(i,j) = min{f(i+1,j),f(i,j+1)} + grid[i][j]
func minPathSum(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])

	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 {
				dp[i][j] = dp[i][j+1] + grid[i][j]
			} else if j == n-1 {
				dp[i][j] = dp[i+1][j] + grid[i][j]
			} else {
				dp[i][j] = min(dp[i+1][j], dp[i][j+1]) + grid[i][j]
			}
		}
	}

	return dp[0][0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	input := [][]int{
		[]int{1, 3, 1},
	}
	fmt.Println(minPathSum(input))
}
