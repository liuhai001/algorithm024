package main

//动态规划：三步走
//找重复性（分治）
//定义状态数组
//dp方程
func uniquePaths(m int, n int) int {
	dp := make([]int, n+1) //start 到{i,j} 的路径和
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if i == 1 {
				dp[j] = 1
			} else {
				dp[j] += dp[j-1]
			}
		}
	}
	return dp[n]
}
