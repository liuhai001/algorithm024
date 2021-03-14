package main

func jump(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		if i == n-1 {
			dp[i] = 0
		} else {
			tmpMin := n
			for j := i + 1; j <= nums[i]+i && j < n; j++ {
				if dp[j] < tmpMin {
					tmpMin = dp[j]
				}
			}
			dp[i] = tmpMin + 1
		}
	}
	return dp[0]
}
