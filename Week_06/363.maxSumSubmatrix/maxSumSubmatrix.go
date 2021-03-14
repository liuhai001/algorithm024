package main

import (
	"fmt"
	"math"
)

//如果行远大于列，就固定左右列边界，相反固定上下行边界
//时间复杂度O(m*n*n)
func maxSumSubmatrix(matrix [][]int, k int) int {
	for len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	maxSum, rows, cols := math.MinInt32, len(matrix), len(matrix[0])
	//时间复杂度O(m*n*n)
	for left := 0; left < cols; left++ {
		rowSums := make([]int, rows)
		// 扩展一列求一次值，并更新小于k的最大值 O(m*n)
		for right := left; right < cols; right++ {
			// 逐行扩展列 O(m)
			for row := 0; row < rows; row++ {
				rowSums[row] += matrix[row][right]
			}
			// 求取新值 O(m)
			maxSum = max(maxSum, maxKSubArray(rowSums, k))
			if maxSum == k {
				return maxSum
			}
		}
	}
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//时间复杂度 最快O(m) m 为matrix 的行数
func maxKSubArray(nums []int, k int) int {
	rollSum, rollMax := nums[0], nums[0]
	// O(rows) 动态规划
	for i := 1; i < len(nums); i++ {
		if rollSum > 0 {
			rollSum += nums[i]
		} else {
			rollSum = nums[i]
		}
		if rollSum > rollMax {
			rollMax = rollSum
		}
	}
	if rollMax <= k {
		return rollMax
	}

	//rollMax > k 情况下 只能暴力求解了
	max := math.MinInt32
	for l := 0; l < len(nums); l++ {
		sum := 0
		for r := l; r < len(nums); r++ {
			sum += nums[r]
			if sum > max && sum <= k {
				max = sum
			}
			if max == k {
				return k
			}
		}
	}
	return max
}

func main() {
	matrix := [][]int{[]int{2, 2, -1}}
	fmt.Println(maxSumSubmatrix(matrix, 0))
}
