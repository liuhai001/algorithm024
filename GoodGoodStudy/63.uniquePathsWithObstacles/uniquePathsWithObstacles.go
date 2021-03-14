package main

//动态规划：三步走
//找重复性（分治）
//定义状态数组
//dp方程
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	if m == 0 {
		return 0
	}
	n := len(obstacleGrid[0])
	dp := make([]int, n+1) //start 到{i,j} 的路径和

	if obstacleGrid[0][0] == 0 {
		dp[1] = 1
	}

	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if obstacleGrid[i-1][j-1] == 1 {
				dp[j] = 0
				continue
			}
			dp[j] += dp[j-1]
		}
	}
	return dp[n]
}
