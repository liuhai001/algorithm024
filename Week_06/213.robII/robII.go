package main

//打家劫舍II
//打家劫舍I的升级版，分为两种情况，不偷第一个房子，不偷最后一个房子
func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	return max(rob_1(nums[1:]), rob_1(nums[:n-1]))
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
