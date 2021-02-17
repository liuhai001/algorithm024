package main

import "fmt"

func generateParenthesis(n int) []string {
	res := make([]string, 0)
	recursion(0, 0, n, &res, "")
	return res
}

func recursion(left, right, n int, res *[]string, s string) {
	//终止条件
	if left == n && right == n {
		*res = append(*res, s)
		return
	}

	//数据处理


	//进入下一层
	//左括号随时可以加,只要别超标
	if left < n {
		recursion(left+1, right, n, res, s+"(")

	}
	//右括号之前一定要有左括号
	if left > right {
		recursion(left, right+1, n, res, s+")")
	}
	//状态恢复
}

func main() {
	fmt.Println(generateParenthesis(3))
}
