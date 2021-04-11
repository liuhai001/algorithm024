package main

//动态规划
//dp[i][j] 表示s前i个字符，p前j个字符 是否匹配上
func isMatch(s string, p string) bool {
	m := len(s)
	n := len(p)
	dp := make([][]bool, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]bool, n+1)
	}
	//特殊情况初始化
	dp[0][0] = true //s,p 都是空串
	for i := 1; i < n+1; i++ {
		if p[i-1] == '*' {
			dp[0][i] = true
		} else {
			break
		}
	}
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if s[i-1] == p[j-1] || p[j-1] == '?' {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				//用* 或者不用* 匹配
				dp[i][j] = dp[i][j-1] || dp[i-1][j]
			}
		}
	}

	return dp[m][n]
}
