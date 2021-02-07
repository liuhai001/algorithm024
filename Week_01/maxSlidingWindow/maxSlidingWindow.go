package main

import (
	"fmt"
	"math"
)

func maxSlidingWindow(nums []int, k int) []int {
	//1.暴力解法 时间复杂度 O((n-k+1)*(k))
	result := make([]int, 0)
	length := len(nums)
	if length == 0 {
		return result
	}
	for i := 0; i < length-k+1; i++ {
		maxNum := math.MinInt32
		for j := i; j < i+k; j++ {
			if nums[j] > maxNum {
				maxNum = nums[j]
			}
		}
		result = append(result, maxNum)
	}
	return result
}

func maxSlidingWindow_1(nums []int, k int) []int {
	//1.暴力解法 时间复杂度 O((n-k+1)*(k))
	length := len(nums)
	if length == 0 {
		return []int{}
	}

	result := make([]int, length-k+1, length-k+1)
	max := nums[0]
	for i := 1; i < k; i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}
	result[0] = max
	for i := 1; i < length-k+1; i++ {
		if nums[i-1] == max {
			max = nums[i]
			for j := i + 1; j < i+k; j++ {
				if nums[j] > max {
					max = nums[j]
				}
			}
		} else {
			if nums[i+k-1] > max {
				max = nums[i+k-1]
			}
		}

		result[i] = max
	}
	return result
}

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	fmt.Println(maxSlidingWindow_1(nums, 3))
}
