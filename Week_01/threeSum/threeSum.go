package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	//1.难点在于去重，排序+双指针
	//排序
	sort.Ints(nums)
	result := make([][]int, 0)
	length := len(nums)
	for i := 0; i < length; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		target := -nums[i]
		left := i + 1
		right := length - 1
		for left < right {
			if nums[left]+nums[right] == target {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if nums[left]+nums[right] > target {
				right--
			} else {
				left++
			}
		}
	}

	return result
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
}
