package main

import "sort"

//排序，双指针
func fourSum(nums []int, target int) [][]int {
	res := make([][]int, 0)
	length := len(nums)
	if length < 4 {
		return res
	}
	//排序
	sort.Ints(nums)
	for i := 0; i <= length-4; i++ {
		//去重
		for i > 0 && i <= length-4 && nums[i] == nums[i-1] {
			i++
		}
		for j := i + 1; j <= length-3; j++ {
			//去重
			for j > i+1 && j <= length-3 && nums[j] == nums[j-1] {
				j++
			}
			tmpTarget := target - (nums[i] + nums[j])
			left := j + 1
			right := length - 1
			for left < right {
				sum := nums[left] + nums[right]
				if sum == tmpTarget {
					res = append(res, []int{nums[i], nums[j], nums[left], nums[right]})
					//继续找
					for left < right && nums[left] == nums[left+1] {
						//去重
						left++
					}
					for left < right && nums[right] == nums[right-1] {
						//去重
						right--
					}
					left++
					right--
				} else if sum < tmpTarget {
					for left < right && nums[left] == nums[left+1] {
						//去重
						left++
					}
					left++
				} else {
					for left < right && nums[right] == nums[right-1] {
						//去重
						right--
					}
					right--
				}
			}

		}
	}

	return res
}
