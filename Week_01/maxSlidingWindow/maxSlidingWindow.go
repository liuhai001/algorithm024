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

//单调队列
func maxSlidingWindow_2(nums []int, k int) []int {
	res := make([]int, 0, len(nums))
	//异常情况
	if len(nums) == 0 || k < 0 || k > len(nums) {
		return res
	}
	length := len(nums)
	//双端队列，里面存的是index
	dequeue := []int{}
	for i := 0; i < length; i++ {
		//窗口滑动
		if len(dequeue) > 0 && dequeue[0] < i-k+1 {
			dequeue = dequeue[1:]
		}
		//保持队列单调
		for len(dequeue) > 0 && nums[i] > nums[dequeue[len(dequeue)-1]] {
			dequeue = dequeue[:len(dequeue)-1]
		}
		dequeue = append(dequeue, i)

		if i >= k-1 {
			res = append(res, nums[dequeue[0]])
		}
	}

	return res
}

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	fmt.Println(maxSlidingWindow_2(nums, 3))
}
