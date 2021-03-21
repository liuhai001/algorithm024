package main

func generateParenthesis(n int) []string {
	res := make([]string, 0)
	var dfs func(left, right int, current string)
	dfs = func(left, right int, current string) {
		//终止条件
		if left == n && right == n {
			res = append(res, current)
			return
		}
		//数据处理

		//下一层
		//剪枝
		if left < n {
			dfs(left+1, right, current+"(")
		}
		//剪枝
		if left <= n && left > right {
			dfs(left, right+1, current+")")
		}
		//恢复状态
	}
	dfs(0, 0, "")
	return res
}
