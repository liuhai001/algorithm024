package main

//动态规划：三步走
//找重复性（分治）
//定义状态数组
//dp方程
//字符串的问题，一般要转化成二维数组来解
func longestCommonSubsequence(text1 string, text2 string) int {
	//行，列
	m := len(text1)
	n := len(text2)
	if m == 0 || n == 0 {
		return 0
	}
	//初始化一个[m+1][n+1]int 二维数组
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
