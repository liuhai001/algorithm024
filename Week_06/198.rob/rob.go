package main

//打家劫舍
//动态规划
//三步走
//重复子问题 a[i] a[i-1] nums[i]
//定义状态：a[i] 0~i 之间打家劫舍的最大值，i 可偷可不偷
//dp方程：dp[i] = max(dp[i-1],dp[i-2]+nums[i])

//时间复杂度O(n) 空间复杂度O(n)
func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}

	dp := make([]int, n)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])

	}
	return dp[n-1]
}

//时间复杂度O(n) 空间复杂度O(1)
func rob_1(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}

	pre := nums[0]
	cur := max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		pre, cur = cur, max(cur, pre+nums[i])

	}
	return cur
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
